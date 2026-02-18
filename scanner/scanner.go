package scanner

import (
	"bufio"
	"os"
	"strings"
	"todoapp/todo"
)

type Scanner struct {
	todoList *todo.List
	events   []Event
}

func NewScanner(todoList *todo.List) *Scanner {
	return &Scanner{
		todoList: todoList,
	}
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
		event := NewEvent(result, inputString)
		s.events = append(s.events, event)
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
	if cmd == "help" {
		return s.cmdHelp(fields)
	}
	if cmd == "events" {
		return s.cmdEvents(fields)
	}
	return unknownCommand
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

	s.todoList.AddTask(task)

	PrintAdd(title)
	return ""
}

func (s *Scanner) cmdList(fields []string) string {
	if len(fields) != 1 {
		return WrongArgs
	}

	tasks := s.todoList.GetAllTasks()

	PrintTaks(tasks)

	return ""
}

func (s *Scanner) cmdDone(fields []string) string {
	if len(fields) != 2 {
		return WrongArgs
	}
	title := fields[1]

	doneTaskResult := s.todoList.Done(title)
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

	delTaskResult := s.todoList.Delete(title)
	if delTaskResult != "" {
		return delTaskResult
	}

	printDel(title)
	return ""
}

func (s *Scanner) cmdHelp(fields []string) string {
	if len(fields) != 1 {
		return WrongArgs
	}
	printHelp()
	return ""
}

func (s *Scanner) cmdEvents(fields []string) string {
	if len(fields) != 1 {
		return WrongArgs
	}
	printEvents(s.events)
	return ""
}
