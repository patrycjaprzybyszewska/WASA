package database

import (
	"database/sql"
)

func (db *appdbimpl) CreateUser(u User) (User, error) {
	res, err := db.c.Exec("INSERT INTO users(name, userphoto) VALUES (?,?)", u.Username)
	if err != nil {
		var user User
	}
	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return u, err
	}
	u.UserIdId = uint64(lastInsertID)
	return u, nil
}