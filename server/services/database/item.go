package database

import "database/sql"

// Item uses to manage single query row.

// Item - item struct.
type Item struct {
	row            *sql.Row
	modelGenerator ModelGenerator
}

// One - applying handler for one row.
func (i *Item) One(h ModelHandler) {
	m := i.modelGenerator()
	s := m.Schema()
	f := s.ResolveColumnFields()

	err := i.row.Scan(f...)

	if err != nil {
		if err.Error() != "sql: no rows in result set" {
			panic(err)
		}
	}

	h(m)
}

// NewItem - instantiate new item.
func NewItem(r *sql.Row, mg ModelGenerator) *Item {
	return &Item{
		row:            r,
		modelGenerator: mg,
	}
}
