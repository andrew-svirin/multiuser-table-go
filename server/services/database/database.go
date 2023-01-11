package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// Database - custom database structure.
type Database struct {
	config *Config
	Conn   *sql.DB
}

// Open - open database
func (db *Database) Open() {
	conn, err := sql.Open(db.config.Driver, db.config.Dsn)

	if err != nil {
		panic(err)
	}

	db.Conn = conn
}

// Close - close database
func (db *Database) Close() {
	err := db.Conn.Close()

	if err != nil {
		panic(err)
	}
}

// NewQuery - create new query.
func (db *Database) NewQuery(r IRepository) Query {
	return NewQuery(db, r.ModelGenerator())
}

// NewDatabase - initialize new database.
func NewDatabase(c *Config) *Database {
	return &Database{
		config: c,
	}
}

// Db - global database instance.
var Db *Database
