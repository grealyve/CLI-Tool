package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/alexeyco/simpletable"
)

const todoFile = "todo.json"

type Item struct {
	Task          string
	Completed     bool
	CreatedTime   time.Time
	CompletedTime time.Time
}

type Todos []Item

func (t *Todos) SaveTasks(todoList Todos) error {
	bytes, err := json.Marshal(todoList)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(todoFile, bytes, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (t *Todos) Load() (Todos, error) {
	file, err := ioutil.ReadFile(todoFile)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, nil
		}
		return nil, err
	}
	var todo Todos

	if len(file) == 0 {
		return nil, err
	}

	err = json.Unmarshal(file, &todo)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (t *Todos) Add(task string) error {
	todo := Item{
		Task:          task,
		Completed:     false,
		CreatedTime:   time.Now(),
		CompletedTime: time.Time{},
	}

	todoList, err := t.Load()
	if err != nil {
		log.Printf("Error loading todo list: %s", err)
	}

	todoList = append(todoList, todo)

	data, err := json.Marshal(todoList)
	if err != nil {
		return err
	}

	err2 := ioutil.WriteFile(todoFile, data, 0644)
	if err2 != nil {
		return err2
	}

	return nil
}

func (t *Todos) Complete(index string, list Todos) (Todos, error) {
	x, _ := strconv.Atoi(index)
	if x > len(list) || x <= 0 {
		return nil, errors.New("invalid index")
	}
	list[x-1].CompletedTime = time.Now()
	list[x-1].Completed = true

	return list, nil
}

func (t *Todos) Delete(index string, list Todos) (Todos, error) {
	x, _ := strconv.Atoi(index)
	if x > len(list) || x <= 0 {
		return nil, errors.New("invalid index")
	}
	list = append(list[:x-1], list[x:]...)
	return list, nil
}

func (t *Todos) Print() {

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
		{Align: simpletable.AlignCenter, Span: 5, Text: red(fmt.Sprintf("You have %d pending todos", t.CountPending()))},
	}}

	table.SetStyle(simpletable.StyleUnicode)

	table.Println()
}

func (t *Todos) CountPending() int {
	total := 0
	for _, item := range *t {
		if !item.Completed {
			total++
		}
	}

	return total
}
