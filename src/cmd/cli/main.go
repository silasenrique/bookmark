package main

import (
	"database/sql"
	"go-read-list/src/business/interface/cli"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func initDb() *sql.DB {
	db, err := sql.Open("sqlite3", "file:../../../db/readlist.db?_fk=true&_case_sensitive_like=true")
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}

func main() {
	db := initDb()
	s := cli.NewWrapperRepository(db)

	cmd := cli.LoadCommands(s)

	if err := cmd.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
