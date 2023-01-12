package models

import (
	"github.com/andrew-svirin/multiuser-table-go/server/services/database"
)

// Cell - cell model struct.
type Cell struct {
	ID     int
	Column string
	Row    int
	Value  []byte
}

// Schema - implements repository schema.
func (c *Cell) Schema() *database.Schema {
	return &database.Schema{
		Table: "cells",
		Columns: []database.Column{
			database.NewAutoColumn("id", &c.ID),
			database.NewColumn("column", &c.Column),
			database.NewColumn("row", &c.Row),
			database.NewColumn("value", &c.Value),
		},
	}
}

// NewCell - instantiate new cell.
func NewCell() *Cell {
	return &Cell{}
}
