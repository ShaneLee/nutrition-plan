package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	http.HandleFunc("/", Index)
	http.ListenAndServe(":7070", nil)
	db := con()

	results, err := db.Query("SELECT goal FROM goals")

	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var goal string

		err = results.Scan(&goal)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println(goal)
	}

	defer db.Close()
}

func con() (db *sql.DB) {
	db, err := sql.Open("mysql", "root:test@tcp(127.0.0.1:3333)/simple_goals")

	if err != nil {
		panic(err.Error())
	}

	return db
}
