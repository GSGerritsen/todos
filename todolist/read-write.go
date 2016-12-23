package main

import (
	"encoding/json"
	//"fmt"
	"io/ioutil"
)

// need to add an init

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func writeToFile(todolist *TodoList) {
	j, err := json.Marshal(todolist)
	check(err)
	error := ioutil.WriteFile("./todos.json", j, 0644)
	check(error)
}
