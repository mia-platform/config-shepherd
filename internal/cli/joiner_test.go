package cli

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnmarshalSplittedMap(t *testing.T) {

	t.Run("Unmarshal Correctly", func(t *testing.T) {
		inputString := "{\"test\": {\"directories\": [\"test-directory\"]}}"
		expectedMap := map[string]splittedMapValue{
			"test": {
				Directories: []string{"test-directory"},
			},
		}
		outputMap, err := unmarshalSplittedMap([]byte(inputString))

		require.Nil(t, err)
		require.Equal(t, expectedMap, outputMap)
	})

	t.Run("Removing unnecessary info", func(t *testing.T) {
		inputString := "{\"test\": {\"directories\": [\"test-directory\"], \"pluto\": 1, \"pippo\": \"test\"}}"
		expectedMap := map[string]splittedMapValue{
			"test": {
				Directories: []string{"test-directory"},
			},
		}
		outputMap, err := unmarshalSplittedMap([]byte(inputString))

		require.Nil(t, err)
		require.Equal(t, expectedMap, outputMap)
	})

	t.Run("Unmarshal Invalid JSON Input", func(t *testing.T) {
		_, err := unmarshalSplittedMap([]byte(""))
		require.NotNil(t, err)
	})
}
