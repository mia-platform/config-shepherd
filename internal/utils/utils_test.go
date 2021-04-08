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
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

const testdata = "testdata/"

func TestExtractFilesFromDir(t *testing.T) {
	folder := filepath.Join(testdata, "folder")

	t.Run("Call ExtractFilesFromDir", func(t *testing.T) {
		filesPath, err := ExtractFilesFromDir(folder)

		expected := map[string]string{
			"pippo.json": "testdata/folder/pippo.json",
			"pluto.json": "testdata/folder/pluto.json",
		}
		require.Nil(t, err)
		require.Equal(t, expected, filesPath)
	})
}

func TestReadFile(t *testing.T) {
	t.Run("read a file ", func(t *testing.T) {
		file, _ := fs.TempFile(".", "tempfile")
		file.WriteString("prova")
		t.Cleanup(func() {
			fs.Remove(file.Name())
		})

		data, err := ReadFile(file.Name())
		require.Equal(t, "prova", string(data))
		require.Nil(t, err, "if paths are returned error must be nil")
	})

	t.Run("read a file that cannot be read", func(t *testing.T) {
		data, err := ReadFile("file-doesnt-exists")
		require.Nil(t, data, "cannot read file in write only")
		require.Error(t, err, "with no file that can be read return the underling error")
	})
}
