package main

import "fmt"

type Todo struct {
	task      string
	completed bool
}

type TodoList struct {
	list []*Todo
}

// breaks DIP
// func FindPendingTodo(t *TodoList) []*Todo {
// 	result := make([]*Todo, 0)

// 	for i, v := range t.list {
// 		if !v.completed {
// 			result = append(result, &t.list[i])
// 		}
// 	}

// 	return result
// }

// fix
type FilterableTodos interface {
	FilterItems(f func(t *Todo) bool) []*Todo
}

func (t *TodoList) FilterItems(f func(t *Todo) bool) []*Todo {
	result := make([]*Todo, 0)

	for i, v := range t.list {
		if f(v) {
			result = append(result, t.list[i])
		}
	}

	return result
}

func IsNotComplete(todo *Todo) bool {
	return todo.completed == false
}

func FindPendingTodo(t FilterableTodos) []*Todo {
	filtered := t.FilterItems(IsNotComplete)

	return filtered
}

func main() {
	todo1 := &Todo{
		"task1",
		false,
	}
	todo2 := &Todo{
		"task2",
		true,
	}
	todo3 := &Todo{
		"task3",
		false,
	}

	todoList := &TodoList{
		list: []*Todo{todo1, todo2, todo3},
	}

	pendingTodo := FindPendingTodo(todoList)

	for i, v := range pendingTodo {
		fmt.Printf("%d) %s\n", i+1, v.task)
	}
}
