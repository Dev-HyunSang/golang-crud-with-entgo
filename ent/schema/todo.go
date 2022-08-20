package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// ToDo holds the schema definition for the ToDo entity.
type ToDo struct {
	ent.Schema
}

// Fields of the ToDo.
func (ToDo) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("todo_uuid", uuid.UUID{}).
			Default(uuid.New),
		field.String("todo").
			Default("unknown"),
		field.Bool("done").
			Default(false),
		field.Time("created_at").
			Default(time.Now()),
		field.Time("edited_at").
			Default(time.Now()),
		field.Time("deleted_at").
			Default(time.Now()),
	}
}

// Edges of the ToDo.
func (ToDo) Edges() []ent.Edge {
	return nil
}
