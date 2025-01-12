package database
import (
	"database/sql"
	"fmt"
)


func (db *appdbimpl) Sendmessage(m Message) (Message, error) {

	res, err := db.c.Exec(`INSERT INTO messages (userId, content, messageDate, messageTime, state, chatId) 
							VALUES (?, ?, ?, ?, ?)`, m.UserId, m.Content, m.MessageDate, m.MessageTime, m.State, m.ChatId)
	if err != nil {
		return m, fmt.Errorf("error inserting message: %w", err)
	}

	
	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return m, fmt.Errorf("error fetching last insert ID: %w", err)
	}


	m.MessageId = uint64(lastInsertID)

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

	// Zapytanie SQL do pobrania wiadomości na podstawie messageId
	query := `SELECT messageId, userId, content, messageDate, messageTime, state, chatId FROM messages WHERE messageId = ?`
	err := db.c.QueryRow(query, messageId).Scan(&message.MessageId, &message.UserId, &message.Content, &message.MessageDate, &message.MessageTime, &message.State, &message.ChatId)
	if err != nil {
		if err == sql.ErrNoRows {
			return Message{}, fmt.Errorf("message with id %d not found", messageId) // Nie znaleziono wiadomości
		}
		return Message{}, fmt.Errorf("could not get message: %v", err)
	}

	return message, nil
}

func (db *appdbimpl) Commentmessage(c Comment) (Comment, error) {

	query := `INSERT INTO comments (messageId,  content) 
	VALUES (?, ?)`
	result, err := db.c.Exec(query, c.MessageId, c.Content)
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
