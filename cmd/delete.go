package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a Lambda function",
	Long: `Deletes a Lambda function `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("this command is not implemented yet!")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
