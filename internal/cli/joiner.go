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
	"encoding/json"
	"fmt"

	"git.tools.mia-platform.eu/platform/devops/config-shepherd/internal/utils"
	"git.tools.mia-platform.eu/platform/devops/config-shepherd/pkg/joiner"

	"github.com/spf13/cobra"
)

type splittedMapValue struct {
	Directories []string `json:"directories"`
}

// ConfigMapJoinerSubcommand add configMapJoiner subcommand to the main command
func ConfigMapJoinerSubcommand(cmd *cobra.Command, options *utils.Options) {
	var inputDirs []string
	var outputDir string

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

			var splittedMap = map[string]splittedMapValue{}
			if options.SplittedMap == "" {
				return fmt.Errorf("configuration not exist")
			}

			unmarshalSplittedMap, err := unmarshalSplittedMap([]byte(options.SplittedMap))
			utils.CheckError(err)
			splittedMap = unmarshalSplittedMap
			for outDirKey, splittedMapValue := range splittedMap {
				inputDirs = splittedMapValue.Directories
				joiner.Run(inputDirs, outDirKey)
			}
			return nil
		},
	}

	configMapJoinerCmd.Flags().StringSliceVar(&inputDirs, "input-dirs", []string{}, "folder paths containing data to join")
	configMapJoinerCmd.Flags().StringVar(&outputDir, "output-dir", "", "folder name where the joined files will be saved")
	cmd.AddCommand(configMapJoinerCmd)
}

func unmarshalSplittedMap(data []byte) (map[string]splittedMapValue, error) {
	var splittedMap = map[string]splittedMapValue{}
	if err := json.Unmarshal(data, &splittedMap); err != nil {
		return nil, err
	}
	return splittedMap, nil
}
