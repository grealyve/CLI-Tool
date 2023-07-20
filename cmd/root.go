/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/grealyve/cli-tool/cmd/info"
	"github.com/grealyve/cli-tool/cmd/net"
	"github.com/grealyve/cli-tool/cmd/todo"
	"github.com/grealyve/cli-tool/cmd/weather"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cli-tool",
	Short: "This is a Command Line Interface (CLI) application that provides weather forecast information for a given city. Additionally, it offers functionalities for pinging domains, checking disk usage, and creating a to-do list.",
	Long: ``,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func addSubcommands() {
	rootCmd.AddCommand(info.InfoCmd)
	rootCmd.AddCommand(net.NetCmd)
	rootCmd.AddCommand(todo.TodoCmd)
	rootCmd.AddCommand(weather.WeatherCmd)
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cli-tool.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	addSubcommands()

}
