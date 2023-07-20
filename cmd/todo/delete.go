/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package todo

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a todo at given index",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
		todos := Todos{}
		list, err := todos.Load()
		if err != nil {
			log.Printf("%v\n", err)
		}
		lst, err := todos.Delete(args[0], list)
		if err != nil {
			log.Fatalf("Delete items: %v\ns", err)
		}
		todos.SaveTasks(lst)

	},
}

func init() {
	TodoCmd.AddCommand(deleteCmd)
}
