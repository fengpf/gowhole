package main

import (
	"database/sql"

	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(192.168.0.105:3306)/wayne?charset=utf8")
	if err != nil {
		panic(err)
	}
	stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,created=?")
	if err != nil {
		fmt.Println(11, err)
		return
	}
	res, err := stmt.Exec("Test", "people", "2017-10-27")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}
