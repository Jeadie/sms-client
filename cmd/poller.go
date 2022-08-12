/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	hi "github.com/Jeadie/SmsClient/pkg/hilink"
	"github.com/spf13/cobra"
)

var (
	push_endpoints []string
	poll_period    uint
	// pollerCmd represents the poller command
	pollerCmd = &cobra.Command{
		Use:   "poller",
		Short: "Poll Huawei Hilink devices for received SMSs.",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			hi.Poll(pollEndpoints, poll_period, func(endpoint string, sms []hi.Sms) {
				for _, e := range push_endpoints {
					hi.PushSms(e, sms)
				}
			})
		},
	}
)

func init() {
	rootCmd.AddCommand(pollerCmd)

	pollerCmd.Flags().StringArrayVar(&push_endpoints, PUSH_ENDPOINTS_FLAG, []string{"127.0.0.1"}, "Endpoint to send new, polled SMSs to.")
	pollerCmd.Flags().UintVar(&poll_period, POLL_PERIOD_FLAG, 10, "Period to poll poll_endpoint, seconds.")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pollerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pollerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
