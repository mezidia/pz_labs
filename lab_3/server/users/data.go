package users

import (
	"database/sql"
	"fmt"
)

type User struct {
	UserId    int      `json:"UserId"`
	UserName  string   `json:"UserName"`
	UserType  int      `json:"UserType"`
	UserMail  string   `json:"UserMail"`
	Password  string   `json:"Password"`
	Interests []string `json:"Interests"`
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

func (s *Store) CreateUser(name string, utype int, mail string, password string, interests []string) error {
	if len(name) < 0 {
		return fmt.Errorf("channel name is not provided")
	}
	fmt.Println(interests)
	var xmlString = "<ITEMS>"
	for _, s := range interests {
		xmlString = xmlString + "<ITEM Interest=\"" + s + "\" />"
	}
	xmlString = xmlString + "</ITEMS>"
	_, err := s.Db.Query("EXEC [InsertUser] @UserType=0, @UserName='" + name + "', @UserMail='" + mail + "', @Password='" + password + "', @Interests='" + xmlString + "'")
	return err
}
