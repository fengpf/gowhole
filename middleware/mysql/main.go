package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var (
	//db
	gDB *sql.DB
	err error
	dsn = "root:fpf@tcp(127.0.0.1:3306)/test?parseTime=true"
)

func init() {
	if gDB, err = sql.Open("mysql", dsn); err != nil {
		panic(err)
	}
	// gDB.SetMaxOpenConns(256)
	// gDB.SetMaxIdleConns(150)

	err = gDB.Ping()
	if err != nil {
		panic(err)
	}
}

type Dao struct {
	db *sql.DB
}

// New init dao
func New(dsn string) (d *Dao) {
	d = &Dao{}
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(256)
	db.SetMaxIdleConns(150)

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	d.db = db
	return
}

// Ping db
func (d *Dao) Ping() (err error) {
	return d.db.Ping()
}

// Close db
func (d *Dao) Close() (err error) {
	return d.db.Close()
}

func main() {
	http.HandleFunc("/mysql_max_conns", mysqlMaxConns)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

type data struct {
	ID int `json:"id"`
	A  int `json:"a"`
	B  int `json:"b"`
	C  int `json:"c"`
}

func mysqlMaxConns(w http.ResponseWriter, r *http.Request) {
	rows, err := New(dsn).db.Query("SELECT * FROM test limit 100") //每次创建db对象

	// rows, err := gDB.Query("SELECT * FROM test limit 3")//使用全局db对象
	defer rows.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	res := make([]*data, 0)
	for rows.Next() {
		a := &data{}
		if err = rows.Scan(&a.ID, &a.A, &a.B, &a.C); err != nil {
			fmt.Printf("rows.Scan error(%v)", err)
			return
		}
		res = append(res, a)
	}
	str, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintln(w, string(str))
}
