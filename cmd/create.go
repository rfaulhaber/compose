package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a new Lambda function",
	Long:  `Creates a new Lambda function using the AWS CreateFunction API call.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("this command is not implemented yet!")
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
