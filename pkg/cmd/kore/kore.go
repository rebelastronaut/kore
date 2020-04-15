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

package kore

import (
	"strings"

	"github.com/appvia/kore/pkg/client"
	"github.com/appvia/kore/pkg/client/config"
	"github.com/appvia/kore/pkg/cmd/kore/apiresources"
	"github.com/appvia/kore/pkg/cmd/kore/apply"
	"github.com/appvia/kore/pkg/cmd/kore/create"
	"github.com/appvia/kore/pkg/cmd/kore/delete"
	"github.com/appvia/kore/pkg/cmd/kore/edit"
	"github.com/appvia/kore/pkg/cmd/kore/get"
	"github.com/appvia/kore/pkg/cmd/kore/kubeconfig"
	"github.com/appvia/kore/pkg/cmd/kore/login"
	"github.com/appvia/kore/pkg/cmd/kore/profiles"
	cmdutils "github.com/appvia/kore/pkg/cmd/utils"
	"github.com/appvia/kore/pkg/utils/render"
	"github.com/appvia/kore/pkg/version"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	root *cobra.Command
)

// NewKoreCommand creates and returns the kore command
func NewKoreCommand(streams cmdutils.Streams) (*cobra.Command, error) {
	// @step: create or read in the client configuration
	config, err := config.GetOrCreateClientConfiguration()
	if err != nil {
		return nil, err
	}
	// we create an client from the configuration
	client, err := client.New(config)
	if err != nil {
		return nil, err
	}

	// @step: create a factory for the commands
	factory, err := cmdutils.NewFactory(client, streams, *config)
	if err != nil {
		return nil, err
	}

	// root represents the base command when called without any subcommands
	root = &cobra.Command{
		Use:          "kore",
		Short:        "kore provides a cli for the " + version.Prog,
		Example:      "kore command [options] [-t|--team]",
		SilenceUsage: true,
		Run:          cmdutils.RunHelp,

		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if cmdutils.GetVerbose(cmd) {
				log.SetLevel(log.DebugLevel)
			}
			if cmdutils.GetDebug(cmd) {
				log.SetLevel(log.TraceLevel)
			}

			log.WithField("profile", config.CurrentProfile).Debug("running with the selected profile")
		},
	}

	flags := root.PersistentFlags()
	flags.Bool("force", false, "is used to force an operation to happen (defaults: false)")
	flags.StringP("team", "t", config.GetCurrentProfile().Team, "the team you are operating within")
	flags.StringP("output", "o", "table", "the output format of the resource ("+strings.Join(render.SupportedFormats(), ",")+")")
	flags.BoolP("no-wait", "", false, "indicates if we should wait for resources to provision")
	flags.BoolP("show-headers", "", true, "indicates we should display headers on table out (defaults: true)")
	flags.BoolP("debug", "", false, "indicates we should use debug / trace logging (defaults: false)")
	flags.Bool("verbose", false, "enables verbose logging for debugging purposes (defaults: false)")

	// @step: add all the commands to the root
	root.AddCommand(
		login.NewCmdLogin(factory),
		login.NewCmdLogout(factory),
		NewCmdCompletion(factory),
		apply.NewCmdApply(factory),
		get.NewCmdGet(factory),
		delete.NewCmdDelete(factory),
		edit.NewCmdEdit(factory),
		profiles.NewCmdProfiles(factory),
		create.NewCmdCreate(factory),
		kubeconfig.NewCmdKubeConfig(factory),
		NewCmdWhoami(factory),
		apiresources.NewCmdAPIResources(factory),
		NewCmdVersion(factory),
	)

	// @step: seriously cobra is pretty damn awesome
	cmdutils.MustRegisterFlagCompletionFunc(root, "team", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.BashCompDirective) {
		list, err := factory.Resources().LookupResourceNames("team", "")
		if err != nil {
			return nil, cobra.BashCompDirectiveError
		}

		return list, cobra.BashCompDirectiveDefault
	})

	cmdutils.MustRegisterFlagCompletionFunc(root, "output", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.BashCompDirective) {
		return render.SupportedFormats(), cobra.BashCompDirectiveDefault
	})

	return root, nil
}
