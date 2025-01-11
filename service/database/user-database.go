package database
import (
	"database/sql"
)

func (db *appdbimpl) CreateLogin(u User) (User, error) {
	res, err := db.c.Exec("INSERT INTO users(UserName, UserPhoto) VALUES (?,?)", u.UserName, u.UserPhoto)
	if err != nil{
		var user User
		if err := db.c.QueryRow(`SELECT UserId, UserName FROM users WHERE username = ?`, u.UserName).Scan(&user.UserId, &user.UserName); err != nil {
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

func (db *appdbimpl) SetUserName(u User, username string) (User, error) {
	res, err := db.c.Exec(`UPDATE users SET UserName=? WHERE UserId=?`, username, u.UserId)
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

