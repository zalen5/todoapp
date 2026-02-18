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

func printHelp() {
	fmt.Println("Список команд:")
	fmt.Println("1. Help")
	fmt.Println("-- эта команда позволяет узнать доступные команды")
	fmt.Println("")
	fmt.Println("2. add {заголовок задачи одно слово} {текст задачи}")
	fmt.Println("-- позволяет добавлять задачи")
	fmt.Println("")
	fmt.Println("3. list")
	fmt.Println("-- позволяет получить полный список")
	fmt.Println("")
	fmt.Println("4. del {заголовок сущетсвующей задачи}")
	fmt.Println("-- эта команда позволяет удалить задачу по её заголовку")
	fmt.Println("")
	fmt.Println("5. done {заголовок существующей задачи}")
	fmt.Println("-- позволяет отменить задачу как вы её выполнили")
	fmt.Println("")
	fmt.Println("6. events")
	fmt.Println("-- позволяет получить список всех событий")
	fmt.Println("")
	fmt.Println("7. exit ")
	fmt.Println("-- позволяет завершить выполнение программы")
	fmt.Println("")
}

func printEvents(events []Event) {
	pp.Println("События:", events)
	fmt.Println("")
}
