// Copyright (c) Gabriel de Quadros Ligneul
// SPDX-License-Identifier: Apache-2.0 (see LICENSE)

package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const TIME_FMT = "15:04:05"

var rootCmd = &cobra.Command{
	Use:   "clockin",
	Short: "clockin is a command-line tool to track work hours",
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
