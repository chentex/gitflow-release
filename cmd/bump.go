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

	"github.com/chentex/gitflow-release/version"
	"github.com/pkg/errors"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// bumpCmd represents the bump command
var bumpCmd = &cobra.Command{
	Use:   "bump",
	Short: "Bumps the version of the project.",
	Long: `Bumps the version of the project. Depends on the flags sent
Examples:
File contains: 0.1.0

sending type flag 'major': bump will result in file containing: 1.0.0
sending type flag 'minor': bump will result in file containing: 0.2.0
sending type flag 'patch': bump will result in file containing: 0.1.1

sending alpha flag: bump will result in file containing: 0.1.0-alpha
sending beta flag:  bump will result in file containing: 0.1.0-beta

It's not possible to send alpha and beta together.
It's also not posible to combine major, minor and patch flags.

You can combine any (major, minor or patch) with (alpha or beta).`,
	RunE: func(cmd *cobra.Command, args []string) error {
		vFile := viper.GetString("versionfile")
		fmt.Println(vFile)
		versionManager := version.NewVersioner()
		alpha, err := cmd.Flags().GetBool("alpha")
		if err != nil {
			return errors.Wrap(err, "while getting alpha flag")
		}
		beta, err := cmd.Flags().GetBool("beta")
		if err != nil {
			return errors.Wrap(err, "while getting beta flag")
		}
		bump, err := cmd.Flags().GetString("type")
		if err != nil {
			return errors.Wrap(err, "while getting bump type flag")
		}
		err = versionManager.BumpVersion(vFile, bump, alpha, beta)
		if err != nil {
			return errors.Wrap(err, "while bumping version")
		}
		//TODO git actions
		return nil
	},
}

func init() {
	rootCmd.AddCommand(bumpCmd)

	bumpCmd.Flags().Bool("alpha", false, "Makes version alpha.")
	bumpCmd.Flags().Bool("beta", false, "Makes version beta.")
	bumpCmd.PersistentFlags().StringP("type", "t", "", "Send type of bump, can be major, minor, patch. Required.")
}
