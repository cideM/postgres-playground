package main

import (
	"database/sql"
	"log"

	_ "embed"

	_ "github.com/lib/pq"
)

//go:embed setup.sql
var sqlCommands string

type substance struct {
	name  string
	drugs []drug
}

type drug struct {
	name string
}

func main() {
	connStr := "user=foo dbname=foo sslmode=disable password=foo port=5555"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(sqlCommands)
	if err != nil {
		log.Fatal(err)
	}
  // Get substances and drugs from the database, so that we end up having two
  // substances as Go values where each substance has two drugs. Experiment
  // with different ways of performing the JOINs
}
