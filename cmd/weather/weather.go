/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package weather

import (
	"fmt"

	"github.com/spf13/cobra"
)

// weatherCmd represents the weather command
var weatherCmd = &cobra.Command{
	Use:   "weather",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("weather called")
	},
}

func init() {
	// https://www.youtube.com/watch?v=zPYjfgxYO7k
}
