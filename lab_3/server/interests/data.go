package forums

import (
	"database/sql"
	"fmt"
)

type Forum struct {
	ForumId   int    `json:"ForumId"`
	ForumName string `json:"ForumName"`
}

type Store struct {
	Db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{Db: db}
}

func (s *Store) ListForums() ([]*Forum, error) {
	rows, err := s.Db.Query("SELECT ForumId, ForumName FROM forum")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []*Forum
	for rows.Next() {
		var c Forum
		if err := rows.Scan(&c.ForumId, &c.ForumName); err != nil {
			return nil, err
		}
		res = append(res, &c)
	}
	if res == nil {
		res = make([]*Forum, 0)
	}
	return res, nil
}

func (s *Store) CreateForum(name string) error {
	if len(name) < 0 {
		return fmt.Errorf("channel name is not provided")
	}
	_, err := s.Db.Exec("INSERT INTO forum (ForumName) VALUES ($1)", name)
	return err
}
