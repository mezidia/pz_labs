package Users

import (
	"database/sql"
	"fmt"
)

type User struct {
	UserId   int    `json:"UserId"`
	UserName string `json:"UserName"`
}

type Store struct {
	Db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{Db: db}
}

func (s *Store) ListUsers() ([]*User, error) {
	rows, err := s.Db.Query("SELECT UserId, UserName FROM User")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []*User
	for rows.Next() {
		var c User
		if err := rows.Scan(&c.UserId, &c.UserName); err != nil {
			return nil, err
		}
		res = append(res, &c)
	}
	if res == nil {
		res = make([]*User, 0)
	}
	return res, nil
}

func (s *Store) CreateUser(name string) error {
	if len(name) < 0 {
		return fmt.Errorf("channel name is not provided")
	}
	_, err := s.Db.Exec("INSERT INTO User (UserName) VALUES ($1)", name)
	return err
}
