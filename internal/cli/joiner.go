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
	"fmt"

	"git.tools.mia-platform.eu/platform/devops/config-shepherd/pkg/joiner"

	"github.com/mia-platform/configlib"
	"github.com/spf13/cobra"
)

// ConfigMapJoinerSubcommand add configMapJoiner subcommand to the main command
func ConfigMapJoinerSubcommand(cmd *cobra.Command) {
	var inputDirs []string
	var outputDir string

	var configPath string
	var configName string

	configMapJoinerCmd := &cobra.Command{
		Use:   "joiner",
		Short: "join splitted files into one single file",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {

			if inputDirs != nil && outputDir != "" {
				joiner.Run(inputDirs, outputDir)
				return nil
			}

			if inputDirs != nil || outputDir != "" {
				return fmt.Errorf("inputDirs or outputDir is missing")
			}

			if configPath != "" && configName != "" {
				config, err := loadConfiguration(configPath, configName, configSchemaPath)
				if err != nil {
					return fmt.Errorf("errors: %s", err.Error())
				}

				for _, splittedMap := range config.SplittedMaps {
					inputDirs = splittedMap.InputMountPaths
					outputDir = splittedMap.OutputMountPath
					joiner.Run(inputDirs, outputDir)
				}
				return nil
			}

			return fmt.Errorf("configuration not passed")

		},
	}

	configMapJoinerCmd.Flags().StringSliceVar(&inputDirs, "input-dirs", nil, "folder paths containing data to join")
	configMapJoinerCmd.Flags().StringVar(&outputDir, "output-dir", "", "folder name where the joined files will be saved")
	configMapJoinerCmd.Flags().StringVar(&configName, "config-name", "", "name of the configuration file")
	configMapJoinerCmd.Flags().StringVar(&configPath, "config-path", "", "path of the configuration file")

	cmd.AddCommand(configMapJoinerCmd)
}

func loadConfiguration(path, filename, configSchemaPath string) (*Configuration, error) {
	jsonSchema, err := configlib.ReadFile(configSchemaPath)
	if err != nil {
		return nil, err
	}

	var config Configuration
	if err := configlib.GetConfigFromFile(filename, path, jsonSchema, &config); err != nil {
		return nil, err
	}

	return &config, err
}
