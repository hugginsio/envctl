// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package cmd

import (
	"context"
	"os"

	goversion "github.com/caarlos0/go-version"
	"github.com/charmbracelet/fang"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "envctl",
	Short: "Connect your environment files to your own secret managers",
}

func Execute() {
	if err := fang.Execute(
		context.Background(),
		rootCmd,
		fang.WithCommit(goversion.GetVersionInfo().GitCommit),
		fang.WithVersion(goversion.GetVersionInfo().GitVersion),
		fang.WithoutCompletions(),
		fang.WithoutManpage(),
	); err != nil {
		os.Exit(1)
	}
}
