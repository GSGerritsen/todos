package main

import (
	"fmt"
	"github.com/fatih/color"
	"text/tabwriter"
)

// Takes a TodoList struct, establishes colors, and spins off one for loop to print the category a set of todos belongs to, and then an
// inner for loop to print the contents of each todo. Nested todos shouldn't be a problem with archiving + the fact that there won't ever be a huge number of todos
// in a given directory
func format(todos TodoList, w *tabwriter.Writer) {

	red := color.New(color.FgHiCyan).SprintFunc()
	yellow := color.New(color.FgHiGreen).SprintFunc()

	for _, todo := range todos.Todos {
		fmt.Fprintf(w, "%s\n", red(todo.Realm))

		for _, entry := range todo.Entries {
			fmt.Fprintf(w, "%s\t%s\t%s\t%s\t\n", yellow(entry.ID), formatIfDone(entry.Done), entry.Duedate, entry.Description)
		}
	}
}

func formatIfDone(todo bool) string {

	//green := color.New(color.FgHiGreen).SprintFunc()
	// Calling green("x") isn't working right now...
	if todo == true {
		doneString := "[ " + "x" + " ]"
		return doneString
	} else {
		return "[ ]"
	}
}
