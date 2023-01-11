package database

import "strings"

// Query uses to make and execute SQL queries to database.

// Query - query struct.
type Query struct {
	m              IModel
	database       *Database
	sql            string
	modelGenerator ModelGenerator
}

// getModel - generate model.
func (q *Query) getModel() IModel {
	if q.m == nil {
		q.m = q.modelGenerator()
	}

	return q.m
}

// Exec - execute sql query.
func (q *Query) Exec() *Collection {
	rows, err := q.database.Conn.Query(q.sql)

	if err != nil {
		panic(err)
	}

	return NewCollection(rows, q.modelGenerator)
}

// BuildSelect - build select query.
func (q *Query) BuildSelect() {
	m := q.getModel()
	cn := ResolveModelColumnNames(m)
	tn := ResolveModelTable(m)

	q.sql = "SELECT `"
	q.sql += strings.Join(cn, "`, `")
	q.sql += "` FROM `" + tn + "`"
}

// NewQuery - initialization function for query.
func NewQuery(db *Database, mg ModelGenerator) Query {
	return Query{
		database:       db,
		modelGenerator: mg,
	}
}
