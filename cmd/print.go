/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/dworthen/gog4gclitest/config"
	"github.com/spf13/cobra"
)

// printCmd represents the print command
var printCmd = &cobra.Command{
	Use:   "print",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		userConfig := config.GetConfig2()
		jsonUserConfig, err := json.Marshal(&userConfig)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(string(jsonUserConfig))
	},
}

func init() {
	configCmd.AddCommand(printCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// printCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// printCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
