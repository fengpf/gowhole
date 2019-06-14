package main

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/qiniu/log"
)

var (
	db  *sql.DB
	err error
	dsn = "root:fpf@tcp(127.0.0.1:3306)/test?parseTime=true"
)

func init() {
	if db, err = sql.Open("mysql", dsn); err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(256)
	db.SetMaxIdleConns(150)

	err = db.Ping()
	if err != nil {
		panic(err)
	}
}

func main() {
	place := make([]string, 0)
	args := make([]interface{}, 0)

	for i := 1; i <= 10000; i++ {
		pwd := _md5(fmt.Sprint(i, time.Now().Unix()))
		uname := _md5(fmt.Sprint(i, time.Now().UnixNano()))

		place = append(place, "(?, ?)")
		args = append(args, pwd, uname)
	}
	placeStr := strings.Join(place, ",")
	sqlStr := fmt.Sprintf("INSERT INTO user_01(username, password)  VALUES %s", placeStr)

	// result, err := db.Exec(sqlStr, args...)
	// if err != nil {
	// 	log.Errorf("db.Query sqlStr(%s) args(%+v) error(%v)", sqlStr, args, err)
	// 	return
	// }
	// rows, _ := result.RowsAffected()
	// log.Println("RowsAffected", rows)

	start := time.Now()
	rows, err := db.Query(sqlStr, args...)
	log.Println("cost", time.Now().Sub(start))

	if err != nil {
		log.Errorf("db.Query sqlStr(%s) args(%+v) error(%v)", sqlStr, args, err)
		return
	}
	rows.Close()

}

func _md5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
