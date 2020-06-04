package application

import (
	"time"
)

type Session struct {
	ID        int       `db:"id"`
	UserId    string    `db:"user_id"`
	Active    bool      `db:"active"`
	CreatedAt time.Time `db:"created_at"`
	Price     int       `db:"price"`
}

type Adv struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

type Showing struct {
	SessionID int `db:"session_id"`
	AdvID     int `db:"adv_id"`
}

type Storage interface {
	GetActiveSession(userId string, price string) (int, error)
	GetAdvPrice(SessionId string, AdvIds []string) (float32, error)
}

type App struct {
	storage Storage
}

func NewApp(storage Storage) *App {
	return &App{
		storage: storage,
	}
}

func (a App) GetSession(userId string, price string) (int, error) {
	return a.storage.GetActiveSession(userId, price)
}

func (a App) GetAdvPrice(SessionId string, AdvIds []string) (float32, error) {
	return a.storage.GetAdvPrice(SessionId, AdvIds)
}
