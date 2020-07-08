package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:test@tcp(127.0.0.1:8080)/simple_goals")

	if err != nil {
		panic(err.Error())
	}

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
