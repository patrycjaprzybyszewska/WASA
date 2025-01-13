package database
import (
	"database/sql"
	"fmt"
)

func (db *appdbimpl) AddUserToChat(chatId int, userId uint64) error {
    var chatExists bool
    err = db.c.QueryRow("SELECT EXISTS(SELECT 1 FROM chats WHERE chatId = ?)", chatId).Scan(&chatExists)
    if err != nil {
        return fmt.Errorf("error checking chat existence: %w", err)
    }
    if !chatExists {
        return fmt.Errorf("chat with ID %d does not exist", chatId)
    }

	_, err = db.c.Exec(
        "INSERT INTO chats (chatId, userId) VALUES (?, ?)",
        chatId, userId,
    )
    if err != nil {
        return fmt.Errorf("error adding user to chat: %w", err)
    }
}
