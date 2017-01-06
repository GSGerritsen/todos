package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func handleBadInput(e error) {
	fmt.Printf("%s\n", e)
	return
}

func (todolist *TodoList) createCategory(todo string) {
	category, err := parseAddCategoryInput(todo)

	checkError(err)

	var emptySlice []*Entries
	newCategory := Todos{Realm: category, Entries: emptySlice}
	todolist.Todos = append(todolist.Todos, &newCategory)

	writeToFile(todolist)
}

func (todolist *TodoList) deleteCategory(todo string) {
	category, err := parseDeleteCategoryInput(todo)

	checkError(err)

	for index, key := range todolist.Todos {
		if key.Realm == category {
			todolist.Todos = append(todolist.Todos[:index], todolist.Todos[index+1:]...)
			break
		}
	}
	writeToFile(todolist)
}

func (todolist *TodoList) addTodo(todo string) {
	newEntry, realm, err := parseAddInput(todo)

	checkError(err)
	for _, key := range todolist.Todos {
		if key.Realm == realm {
			key.Entries = append(key.Entries, newEntry)
		}
	}
	todolist.applyIdOrdering()
	writeToFile(todolist)
}

func (todolist *TodoList) deleteTodo(todo string) {

	id, category, err := parseDeleteInput(todo)
	checkError(err)
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
	writeToFile(todolist)
}

func (todolist *TodoList) markDone(todo string) {
	id, category, err := parseMarkDoneInput(todo)
	checkError(err)
	i, _ := strconv.Atoi(id)

	for _, key := range todolist.Todos {
		if key.Realm == category {

			for _, k := range key.Entries {
				if k.ID == i {
					k.markDone()
				}
			}
		}
	}
	writeToFile(todolist)
}

func (todolist *TodoList) unmarkDone(todo string) {
	id, category, err := parseUnmarkDoneInput(todo)
	checkError(err)
	i, _ := strconv.Atoi(id)

	for _, key := range todolist.Todos {
		if key.Realm == category {

			for _, k := range key.Entries {
				if k.ID == i {
					k.unmarkDone()
				}
			}
		}
	}
	writeToFile(todolist)
}

func (todolist *TodoList) purge() {
	for _, key := range todolist.Todos {
		for i := len(key.Entries) - 1; i >= 0; i-- {
			if key.Entries[i].Done == true {
				key.Entries = append(key.Entries[:i], key.Entries[i+1:]...)
			}
		}
	}
	writeToFile(todolist)
}

func (entry *Entries) markDone() {
	entry.Done = true
}

func (entry *Entries) unmarkDone() {
	entry.Done = false
}

// example input: todos addC CPSC_322
func parseAddCategoryInput(todo string) (string, error) {
	var categoryRegex = regexp.MustCompile(`addC\s(.*)$`)

	var category string

	match := categoryRegex.FindStringSubmatch(todo)
	if len(match) > 1 {
		category = match[1]
		return category, nil
	} else {
		return "", errors.New("No category provided! Try addC <category>")
	}

}

func parseDeleteCategoryInput(todo string) (string, error) {
	var categoryRegex = regexp.MustCompile(`deleteC\s(.*)$`)

	var category string

	match := categoryRegex.FindStringSubmatch(todo)
	if len(match) > 1 {
		category = match[1]
		return category, nil
	} else {
		return "", errors.New("No category provided! Try deleteC <category>")
	}

}

// example input: todos add CPSC_304 => configure sql monkey due Thursday 2pm (use os.Args[1] to get the action word {add, delete, update, done, purge}
// add\s(.+)\:   		matches what is between 'add' and '=>'  = REALM
// :\s(.+)\ due   		matches what is between '=>' and 'due'  = DESCRIPTION
// due\s(.+)            matches everything after 'due'          = DUEDATE

// returning an entry struct, created from parsing the input, and the realm, to be used appropriately in addTodo.
func parseAddInput(todo string) (*Entries, string, error) {
	var realmRegex = regexp.MustCompile(`(add|a)\s(.+)\:`)
	var descriptionRegex = regexp.MustCompile(`:\s(.+)\ due`)
	var dueDateRegex = regexp.MustCompile(`due\s(.+)`)

	// FindStringSubmatch returns the leftmost match of the expression as the first element of the return slice of strings, and any matched subexpressions (capture groups)
	// as elements 1 and up

	var realm, description, dueDate string

	realmMatch := realmRegex.FindStringSubmatch(todo)
	descriptionMatch := descriptionRegex.FindStringSubmatch(todo)
	dueDateMatch := dueDateRegex.FindStringSubmatch(todo)
	if len(realmMatch) > 1 && len(descriptionMatch) > 1 && len(dueDateMatch) > 1 {
		realm = realmMatch[2]
		description = descriptionMatch[1]
		dueDate = dueDateMatch[1]
		newEntry := Entries{ID: 0, Description: strings.TrimSpace(description), Duedate: strings.TrimSpace(dueDate), Done: false}
		return &newEntry, realm, nil
	}
	return nil, "", errors.New("Add todo input error")

}

// example input: todos delete CPSC_304 1
func parseDeleteInput(todo string) (string, string, error) {
	var categoryRegex = regexp.MustCompile(`delete\s(.+)\ \d`)
	var IdRegex = regexp.MustCompile(`.*?([\d]+)$`)

	var category, id string

	categoryMatch := categoryRegex.FindStringSubmatch(todo)
	idMatch := IdRegex.FindStringSubmatch(todo)

	if len(categoryMatch) > 1 && len(idMatch) > 1 {
		category = categoryMatch[1]
		id = idMatch[1]
		return id, category, nil
	}

	return "", "", errors.New("Delete todo input error. Missing category or ID")

}

func parseMarkDoneInput(todo string) (string, string, error) {
	var categoryRegex = regexp.MustCompile(`(done|d)\s(.+)\ \d`)
	var IdRegex = regexp.MustCompile(`.*?([\d]+)$`)

	var id, category string

	categoryMatch := categoryRegex.FindStringSubmatch(todo)
	idMatch := IdRegex.FindStringSubmatch(todo)

	if len(categoryMatch) > 1 && len(idMatch) > 1 {
		id = idMatch[1]
		category = categoryMatch[2]
		return id, category, nil
	}

	return "", "", errors.New("Mark done input error. Try done <category> <ID>")
}

func parseUnmarkDoneInput(todo string) (string, string, error) {
	var categoryRegex = regexp.MustCompile(`undo\s(.+)\ \d`)
	var IdRegex = regexp.MustCompile(`.*?([\d]+)$`)

	var id, category string

	categoryMatch := categoryRegex.FindStringSubmatch(todo)
	idMatch := IdRegex.FindStringSubmatch(todo)

	if len(categoryMatch) > 1 && len(idMatch) > 1 {
		id = idMatch[1]
		category = categoryMatch[1]
		return id, category, nil
	}

	return "", "", errors.New("Mark undone input error. Try undo <category> <ID>")
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
