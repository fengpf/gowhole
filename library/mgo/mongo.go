package main

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/akkuman/parseConfig"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type SecuCode struct {
	InnerCode    string `bson:"InnerCode"`
	CompanyCode  string `bson:"CompanyCode"`
	SecuCode     string `bson:"SecuCode"`
	SecuAbbr     string `bson:"SecuAbbr"`
	SecuMarket   string `bson:"SecuMarket"`
	SecuCategory string `bson:"SecuCategory"`
	ChiName      string `bson:"ChiName"`
	ChiNameAbbr  string `bson:"ChiNameAbbr"`
	EngName      string `bson:"EngName"`
	EngNameAbbr  string `bson:"EngNameAbbr"`
	ChiSpelling  string `bson:"ChiSpelling"`
	ListedSector string `bson:"ListedSector"`
	ListedState  string `bson:"ListedState"`
}

type Mongo struct {
	url string
}

func (m Mongo) getFilePath(str string) string {
	config := parseConfig.New("./filepath.json")
	path := config.Get(str)
	return path.(string)
}

func (u Mongo) GetConfig(path, host string) string {
	config := parseConfig.New(path)
	get := config.Get(host)

	return get.(string)
}

func MongoCon(coll, table string, ms interface{}) (secucode []SecuCode) {
	r := Mongo{}
	filePath := r.getFilePath("mongopath")
	session, err := mgo.Dial(r.GetConfig(filePath, "host"))
	if nil != err {
		panic(err)
	}
	defer session.Close()
	collection := session.DB(coll).C(table)
	collection.Find(ms).All(&secucode)
	fmt.Println(secucode)
	return secucode
}

func main() {

	ms := bson.M{"SecuCategory": 1, "ListedDate": bson.M{"$ne": nil}, "ListedState": bson.M{"$in": []int{1, 3, 9}}, "SecuMarket": bson.M{"$in": []int{83, 90}}}
	MongoCon("sqlData", "SecuMain", ms)

	//start := time.Now()
	//ch := make(chan int, 10)
	//for _, v := range con {
	//	fmt.Println(v.SecuCode, v.ChiName)
	//	//ch <- 1
	//
	//	//go Fetch(index, v.SecuCode, ch)
	//	//continue
	//}
	////<-ch
	//defer close(ch)
	//fmt.Println(time.Since(start))
}

func Fetch(index int, secucode string, ch chan int) {
	var buf bytes.Buffer
	url := "http://goldeye.cfbond.com/cattle/complex_score?secu_code="
	buf.WriteString(url)
	buf.WriteString(secucode)
	resp, err := http.Get(buf.String())
	if err != nil {
		fmt.Println(err)
	}
	ints := <-ch

	fmt.Println("拼接后的结果为-->", buf.String(), index, resp.StatusCode, ints)
}
