package main

import (
	"encoding/json"
	//"fmt"
	"io/ioutil"
	"os"
	"text/tabwriter"
)

type TodoList struct {
	Todos []*Todos `json:"todos"`
}

type Todos struct {
	Realm   string     `json:"realm"`
	Entries []*Entries `json:"entries"`
}

type Entries struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Duedate     string `json:"duedate"`
	Done        bool   `json:"done"`
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	// Move this stuff into read-write.go
	var todolist TodoList
	data, err := ioutil.ReadFile("./todos.json")
	checkError(err)

	unmarshalErr := json.Unmarshal(data, &todolist)
	checkError(unmarshalErr)

	w := tabwriter.NewWriter(os.Stdout, 30, 35, 15, ' ', tabwriter.TabIndent)

	parseInput(todolist)
	handleClear()
	updatedTodolist := readUpdatedFile()
	format(updatedTodolist, w)
	w.Flush()

}
