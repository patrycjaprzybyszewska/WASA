package database

import (
	"fmt"
	"database/sql"
)

func (db *appdbimpl) CreateLogin(u User) (User, error) {
	var user User
	err := db.c.QueryRow(`SELECT UserId, UserName, UserPhoto FROM users WHERE userName = ?`, u.UserName).Scan(&user.UserId, &user.UserName, &user.UserPhoto)
	if err != nil {
		if err == sql.ErrNoRows {
			res, err := db.c.Exec("INSERT INTO users(UserName, UserPhoto) VALUES (?,?)", u.UserName, u.UserPhoto)

			lastInsertID, err := res.LastInsertId()
			if err != nil {
				return u, err
			}
			u.UserId = uint64(lastInsertID)
			return u, nil
		}
		return user, nil
	}
	return user, nil

}
func (db *appdbimpl) GetUserPhotoById(userId uint64) (string, error) {
	var userPhoto string
	err := db.c.QueryRow(`SELECT UserPhoto FROM users WHERE userId = ?`, userId).Scan(&userPhoto)
	if err != nil {
		return "", err
	}
	return userPhoto, nil
}
func (db *appdbimpl) GetUserNameById(userId uint64) (string, error) {
	var userName string
	err := db.c.QueryRow(`SELECT UserName FROM users WHERE userId = ?`, userId).Scan(&userName)
	if err != nil {
		return "", err
	}
	return userName, nil
}
func (db *appdbimpl) SetUsername(u User, username string) (User, error) {
	var userId uint64
	err := db.c.QueryRow(`SELECT userName FROM users WHERE userId = ?`, userId).Scan(&username)
	if userId != nil{
		  return u, fmt.Errorf("UserName is not aviable, cannot be changed")}
	
	res, err := db.c.Exec(`UPDATE users SET userName=?, userPhoto=? WHERE userId=?`, username, u.UserPhoto, u.UserId)
	if err != nil {
		return u, err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return u, err
	} else if affected == 0 {
		return u, err
	}
	return u, nil
}



func (db *appdbimpl) SetUserphoto(u User, photo string) (User, error) {
	res, err := db.c.Exec(`UPDATE users SET userPhoto=?, userName=? WHERE userId=?`, photo, u.UserName, u.UserId)
	if err != nil {
		return u, err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return u, err
	} else if affected == 0 {
		return u, err
	}
	return u, nil
}
