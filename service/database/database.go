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
	SenderId	uint64 `json:"senderId"`
	ChatId      uint64 `json:"chatId"`
	Content     string `json:"content"`
	MessageDate string `json:"messageDate"`
	State       string `json:"state"`
	MessageTime string `json:"messageTime"`
}
type Comment struct{
	CommentId uint64  `json:"commentId"`
	MessageId uint64 `json:"messageId"`
	Content   string  `json:"content"`
}

type Chat struct{
	ChatId		 uint64 `json:"chatId"`
	ChatName 	 string `json:"chatName"`
	ChatPhoto 	 string `json:"chatPhoto"`//trzeba dodac uzytkownikow
}

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	GetName() (string, error)
	SetUsername(User, string) (User, error)
	SetUserphoto(User, string) (User, error)
	GetUserPhotoById(uint64) (string, error)
	GetUserNameById(uint64) (string, error)
	GetChatPhotoById(uint64) (string, error)
	GetChatNameById(uint64) (string, error)
	CreateLogin(User) (User, error)
	GetMessageById(uint64) (Message, error)
	GetCommentById(uint64) (Comment, error)
	Sendmessage(Message) (Message, error)
	Removemessage(uint64) error
	Commentmessage(Comment) (Comment, error)
	AddUserToChat(uint64, uint64) error 
	LeaveGroup(uint64, uint64) error 
	SetGroupName(Chat, string) (Chat, error)
	SetGroupPhoto(Chat, string) (Chat, error)
	GetConversation(uint64) ([]Message, error)
	GetChats()([]Chat, error)
	Removecomment(uint64) error
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
	
		sqlStmt := `
		DROP TABLE users
		CREATE TABLE IF NOT EXISTS users (
			userId INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			userName STRING NOT NULL, 
			userPhoto STRING	
		);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating 'users' table: %w", err)
		}

		sqlStmt = `
		DROP TABLE chats
		CREATE TABLE IF NOT EXISTS chats (
			chatId INTEGER PRIMARY KEY,
			chatName STRING,
			chatPhoto STRING
		);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating 'chats' table: %w", err)
		}
	
		sqlStmt = `
		DROP TABLE chat_users
		CREATE TABLE IF NOT EXISTS chat_users (
  			chatId INTEGER,
  			userId INTEGER,
   			PRIMARY KEY (chatId, userId),
   			FOREIGN KEY (chatId) REFERENCES chats(chatId),
    		FOREIGN KEY (userId) REFERENCES users(userId)
		);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating 'chat_users' table: %w", err)
		}
	
		sqlStmt = `
		DROP TABLE messages
		CREATE TABLE IF NOT EXISTS messages (
			messageId INTEGER NOT NULL PRIMARY KEY,
			senderId INTEGER,
			content TEXT,
			messageDate TEXT,
			messageTime TEXT,
			state TEXT,
			chatId INTEGER,
			FOREIGN KEY (chatId) REFERENCES chats(chatId)
		);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating 'messages' table: %w", err)
		}
	
	
		sqlStmt = `
		DROP TABLE comments
		CREATE TABLE IF NOT EXISTS comments (
			commentId INTEGER NOT NULL PRIMARY KEY,
			content STRING,
			messageId INTEGER,
			FOREIGN KEY (messageId) REFERENCES messages(messageId)
	);`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, fmt.Errorf("error creating 'comments' table: %w", err)
	}

	return &appdbimpl{
		c: db,
	}, nil
}
func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
