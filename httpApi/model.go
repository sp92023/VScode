package main

import (
	"database/sql"
	"fmt"
	"log"
)

type user struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Birthday string `json:"birthday"`
	Gender   string `json:"gender"`
}

func getUsers(db *sql.DB, start, count int) ([]user, error) {
	statement := fmt.Sprintf("select ID,Name,Birthday,Gender from student limit %d offset %d", count, start)
	rows, err := db.Query(statement)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	users := []user{}

	for rows.Next() {
		var u user
		if err := rows.Scan(&u.ID, &u.Name, &u.Birthday, &u.Gender); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (u *user) createUser(db *sql.DB) error {
	statement := fmt.Sprintf("insert into student(ID,Name,Birthday,Gender)values(%d,'%s','%s','%s')", u.ID, u.Name, u.Birthday, u.Gender)
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

func (u *user) getUser(db *sql.DB) error {
	statement := fmt.Sprintf("select Name,Birthday,Gender from student where ID=%d", u.ID)
	return db.QueryRow(statement).Scan(&u.Name, &u.Birthday, &u.Gender)
}

func (u *user) updateUser(db *sql.DB) error {
	statement := fmt.Sprintf("update student set Name='%s',Birthday='%s',Gender='%s' where ID=%d", u.Name, u.Birthday, u.Gender, u.ID)
	_, err := db.Exec(statement)
	return err
}

func (u *user) deleteUser(db *sql.DB) error {
	statement := fmt.Sprintf("delete from student where ID=%d", u.ID)
	_, err := db.Exec(statement)
	return err
}
