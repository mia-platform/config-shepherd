package joiner

import (
	"join-config-map/internal/utils"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

const testdata = "testdata/"
const outdir = "outdir/"
const filePluto = "pluto.json"
const filePippo = "pippo.json"

var split0 = filepath.Join(testdata, "configmap-1-split-0")
var split1 = filepath.Join(testdata, "configmap-1-split-1")

var split0AbsPath, _ = filepath.Abs(split0)
var split1AbsPath, _ = filepath.Abs(split1)

func TestExtractFilesParts(t *testing.T) {
	filesDirPath := []string{
		split0,
		split1,
	}

	t.Run("extract files parts", func(t *testing.T) {
		mapping, err := extractFilesParts(filesDirPath)

		expected := map[string][]string{
			filePippo: {
				filepath.Join(split0AbsPath, filePippo),
				filepath.Join(split1AbsPath, filePippo),
			},
			filePluto: {
				filepath.Join(split1AbsPath, filePluto),
			},
		}

		require.Nil(t, err, "can not extract files parts")
		require.Equal(t, expected, mapping)
	})
}

func TestJoinFileParts(t *testing.T) {
	filesPath := []string{
		filepath.Join(split0AbsPath, filePippo),
		filepath.Join(split1AbsPath, filePippo),
	}

	t.Run("join files parts", func(t *testing.T) {
		content, err := joinFileParts(filesPath)
		require.Nil(t, err, "can not join files parts")

		expected := []byte{
			0x7b, 0xa, 0x20, 0x20, 0x20, 0x20,
			0x22, 0x63, 0x69, 0x61, 0x6f, 0x22,
			0x3a, 0x20, 0x22, 0x70, 0x69, 0x70,
			0x70, 0x6f, 0x22, 0xa, 0x7d,
		}
		require.Equal(t, expected, content)
	})
}

func TestJoinAllFiles(t *testing.T) {
	pippoFilesPath := []string{
		filepath.Join(split0AbsPath, filePippo),
		filepath.Join(split1AbsPath, filePippo),
	}

	plutoFilesPath := []string{
		filepath.Join(split1AbsPath, filePluto),
	}

	mapping := map[string][]string{
		filePluto: plutoFilesPath,
		filePippo: pippoFilesPath,
	}

	t.Run("join all files parts and write", func(t *testing.T) {
		testOutdir := filepath.Join(testdata, outdir)
		err := utils.MkdirAll(testOutdir)
		require.Nil(t, err)

		err = joinAllFiles(mapping, testOutdir)
		require.Nil(t, err)

		pathMapping, err := utils.ExtractFilesFromDir(testOutdir)
		require.Nil(t, err)

		filePlutoPath := filepath.Join(testOutdir, filePluto)
		filePippoPath := filepath.Join(testOutdir, filePippo)

		expected := map[string]string{
			filePluto: filePlutoPath,
			filePippo: filePippoPath,
		}
		require.Equal(t, expected, pathMapping)

		content, err := utils.ReadFile(filePlutoPath)
		require.Nil(t, err)
		require.Equal(t, []byte{
			0x7b, 0xa, 0x20, 0x20, 0x20,
			0x20, 0x22, 0x70, 0x6c, 0x75,
			0x74, 0x6f, 0x22, 0x3a, 0x20,
			0x22, 0x70, 0x6c, 0x75, 0x74,
			0x6f, 0x22, 0xa, 0x7d,
		}, content)
		content, err = utils.ReadFile(filePippoPath)
		require.Nil(t, err)
		require.Equal(t, []byte{
			0x7b, 0xa, 0x20, 0x20, 0x20,
			0x20, 0x22, 0x63, 0x69, 0x61,
			0x6f, 0x22, 0x3a, 0x20, 0x22,
			0x70, 0x69, 0x70, 0x70, 0x6f,
			0x22, 0xa, 0x7d,
		}, content)

		err = utils.RemoveAll(testOutdir)
		require.Nil(t, err)
	})
}
