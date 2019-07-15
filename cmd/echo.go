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
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(echoCmd)
}

var echoCmd = &cobra.Command{
	Use:   "echo",
	Short: "Echo something",
	Long:  `Echo something you want`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(strings.Join(args, "-"))
	},
}
