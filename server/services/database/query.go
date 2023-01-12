package database

import (
	"strings"
)

// Query uses to build and execute SQL queries to database.

// Query - query struct.
type Query struct {
	m              IModel
	database       *Database
	sql            string
	modelGenerator ModelGenerator
}

// getModel - resolve model schema.
func (q *Query) modelSchema() *Schema {
	if q.m == nil {
		q.m = q.modelGenerator()
	}

	return q.m.Schema()
}

// Select - execute select one sql query.
func (q *Query) Select(args ...any) *Item {
	row := q.database.Conn.QueryRow(q.sql, args...)

	return NewItem(row, q.modelGenerator)
}

// SelectAll - execute select all sql query.
func (q *Query) SelectAll() *Collection {
	rows, err := q.database.Conn.Query(q.sql)

	if err != nil {
		panic(err)
	}

	return NewCollection(rows, q.modelGenerator)
}

// BuildSelect - build select sql query.
func (q *Query) BuildSelect() {
	s := q.modelSchema()
	cn := s.ResolveColumnNames()
	tn := s.Table

	q.sql = "SELECT "
	q.sql += "`" + strings.Join(cn, "`, `") + "` "
	q.sql += "FROM `" + tn + "`"
}

// Where - adds where conditions to sql query.
func (q *Query) Where(wc WhereCondition) {
	q.sql += " WHERE "
	q.sql += strings.Join(wc.Conditions, " "+wc.Op+" ")
}

// BuildInsert - build insert sql query.
func (q *Query) BuildInsert() {
	s := q.modelSchema()
	cn := s.ResolveMutableColumnNames()
	tn := s.Table

	q.sql = "INSERT "
	q.sql += "INTO `" + tn + "` "
	q.sql += "(`" + strings.Join(cn, "`, `") + "`) "
	q.sql += "VALUES (" + strings.Repeat("?, ", len(cn)-1) + "?)"
}

// Insert - execute insert one sql query.
func (q *Query) Insert(m IModel) int64 {
	s := m.Schema()
	fs := s.ResolveMutableColumnFields()

	ir, err := q.database.Conn.Exec(q.sql, fs...)

	if err != nil {
		panic(err)
	}

	id, err := ir.LastInsertId()

	if err != nil {
		panic(err)
	}

	return id
}

// BuildUpdate - build update sql query.
func (q *Query) BuildUpdate() {
	s := q.modelSchema()
	cn := s.ResolveMutableColumnNames()
	tn := s.Table

	q.sql = "UPDATE "
	q.sql += "`" + tn + "` "
	q.sql += "SET `" + strings.Join(cn, "`=?, `") + "`=?"
}

// Update - execute update one sql query.
func (q *Query) Update(id *int, m IModel) {
	s := m.Schema()
	fs := s.ResolveMutableColumnFields()

	_, err := q.database.Conn.Exec(q.sql, append(fs, id)...)

	if err != nil {
		panic(err)
	}
}

// NewQuery - initialization function for query.
func NewQuery(db *Database, mg ModelGenerator) Query {
	return Query{
		database:       db,
		modelGenerator: mg,
	}
}

// WhereCondition - where condition struct.
type WhereCondition struct {
	Conditions []string
	Op         string
}

// BuildWhere - build where expression
func (wc *WhereCondition) BuildWhere() string {
	return strings.Join(wc.Conditions, " "+wc.Op+" ")
}
