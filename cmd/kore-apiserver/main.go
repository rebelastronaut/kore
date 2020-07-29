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

package main

import (
	"fmt"
	"os"

	"github.com/appvia/kore/cmd/kore-apiserver/options"
	"github.com/appvia/kore/pkg/cmd"
	"github.com/appvia/kore/pkg/version"
	"github.com/urfave/cli/v2"
)

func init() {
	cmd.DefaultLogging()
}

func main() {
	app := &cli.App{
		Name: "kore-apiserver",
		Authors: []*cli.Author{
			{
				Name:  version.Author,
				Email: version.Email,
			},
		},
		Flags:                options.Options(),
		Usage:                "Kore API provides the frontend api services",
		Version:              version.Version(),
		EnableBashCompletion: true,

		OnUsageError: func(context *cli.Context, err error, _ bool) error {
			fmt.Fprintf(os.Stderr, "[error] invalid options %s\n", err)
			return err
		},

		Action: func(ctx *cli.Context) error {
			return invoke(ctx)
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "[error] %s\n", err)

		os.Exit(1)
	}
}
