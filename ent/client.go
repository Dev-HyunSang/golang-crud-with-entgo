// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/dev-hyunsang/golang-crud-with-entgo/ent/migrate"

	"github.com/dev-hyunsang/golang-crud-with-entgo/ent/todo"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// ToDo is the client for interacting with the ToDo builders.
	ToDo *ToDoClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.ToDo = NewToDoClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:    ctx,
		config: cfg,
		ToDo:   NewToDoClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:    ctx,
		config: cfg,
		ToDo:   NewToDoClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		ToDo.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.ToDo.Use(hooks...)
}

// ToDoClient is a client for the ToDo schema.
type ToDoClient struct {
	config
}

// NewToDoClient returns a client for the ToDo from the given config.
func NewToDoClient(c config) *ToDoClient {
	return &ToDoClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `todo.Hooks(f(g(h())))`.
func (c *ToDoClient) Use(hooks ...Hook) {
	c.hooks.ToDo = append(c.hooks.ToDo, hooks...)
}

// Create returns a builder for creating a ToDo entity.
func (c *ToDoClient) Create() *ToDoCreate {
	mutation := newToDoMutation(c.config, OpCreate)
	return &ToDoCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of ToDo entities.
func (c *ToDoClient) CreateBulk(builders ...*ToDoCreate) *ToDoCreateBulk {
	return &ToDoCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for ToDo.
func (c *ToDoClient) Update() *ToDoUpdate {
	mutation := newToDoMutation(c.config, OpUpdate)
	return &ToDoUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ToDoClient) UpdateOne(td *ToDo) *ToDoUpdateOne {
	mutation := newToDoMutation(c.config, OpUpdateOne, withToDo(td))
	return &ToDoUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ToDoClient) UpdateOneID(id int) *ToDoUpdateOne {
	mutation := newToDoMutation(c.config, OpUpdateOne, withToDoID(id))
	return &ToDoUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for ToDo.
func (c *ToDoClient) Delete() *ToDoDelete {
	mutation := newToDoMutation(c.config, OpDelete)
	return &ToDoDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ToDoClient) DeleteOne(td *ToDo) *ToDoDeleteOne {
	return c.DeleteOneID(td.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *ToDoClient) DeleteOneID(id int) *ToDoDeleteOne {
	builder := c.Delete().Where(todo.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ToDoDeleteOne{builder}
}

// Query returns a query builder for ToDo.
func (c *ToDoClient) Query() *ToDoQuery {
	return &ToDoQuery{
		config: c.config,
	}
}

// Get returns a ToDo entity by its id.
func (c *ToDoClient) Get(ctx context.Context, id int) (*ToDo, error) {
	return c.Query().Where(todo.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ToDoClient) GetX(ctx context.Context, id int) *ToDo {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *ToDoClient) Hooks() []Hook {
	return c.hooks.ToDo
}