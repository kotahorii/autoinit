package cmd

import (
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

func execCommand(command string) error {
	chars := strings.Split(command, " ")
	if err := exec.Command(chars[0], chars[1:]...).Run(); err != nil {
		return err
	}
	return nil
}

var pythonCmd = &cobra.Command{
	Use:   "python",
	Short: "create python environment",
	Long:  `create venv and pip install black flake8.`,
	Run: func(cmd *cobra.Command, args []string) {
		execCommand("python3 -m venv venv")
	},
}

func init() {
	rootCmd.AddCommand(pythonCmd)
}
