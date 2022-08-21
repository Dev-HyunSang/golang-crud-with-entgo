package cmd

import (
	"time"

	"github.com/google/uuid"
)

type ToDos struct {
	ToDoUUID  uuid.UUID `json:"todo_uuid"`
	Done      bool      `json:"bone"`
	ToDo      string    `json:"todo"`
	CreatedAt time.Time `json:"created_at"`
	EditedAt  time.Time `json:"edited_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
