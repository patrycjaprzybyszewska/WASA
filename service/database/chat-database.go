package database
import (

	"fmt"
)

func (db *appdbimpl) AddUserToChat(chatId uint64, userId uint64) error {
    var chatExists bool
	var err error
    err = db.c.QueryRow("SELECT EXISTS(SELECT 1 FROM chats WHERE chatId = ?)", chatId).Scan(&chatExists)
    if err != nil {
        return fmt.Errorf("error checking chat existence: %w", err)
    }
    if !chatExists {
        return fmt.Errorf("chat with ID %d does not exist", chatId)
    }
  
	_, err = db.c.Exec("INSERT INTO chats (chatId, chatUsers) VALUES (?, ?)", chatId, userId)
    if err != nil {
        return fmt.Errorf("error adding user to chat: %w", err)
    }
	return nil
}
func (db *appdbimpl) LeaveGroup(chatId uint64, userId uint64) error {
    var chatExists bool
	var err error
    err = db.c.QueryRow("SELECT EXISTS(SELECT 1 FROM chats WHERE chatId = ?)", chatId).Scan(&chatExists)
    if err != nil {
        return fmt.Errorf("error checking chat existence: %w", err)
    }
    if !chatExists {
        return fmt.Errorf("chat with ID %d does not exist", chatId)
    }

	_, err = db.c.Exec(
        `DELETE FROM chats WHERE chatId = ? AND userId = ? `, chatId, userId)
        if err != nil {
            return fmt.Errorf("error removing user from chat: %w", err)
        }
  
	return nil
}
func (db *appdbimpl) SetGroupPhoto(ChatPhoto string, chatId uint64) error {
  
	var err error

	res, err := db.c.Exec(`UPDATE chats SET GroupPhoto=? WHERE ChatId=?`, ChatPhoto, chatId)
	if err != nil {
		return err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {
		return err
	}
	return nil
}
func (db *appdbimpl) SetGroupName(ChatName uint64, chatId uint64) error {
  
	var err error

	res, err := db.c.Exec(`UPDATE chats SET GroupName=? WHERE ChatId=?`, ChatName, chatId)
	if err != nil {
		return err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {
		return err
	}
	return nil
}


func (db *appdbimpl) GetChats() ([]Chat, error) {
    var chats []Chat

    query := `
        SELECT c.chatId, c.chatName, c.chatPhoto 
        FROM chats c
    `
    
    rows, err := db.c.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var chat Chat
        
        err := rows.Scan(&chat.ChatId, &chat.ChatName, &chat.ChatPhoto)
        if err != nil {
            return nil, err
        }
        
        chats = append(chats, chat)
    }
    
    if err := rows.Err(); err != nil {
        return nil, err
    }

    return chats, nil
}

