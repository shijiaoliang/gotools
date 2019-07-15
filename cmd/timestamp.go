/**
 *******************************************slade********************************************
 * Copyright (c)  slade
 * Created by slade.
 * User: 605724193@qq.com
 * Date: 2019/07/15
 * Time: 11:18
 ********************************************************************************************
 */

package cmd

import (
	"fmt"
	"time"
	"strconv"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(timestampCmd)
}

var timestampCmd = &cobra.Command{
	Use:   "timestamp",
	Short: "Time string to timestamp",
	Long:  `Time string to timestamp`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, item := range args {
			timestamp, _ := time.Parse("2006-01-02 15:04:05", item)
			fmt.Println("[" + item + "]:" + strconv.Itoa(int(timestamp.Unix())))
		}
	},
}
