/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"os/exec"
	"unsafe"

	"github.com/spf13/cobra"
)

var pythonCmd = &cobra.Command{
	Use:   "python",
	Short: "create python environment",
	Long:  `create venv and pip install black flake8.`,
	Run: func(cmd *cobra.Command, args []string) {
		if b, err := exec.Command("python3", "--version").Output(); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(*(*string)(unsafe.Pointer(&b)))
		}
	},
}

func init() {
	rootCmd.AddCommand(pythonCmd)
}
