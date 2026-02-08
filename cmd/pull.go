// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/hugginsio/modeline"
	"github.com/spf13/cobra"
)

var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pull secrets to the working directory",
	RunE: func(cmd *cobra.Command, args []string) error {
		var envFilePath string

		flagEnv := cmd.Flag(FlagEnvironment)
		flagFile := cmd.Flag(FlagFile)

		if flagEnv.Changed {
			envFilePath = flagEnv.Value.String() + ".env"
		} else if flagFile.Changed {
			envFilePath = filepath.Clean(flagFile.Value.String())
			if !filepath.IsAbs(envFilePath) {
				wd, err := os.Getwd()
				if err != nil {
					return err
				}

				envFilePath = filepath.Join(wd, envFilePath)
			}
		}

		fmt.Println(envFilePath)

		modelines, err := modeline.ScanFile("local.env")
		if err != nil {
			return err
		}

		bytes, err := json.MarshalIndent(modelines, "", "  ")
		if err != nil {
			return err
		}

		fmt.Println(string(bytes))

		return nil
	},
}

func init() {
	pullCmd.Flags().StringP(FlagEnvironment, "e", "local", "The environment to reference")
	pullCmd.Flags().StringP(FlagFile, "f", filepath.Join(".", "local.env"), "Path to env file")

	pullCmd.MarkFlagsMutuallyExclusive(FlagEnvironment, FlagFile)

	rootCmd.AddCommand(pullCmd)
}
