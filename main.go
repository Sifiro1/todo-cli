package main

import (
	"fmt"
	"os"

	"todo-cli/todo"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Использование: todo add <дело> | list | done <номер>")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Ошибка: укажи текст задачи")
			return
		}
		task := os.Args[2]
		todo.Add(task)
	case "list":
		todo.List()
	case "done":
		if len(os.Args) < 3 {
			fmt.Println("Ошибка: укажи номер задачи")
			return
		}
		todo.Done(os.Args[2])
case "help":
	fmt.Println(`
Доступные команды:
  add <текст>     - добавить задачу
  list            - показать список
  done <номер>    - отметить выполненное
  clear           - очистить все задачи
  help            - показать это сообщение
`)
	default:
		fmt.Println("Неизвестная команда:", command)
        case "clear":
	todo.Clear()
	}
}