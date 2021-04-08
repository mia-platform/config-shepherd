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
	"os"
	"flag"
	"fmt"
	"runtime"

	"join-config-map/internal/cli"
	"github.com/spf13/cobra"
)

var options = cli.New()

func main() {
	rootCmd := &cobra.Command{
		Short: "mia-init-container CLI",
		Long: "Join Splitted Configmaps into one",
		Use: "mia-init-container", // da capire: mlp usa mlp

		SilenceErrors: true,
		SilenceUsage:  true,
		Example: "", // da capire: mlp usa heredoc: "github.com/MakeNowJust/heredoc/v2"
	}


	versionOutput := versionFormat(cli.Version, cli.BuildDate)
	// Version subcommand
	rootCmd.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "Show mia-init-container version",
		Long:  "Show mia-init-container version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(versionOutput)
		},
	})

	cli.AddGlobalFlags(rootCmd, options)
	cli.ConfigMapJoinerSubcommand(rootCmd, options)
	
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

	version = fmt.Sprintf("mia-init-container version: %s", version)
	// don't return GoVersion during a test run for consistent test output
	if flag.Lookup("test.v") != nil {
		return version
	}

	return fmt.Sprintf("%s, Go Version: %s", version, runtime.Version())
}
