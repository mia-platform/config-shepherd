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
	// "os"

	"join-config-map/internal/utils"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

// New create a new options struct
func New() *utils.Options {
	options := &utils.Options{
		// Kubeconfig: os.Getenv("KUBECONFIG"),
	}

	// bind to kubernetes config flags
	options.Config = &genericclioptions.ConfigFlags{
		// CAFile:       &options.CertificateAuthority,
	}

	return options
}

// AddGlobalFlags add to the cobra command all the global flags
func AddGlobalFlags(cmd *cobra.Command, options *utils.Options) {

	// flags := cmd.PersistentFlags()
	// flags.StringVarP(&options.CertificateAuthority, "certificate-authority", "", "", "Path to a cert file for the certificate authority")
}
