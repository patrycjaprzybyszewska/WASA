package database
import (
	"database/sql"
)


func (db *appdbimpl) Sendmessage(m Message) (Message, error) {

	res, err := db.c.Exec(`INSERT INTO messages (userId, content, messageDate, messageTime, state) 
							VALUES (?, ?, ?, ?, ?)`, m.UserId, m.Content, m.MessageDate, m.MessageTime, m.State)
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

