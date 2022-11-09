# golang-crud-with-entgo
기존의 사용하던 GORM에서 entgo로 넘어가기 위해서 공부하는 프로젝트입니다.  

## What is the entgo?
>ent is a simple, yet powerful entity framework for Go, that makes it easy to build and maintain applications with large data-models and sticks with the following principles:  
> - Easily model database schema as a graph structure.
> - Define schema as a programmatic Go code.
> - Static typing based on code generation.
> - Database queries and graph traversals are easy to write.
> - Simple to extend and customize using Go templates.

## Stack
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

## ToDos:
- **Database:**
  - [X] Create Schema
  - [X] AutoMigrate(SQLite table)
- **REST API:**
  - [X] Create ToDos
  - [X] Read ToDos
  - [X] Update ToDos
  - [X] Delete ToDos

## Docs:
### POST `/create`
**Reqeust:**
```json
{
    "todo": "hello! hyunsang"
}
```

**Response:**
```json
{
    "message": "성공적으로 새로운 항목을 추가하였습니다.",
    "status": "success",
    "success": true,
    "time": "2022-08-20T21:24:47.336995+09:00"
}
```

### POST `/read`
**Reqeust:**
```json
NULL
```

**Response:**
```json
{
    "datas": [
        {
            "id": 1,
            "todo_uuid": "14653b7c-2083-11ed-bca8-acde48001122",
            "todo": "hello! hyunsang",
            "created_at": "2022-08-20T21:24:47.333822+09:00",
            "edited_at": "2022-08-20T21:24:47.333822+09:00",
            "deleted_at": "2022-08-20T21:24:47.333822+09:00"
        }
    ],
    "message": "성공적으로 할 일을 가지고 왔습니다.",
    "status": "success",
    "success": true,
    "time": "2022-08-20T21:25:24.606855+09:00"
}
```

### POST `/update`
**Reqeust:**
```json
{
    "todo_uuid": "14653b7c-2083-11ed-bca8-acde48001122",
    "todo": "Hello! HyunSang's"
}
```

**Response:**
```json
{
    "message": "정상적으로 할 일을 업로드 하였습니다.",
    "status": "sucess",
    "sucess": true,
    "time": "2022-08-20T21:32:39.879555+09:00"
}
```

### DELETE `/delete`
**Reqeust::**
```json
{
    "todo_uuid": "14653b7c-2083-11ed-bca8-acde48001122"
}
```

**Response:**
```json
{
    "message": "성공적으로 할 일을 삭제하였습니다.",
    "status": "success",
    "success": true,
    "time": "2022-08-20T21:48:56.732876+09:00"
}
```
