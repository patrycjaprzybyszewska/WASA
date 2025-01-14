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
  
	_, err = db.c.Exec("INSERT INTO chat_users (chatId, userId) VALUES (?, ?)", chatId, userId)
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
        `DELETE FROM chat_users WHERE chatId = ? AND userId = ? `, chatId, userId)
        if err != nil {
            return fmt.Errorf("error removing user from chat: %w", err)
        }
  
	return nil
}

func (db *appdbimpl) SetGroupPhoto(ch Chat, chatPhoto string) (Chat, error) {
  
	var err error

	res, err := db.c.Exec(`UPDATE chats SET ChatPhoto=?, ChatName=? WHERE ChatId=?`, chatPhoto, ch.ChatName, ch.ChatId)
	if err != nil {
		return err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {
		return err
	}
	return ch, nil
}
func (db *appdbimpl) SetGroupName(ch Chat, chatName string) (Chat, error) {
  
	var err error

	res, err := db.c.Exec(`UPDATE chats SET ChatName=?, ChatPhoto=? WHERE ChatId=?`, chatName, ch.ChatPhoto, ch.ChatId)
	if err != nil {
		return err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {
		return err
	}
	return ch, nil
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

func (db *appdbimpl) GetChatPhotoById(chatId uint64) (string, error) {
    var userPhoto string
    err := db.c.QueryRow(`SELECT ChatPhoto FROM chats WHERE ChatId = ?`, chatId).Scan(&chatPhoto)
    if err != nil {
        return "", err
    }
    return userPhoto, nil
}
func (db *appdbimpl) GetChatNameById(userId uint64) (string, error) {
    var userName string
    err := db.c.QueryRow(`SELECT ChatName FROM chats WHERE ChatId = ?`, chatId).Scan(&chatName)
    if err != nil {
        return "", err
    }
    return userName, nil
}