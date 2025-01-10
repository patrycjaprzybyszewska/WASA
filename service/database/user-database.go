package database
import (
	"database/sql"
)

func (db *appdbimpl) CreateLogin(u User) (User, error) {
	res, err := db.c.Exec("INSERT INTO users(name, userphoto) VALUES (?,?)", u.UserName)
	if err != nil{
		var user User
		if err := db.c.QueryRow(`SELECT UserId, username FROM users WHERE username = ?`, u.UserName).Scan(&user.UserId, &user.UserName); err != nil {
			if err == sql.ErrNoRows{
				return user, err
			}
		}
		return user, nil

	}
	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return u, err
	}
	u.UserId = uint64(lastInsertID)
	return u, nil
}

