package scanner

import (
	"bufio"
	"os"
	"strings"
	"todoapp/todo"
)

type Scanner struct {
	todo *todo.List
}

func (s *Scanner) Start() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		NewPromt()

		ok := scanner.Scan()
		if !ok {
			return
		}

		inputString := scanner.Text()

		result := s.process(inputString)
		if result != "" {
			if result == NeedExit {
				Exit()
				return
			}
		}
	}
}

func (s *Scanner) process(inputString string) string {
	fields := strings.Fields(inputString)

	if len(fields) == 0 {
		return EmptyInput
	}

	cmd := fields[0]

	if cmd == "exit" {
		return NeedExit
	}
	if cmd == "add" {
		return s.cmdAdd(fields)
	}
	if cmd == "list" {
		return s.cmdList(fields)
	}
	if cmd == "done" {
		return s.cmdDone(fields)
	}
	if cmd == "del" {
		return s.cmdDel(fields)
	}
}

func (s *Scanner) cmdAdd(fields []string) string {
	if len(fields) < 3 {
		return WrongArgs
	}

	title := fields[1]

	taskText := ""

	for i := 2; i < len(fields); i++ {
		taskText += fields[i]

		if i != len(fields)-1 {
			taskText += " "
		}
	}
	task := todo.NewTask(title, taskText)

	s.todo.AddTask(task)

	PrintAdd(title)

}

func (s *Scanner) cmdList(fields []string) string {
	if len(fields) != 1 {
		return WrongArgs
	}

	tasks := s.todo.GetAllTasks()

	PrintTaks(tasks)

	return ""
}

func (s *Scanner) cmdDone(fields []string) string {
	if len(fields) != 2 {
		return WrongArgs
	}
	title := fields[1]

	doneTaskResult := s.todo.Done(title)
	if doneTaskResult != "" {
		return doneTaskResult
	}

	printDone(title)
	return ""
}

func (s *Scanner) cmdDel(fields []string) string {
	if len(fields) != 2 {
		return WrongArgs
	}

	title := fields[1]

	delTaskResult := s.todo.Delete(title)
	if delTaskResult != "" {
		return delTaskResult
	}

	printDel(title)
	return ""
}
