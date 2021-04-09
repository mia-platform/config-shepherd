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

package utils

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/afero"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

// Options global option for the cli that can be passed to all commands
type Options struct {
	Config      *genericclioptions.ConfigFlags
	SplittedMap string
}

// fs return the file system to use by default (override it for tests)
var fs = &afero.Afero{Fs: afero.NewOsFs()}

// CheckError default error handling function
func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// ExtractFilesFromDir extracts from a directory the global path to files contained in it.
// This function does not look into subdirs.
func ExtractFilesFromDir(directoryPath string) (map[string]string, error) {
	filesPath := map[string]string{}

	files, err := fs.ReadDir(directoryPath)
	if err != nil {
		return nil, err
	}
	for _, path := range files {
		if !path.IsDir() {
			filesPath[path.Name()] = filepath.Join(directoryPath, path.Name())
		}
	}
	return filesPath, nil
}

// IsADir check if a path is a directory or not
func IsADir(path string) (bool, error) {
	return fs.IsDir(path)
}

// ReadFile read a file from the file system
func ReadFile(filename string) ([]byte, error) {
	return fs.ReadFile(filename)
}

// MkdirAll create a folder
func MkdirAll(name string) error {
	return fs.MkdirAll(name, os.FileMode(0755))
}

// RemoveAll removes a directory path and any children it contains.
func RemoveAll(path string) error {
	return fs.RemoveAll(path)
}

// CreateFile create a new file in path
func CreateFile(path string) (afero.File, error) {
	return fs.Create(path)
}

// WriteFile write data to file
func WriteFile(filename string, data []byte) error {
	return fs.WriteFile(filename, data, os.FileMode(0644))
}
