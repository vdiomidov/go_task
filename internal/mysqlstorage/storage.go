package mysqlstorage

import (
	"database/sql"
	"fmt"

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
	var sessionId int

	err := s.db.
		QueryRow("SELECT id FROM sessions WHERE date_add(created_at, INTERVAL 30 MINUTE) > NOW() AND active = 1 AND user_id = ?;", userId).
		Scan(&sessionId)
	switch err {
	case sql.ErrNoRows:
	case nil:
		return sessionId, nil
	default:
		return 0, fmt.Errorf("query row: %w", err)
	}

	_, err = s.db.Exec("UPDATE sessions SET active=0 WHERE user_id = ?;", userId)

	res, err := s.db.Exec("INSERT INTO sessions (user_id, price) VALUES(?,?)", userId, price)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("last insert id: %w", err)
	}

	return int(id), nil
}

func (s Storage) GetAdvPrice(SessionId string, AdvIds []string) (float32, error) {
	var count int

	err := s.db.
		QueryRow("SELECT count(id) FROM sessions WHERE date_add(created_at, INTERVAL 30 MINUTE) > NOW() AND id = ?;", SessionId).
		Scan(&count)
	switch err {
	case nil:
	default:
		return 0, fmt.Errorf("query row: %w", err)
	}
	if count == 0 {
		return 0, fmt.Errorf("wrong session: %s", SessionId)
	}

	for _, advId := range AdvIds {
		_, err := s.db.Exec("INSERT INTO showing (session_id, adv_id) VALUES(?,?)", SessionId, advId)
		if err != nil {
			return 0, err
		}
	}

	var price float32

	err = s.db.
		QueryRow("SELECT (s.price/COUNT(session_id)) cnt FROM showing sh INNER JOIN sessions s ON sh.session_id = s.id WHERE session_id=? LIMIT 1;", SessionId).
		Scan(&price)
	switch err {
	case sql.ErrNoRows:
		return 0, fmt.Errorf("wrong session: %s", SessionId)
	case nil:
	default:
		return 0, fmt.Errorf("query row: %w", err)
	}

	return price, nil
}
