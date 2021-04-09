package joiner

import (
	"path/filepath"
	"testing"

	"git.tools.mia-platform.eu/platform/devops/config-shepherd/internal/utils"

	"github.com/stretchr/testify/require"
)

const testdata = "testdata/"
const outdir = "outdir/"
const fileBar = "bar.json"
const fileFoo = "foo.json"

var split0 = filepath.Join(testdata, "configmap-1-split-0")
var split1 = filepath.Join(testdata, "configmap-1-split-1")

var split0AbsPath, _ = filepath.Abs(split0)
var split1AbsPath, _ = filepath.Abs(split1)

func TestExtractFilesParts(t *testing.T) {
	t.Run("extract files parts", func(t *testing.T) {
		filesDirPath := []string{
			split0,
			split1,
		}
		mapping, err := extractFilesParts(filesDirPath)

		expected := map[string][]string{
			fileFoo: {
				filepath.Join(split0AbsPath, fileFoo),
				filepath.Join(split1AbsPath, fileFoo),
			},
			fileBar: {
				filepath.Join(split1AbsPath, fileBar),
			},
		}

		require.Nil(t, err, "can not extract files parts")
		require.Equal(t, expected, mapping)
	})

	t.Run("with an invalid path", func(t *testing.T) {
		filesDirPath := []string{
			filepath.Join(split0, "invalidPath"),
			split1,
		}
		mapping, err := extractFilesParts(filesDirPath)

		expected := map[string][]string{
			fileFoo: {
				filepath.Join(split1AbsPath, fileFoo),
			},
			fileBar: {
				filepath.Join(split1AbsPath, fileBar),
			},
		}

		require.Nil(t, err, "can not extract files parts")
		require.Equal(t, expected, mapping)
	})
}

func TestJoinFileParts(t *testing.T) {
	filesPath := []string{
		filepath.Join(split0AbsPath, fileFoo),
		filepath.Join(split1AbsPath, fileFoo),
	}

	t.Run("join files parts", func(t *testing.T) {
		content, err := joinFileParts(filesPath)
		require.Nil(t, err, "can not join files parts")

		expected := "{\n    \"hello\": \"foo\"\n}"
		require.Equal(t, expected, string(content[:]))
	})
}

func TestJoinAllFiles(t *testing.T) {
	pippoFilesPath := []string{
		filepath.Join(split0AbsPath, fileFoo),
		filepath.Join(split1AbsPath, fileFoo),
	}

	plutoFilesPath := []string{
		filepath.Join(split1AbsPath, fileBar),
	}

	mapping := map[string][]string{
		fileBar: plutoFilesPath,
		fileFoo: pippoFilesPath,
	}

	t.Run("join all files parts and write", func(t *testing.T) {
		testOutdir := filepath.Join(testdata, outdir)
		err := utils.MkdirAll(testOutdir)
		require.Nil(t, err)

		err = joinAllFiles(mapping, testOutdir)
		require.Nil(t, err)

		pathMapping, err := utils.ExtractFilesFromDir(testOutdir)
		require.Nil(t, err)

		fileBarPath := filepath.Join(testOutdir, fileBar)
		fileFooPath := filepath.Join(testOutdir, fileFoo)

		expected := map[string]string{
			fileBar: fileBarPath,
			fileFoo: fileFooPath,
		}
		require.Equal(t, expected, pathMapping)

		content, err := utils.ReadFile(fileBarPath)
		require.Nil(t, err)
		require.Equal(t, "{\n    \"bar\": \"bar\"\n}", string(content[:]))
		content, err = utils.ReadFile(fileFooPath)
		require.Nil(t, err)
		require.Equal(t, "{\n    \"hello\": \"foo\"\n}", string(content[:]))

		err = utils.RemoveAll(testOutdir)
		require.Nil(t, err)
	})
}
