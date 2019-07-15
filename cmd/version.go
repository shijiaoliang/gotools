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

	"github.com/spf13/cobra"
)

const VERSION = "v0.1"

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "version of gotools",
	Long:  `All software has versions. This is gotools's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(VERSION)
	},
}
