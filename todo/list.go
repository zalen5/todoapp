package todo

type List struct {
	tasks map[string]Task
}

func NewList() *List {
	return &List{
		tasks: make(map[string]Task),
	}
}

func (l *List) AddTask(task Task) {
	l.tasks[task.Title] = task
}

func (l *List) GetAllTasks() map[string]Task {
	return l.tasks
}

func (l *List) Done(title string) string {
	task, ok := l.tasks[title]
	if !ok {
		return TaskNotFound
	}

	l.Done(title)

	l.tasks[title] = task

	return ""
}

func (l *List) Delete(title string) string {
	_, ok := l.tasks[title]
	if !ok {
		return TaskNotFound
	}
	delete(l.tasks, title)

	return ""
}
