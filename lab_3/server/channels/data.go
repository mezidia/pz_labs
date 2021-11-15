package channels

import (
	"database/sql"
	"fmt"
)

type Channel struct {
	ForumId   int64  `json:"ForumId"`
	ForumName string `json:"ForumName"`
}

type Store struct {
	Db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{Db: db}
}

func (s *Store) ListChannels() ([]*Channel, error) {
	rows, err := s.Db.Query("SELECT ForumId, ForumName FROM forum")
	if err != nil {
		return nil, err
	}
	fmt.Print(rows)
	defer rows.Close()

	var res []*Channel
	for rows.Next() {
		var c Channel
		if err := rows.Scan(&c.ForumId, &c.ForumName); err != nil {
			return nil, err
		}
		res = append(res, &c)
	}
	if res == nil {
		res = make([]*Channel, 0)
	}
	return res, nil
}

func (s *Store) CreateChannel(name string) error {
	if len(name) < 0 {
		return fmt.Errorf("channel name is not provided")
	}
	_, err := s.Db.Exec("INSERT INTO channels (name) VALUES ($1)", name)
	return err
}
