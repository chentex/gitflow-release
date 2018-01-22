// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
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

package cmd

import (
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/pkg/errors"

	"github.com/spf13/cobra"
)

var (
	initialVersion string
	versionFile    string
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes the gitflow-release configuration.",
	Long: `Init will create the following files:

VERSION (default name, with the current version of the project.)
.gitflow-release (with the configuration for this project.)`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Validate if repo already initialized
		force, err := strconv.ParseBool(cmd.Flag("force").Value.String())
		if err != nil {
			return errors.Wrapf(err, "Error parsing flag.")
		}

		if _, err = ioutil.ReadFile(versionFile); err == nil && !force {
			return errors.Errorf("Repository already initialized with version file: %s", versionFile)
		}

		yaml := fmt.Sprintf(`versionfile: %s`, versionFile)
		err = ioutil.WriteFile(cfgFile, []byte(yaml), 0644)
		if err != nil {
			return errors.Wrapf(err, "Creating configuration file: %s", cfgFile)
		}

		err = ioutil.WriteFile(versionFile, []byte(initialVersion), 0644)
		if err != nil {
			return errors.Wrapf(err, "Creating version file: %s", versionFile)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.PersistentFlags().StringVarP(&initialVersion, "initialversion", "v", "0.1.0", "Current version of the project.")
	initCmd.PersistentFlags().StringVarP(&versionFile, "versionfile", "f", "VERSION", "Name of the file that will be used to track the version.")
	initCmd.Flags().Bool("force", false, "When sending this flag true it will override previous gitflow-release configuration.")
}
