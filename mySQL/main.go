package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "michael:michael@tcp(10.91.4.96:3306)/YiChun?charset=utf8")

	if err != nil {
		log.Fatalf("error: %s", err)
	}

	var (
		ID       int
		name     string
		birthday string
		gender   string
	)
	/********************利用條件查詢(1)******************/
	/*stmt, err := db.Prepare("select ID, name, birthday, gender from student where ID = ?")
	if err != nil {
		log.Fatal(err)
	}
	//var name string
	err = stmt.QueryRow(1041419).Scan(&ID, &name, &birthday, &gender)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ID, name, birthday, gender)

	err = stmt.QueryRow(1041108).Scan(&ID, &name, &birthday, &gender)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ID, name, birthday, gender)*/
	/**************************************/

	/********************利用條件查詢(2)******************/
	/*rows, err := db.Query("select name from student where gender=?", "female")
	if err != nil {
		log.Fatal(err)
	}
	var count = 0
	for rows.Next() {
		err = rows.Scan(&name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(name)
		count++
	}
	fmt.Printf("count female has %d\n", count)*/
	/**************************************/

	/*********************查詢全部資料*****************/
	rows, err := db.Query("select * from student")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		err = rows.Scan(&ID, &name, &birthday, &gender)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(ID, name, birthday, gender)
	}
	/**************************************/

	/**********************加入資料****************/
	/*insert, er := db.Prepare("insert into student value(?,?,?,?)")
	if er != nil {
		log.Fatal(er)
	}
	result, ers := insert.Exec(1040000, "Polly", "19960101", "female")
	if ers != nil {
		log.Fatal(ers)
	}
	rowCount, ers := result.RowsAffected()
	if ers != nil {
		log.Fatal(ers)
	}
	fmt.Printf("insert %d of data", rowCount)*/
	/**************************************/

	/**********************刪除資料****************/
	/*del, errs := db.Prepare("delete from student where ID=?")
	if errs != nil {
		log.Fatal(errs)
	}
	res, errs := del.Exec(1040000)
	if errs != nil {
		log.Fatal(errs)
	}
	affect, errs := res.RowsAffected() //刪除幾筆資料
	if errs != nil {
		log.Fatal(errs)
	}
	fmt.Printf("delete %d of data", affect)*/
	/**************************************/

	/***********************資料更新***************/
	/*update, errr := db.Prepare("update student set birthday=? where ID=?")
	if errr != nil {
		log.Fatal(errr)
	}
	_, errr = update.Exec("19960202", 1040000)
	if errr != nil {
		log.Fatal(errr)
	}*/
	/**************************************/

	defer db.Close()
}
