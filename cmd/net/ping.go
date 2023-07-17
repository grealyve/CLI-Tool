/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package net

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var domainName string

func ping(domain string) {
	cmd := exec.Command("ping", domain)
	resp, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(resp))

}

// pingCmd represents the ping command
var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "This command pings a remote URL",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		ping(domainName)
	},
}

func init() {
	pingCmd.Flags().StringVarP(&domainName, "domain", "d", "", "Write down a domain")

	if err := pingCmd.MarkFlagRequired("domain"); err != nil {
		fmt.Println(err)
	}

	NetCmd.AddCommand(pingCmd)
}
