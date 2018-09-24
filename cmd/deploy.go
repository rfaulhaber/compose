package cmd

import (
	"github.com/pkg/errors"
	"github.com/rfaulhaber/compose/pkg/compose"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"os/exec"
	"strings"
)

// deployCmd represents the deploy command
var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Runs build commands and deploys function(s) to AWS",
	Long: `deploy runs all the specified build commands, compresses the build,
and uploads the new code to AWS.`,
	// TODO display errors usefully
	Run: func(cmd *cobra.Command, args []string) {
		if err := RunDeploy(cmd, args); err != nil {
			stderr.Fatalln(err)
		}
	},
	Args: cobra.MinimumNArgs(1),
}

func init() {
	rootCmd.AddCommand(deployCmd)
}

// 1. build function
// 2. package function
// 3. compress function
// 3. deploy function
// deploy to S3 if necessary
func RunDeploy(cmd *cobra.Command, args []string) error {
	var compFile compose.File

	err := viper.Unmarshal(&compFile)

	if err != nil {
		return err
	}

	for _, fun := range compFile.Functions {
		if err := buildFunc(fun); err != nil {
			return err
		}
	}

	return nil
}

// TODO refactor into "function builder"?
func buildFunc(function compose.Function) error {
	for n, step := range function.Build {
		if err := runCmd(step); err != nil {
			return errors.Wrap(err, "build step "+string(n+1)+" for function "+function.Name+"failed: ")
		}
	}

	return nil
}

func runCmd(cmdStr string) error {
	cmdLine := strings.Split(cmdStr, " ")

	command := cmdLine[0]
	args := cmdLine[1:]

	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
