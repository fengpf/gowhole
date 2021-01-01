package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB
var dataBase = "test:test@tcp(172.16.33.205:3308)/test?timeout=1s&readTimeout=6s"

// iptables -A OUTPUT -p tcp --dport 3308 -d 172.16.33.205 -j DROP

// iptables -A OUTPUT -p tcp --dport 3308 -d 172.16.33.205 -j REJECT
// [mysql] 2021/01/01 12:17:41 connector.go:48: net.Error from Dial()': dial tcp 172.16.33.205:3308: i/o timeout
// [mysql] 2021/01/01 12:17:42 connector.go:48: net.Error from Dial()': dial tcp 172.16.33.205:3308: i/o timeout
// [mysql] 2021/01/01 12:17:43 connector.go:48: net.Error from Dial()': dial tcp 172.16.33.205:3308: i/o timeout
// 2021/01/01 12:17:43 query failed: driver: bad connection

func mysqlInit() {
	var err error
	DB, err = sql.Open("mysql", dataBase)
	if err != nil {
		log.Fatalln("open db fail:", err)
	}

	DB.SetMaxOpenConns(3)
	DB.SetMaxIdleConns(3)
}

func main() {
	mysqlInit()

	for {
		log.Println("start")
		execSql()
		time.Sleep(3 * time.Second)
	}
}

func execSql() {
	var value int
	err := DB.QueryRow("select 1").Scan(&value)
	if err != nil {
		log.Println("query failed:", err)
		return
	}

	log.Println("value:", value)
}
