package database

import (
	. "bot/config"
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

type DBEntry struct {
	Name string
	Fields []string
}

var DBEntries []*DBEntry

func init() {
	db, err := sql.Open("sqlite3", Config.DB)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	
	DB = db

	for _, entry := range DBEntries {
		_, err = DB.Exec(
			fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s)",
				entry.Name, strings.Join(entry.Fields, ", ")))
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Println("DB", entry.Name, "ready")
	}
}
