package database

import (
	"database/sql"
)

// Collection uses to manage query rows.

// Collection - collection struct.
type Collection struct {
	rows           *sql.Rows
	modelGenerator ModelGenerator
}

// Each - walker over rows and applying handler.
func (c *Collection) Each(h ModelHandler) {
	for c.rows.Next() {
		m := c.modelGenerator()
		s := m.Schema()
		f := s.ResolveColumnFields()

		err := c.rows.Scan(f...)

		if err != nil {
			panic(err)
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
