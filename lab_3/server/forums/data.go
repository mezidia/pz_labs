package forums

import (
	"database/sql"
	"fmt"
	"sort"
)

type Forum struct {
	ForumId      int    `json:"ForumId"`
	ForumName    string `json:"ForumName"`
	InterestName string `json:"InterestName"`
	UserName     string `json:"UserName"`
}

type ForumOut struct {
	ForumName    string   `json:"ForumName"`
	InterestName string   `json:"InterestName"`
	Users        []string `json:"UserName"`
}

type Store struct {
	Db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{Db: db}
}

func (s *Store) ListForums() ([]ForumOut, error) {
	rows, err := s.Db.Query("SELECT DISTINCT  Forum.ForumId, Forum.ForumName, Theme.ThemeName, [User].UserName FROM  Forum INNER JOIN	ForumUser ON Forum.ForumID = ForumUser.ForumID INNER JOIN [User] ON ForumUser.UserID = [User].UserID INNER JOIN ThemeForum ON Forum.ForumID = ThemeForum.ForumID INNER JOIN Theme ON ThemeForum.ThemeID = Theme.ThemeID")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var res []*Forum
	for rows.Next() {
		var f Forum
		if err := rows.Scan(&f.ForumId, &f.ForumName, &f.InterestName, &f.UserName); err != nil {
			return nil, err
		}
		res = append(res, &f)
	}
	if res == nil {
		res = make([]*Forum, 0)
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i].ForumId < res[j].ForumId
	})
	var output []ForumOut
	var fo ForumOut
	for i, _ := range res {
		if i != 0 {
			if res[i].ForumId == res[i-1].ForumId {
				fo.Users = append(fo.Users, res[i].UserName)
			} else {
				output = append(output, fo)
				fo = ForumOut{}
				fo.ForumName = res[i].ForumName
				fo.InterestName = res[i].InterestName
				fo.Users = append(fo.Users, res[i].UserName)
			}
		} else {
			fo.ForumName = res[i].ForumName
			fo.InterestName = res[i].InterestName
			fo.Users = append(fo.Users, res[i].UserName)
		}
	}
	output = append(output, fo)
	return output, nil
}

func (s *Store) CreateForum(name string) error {
	if len(name) < 1 {
		return fmt.Errorf("channel name is not provided")
	}
	_, err := s.Db.Exec("INSERT INTO forum (ForumName) VALUES ('" + name + "')")
	return err
}
