// Copyright 2021 Mia srl
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cli

import (
	"join-config-map/internal/utils"
	"join-config-map/pkg/joiner"

	"github.com/spf13/cobra"
)

// ConfigMapJoinerSubcommand add configMapJoiner subcommand to the main command
func ConfigMapJoinerSubcommand(cmd *cobra.Command, options *utils.Options) {
	var inputDirs []string
	var outputDir string

	configMapJoinerCmd := &cobra.Command{
		Use:   "joiner",
		Short: "join splitted configmaps into one single configmap",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			joiner.Run(inputDirs, outputDir, options)
		},
	}

	configMapJoinerCmd.Flags().StringSliceVar(&inputDirs, "input-dirs", []string{}, "file and/or folder paths containing data to interpolate")
	configMapJoinerCmd.Flags().StringVar(&outputDir, "output-dir", "", "file and/or folder paths containing data to interpolate")
	cmd.AddCommand(configMapJoinerCmd)
}
