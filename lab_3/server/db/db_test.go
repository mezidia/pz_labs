package db

import (
	"fmt"
	"testing"
)

func TestDbConnection_ConnectionURL(t *testing.T) {
	conn := &Connection{
		DbName:     "ForumLab",
		User:       "sichkartetiana_SQLLogin_1",
		Password:   "y91hx58s5h",
		Host:       "ForumLab.mssql.somee.com",
		DisableSSL: true,
	}
	fmt.Println(conn.ConnectionURL())
	if conn.ConnectionURL() != "sqlserver://sichkartetiana_SQLLogin_1:y91hx58s5h@ForumLab.mssql.somee.com/ForumLab?sslmode=disable" {
		t.Error("Unexpected connection string")
	}
}
