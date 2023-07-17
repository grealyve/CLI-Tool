/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package info

import (
	"fmt"

	"github.com/ricochet2200/go-disk-usage/du"
	"github.com/spf13/cobra"
)

var (
	path string
	KB   = uint64(1024)
)

func showUsage(path string) {
	usage := du.NewDiskUsage(path)

	fmt.Println("Free:", usage.Free()/(KB*KB))
	fmt.Println("Available:", usage.Available()/(KB*KB))
	fmt.Println("Size:", usage.Size()/(KB*KB))
	fmt.Println("Used:", usage.Used()/(KB*KB))
	fmt.Println("Usage:", usage.Usage()*100, "%")
}

// diskUsageCmd represents the diskUsage command
var diskUsageCmd = &cobra.Command{
	Use:   "diskUsage",
	Short: "Disk usage information",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		showUsage(path)
	},
}

func init() {
	diskUsageCmd.Flags().StringVarP(&path, "path", "p", "", "Path to check disk usage")

	if err := diskUsageCmd.MarkFlagRequired("path"); err != nil {
		fmt.Println(err)
	}

	InfoCmd.AddCommand(diskUsageCmd)
}
