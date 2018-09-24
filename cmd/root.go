package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var stdout = log.New(os.Stdout, "compose: ", 0)
var stderr = log.New(os.Stderr, "compose: ", 0)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "compose",
	Short: "CLI tool for managing AWS Lambda functions",
	Long:  `CLI tool for managing AWS Lambda functions using the AWS API.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "compose-file", "", "compose file (default is ./compose.yaml)")
	rootCmd.PersistentFlags().String("profile", "default", "AWS profile in your config file")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName("./compose")
	}

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil && os.Args[1] != "version" {
		stderr.Fatalln("no compose.yaml file found")
	}
}
