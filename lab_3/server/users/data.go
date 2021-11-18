package users

import (
	"database/sql"
	"fmt"
)

type User struct {
	UserId   int    `json:"UserId"`
	UserName string `json:"UserName"`
	UserType int    `json:"UserType"`
	UserMail string `json:"UserMail"`
	Password string `json:"Password"`
}

type Store struct {
	Db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{Db: db}
}

func (s *Store) ListUsers() ([]*User, error) {
	rows, err := s.Db.Query("SELECT UserId, [UserName], [UserType], [UserMail] FROM [User]")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []*User
	for rows.Next() {
		var c User
		if err := rows.Scan(&c.UserId, &c.UserName, &c.UserType, &c.UserMail); err != nil {
			return nil, err
		}
		res = append(res, &c)
	}
	if res == nil {
		res = make([]*User, 0)
	}
	return res, nil
}

func (s *Store) CreateUser(name string, utype int, mail string, password string) error {
	if len(name) < 0 {
		return fmt.Errorf("channel name is not provided")
	}
	_, err := s.Db.Exec("INSERT INTO [User] (UserName, UserType, UserMail, Password, IsDeleted) VALUES ($1, $2, $3, $4, 0)", name, utype, mail, password)
	return err
}
