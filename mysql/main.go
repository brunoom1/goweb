package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	id         uint
	username   string
	password   string
	created_at time.Time
}

func ShowUser(u *User) {
	fmt.Printf("ID: %d, Username: %s, Password: %s, Updated_AT: %s\n", u.id, u.username, u.password, u.created_at)
}

func InsertUser(db *sql.DB, u *User) (*User, error) {
	result, err := db.Exec(`insert into users (username, password, created_at) values (?, ?, ?)`, u.username, u.password, u.created_at)

	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return nil, err
	}
	u.id = uint(id)

	return u, nil
}

func FetchUsers(db *sql.DB) ([]User, error) {
	query := "select * from users"
	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}

	var users []User

	for rows.Next() {
		var u User
		rows.Scan(&u.id, &u.username, &u.password, &u.created_at)
		users = append(users, u)
	}

	err = rows.Err()

	if err != nil {
		return nil, err
	}

	return users, nil
}

func main() {
	db, err := sql.Open("mysql", "root:123456@(127.0.0.1:3306)/tasks?parseTime=true")

	if err != nil {
		panic(fmt.Sprintf("Mysql error: %s", err.Error()))
	}

	pingErr := db.Ping()

	if pingErr != nil {
		panic("Não foi possível acessa base de dados")
	}

	users, err := FetchUsers(db)

	if err != nil {
		panic(err.Error())
	}

	for _, user := range users {
		ShowUser(&user)
	}

	user, err := InsertUser(db, &User{
		username:   "Gabriel Mendonça",
		password:   "123456",
		created_at: time.Now(),
	})

	if err != nil {
		panic(err.Error())
	}

	if user != nil {
		ShowUser(user)
	}

}
