package todo

import (
	"fmt"
	"os"
	"strconv"
)

var tasks []string
var completed = make(map[int]bool)

func Add(task string) {
	tasks = append(tasks, task)
	fmt.Println("Добавлено:", task)
	saveToFile()
}

func Clear() {
	os.Remove("tasks.txt")
	tasks = nil
	completed = make(map[int]bool)
	fmt.Println("Все задачи удалены.")
}

func List() {
	loadFromFile()
	if len(tasks) == 0 {
		fmt.Println("Список пуст 😊")
		return
	}
	for i, task := range tasks {
		status := " "
		if completed[i] {
			status = "✔"
		}
		fmt.Printf("%d. [%s] %s\n", i+1, status, task)
	}
}

func Done(numStr string) {
	loadFromFile()
	num, err := strconv.Atoi(numStr)
	if err != nil || num < 1 || num > len(tasks) {
		fmt.Println("Неверный номер задачи")
		return
	}
	completed[num-1] = true
	fmt.Println("Отмечено как выполненное:", tasks[num-1])
	saveToFile()
}

func saveToFile() {
	file, err := os.Create("tasks.txt")
	if err != nil {
		fmt.Println("Ошибка записи:", err)
		return
	}
	defer file.Close()

	for i, task := range tasks {
		status := "0"
		if completed[i] {
			status = "1"
		}
		fmt.Fprintf(file, "%s|%s\n", status, task)
	}
}

func loadFromFile() {
	tasks = nil
	completed = make(map[int]bool)

	data, err := os.ReadFile("tasks.txt")
	if err != nil {
		return
	}

	lines := string(data)
	for i, line := range splitLines(lines) {
		if len(line) < 3 {
			continue
		}
		status := line[0]
		task := line[2:]
		tasks = append(tasks, task)
		if status == '1' {
			completed[i] = true
		}
	}
}

func splitLines(s string) []string {
	var lines []string
	curr := ""
	for _, r := range s {
		if r == '\n' {
			lines = append(lines, curr)
			curr = ""
		} else {
			curr += string(r)
		}
	}
	if curr != "" {
		lines = append(lines, curr)
	}
	return lines
}