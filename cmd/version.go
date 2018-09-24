package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

const currentVersion = "0.1"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the current version and exits",
	Long:  `Prints the current version and exits`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("compose v" + currentVersion)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
