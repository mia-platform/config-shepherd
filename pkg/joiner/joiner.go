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

package joiner

import (
	"fmt"
	"path/filepath"

	"git.tools.mia-platform.eu/platform/devops/config-shepherd/internal/utils"
)

// Run execute the joiner command from cli
func Run(inputDirs []string, outDir string) {
	absOutDir, err := filepath.Abs(outDir)
	utils.CheckError(err)

	err = utils.MkdirAll(absOutDir)
	utils.CheckError(err)

	filesMapping, err := extractFilesParts(inputDirs)
	utils.CheckError(err)

	err = joinAllFiles(filesMapping, absOutDir)
	utils.CheckError(err)
}

// ExtractAllFiles return a map of all filenames from array of directories
func extractFilesParts(paths []string) (map[string][]string, error) {
	fmt.Printf("EXTRACTING files parts from %s\n", paths)
	filesMapping := map[string][]string{}
	// Extract files from directories
	for _, path := range paths {
		// get absolute path for good measure
		globalPath, err := filepath.Abs(path)
		if err != nil {
			return nil, err
		}

		isADirectory, err := utils.IsADir(globalPath)
		if err != nil {
			fmt.Printf("WARN: can't read input file at path %s\n", globalPath)
			continue
		}

		if isADirectory {
			pathsInDirectory, err := utils.ExtractFilesFromDir(globalPath)
			if err != nil {
				return nil, err
			}
			for key, path := range pathsInDirectory {
				if val, ok := filesMapping[key]; ok {
					filesMapping[key] = append(val, path)
				} else {
					filesMapping[key] = []string{path}
				}
			}
		}
	}
	return filesMapping, nil
}

// joinAllFiles join all parts and write them in a directory
func joinAllFiles(filesMapping map[string][]string, outDir string) error {
	for key, paths := range filesMapping {
		finalContent, err := joinFileParts(paths)
		utils.CheckError(err)

		filePath := filepath.Join(outDir, key)
		finalGlobalPath, err := filepath.Abs(filePath)
		utils.CheckError(err)

		fmt.Printf("CREATING file %s\n", filePath)
		_, err = utils.CreateFile(finalGlobalPath)
		utils.CheckError(err)

		err = utils.WriteFile(finalGlobalPath, finalContent)
		utils.CheckError(err)
	}
	return nil
}

// joinFileParts join the files parts into a single file
func joinFileParts(paths []string) ([]byte, error) {
	var finalContents []byte
	for _, path := range paths {
		content, err := utils.ReadFile(path)
		utils.CheckError(err)

		finalContents = append(finalContents, content...)
	}
	return finalContents, nil
}
