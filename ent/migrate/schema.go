// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ToDosColumns holds the columns for the "to_dos" table.
	ToDosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "todo_uuid", Type: field.TypeUUID},
		{Name: "todo", Type: field.TypeString, Default: "unknown"},
		{Name: "done", Type: field.TypeBool, Default: false},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "edited_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime},
	}
	// ToDosTable holds the schema information for the "to_dos" table.
	ToDosTable = &schema.Table{
		Name:       "to_dos",
		Columns:    ToDosColumns,
		PrimaryKey: []*schema.Column{ToDosColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ToDosTable,
	}
)

func init() {
}
