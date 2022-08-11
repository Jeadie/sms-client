/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// pollerCmd represents the poller command
var pollerCmd = &cobra.Command{
	Use:   "poller",
	Short: "Poll Huawei Hilink devices for received SMSs.",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("poller called")
	},
}

func init() {
	rootCmd.AddCommand(pollerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pollerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pollerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
