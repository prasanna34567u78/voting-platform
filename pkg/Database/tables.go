package database

import (
	"fmt"
	"log"
)

func UserTable() {
	newcontroller, err := NewMyController()
	if err != nil {
		log.Fatal()
	}
	query := `CREATE TABLE IF NOT EXISTS users (
		
		email text primary key,
		password text,
		
	)`
	newcontroller.ExcuteQuery(query)
}

// creating the tables
// function to create voter table function name :VoterTable

func VoterTable() {
	newcontroller, err := NewMyController()
	if err != nil {
		log.Fatal()
	}
	query := `CREATE TABLE IF NOT EXISTS voter (
		id UUID ,
		firstname text,
		lastname text,
		email text ,
		password text,
		verified boolean,
		voted boolean,
		primary key(id,email)
	)`
	newcontroller.ExcuteQuery(query)
}

// function to create Position table function name :PositonTable
func PositionTable() {
	newcontroller, err := NewMyController()
	if err != nil {
		log.Fatal()
	}
	query := `CREATE TABLE IF NOT EXISTS position (
		id UUID primary key,
		name text ,
		priority int
	)`
	newcontroller.ExcuteQuery(query)
}

// function to create candidate table function name :CandidatesTable
func CandidatesTable() {
	newcontroller, err := NewMyController()
	if err != nil {
		log.Fatal()
	}
	query := `CREATE TABLE IF NOT EXISTS candidate (
		id UUID primary key,
		fullname text ,
		bio text,
		position text
	)`
	newcontroller.ExcuteQuery(query)
}

// function to create votes table function name :VotesTable
func VotesTable() {
	newcontroller, err := NewMyController()
	if err != nil {
		log.Fatal()
	}
	query := `CREATE TABLE IF NOT EXISTS votes (
		votes UUID primary key,
		candidate text,
		position text
	)`
	newcontroller.ExcuteQuery(query)
}

/*
Function to create setupdb
function name SetUpDB
*/
func SetUpDB() {
	UserTable()
	VoterTable()
	CandidatesTable()
	PositionTable()
	VotesTable()
	fmt.Println("comment the database set line tables are created and comment the import section also")
}
