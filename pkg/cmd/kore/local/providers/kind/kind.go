/**
 * Copyright 2020 Appvia Ltd <info@appvia.io>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package kind

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"text/template"
	"time"

	"github.com/appvia/kore/pkg/client/config"
	"github.com/appvia/kore/pkg/cmd/kore/local/providers"
	"github.com/appvia/kore/pkg/utils"

	"github.com/spf13/cobra"
)

const (
	// KindURL is the release URL for kind
	KindURL = "https://github.com/kubernetes-sigs/kind/releases/download/v0.8.1/kind-linux-%s"
)

type providerImpl struct {
	utils.Logger
	// path is the file path to the kind binary
	path string
	// options are the configurables
	options providers.CreateOptions
}

var (
	// KindConfiguration is the kind configuration
	KindConfiguration = `
apiVersion: kind.x-k8s.io/v1alpha4
kind: Cluster
nodes:
- role: control-plane
  image: {{ .Image }}
  extraPortMappings:
	{{- if not .DisableUI }}
  - containerPort: 3000
    hostPort: 3000
    protocol: TCP
	{{- end }}
  - containerPort: 10080
    hostPort: 10080
    protocol: TCP
  - containerPort: 10443
    hostPort: 10443
    protocol: TCP
`
)

var (
	// kindVersion is the version of kind image
	kindVersion = "kindest/node:v1.16.9@sha256:7175872357bc85847ec4b1aba46ed1d12fa054c83ac7a8a11f5c268957fd5765"
	// loadImagess is a collection of images to load after creating the cluster
	loadImages []string
)

// AddProviderFlags allows kind to add provider specific flags
func AddProviderFlags(cmd *cobra.Command) {
	flags := cmd.Flags()
	flags.StringVar(&kindVersion, "kind-image", kindVersion, "version of the kind image to use")
	flags.StringSliceVar(&loadImages, "kind-load-image", []string{}, "collection of images to load after creating cluster")
}

// New creates and returns a kind provider
func New(logger utils.Logger) (providers.Interface, error) {
	return &providerImpl{Logger: logger}, nil
}

// Destroy is responsible for deleting the cluster
func (p *providerImpl) Destroy(ctx context.Context, name string) error {
	found, err := p.Has(ctx, name)
	if err != nil {
		return err
	}
	if !found {
		return nil
	}
	args := []string{
		"delete",
		"cluster",
		"--name", name,
	}
	p.Infof("Deleting the kind cluster: %q\n", name)

	return exec.CommandContext(ctx, p.path, args...).Run()
}

// Create is responsible for provisioning a kind cluster
func (p *providerImpl) Create(ctx context.Context, name string, options providers.CreateOptions) error {
	// @step: adding to give some context for longer up times
	found, err := p.Has(ctx, name)
	if err != nil {
		return err
	}
	if found {
		p.Infof("Kind cluster: %q already exists, skipping creation\n", name)

		return p.ensureRunning(ctx, name)
	}

	p.Infof("Using Kind image: %q\n", KindImageShortName())

	// @step: adding to give some context for longer up times
	found, err = hasDockerImage(ctx, KindImageShortName())
	if err != nil {
		return err
	}
	if !found {
		p.Infof("Kind image not found, using docker pull (takes a while depending network)\n")
	}

	start := time.Now()

	args := []string{
		"create",
		"cluster",
		"--name", name,
		"--wait", "10m",
		"--config=-",
	}
	p.Infof("Provisioning a kind cluster: %q (usually takes 1-3mins)\n", name)

	cmd := exec.CommandContext(ctx, p.path, args...)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	config, err := GetKindConfiguration(options)
	if err != nil {
		return err
	}

	if _, err := io.WriteString(stdin, config); err != nil {
		return err
	}
	stdin.Close()

	// @note: largely for user feedback so they are not left with an empty prompt
	cutoff, cancel := context.WithCancel(ctx)
	defer cancel()

	go func() {
		interval := 20 * time.Second
		start := time.Now()
		time.Sleep(interval)

		_ = utils.RetryWithTimeout(cutoff, 10*time.Minute, interval, func() (bool, error) {
			p.Infof("Still building the kind cluster %q: %s\n", name, utils.HumanDuration(time.Since(start)))

			return false, nil
		})
	}()

	if combined, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("%s", combined)
	}

	p.Infof("Built local kind cluster in %s\n", utils.HumanDuration(time.Since(start)))

	return p.ensureImages(ctx, name)
}

// Export is responsible for exporting the kind kubeconfig
func (p *providerImpl) Export(ctx context.Context, name string) (string, error) {
	args := []string{
		"export",
		"kubeconfig",
		"--name", name,
	}
	p.Infof("Exporting kubeconfig from kind cluster: %q\n", name)

	combined, err := exec.CommandContext(ctx, p.path, args...).CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("trying to export kubeconfig: %s", combined)
	}

	contextName := fmt.Sprintf("kind-%s", name)

	return contextName, nil
}

// Has checks if a kind cluster already exists
func (p *providerImpl) Has(ctx context.Context, name string) (bool, error) {
	args := []string{
		"get",
		"clusters",
	}
	p.Infof("Checking if kind cluster: %q already exists\n", name)

	combined, err := exec.CommandContext(ctx, p.path, args...).CombinedOutput()
	if err != nil {
		return false, err
	}

	scanner := bufio.NewScanner(bytes.NewReader(combined))
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), name) {
			return true, nil
		}
	}

	return false, nil
}

// Stop is called to stop the kind instance
func (p *providerImpl) Stop(ctx context.Context, name string) error {
	found, err := p.Has(ctx, name)
	if err != nil {
		return err
	}
	if !found {
		return nil
	}

	args := []string{
		"stop",
		"kore-control-plane",
	}
	p.Infof("Ensuring the kind cluster: %q is stopped\n", name)

	path, err := exec.LookPath("docker")
	if err != nil {
		return errors.New("missing binary: docker in $PATH")
	}

	return exec.CommandContext(ctx, path, args...).Run()
}

func (p *providerImpl) Preflight(ctx context.Context) error {
	path, err := exec.LookPath("kind")
	if err != nil {
		path = filepath.Join(filepath.Join(config.GetClientPath(), "build"), "kind")

		found, err := utils.FileExists(path)
		if err != nil {
			return err
		}
		if found {
			p.Infof("Kind binary requirement found\n")
			p.path = path

			return nil
		}
		p.Infof("Kind binary not found in $PATH\n")

		if p.options.AskConfirmation {
			p.Infof("Download: %s (%s) (y/N)? ", getReleaseURL(), path)
			if ok := utils.AskForConfirmation(os.Stdin); !ok {
				return errors.New(`missing binary: "kind" in $PATH`)
			}
		}
		p.Infof("Attempting to download the kind binary\n")

		if err := utils.DownloadFile(ctx, path, getReleaseURL()); err != nil {
			return err
		}

		if err := os.Chmod(path, os.FileMode(0500)); err != nil {
			return err
		}
	} else {
		p.Infof("Kind binary requirement found in $PATH\n")
	}

	p.path = path

	_, err = exec.LookPath("docker")
	if err != nil {
		return errors.New("missing binary: docker in $PATH")
	}
	p.Infof("Docker binary requirement found in $PATH\n")

	return nil
}

// GetKindConfiguration returns the kind config
func GetKindConfiguration(options providers.CreateOptions) (string, error) {
	tmpl, err := template.New("main").Parse(KindConfiguration)
	if err != nil {
		return "", err
	}
	values := map[string]interface{}{
		"DisableUI": options.DisableUI,
		"Image":     kindVersion,
	}
	b := &bytes.Buffer{}
	if err := tmpl.Execute(b, values); err != nil {
		return "", err
	}

	return b.String(), nil
}

func (p *providerImpl) ensureRunning(ctx context.Context, name string) error {
	args := []string{
		"start",
		"kore-control-plane",
	}
	p.Infof("Ensuring the kind cluster: %q is running\n", name)

	path, err := exec.LookPath("docker")
	if err != nil {
		return err
	}

	if err := exec.CommandContext(ctx, path, args...).Run(); err != nil {
		return err
	}

	return p.ensureImages(ctx, name)
}

// KindImage returns the full image
func KindImage() string {
	return kindVersion
}

// KindImageShortName return image minus any SHA
func KindImageShortName() string {
	return strings.Split(kindVersion, "@")[0]
}

// hasDockerImage checks if the image exists
func hasDockerImage(ctx context.Context, image string) (bool, error) {
	args := []string{
		"images",
		"--format",
		`"{{ .Repository }}:{{ .Tag }}"`,
	}
	image = strings.Split(image, "@")[0]

	path, err := exec.LookPath("docker")
	if err != nil {
		return false, errors.New("missing binary: docker in $PATH")
	}

	combined, err := exec.CommandContext(ctx, path, args...).CombinedOutput()
	if err != nil {
		return false, nil
	}

	scanner := bufio.NewScanner(bytes.NewReader(combined))
	for scanner.Scan() {
		if strings.HasPrefix(strings.ReplaceAll(scanner.Text(), `"`, ""), image) {
			return true, nil
		}
	}

	return false, nil
}

func (p *providerImpl) ensureImages(ctx context.Context, name string) error {
	if len(loadImages) == 0 {
		return nil
	}

	for _, image := range loadImages {
		p.Infof("Attempting to load docker image: %s into cluster\n", image)

		err := utils.RetryWithTimeout(ctx, 2*time.Minute, 5*time.Second, func() (bool, error) {
			args := []string{
				"load",
				"docker-image", image,
				"--name", name,
			}
			if err := exec.CommandContext(ctx, p.path, args...).Run(); err != nil {
				return false, nil
			}

			return true, nil
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func getReleaseURL() string {
	return fmt.Sprintf(KindURL, runtime.GOARCH)
}
