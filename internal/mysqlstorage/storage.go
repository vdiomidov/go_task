package mysqlstorage

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Storage struct {
	db *sql.DB
}

func New(connStr string) (*Storage, error) {
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return nil, err
	}

	return &Storage{
		db: db,
	}, nil
}

func (s Storage) GetActiveSession(userId string, price string) (int, error) {
	var userId1 = "111"
	var sessionId int
	log.Printf(userId)

	err := s.db.
		QueryRow("SELECT id FROM sessions WHERE date_add(created_at, INTERVAL 20 MINUTE) < NOW() AND active = 1 AND user_id = ?;", userId1).
		Scan(&sessionId)
		//
	switch err {
	case sql.ErrNoRows:
	case nil:
		return sessionId, nil
	default:
		return 0, fmt.Errorf("query row: %w", err)
	}

	// update prev session
	//_, err := db.Exec(`UPDATE foo VALUES("bar", ?))`, someParam)

	res, err := s.db.Exec(`INSERT INTO sessions VALUES("bar", ?))`, userId, price)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("last insert id: %w", err)
	}

	return int(id), nil
}
