/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package todo

import (
	"log"

	"github.com/spf13/cobra"
)

var (
	lst bool
)

// todoCmd represents the todo command
var TodoCmd = &cobra.Command{
	Use:   "todo",
	Short: "Special chart for your todo list",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		todos := Todos{}

		loadedTodos, err := todos.Load()
		if err != nil {
			log.Fatalf("Error loading todos: %s", err)
		}
		todos = loadedTodos
		todos.Print()
	},
}

func init() {
	TodoCmd.Flags().BoolVarP(&lst, "list", "l", false, "list all todos")

}
