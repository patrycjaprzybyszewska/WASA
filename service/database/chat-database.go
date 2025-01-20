package database

import (
	"errors"
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
		return ch, err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return ch, err
	} else if affected == 0 {
		return ch, err
	}
	return ch, nil
}
func (db *appdbimpl) SetGroupName(ch Chat, chatName string) (Chat, error) {

	var err error

	res, err := db.c.Exec(`UPDATE chats SET ChatName=?, ChatPhoto=? WHERE ChatId=?`, chatName, ch.ChatPhoto, ch.ChatId)
	if err != nil {
		return ch, err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return ch, err
	} else if affected == 0 {
		return ch, err
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
	var chatPhoto string
	err := db.c.QueryRow(`SELECT ChatPhoto FROM chats WHERE ChatId = ?`, chatId).Scan(&chatPhoto)
	if err != nil {
		return "", errors.New("No chat found")
	}
	return chatPhoto, nil
}
func (db *appdbimpl) GetChatNameById(chatId uint64) (string, error) {
	var chatName string
	err := db.c.QueryRow(`SELECT ChatName FROM chats WHERE ChatId = ?`, chatId).Scan(&chatName)
	if err != nil {
		return "", errors.New("No chat found")
	}
	return chatName, nil
}
func (db *appdbimpl) GetChatIdbyName(chatName string) (uint64, error) {
	var chatId uint64
	var userId uint64
	err := db.c.QueryRow(`SELECT ChatId FROM chats WHERE ChatName = ?`, chatName).Scan(&chatId)
	if err != nil {
		err = db.c.QueryRow(`SELECT UserId FROM users WHERE Username = ?`, chatName).Scan(&userId)
		if err == nil{ 
		res, err := db.c.Exec(`INSERT INTO chats (ChatName, ChatPhoto) VALUES (?, ?)`, chatName, "")
		if err != nil {
			return 0, fmt.Errorf("failed to create chat: %v", err)
		}
		
		lastInsertId, err := res.LastInsertId()
		if err != nil {
			return 0, fmt.Errorf("failed to create chat id: %v", err)
		}
		chatId = uint64(lastInsertId)
		err = db.c.AddUserToChat(chatId, userId)
			if err != nil {
				return 0, fmt.Errorf("failed to create chat id: %v", err)
			}
		}
		else{
			fmt.Println("No chat found, creating new one")
			res, err := db.c.Exec(`INSERT INTO chats (ChatName, ChatPhoto) VALUES (?, ?)`, chatName, "")
			lastInsertId, err := res.LastInsertId()
			if err != nil {
				return 0, fmt.Errorf("failed to retrieve last insert id: %v", err)
			}
			chatId = uint64(lastInsertId)

		}
	}
	return chatId, nil
}
