package cli

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadConfiguration(t *testing.T) {
	t.Run("Valid Configuration", func(t *testing.T) {
		configPath := "./testdata"
		configName := "valid.test"
		jsonSchemaPath := "../../config.schema.json"
		config, err := loadConfiguration(configPath, configName, jsonSchemaPath)

		require.Nil(t, err)

		splittedMaps := config.SplittedMaps
		require.Equal(t, 1, len(splittedMaps))

		splittedMap := splittedMaps[0]
		require.Equal(t, "./", splittedMap.OutputMountPath)
		require.Equal(t, []string{
			"./pkg/joiner/testdata/configmap-1-split-0",
			"./pkg/joiner/testdata/configmap-1-split-1",
		}, splittedMap.InputMountPaths)
	})

	t.Run("Invalid Properties Configuration", func(t *testing.T) {
		configPath := "./testdata"
		configName := "invalid.properties.test"
		jsonSchemaPath := "../../config.schema.json"
		_, err := loadConfiguration(configPath, configName, jsonSchemaPath)

		require.NotNil(t, err)
		require.Equal(
			t,
			"configuration not valid: json schema validation errors: [splittedMaps.0: Additional property output is not allowed]",
			err.Error(),
		)
	})

	t.Run("Missing Required Properties Configuration", func(t *testing.T) {
		configPath := "./testdata"
		configName := "missing.required.test"
		jsonSchemaPath := "../../config.schema.json"
		_, err := loadConfiguration(configPath, configName, jsonSchemaPath)

		require.NotNil(t, err)
		require.Equal(
			t,
			"configuration not valid: json schema validation errors: [splittedMaps.0: outputMountPath is required]",
			err.Error(),
		)
	})

	t.Run("Empty Configuration", func(t *testing.T) {
		configPath := "./testdata"
		configName := "empty.test"
		jsonSchemaPath := "../../config.schema.json"
		_, err := loadConfiguration(configPath, configName, jsonSchemaPath)

		require.NotNil(t, err)
		require.Equal(
			t,
			"configuration not valid: json schema validation errors: [(root): splittedMaps is required]",
			err.Error(),
		)
	})

	t.Run("Not a Json Configuration", func(t *testing.T) {
		configPath := "./testdata"
		configName := "notJson.test"
		jsonSchemaPath := "../../config.schema.json"
		_, err := loadConfiguration(configPath, configName, jsonSchemaPath)

		require.NotNil(t, err)
		require.Equal(
			t,
			"error loading config file: invalid character 'T' looking for beginning of value",
			err.Error(),
		)
	})
}
