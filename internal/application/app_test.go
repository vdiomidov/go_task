package application

import (
	"task/internal/application/mock"
	"testing"
)

func TestApp_GetSession(t *testing.T) {
	const expectedSessionId = 123

	s := mock.Storage{
		FnGetActiveSession: func(userId string, price int) (int, error) {
			return expectedSessionId, nil
		},
	}

	app := NewApp(s)

	res, err := app.GetSession("1", 500)
	if err != nil {
		t.Fatal(err)
	}
	if res != expectedSessionId {
		t.Fatalf("wrong result: %d", res)
	}
}
