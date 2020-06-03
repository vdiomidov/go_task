package mock

type Storage struct {
	FnGetActiveSession func(userId string, price int) (int, error)
}

func (s Storage) GetActiveSession(userId string, price int) (int, error) {
	return s.FnGetActiveSession(userId, price)
}
