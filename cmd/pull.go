// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/hugginsio/modeline"
	"github.com/spf13/cobra"
)

var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pull secrets to the working directory",
	RunE: func(cmd *cobra.Command, args []string) error {
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
	pullCmd.Flags().StringP(FlagEnvironment, "e", "local", "The env file to reference")

	rootCmd.AddCommand(pullCmd)
}
