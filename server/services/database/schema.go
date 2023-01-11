package database

// Schema helps to map model with database table.

// Schema - schema struct.
type Schema struct {
	Table   string
	Columns []Column
}

// Column - schema column struct.
type Column struct {
	name  string
	field interface{}
}

// NewColumn - instantiate new column.
func NewColumn(name string, field interface{}) Column {
	return Column{name: name, field: field}
}
