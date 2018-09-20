package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// aliasCmd represents the alias command
var aliasCmd = &cobra.Command{
	Use:   "alias",
	Short: "Adds alias to function, or updates alias for function",
	Long: `Adds or updates alias for function`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("alias called")
	},
}

func init() {
	rootCmd.AddCommand(aliasCmd)
}
