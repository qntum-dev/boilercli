/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// cwdCmd represents the cwd command
var cwdCmd = &cobra.Command{
	Use:   "cwd",
	Short: "current working directory",
	Long:  `print current working directory`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cwd called")

		exePath, err := os.Executable()
		if err != nil {
			fmt.Println(err)
		}
		exeDir := filepath.Dir(exePath)

		fmt.Println(exeDir)
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(cwd)
	},
}

func init() {
	rootCmd.AddCommand(cwdCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cwdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cwdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
