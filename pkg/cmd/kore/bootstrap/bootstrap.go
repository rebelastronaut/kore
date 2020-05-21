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

package bootstrap

import (
	"fmt"
	"strings"

	"github.com/appvia/kore/pkg/cmd/kore/bootstrap/providers"
	"github.com/appvia/kore/pkg/cmd/kore/bootstrap/providers/kind"
	cmdutil "github.com/appvia/kore/pkg/cmd/utils"

	"github.com/spf13/cobra"
)

var (
	usage = `
Bootstrap provides an experimental means of bootstrapping a local Kore installation. At
present the local installation use "kind" https://github.com/kubernetes-sigs/kind.

Unless specified otherwise it will deploy an official tagged release from Github, though
this can be overriden using the --release flag. Note the installation is performed via
helm with a local values.yaml is generated in the directory; so if you wish to change
any of the values post a installation can change this file and run the 'up' command.

Note, the local installtions data persistency is tied to the provider. For kind as long
as the container is not delete the data is kept.
`
	examples = `
# Provision a local kore instance called 'kore' (defaults to kind)
$ kore alpha bootstrap up

# Override the release and use a local chart
$ kore alpha bootstrap up --release ./charts
$ kore alpha bootstrap up --release https://URL

# Destroy the local installtion
$ kore alpha bootstrap destroy

# To stop the local installed without deleting the data
$ kore alpha bootstrap stop

The application should be available on http://127.0.0.1:3000. You can provision the
CLI via.

$ kore login -a http://127.0.0.1:10080 local

Post the command your Kubectl context is switched to the kind installation:

$ kubectl config current-context
`
)

var (
	// GithubRelease is the link to release
	GithubRelease = "https://github.com/appvia/kore/releases/download/%s/kore-helm-chart-%s.tgz"
	// ClusterName is the name of the cluster to create
	ClusterName = "kore"
)

const (
	// Kubectl is the name of the binary
	Kubectl = "kubectl"
)

// NewCmdBootstrap creates and returns the delete command
func NewCmdBootstrap(factory cmdutil.Factory) *cobra.Command {
	command := &cobra.Command{
		Use:     "bootstrap",
		Short:   "Provides the provision of local installation for Kore for testing",
		Long:    usage,
		Example: examples,
		Run:     cmdutil.RunHelp,
	}

	command.AddCommand(
		NewCmdBootstrapDestroy(factory),
		NewCmdBootstrapUp(factory),
		NewCmdBootstrapStop(factory),
	)

	return command
}

// GetHelmReleaseURL returns the helm release for kore
func GetHelmReleaseURL(release string) string {
	if strings.HasPrefix(release, "v") {
		return fmt.Sprintf(GithubRelease, release, release)
	}

	return release
}

// GetProvider returns the provider implementation
func GetProvider(f cmdutil.Factory, name string) (providers.Interface, error) {
	switch name {
	case "kind":
		return kind.New(newProviderLogger(f))
	default:
		return nil, fmt.Errorf("unknown provider: %s", name)
	}
}
