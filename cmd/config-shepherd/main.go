/*
 * Copyright 2021 Mia srl
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
	"flag"
	"fmt"
	"os"
	"runtime"

	"git.tools.mia-platform.eu/platform/devops/config-shepherd/internal/cli"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Short: "config-shepherd CLI",
		Long:  "Configure a container before lauching it",
		Use:   "config-shepherd",

		SilenceErrors: true,
		SilenceUsage:  true,
		Example:       `$ config-shepherd joiner --input-dirs './folder, ./folder2' --output-dirs 'output' --root-dir './rootDir'`,
	}

	versionOutput := versionFormat(cli.Version, cli.BuildDate)
	// Version subcommand
	rootCmd.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "Show config-shepherd version",
		Long:  "Show config-shepherd version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(versionOutput)
		},
	})

	cli.ConfigMapJoinerSubcommand(rootCmd)
	expandedArgs := []string{}
	if len(os.Args) > 0 {
		expandedArgs = os.Args[1:]
	}
	rootCmd.SetArgs(expandedArgs)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// versionFormat return the version string nicely formatted
func versionFormat(version, buildDate string) string {
	if buildDate != "" {
		version = fmt.Sprintf("%s (%s)", version, buildDate)
	}

	version = fmt.Sprintf("config-shepherd version: %s", version)
	// don't return GoVersion during a test run for consistent test output
	if flag.Lookup("test.v") != nil {
		return version
	}

	return fmt.Sprintf("%s, Go Version: %s", version, runtime.Version())
}
