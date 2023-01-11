package database

import (
	"database/sql"
	"log"
)

// Collection uses to manage query rows.

// ModelHandler - model handler type.
type ModelHandler func(IModel)

// Collection - collection struct.
type Collection struct {
	rows           *sql.Rows
	modelGenerator ModelGenerator
}

// Each - walker over rows.
func (c *Collection) Each(h ModelHandler) {
	for c.rows.Next() {
		m := c.modelGenerator()

		f := ResolveModelColumnFields(m)

		err := c.rows.Scan(f...)

		if err != nil {
			log.Fatal(err)
		}

		h(m)
	}
}

// NewCollection - instantiate new collection.
func NewCollection(rs *sql.Rows, mg ModelGenerator) *Collection {
	return &Collection{
		rows:           rs,
		modelGenerator: mg,
	}
}
