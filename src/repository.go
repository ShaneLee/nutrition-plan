package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Data struct {
	Fields []string
}

func findAllMeals() Data {

	db := con()

	results, err := db.Query("SELECT name FROM plan")

	if err != nil {
		panic(err.Error())
	}

	names := make([]string, 4)

	for results.Next() {
		var name string

		err = results.Scan(&name)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println(name)
		names = append(names, name)

	}

	defer db.Close()

	data := Data{Fields: names}
	return data
}

func con() (db *sql.DB) {
	db, err := sql.Open("mysql", "root:test@tcp(127.0.0.1:3333)/nutrition")

	if err != nil {
		panic(err.Error())
	}

	return db
}
