package mapdb

import (
	"database/sql"
	"fmt"
	_ "github.com/ziutek/mymysql/godrv"
)

const (
	DB_NAME = "rpgmap"
	DB_USER = "root"
	DB_PASS = "root"
	// DB_ADDR = "localhost"
)

var (
	Con *sql.DB
)

func GetDbConn() *sql.DB {

	db, err := sql.Open("mymysql", fmt.Sprintf("%s/%s/%s", DB_NAME, DB_USER, DB_PASS))
	if err != nil {
		fmt.Printf("Error opening DB Connectin %s", err)
	}
	Con = db
	return db
}

func AddUser(username, password string) error {
	fmt.Printf("Adding %s %s", username, password)
	_, err := Con.Exec("insert into users (user,password) values (?,?)", username, password)
	if err != nil {
		fmt.Printf("Error adding user from AddUser: %s", err)
	}
	return err
}

func AuthorizeUser(username, password string, con *sql.DB) bool {
	user := con.QueryRow("select * from user where user=?", username)

	if user != nil {
		return true
	}
	return false

}
