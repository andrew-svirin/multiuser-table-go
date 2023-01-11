package database

// Model uses to represent custom models.

// IModel - database model interface.
type IModel interface {
	Schema() Schema
}

// ModelGenerator - model generator type.
type ModelGenerator func() IModel

// ResolveModelColumnFields - resolve model column fields.
func ResolveModelColumnFields(m IModel) []interface{} {
	scf := make([]interface{}, len(m.Schema().Columns))

	for k, col := range m.Schema().Columns {
		scf[k] = col.field
	}

	return scf
}

// ResolveModelColumnNames - resolve model column names.
func ResolveModelColumnNames(m IModel) []string {
	scn := make([]string, len(m.Schema().Columns))

	for k, col := range m.Schema().Columns {
		scn[k] = col.name
	}

	return scn
}

// ResolveModelTable - resolve model table.
func ResolveModelTable(m IModel) string {
	return m.Schema().Table
}
