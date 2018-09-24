package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// rollbackCmd represents the rollback command
var rollbackCmd = &cobra.Command{
	Use:   "rollback",
	Short: "Rolls back specified function to previous version",
	Long:  `Rolls back function to specified alias or redeploys to previous version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("this command is not implemented yet!")
	},
}

func init() {
	rootCmd.AddCommand(rollbackCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rollbackCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rollbackCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
