package main

import (
	"fmt"
	"os"
	"strings"
)

func parseInput(todolist TodoList) {

	// if no arguments supplied, just return and execute format in main
	if len(os.Args) == 1 {
		return
	}

	command := os.Args[1]
	input := strings.Join(os.Args[1:], " ")

	switch command {
	case "add":
		todolist.addTodo(input)
	case "addC":
		todolist.createCategory(input)
	case "deleteC":
		todolist.deleteCategory(input)
	case "delete":
		todolist.deleteTodo(input)
	case "done":
		todolist.markDone(input)
	case "undo":
		todolist.unmarkDone(input)
	case "purge":
		todolist.purge()
	default:
		fmt.Println("Something went wrong noooooo.")
	}

}

// chain of calls: parseInput -> createCategory -> parseAddCategoryInput -> writeToFile -> handleClear -> format
//
