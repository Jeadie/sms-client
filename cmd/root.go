package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var (
	reconstruct_packets bool
	pollEndpoints       []string
	rootCmd             = &cobra.Command{
		Use:   "SmsClient",
		Short: "A brief description of your application",
		Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		Run: func(cmd *cobra.Command, args []string) { pollerCmd.Run(cmd, args) },
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringArrayVar(&pollEndpoints, POLL_ENDPOINTS_FLAG, GetDefaultPollEndpoints(), "Endpoint to expect a Hilink device server to be running")
	rootCmd.Flags().BoolVar(&reconstruct_packets, RECONSTRUCT_PACKETS_FLAG, GetDefaultReconstructPackets(), "Attempt to reconstruct Smss from Sms fragments if found.")
}
