package main

import (
	"database/sql"
	"fmt"
	"log"
)

type user struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (u *user) getUser(db *sql.DB) error {
	//return errors.New("Not implemented")
	statement := fmt.Sprintf("select name, age from person where id=%d", u.ID)
	return db.QueryRow(statement).Scan(&u.Name, &u.Age)
}

func (u *user) updateUser(db *sql.DB) error {
	//return errors.New("Not implemented")
	statement := fmt.Sprintf("update person set name='%s',age=%d where id=%d", u.Name, u.Age, u.ID)
	_, err := db.Exec(statement)
	return err
}

func (u *user) deleteUser(db *sql.DB) error {
	//return errors.New("Not implemented")
	statement := fmt.Sprintf("delete from person where id=%d", u.ID)
	_, err := db.Exec(statement)
	return err
}

func (u *user) createUser(db *sql.DB) error {
	//return errors.New("Not implemented")
	statement := fmt.Sprintf("insert into person(id,name,age)values(%d,'%s',%d)", u.ID, u.Name, u.Age)
	_, err := db.Exec(statement)

	if err != nil {
		return err
	}
	err = db.QueryRow("select LAST_INSERT_ID()").Scan(&u.ID)

	if err != nil {
		return err
	}

	return nil
}

func getUsers(db *sql.DB, start, count int) ([]user, error) {
	//return nil,errors.New("Not implemented")

	statement := fmt.Sprintf("select id,name,age from person limit %d offset %d", count, start)
	rows, err := db.Query(statement)

	if err != nil {
		//return nil, err
		log.Fatal(err)
	}

	defer rows.Close()

	users := []user{}

	for rows.Next() {
		var u user
		if err := rows.Scan(&u.ID, &u.Name, &u.Age); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}
