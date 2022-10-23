package cmd

import (
	"fmt"
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
			smsChan := hi.ReceiveSms(pollEndpoints, reconstruct_packets)
			for sms := range smsChan {
				fmt.Println(sms)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(receiveCmd)
}
