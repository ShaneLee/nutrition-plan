package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func findAllMeals() {

	db := con()

	results, err := db.Query("SELECT name FROM meals")

	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var name string

		err = results.Scan(&goal)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println(name)
	}

	defer db.Close()
}

func con() (db *sql.DB) {
	db, err := sql.Open("mysql", "root:test@tcp(127.0.0.1:3333)/nutrition")

	if err != nil {
		panic(err.Error())
	}

	return db
}
