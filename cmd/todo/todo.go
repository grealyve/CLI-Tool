/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	"github.com/alexeyco/simpletable"

	"github.com/spf13/cobra"
)

const todoFile = "E:\\VisualStudio\\Go_Practice\\CLI\\todo.json"

var (
	add      string
	del      string
	complete string
	lst      string
)

type item struct {
	Task          string
	Completed     bool
	CreatedTime   time.Time
	CompletedTime time.Time
}

type Todos []item

func (t *Todos) complete(index string) error {
	list := *t
	x, _ := strconv.Atoi(index)
	if x > len(list) || x <= 0 {
		return errors.New("invalid index")
	}
	list[x-1].CompletedTime = time.Now()
	list[x-1].Completed = true

	return nil
}

func (t *Todos) add(task string) error {
	todo := item{
		Task:          task,
		Completed:     false,
		CreatedTime:   time.Now(),
		CompletedTime: time.Time{},
	}

	*t = append(*t, todo)

	data, err := json.Marshal(t)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(todoFile, data, 0644)
}

func (t *Todos) delete(index string) error {
	list := *t
	x, _ := strconv.Atoi(index)
	if x > len(list) || x <= 0 {
		return errors.New("invalid index")
	}
	*t = append(list[:x-1], list[x:]...)
	return nil
}

func (t *Todos) load() error {
	file, err := ioutil.ReadFile(todoFile)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return err
	}
	err = json.Unmarshal(file, t)
	if err != nil {
		return err
	}

	return nil
}

func (t *Todos) print() {

	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Task"},
			{Align: simpletable.AlignCenter, Text: "Done?"},
			{Align: simpletable.AlignRight, Text: "CreatedAt"},
			{Align: simpletable.AlignRight, Text: "CompletedAt"},
		},
	}

	var cells [][]*simpletable.Cell

	for idx, item := range *t {
		idx++
		task := blue(item.Task)
		done := blue("no")
		if item.Completed {
			task = green(fmt.Sprintf("\u2705 %s", item.Task))
			done = green("yes")
		}
		cells = append(cells, []*simpletable.Cell{
			{Text: fmt.Sprintf("%d", idx)},
			{Text: task},
			{Text: done},
			{Text: item.CreatedTime.Format(time.RFC822)},
			{Text: item.CompletedTime.Format(time.RFC822)},
		})
	}

	table.Body = &simpletable.Body{Cells: cells}

	table.Footer = &simpletable.Footer{Cells: []*simpletable.Cell{
		{Align: simpletable.AlignCenter, Span: 5, Text: red(fmt.Sprintf("You have %d pending todos", t.countPending()))},
	}}

	table.SetStyle(simpletable.StyleUnicode)

	table.Println()
}

func (t *Todos) countPending() int {
	total := 0
	for _, item := range *t {
		if !item.Completed {
			total++
		}
	}

	return total
}

// todoCmd represents the todo command
var TodoCmd = &cobra.Command{
	Use:   "todo",
	Short: "Special chart for your todo list",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args[0])
		todos := Todos{}
		if err := todos.load(); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		todos.print()
		// switch args[0] {
		// case "-d":
		// 	todos.delete(args[0])
		// case "-a":
		// 	todos.add(args[0])
		// case "-l":
		// 	todos.print()
		// case "-c":
		// 	todos.complete(args[0])

		// default:
		// 	fmt.Println("Invalid command")
		// }
	},
}

func init() {
	// TodoCmd.Flags().StringVarP(&del, "delete", "d", "", "del a todo")
	// TodoCmd.Flags().StringVarP(&add, "add", "a", "", "add a todo")
	// TodoCmd.Flags().StringVarP(&complete, "complete", "c", "", "complete a todo")
	TodoCmd.Flags().StringVarP(&lst, "list", "l", "", "list all todos")

}
