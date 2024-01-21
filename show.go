// Copyright (c) Gabriel de Quadros Ligneul
// SPDX-License-Identifier: Apache-2.0 (see LICENSE)

package main

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "show the clock-in and clock-out hours for the day",
	Run: func(cmd *cobra.Command, args []string) {
		model, err := NewModel()
		cobra.CheckErr(err)
		workHours, ok := model.Show(time.Now())
		if !ok {
			fmt.Println("No work hours for this date")
			return
		}
		fmt.Println("Clock In: ", workHours.ClockIn.Format(TIME_FMT))
		fmt.Println("Clock Out:", workHours.ClockOut.Format(TIME_FMT))
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}
