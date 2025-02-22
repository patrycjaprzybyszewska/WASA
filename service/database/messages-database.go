package database

import (
	"database/sql"
	"fmt"
)

func (db *appdbimpl) Sendmessage(m Message) (Message, error) {
	// trzeba dodac tworzenie czatu tutaj

	res, err := db.c.Exec(`INSERT INTO messages (senderName, senderId, chatId, content, messageDate, messageTime, state) VALUES (?, ?, ?, ?, ?, ?, ?)`, m.SenderName, m.SenderId, m.ChatId, m.Content, m.MessageDate, m.MessageTime, m.State)
	if err != nil {
		return m, fmt.Errorf("error inserting chat: %w", err)
	}

	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return m, fmt.Errorf("error fetching last insert ID: %w", err)
	}

	m.MessageId = uint64(lastInsertID)
	var count int
	err = db.c.QueryRow("SELECT COUNT(*) FROM chat_users WHERE chatId = ? AND userId = ?", m.ChatId, m.SenderId).Scan(&count)
	if err != nil {
		return m, fmt.Errorf("error checking user participation in chat: %w", err)
	}

	if count == 0 {
		_, err = db.c.Exec("INSERT INTO chat_users (chatId, userId, read) VALUES (?, ?, ?)", m.ChatId, m.SenderId, 0)
		if err != nil {
			return m, fmt.Errorf("error adding user to chat: %w", err)
		}
	}

	return m, nil
}

func (db *appdbimpl) Removemessage(messageId uint64) error {
	_, err := db.c.Exec(`DELETE FROM messages WHERE messageId=?`, messageId)
	if err != nil {
		return err
	}
	return nil

}

func (db *appdbimpl) GetMessageById(messageId uint64) (Message, error) {
	var message Message

	query := `SELECT messageId, senderName, senderId, chatId, content, messageDate, messageTime, state FROM messages WHERE messageId = ?`
	err := db.c.QueryRow(query, messageId).Scan(&message.MessageId, &message.SenderName, &message.SenderId, &message.ChatId, &message.Content, &message.MessageDate, &message.MessageTime, &message.State)
	if err != nil {
		if err == sql.ErrNoRows {
			return Message{}, fmt.Errorf("message with id %d not found", messageId) 
		}
		return Message{}, fmt.Errorf("could not get message: %w", err)
	}

	return message, nil
}
func (db *appdbimpl) CheckMessageById(messageId uint64) error {
	var message Message

	query := `SELECT messageId, senderName, senderId, chatId, content, messageDate, messageTime, state FROM messages WHERE messageId = ?`
	err := db.c.QueryRow(query, messageId).Scan(&message.MessageId, &message.SenderName, &message.SenderId, &message.ChatId, &message.Content, &message.MessageDate, &message.MessageTime, &message.State)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("message with id %d not found", messageId) 
		}
		return fmt.Errorf("could not get message: %w", err)
	}

	return nil
}

func (db *appdbimpl) Commentmessage(c Comment) (Comment, error) {

	query := `INSERT INTO comments (messageId,  content, commentername) VALUES (?, ?, ?)`
	result, err := db.c.Exec(query, c.MessageId, c.Content, c.Commentername)
	if err != nil {
		return c, err
	}

	commentId, err := result.LastInsertId()
	if err != nil {
		return c, err
	}

	c.CommentId = uint64(commentId)
	return c, nil
}

func (db *appdbimpl) Removecomment(commentId uint64) error {
	_, err := db.c.Exec(`DELETE FROM comments WHERE commentId=?`, commentId)
	if err != nil {
		return err
	}
	return nil

}

func (db *appdbimpl) GetConversation(chatId uint64, userId uint64) ([]MessageandComments, error) {
	var read int = 1
	_, err := db.c.Exec(`UPDATE chat_users SET read = ? WHERE chatId = ? AND userId = ?`, read, chatId, userId)
	if err != nil {
		return nil, fmt.Errorf("error %w", err)
	}
	var count int
	err = db.c.QueryRow(`SELECT COUNT(*) FROM chat_users WHERE chatId = ? AND read = 0`, chatId).Scan(&count)
	if err != nil {
		return nil, fmt.Errorf("error %w", err)
	}
	if count == 0 {
		_, err = db.c.Exec(`UPDATE messages SET state = 'read' WHERE chatId = ?`, chatId)
		if err != nil {
			return nil, fmt.Errorf("error %w", err)
		}
	}

	rows, err := db.c.Query(`SELECT messageId, senderName, senderId, content, messageDate, messageTime, state, chatId 
								FROM messages 
								WHERE chatId = ? 
								ORDER BY messageDate, messageTime`, chatId)
	if err != nil {
		return nil, fmt.Errorf("error fetching messages: %w", err)
	}
	defer rows.Close()

	var conversation []MessageandComments

	for rows.Next() {
		var message Message
		if err := rows.Scan(&message.MessageId, &message.SenderName, &message.SenderId, &message.Content, &message.MessageDate, &message.MessageTime, &message.State, &message.ChatId); err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		comments, err := db.GetCommentsById(message.MessageId)
		if err != nil {
			return nil, fmt.Errorf("error for message %d: %w", message.MessageId, err)
		}
		messageandComments := MessageandComments{
			Message:  message,
			Comments: comments,
		}
		conversation = append(conversation, messageandComments)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return conversation, nil
}

func (db *appdbimpl) GetCommentById(commentId uint64) (Comment, error) {
	var comment Comment
	query := `SELECT commentId, messageId, content, commentername FROM comments WHERE commentId = ?`
	err := db.c.QueryRow(query, commentId).Scan(&comment.CommentId, &comment.MessageId, &comment.Content, &comment.Commentername)
	if err != nil {
		if err == sql.ErrNoRows {
			return Comment{}, fmt.Errorf("comment with id %d not found", commentId) 
		}
		return Comment{}, fmt.Errorf("could not get comment: %w", err)
	}

	return comment, nil
}
func (db *appdbimpl) CheckCommentById(commentId uint64) error {
	var comment Comment
	query := `SELECT commentId, messageId, content FROM comments WHERE commentId = ?`
	err := db.c.QueryRow(query, commentId).Scan(&comment.CommentId, &comment.MessageId, &comment.Content)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("comment with id %d not found", commentId) 
		}
		return fmt.Errorf("could not get comment: %w", err)
	}

	return nil
}

func (db *appdbimpl) GetCommentsById(messageId uint64) ([]Comment, error) {
	var comments []Comment

	query := `SELECT commentId, messageId, content, commentername FROM comments WHERE messageId = ?`

	rows, err := db.c.Query(query, messageId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var comment Comment

		err := rows.Scan(&comment.CommentId, &comment.MessageId, &comment.Content, &comment.Commentername)
		if err != nil {
			return nil, err
		}

		comments = append(comments, comment)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return comments, nil
}
