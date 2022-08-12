/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	hi "github.com/Jeadie/SmsClient/pkg/hilink"
	"github.com/spf13/cobra"
)

// receiveCmd represents the receive command
var (
	receiveCmd = &cobra.Command{
		Use:   "receive",
		Short: "Receive all SMS messages from Hilink modem device",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			hi.ReceiveSms(pollEndpoints)
		},
	}
)

func init() {
	rootCmd.AddCommand(receiveCmd)
}
