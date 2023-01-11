package app

import (
	"github.com/andrew-svirin/multiuser-table-go/server/services/database"
	"log"
	"time"
)

// initDb - init database.
func (a *App) initDb() {
	database.Db = database.NewDatabase(&a.config.Db)
}

// openDb - open database
func (a *App) openDb() {
	log.Println("DB Open")
	database.Db.Open()

	// See "Important settings" section.
	database.Db.Conn.SetConnMaxLifetime(time.Minute * 3)
	database.Db.Conn.SetMaxOpenConns(10)
	database.Db.Conn.SetMaxIdleConns(10)
}

// closeDb - close database
func (a *App) closeDb() {
	log.Println("DB Close")
	database.Db.Close()
}
