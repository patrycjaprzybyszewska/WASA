package api

import (
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"regexp"
	"strconv"
)

type User struct {
	UserId    uint64 `json:"userId"`
	UserName  string `json:"name"`
	UserPhoto string `json:"userPhoto"`
}

type Message struct {
	MessageId   uint64 `json:"messageId"`
	SenderName  string `json:"senderName"`
	SenderId    uint64 `json:"senderId"`
	ChatId      uint64 `json:"chatId"`
	Content     string `json:"content"`
	MessageDate string `json:"messageDate"`
	State       string `json:"state"`
	MessageTime string `json:"messageTime"`
}

type Comment struct {
	CommentId uint64 `json:"commentId"`
	MessageId uint64 `json:"messageId"`
	Content   string `json:"content"`
	Commentername string `json:"commentername"`
}
type Chat struct {
	ChatId    uint64 `json:"chatId"`
	ChatName  string `json:"chatName"`
	ChatPhoto string `json:"chatPhoto"` // trzeba dodac uzytkownikow
}

func (u *User) FromDatabase(user database.User) {
	u.UserId = user.UserId
	u.UserName = user.UserName
	u.UserPhoto = user.UserPhoto
}

func (u *User) ToDatabase() database.User {
	return database.User{
		UserId:    u.UserId,
		UserPhoto: u.UserPhoto,
		UserName:  u.UserName,
	}
}

func (m *Message) MessageFromDatabase(message database.Message) {
	m.MessageId = message.MessageId
	m.SenderName = message.SenderName
	m.SenderId = message.SenderId
	m.ChatId = message.ChatId
	m.Content = message.Content
	m.MessageDate = message.MessageDate
	m.State = message.State
	m.MessageTime = message.MessageTime
}

func (m *Message) MessageToDatabase() database.Message {
	return database.Message{
		MessageId:   m.MessageId,
		SenderName:  m.SenderName,
		SenderId:    m.SenderId,
		ChatId:      m.ChatId,
		Content:     m.Content,
		MessageDate: m.MessageDate,
		State:       m.State,
		MessageTime: m.MessageTime,
	}
}

func (c *Comment) CommentFromDatabase(comment database.Comment) {
	c.MessageId = comment.MessageId
	c.Content = comment.Content
	c.CommentId = comment.CommentId
	c.Commentername = comment.Commentername // mozliwe ze teho tu nike trzeba

}

func (c *Comment) CommentToDatabase() database.Comment {
	return database.Comment{
		MessageId: c.MessageId,
		Content:   c.Content,
		Commentername: c.Commentername,
	}
}

func (ch *Chat) ChatToDatabase() database.Chat {
	return database.Chat{
		ChatId:    ch.ChatId,
		ChatName:  ch.ChatName,
		ChatPhoto: ch.ChatPhoto,
	}
}
func (ch *Chat) ChatFromDatabase(chat database.Chat) {
	ch.ChatId = chat.ChatId
	ch.ChatName = chat.ChatName
	ch.ChatPhoto = chat.ChatPhoto
}
func auth(authHeader string) (uint64, error) {
	re := regexp.MustCompile(`\d+`)
	stringToken := re.FindString(authHeader)
	authid, err := strconv.ParseUint(stringToken, 10, 64)
	if err != nil {
		return 0, err
	}
	return authid, nil
}
