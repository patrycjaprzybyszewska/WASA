package database
import (
	"database/sql"
)

func (db *appdbimpl) CreateLogin(u User) (User, error) {
	var user User
	err := db.c.QueryRow(`SELECT UserId, UserName, UserPhoto FROM users WHERE username = ?`, u.UserName).Scan(&user.UserId, &user.UserName, &user.UserPhoto); 
	if err != nil {
		if err == sql.ErrNoRows{
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

func (db *appdbimpl) SetUsername(u User, username string) (User, error) {

	var currentUser User
    err := db.c.QueryRow(`SELECT UserId, UserName, UserPhoto FROM users WHERE UserId = ?`, u.UserId).Scan(&currentUser.UserId, &currentUser.UserName, &currentUser.UserPhoto)
    if err != nil {
        return u, err
    }
	res, err := db.c.Exec(`UPDATE users SET UserName=?, UserPhoto =? WHERE UserId=?`, username, currentUser.UserPhoto, u.UserId)
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
	var currentUser User
    err := db.c.QueryRow(`SELECT UserId, UserName, UserPhoto FROM users WHERE UserId = ?`, u.UserId).Scan(&currentUser.UserId, &currentUser.UserName, &currentUser.UserPhoto)
    if err != nil {
        return u, err
    }
	res, err := db.c.Exec(`UPDATE users SET UserName=?, UserPhoto =? WHERE UserId=?`, currentUser.UserName, photo, u.UserId)
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