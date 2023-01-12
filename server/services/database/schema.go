package database

// Schema helps to map model with database table.

// Schema - schema struct.
type Schema struct {
	// Table - table name in database.
	Table   string
	Columns []Column
}

// ResolveColumnFields - resolve column fields.
func (s *Schema) ResolveColumnFields() []any {
	scf := make([]any, len(s.Columns))

	for k, col := range s.Columns {
		scf[k] = col.field
	}

	return scf
}

// ResolveColumnNames - resolve column names.
func (s *Schema) ResolveColumnNames() []string {
	scn := make([]string, len(s.Columns))

	for k, col := range s.Columns {
		scn[k] = col.name
	}

	return scn
}

// ResolveMutableColumnFields - resolve mutable column fields.
func (s *Schema) ResolveMutableColumnFields() []any {
	var scf []any

	for _, col := range s.Columns {
		if col.IsMutable() {
			scf = append(scf, col.field)
		}
	}

	return scf
}

// ResolveMutableColumnNames - resolve mutable column names.
func (s *Schema) ResolveMutableColumnNames() []string {
	var scn []string

	for _, col := range s.Columns {
		if col.IsMutable() {
			scn = append(scn, col.name)
		}
	}

	return scn
}

// Column - schema column struct.
type Column struct {
	// name - name of field in database.
	name string
	// field - pointer on model field.
	field any
	// auto - auto is the autofilled flag
	auto bool
}

// IsMutable - is mutable column?
func (c *Column) IsMutable() bool {
	return c.auto != true
}

// NewColumn - instantiate new column.
func NewColumn(name string, field any) Column {
	return Column{name: name, field: field}
}

// NewAutoColumn - instantiate new auto column.
func NewAutoColumn(name string, field any) Column {
	c := NewColumn(name, field)
	c.auto = true

	return c
}
