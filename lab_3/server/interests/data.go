package interests

import (
	"database/sql"
	"fmt"
	"strconv"
)

type Interest struct {
	InterestId   int    `json:"InterestId"`
	InterestName string `json:"InterestName"`
	ForumName    string `json:"ForumName"`
	ForumId      int    `json:"ForumId"`
}

type Store struct {
	Db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{Db: db}
}

func (s *Store) ListInterests() ([]*Interest, error) {
	rows, err := s.Db.Query("SELECT Theme.ThemeID, Theme.ThemeName, Forum.ForumName, ThemeForum.ForumID  FROM Forum INNER JOIN ThemeForum ON Forum.ForumID = ThemeForum.ForumID INNER JOIN Theme ON ThemeForum.ThemeID = Theme.ThemeID")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []*Interest
	for rows.Next() {
		var c Interest
		if err := rows.Scan(&c.InterestId, &c.InterestName, &c.ForumName, &c.ForumId); err != nil {
			return nil, err
		}
		res = append(res, &c)
	}
	if res == nil {
		res = make([]*Interest, 0)
	}
	return res, nil
}

func (s *Store) CreateInterest(name string, forumId int) error {
	if len(name) < 1 || forumId < 1 {
		return fmt.Errorf("interest name or forum ID is not provided")
	}
	fmt.Println(name)
	_, err := s.Db.Query("EXEC [InsertInterest] @InterestName='" + name + "', @ForumId=" + strconv.Itoa(forumId))
	return err
}
