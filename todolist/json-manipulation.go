package main

import (
	//"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Think about edge cases? Regex sanitizing?
func (todolist *TodoList) createCategory(category string) {

	var emptySlice []*Entries
	newCategory := &Todos{Realm: category, Entries: emptySlice}
	todolist.Todos = append(todolist.Todos, newCategory)

}

// Think about edge cases? Deleting non-existent category?
func (todolist *TodoList) deleteCategory(category string) {
	for index, key := range todolist.Todos {
		if key.Realm == category {
			todolist.Todos = append(todolist.Todos[:index], todolist.Todos[index+1:]...)
		}
	}
}

func (todolist *TodoList) addTodo(todo string) {
	newEntry, realm := parseAddInput(todo)
	for _, key := range todolist.Todos {
		if key.Realm == realm {
			key.Entries = append(key.Entries, newEntry)
		}
	}
	todolist.applyIdOrdering()
}

func (todolist *TodoList) deleteTodo(todo string) {

	id, category := parseDeleteInput(todo)
	i, _ := strconv.Atoi(id)

	for _, key := range todolist.Todos {
		if key.Realm == category {

			for ind, k := range key.Entries {
				if k.ID == i {
					key.Entries = append(key.Entries[:ind], key.Entries[ind+1:]...)
				}
			}
		}
	}
	todolist.applyIdOrdering()
}

// example input: todos add CPSC_304 => configure sql monkey due Thursday 2pm (use os.Args[1] to get the action word {add, delete, update, done, purge}
// add\s(.+)\ =>   		matches what is between 'add' and '=>'  = REALM
// =>\s(.+)\ due   		matches what is between '=>' and 'due'  = DESCRIPTION
// due\s(.+)            matches everything after 'due'          = DUEDATE

// returning an entry struct, created from parsing the input, and the realm, to be used appropriately in addTodo.
func parseAddInput(todo string) (*Entries, string) {
	var realmRegex = regexp.MustCompile(`add\s(.+)\:`)
	var descriptionRegex = regexp.MustCompile(`:\s(.+)\ due`)
	var dueDateRegex = regexp.MustCompile(`due\s(.+)`)

	// FindStringSubmatch returns the leftmost match of the expression as the first element of the return slice of strings, and any matched subexpressions (capture groups)
	// as elements 1 and up

	realm := realmRegex.FindStringSubmatch(todo)[1]

	description := descriptionRegex.FindStringSubmatch(todo)[1]

	dueDate := dueDateRegex.FindStringSubmatch(todo)[1]

	newEntry := Entries{ID: 0, Description: strings.TrimSpace(description), Duedate: strings.TrimSpace(dueDate), Done: false}
	return &newEntry, realm
}

// example input: todos delete CPSC_304 1
func parseDeleteInput(todo string) (string, string) {
	var categoryRegex = regexp.MustCompile(`delete\s(.+)\ \d`)
	var IdRegex = regexp.MustCompile(`.*?([\d]+)$`)

	category := categoryRegex.FindStringSubmatch(todo)[1]
	id := IdRegex.FindStringSubmatch(todo)[1]

	return id, category

}

func (todos *Todos) putIdsInOrder() {
	for index, key := range todos.Entries {
		key.ID = index + 1
	}
}

func (todolist *TodoList) applyIdOrdering() {
	for _, key := range todolist.Todos {
		key.putIdsInOrder()
	}
}
