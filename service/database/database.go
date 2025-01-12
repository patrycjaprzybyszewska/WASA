/*
Package database is the middleware between the app database and the code. All data (de)serialization (save/load) from a
persistent database are handled here. Database specific logic should never escape this package.

To use this package you need to apply migrations to the database if needed/wanted, connect to it (using the database
data source name from config), and then initialize an instance of AppDatabase from the DB connection.

For example, this code adds a parameter in `webapi` executable for the database data source name (add it to the
main.WebAPIConfiguration structure):

	DB struct {
		Filename string `conf:""`
	}

This is an example on how to migrate the DB and connect to it:

	// Start Database
	logger.Println("initializing database support")
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		logger.WithError(err).Error("error opening SQLite DB")
		return fmt.Errorf("opening SQLite: %w", err)
	}
	defer func() {
		logger.Debug("database stopping")
		_ = db.Close()
	}()

Then you can initialize the AppDatabase and pass it to the api package.
*/
package database

import (
	"database/sql"
	"errors"
	"fmt"
)
type User struct {
	UserId 		uint64 	`json:"userId"`
	UserName 	string 	`json:"name"`
	UserPhoto 	string 	`json:"userPhoto"`
}
type Message struct{
	MessageId   uint64 `json:"messageId"`
	UserId      uint64 `json:"userId"`
	Content     string `json:"content"`
	MessageDate string `json:"messageDate"`
	State       string `json:"state"`
	MessageTime string `json:"messageTime"`
	ChatId      uint64 `json:"ChatId"`
}
type Comment struct{
	CommentId uint64  `json:"commentId"`
	MessageId uint64 `json:"messageId"`
	Content   string  `json:"content"`
}

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	GetName() (string, error)
	SetUsername(User, string) (User, error)
	SetUserphoto(User, string) (User, error)
	CreateLogin(User) (User, error)
	GetMessageById(uint64) (Message, error)
	Sendmessage(Message) (Message, error)
	Removemessage(uint64) error
	Commentmessage(Comment) (Comment, error)
	Ping() error
}

type appdbimpl struct {
	c *sql.DB
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}
	_, err := db.Exec("PRAGMA foreign_keys = ON")
	if err != nil {
		return nil, err
	}
	// Check if table exists. If not, the database is empty, and we need to create the structure
	var tableName string
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='users';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		usersDatabase := `CREATE TABLE users (
			UserId INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			UserName TEXT NOT NULL, 
			UserPhoto BLOB	
			);`
			messagesDatabase := `CREATE TABLE messages (
			messageId INTEGER NOT NULL PRIMARY KEY,
			content TEXT,
			messageDate TEXT,
			messageTime TEXT,
			state TEXT,
			userId INTEGER,
			FOREIGN KEY (userId) REFERENCES users(userId)
			);`
			commentsDatabase := `CREATE TABLE comments (
			commentId INTEGER NOT NULL PRIMARY KEY,
			content TEXT,
			FOREIGN KEY (messageId) REFERENCES messages(messageId)
			);`
		_, err = db.Exec(usersDatabase)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
		_, err = db.Exec(messagesDatabase)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
		_, err = db.Exec(commentsDatabase)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
		
		
	}

	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
