package interests

import (
	"database/sql"
	"fmt"
)

type Interest struct {
	InterestId   int    `json:"InterestId"`
	InterestName string `json:"InterestName"`
}

type Store struct {
	Db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{Db: db}
}

func (s *Store) ListInterests() ([]*Interest, error) {
	rows, err := s.Db.Query("SELECT ThemeId, ThemeName FROM [Theme]")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []*Interest
	for rows.Next() {
		var c Interest
		if err := rows.Scan(&c.InterestId, &c.InterestName); err != nil {
			return nil, err
		}
		res = append(res, &c)
	}
	if res == nil {
		res = make([]*Interest, 0)
	}
	return res, nil
}

func (s *Store) CreateInterest(name string) error {
	if len(name) < 0 {
		return fmt.Errorf("interest name is not provided")
	}
	fmt.Println(name)
	_, err := s.Db.Exec("INSERT INTO [Theme] (ThemeName) VALUES ('" + name + "')")
	return err
}
