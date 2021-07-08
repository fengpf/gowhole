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
	dsn = "root:fpf123@tcp(127.0.0.1:3306)/test?parseTime=true"
)

func init() {
	if db, err = sql.Open("mysql", dsn); err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(3)

	err = db.Ping()
	if err != nil {
		panic(err)
	}
}

func main() {
	place := make([]string, 0)
	args := make([]interface{}, 0)

	for i := 1; i <= 1000; i++ {
		uname := fmt.Sprint(i)
		//addr := _md5(fmt.Sprint(i, time.Now().Unix()))
		addr := "addr_" + fmt.Sprint(i)
		place = append(place, "(?, ?)")
		args = append(args, uname, addr)
	}
	placeStr := strings.Join(place, ",")
	sqlStr := fmt.Sprintf("INSERT INTO t1(name, addr)  VALUES %s", placeStr)

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
		log.Errorf("db.Query sqlStr(%s) args(%+v)\n error(%v) \n", sqlStr, args, err)
		return
	}
	rows.Close()

}

func _md5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
