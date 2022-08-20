package database

import (
	"github.com/dev-hyunsang/golang-crud-with-entgo/ent"
	_ "github.com/mattn/go-sqlite3"
)

func ConnectionSQLite() (*ent.Client, error) {
	client, err := ent.Open("sqlite3", "file:./golang-crud-with-entgo.db?_fk=1")
	if err != nil {
		return nil, err
	}

	return client, nil
}
