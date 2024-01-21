// Copyright (c) Gabriel de Quadros Ligneul
// SPDX-License-Identifier: Apache-2.0 (see LICENSE)

package main

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "clock-in for the day now",
	Run: func(cmd *cobra.Command, args []string) {
		model, err := NewModel()
		cobra.CheckErr(err)
		workHours, err := model.ClockIn(time.Now())
		cobra.CheckErr(err)
		fmt.Println("Clock In: ", workHours.ClockIn.Format(TIME_FMT))
		fmt.Println("Clock Out:", workHours.ClockOut.Format(TIME_FMT))
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
