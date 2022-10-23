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
			for msg := range hi.Poll(pollEndpoints, poll_period, false) {
				for _, e := range push_endpoints {
					hi.PushSms(e, msg)
				}
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(pollerCmd)

	pollerCmd.Flags().StringArrayVar(&push_endpoints, PUSH_ENDPOINTS_FLAG, GetDefaultPushEndpoints(), "Endpoint to send new, polled SMSs to.")
	pollerCmd.Flags().UintVar(&poll_period, POLL_PERIOD_FLAG, GetDefaultPollPeriod(), "Period to poll poll_endpoint, seconds.")
}
