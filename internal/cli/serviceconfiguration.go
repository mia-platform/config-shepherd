/*
 * Copyright © 2021-present Mia s.r.l.
 * All rights reserved
 */

package cli

// configSchemaPath is the local path to the static JSON schema file.
const configSchemaPath = "./config.schema.json"

type SplittedMap struct {
	OutputMountPath string   `json:"outputMountPath" koanf:"outputMountPath"`
	InputMountPaths []string `json:"inputMountPaths" koanf:"inputMountPaths"`
}

// Configuration is read from provided config files and holds required configuration for the script.
type Configuration struct {
	// List of splitted configmaps to be rejoined
	SplittedMaps []SplittedMap `json:"splittedMaps" koanf:"splittedMaps"`
}
