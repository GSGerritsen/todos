package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

// need to add an init: create empty TodoList struct, create a todos.json, marshal empty TodoList into it. Make it conditional
// on whether or not a todos.json already exists. Call in the beginning of main

func initialize() {
	if _, e := os.Stat("./todos.json"); os.IsNotExist(e) {
		var emptySlice []*Todos
		todolist := TodoList{Todos: emptySlice}
		j, err := json.Marshal(todolist)
		checkError(err)
		error := ioutil.WriteFile("./todos.json", j, 0644)
		checkError(error)
		fmt.Println("Directory intialized, todos.json file created.")
		return
	}
	return
}

func writeToFile(todolist *TodoList) {
	j, err := json.Marshal(todolist)
	checkError(err)
	error := ioutil.WriteFile("./todos.json", j, 0644)
	checkError(error)
}

func readUpdatedFile() TodoList {
	var todolist TodoList
	data, err := ioutil.ReadFile("./todos.json")
	checkError(err)

	unmarshalErr := json.Unmarshal(data, &todolist)
	checkError(unmarshalErr)
	return todolist
}
