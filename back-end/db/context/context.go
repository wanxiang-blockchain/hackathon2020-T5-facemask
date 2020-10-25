package context

import (
	"data-manager/db"
)

type Context struct {
	*db.DB
}

func New(db *db.DB) *Context {
	return &Context{DB: db}
}
