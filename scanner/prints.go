package scanner

import (
	"fmt"
	"todoapp/todo"

	"github.com/k0kubun/pp"
)

func NewPromt() {
	fmt.Print("Введите команду:")
}

func Exit() {
	fmt.Println("Я завершаюсь")
}

func PrintAdd(title string) {
	fmt.Println("Задача '" + title + "' успешно добавлена")
	fmt.Println("")
}

func PrintTaks(tasks map[string]todo.Task) {
	pp.Println("Список дел:", tasks)
	fmt.Println("")
}

func printDone(title string) {
	fmt.Println("Задача'" + title + "'помечена как выполненная")
	fmt.Println("")

}

func printDel(title string) {
	fmt.Println("Задача'" + title + "'успешно удалена")
	fmt.Println("")
}
