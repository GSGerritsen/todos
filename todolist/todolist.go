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

	var todolist TodoList
	data, err := ioutil.ReadFile("./test.json")
	checkError(err)

	unmarshalErr := json.Unmarshal(data, &todolist)
	checkError(unmarshalErr)

	w := tabwriter.NewWriter(os.Stdout, 20, 25, 15, ' ', tabwriter.TabIndent)
	todolist.addTodo("todos add CPSC_304 => do aaaaaaaaaaaaa other thing due Thursday 9am")

	format(todolist, w)
	w.Flush()

}
