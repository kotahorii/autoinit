package cmd

import (
	"fmt"
	"log"
	"os/exec"
	"unsafe"

	"github.com/spf13/cobra"
)

var goCmd = &cobra.Command{
	Use:   "go",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if b, err := exec.Command("python3", "--version").Output(); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(*(*string)(unsafe.Pointer(&b)))
		}
	},
}

func init() {
	rootCmd.AddCommand(goCmd)
}
