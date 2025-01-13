package api

import (

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
)


type User struct{
	UserId 		uint64 	`json:"userId"`
	UserName 	string	`json:"name"`
	UserPhoto 	string 	`json:"userPhoto"`
}

type Message struct{
	MessageId   uint64 `json:"messageId"`
	UserId      uint64 `json:"userId"`
	Content     string `json:"content"`
	MessageDate string `json:"messageDate"`
	State       string `json:"state"`
	MessageTime string `json:"messageTime"`
	ChatId      uint64 `json:"chatId"`
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
	ChatUsers	 uint64 `json:"chatUsers`
}

func (u *User) FromDatabase(user database.User){
	u.UserId = user.UserId
	u.UserName = user.UserName
	u.UserPhoto = user.UserPhoto
}

func (u *User) ToDatabase() database.User{
	return database.User{
		UserId: u.UserId,
		UserPhoto: u.UserPhoto,
		UserName: u.UserName,
	}
}

func(m *Message) MessageFromDatabase(message database.Message){
	m.MessageId = message.MessageId
	m.UserId = message.UserId
	m.Content = message.Content
	m.MessageDate = message.MessageDate
	m.State = message.State
	m.MessageTime = message.MessageTime
}

func (m *Message) MessageToDatabase() database.Message {
	return database.Message{
		MessageId:  	m.MessageId,
		UserId:	 		m.UserId,
		Content:		m.Content,
		MessageDate:	m.MessageDate,
		State:			m.State,
		MessageTime: 	m.MessageTime,
	}
}

func(c *Comment) CommentFromDatabase(comment database.Comment){
	c.MessageId = comment.MessageId
	c.Content = comment.Content
	c.CommentId = comment.CommentId//mozliwe ze teho tu nie trzeba

}

func (c *Comment) CommentToDatabase() database.Comment {
	return database.Comment{
		MessageId:  	c.MessageId,
		Content:		c.Content,
	}
}