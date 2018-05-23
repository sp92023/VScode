package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "michael:michael@tcp(10.91.4.96:22)/YiChun?charset=utf8")
	//db, err := sql.Open("mysql", "michael:michael@/YiChun?charset=utf8")

	if err != nil {
		log.Fatalf("error: %s", err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	/*var (
		ID       int
		name     string
		birthday string
		gender   string
	)
	ID = 1041419
	rows, err := db.Query("SELECT", ID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&ID, &name, &birthday, &gender); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d,%s,%s,%s", ID, name, birthday, gender)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}*/

	db.Exec("insert", 1041419, "Polly", "19960902", "female")
}
