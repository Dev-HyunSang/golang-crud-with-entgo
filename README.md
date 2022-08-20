# golang-crud-with-entgo
기존의 사용하던 GORM에서 entgo로 넘어가기 위해서 공부하는 프로젝트입니다.  

## What is the entgo?
>ent is a simple, yet powerful entity framework for Go, that makes it easy to build and maintain applications with large data-models and sticks with the following principles:  
> - Easily model database schema as a graph structure.
> - Define schema as a programmatic Go code.
> - Static typing based on code generation.
> - Database queries and graph traversals are easy to write.
> - Simple to extend and customize using Go templates.

## 🛠 Stack
- [`gofiber/fiber`](https://gofiber.io/)
- [`ent/ent` with Facebook](https://entgo.io/)
- [`gofrs/uuid`](github.com/gofrs/uuid)
- [SQLite](https://www.sqlite.org/index.html)
  - Dirver. [`mattn/go-sqlite3`](https://pkg.go.dev/github.com/mattn/go-sqlite3@v1.14.15?utm_source=gopls)

## Installation
```shell
$ go install entgo.io/ent/cmd/ent@latest
```

## Create Schema
```shell
$ go run entgo.io/ent/cmd/ent init <Schema Name>
```

## Generate
```shell
$ go generate ./ent
```

## [Types](https://entgo.io/docs/schema-fields/)
- All Go numeric types. Like int, uint8, float64, etc.
- bool
- string
- time.Time
- UUID
- []byte (SQL only).
- JSON (SQL only).
- Enum (SQL only).
- Other (SQL only).