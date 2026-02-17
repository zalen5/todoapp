package todo

import "time"

type Task struct {
	Title     string
	Text      string
	IsDone    bool
	CreatedAt time.Time
	DoneAt    *time.Time
}
