package database


func (db *appdbimpl) CreateLogin(u User) (User, error) {
	res, err := db.c.Exec("INSERT INTO users(name, userphoto) VALUES (?,?)", u.UserName)

	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return u, err
	}
	u.UserId = uint64(lastInsertID)
	return u, nil
}