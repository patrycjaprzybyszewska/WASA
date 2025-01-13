package database
import (
	"database/sql"
	"errors"
	"fmt"
)


func (db *appdbimpl) Sendmessage(m Message) (Message, error) {
////trzeba dodac tworzenie czatu tutaj
	var existingChatId uint64
	err := db.c.QueryRow(`SELECT chatId FROM chats WHERE chatId = ?`, m.ChatId).Scan(&existingChatId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
		
			res, err := db.c.Exec(`INSERT INTO chats (chatName, chatPhoto) VALUES (?, ?)`, "New Chat", "")
			if err != nil {
				return m, fmt.Errorf("error creating new chat: %w", err)
			}

			lastInsertID, err := res.LastInsertId()
			if err != nil {
				return m, fmt.Errorf("error fetching new chat ID: %w", err)
			}

			m.ChatId = uint64(lastInsertID)
		} 
	}
	
	res, err := db.c.Exec(`INSERT INTO messages (chatId, content, messageDate, messageTime, state) 
                        VALUES (?, ?, ?, ?, ?)`, m.ChatId, m.Content, m.MessageDate, m.MessageTime, m.State)


	
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
	query := `SELECT messageId, chatId, content, messageDate, messageTime, state, chatId FROM messages WHERE messageId = ?`
	err := db.c.QueryRow(query, messageId).Scan(&message.MessageId, &message.ChatId, &message.Content, &message.MessageDate, &message.MessageTime, &message.State, &message.ChatId)
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

func (db *appdbimpl) GetConversation(chatId uint64) ([]Message, error) {

	rows, err := db.c.Query(`SELECT messageId, content, messageDate, messageTime, state 
								FROM messages 
								WHERE chatId = ? 
								ORDER BY messageDate, messageTime`, chatId)
	if err != nil {
		return nil, fmt.Errorf("error fetching messages: %w", err)
	}
	defer rows.Close()

	
	var conversation []Message

	for rows.Next() {
		var message Message
		if err := rows.Scan(&message.MessageId, &message.ChatId, &message.Content, &message.MessageDate, &message.MessageTime, &message.State); err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		conversation = append(conversation, message)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}
	return conversation, nil
}

