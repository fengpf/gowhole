package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"testing"
	"time"

	"github.com/tsuna/gohbase/filter"
	"github.com/tsuna/gohbase/hrpc"

	"github.com/tsuna/gohbase"
)

// Name of the meta region.
const (
	host          = "172.18.33.75"
	metaTableName = "hbase:meta"
)

// Info family
var infoFamily = map[string][]string{
	"info": nil,
}

var cFamilies = map[string]map[string]string{
	"cf":  nil,
	"cf2": nil,
}

func getTimestampString() string {
	return time.Now().Format("20060102150405")
}

func TestCreateTable(t *testing.T) {
	testTableName := "test1_" + getTimestampString()
	t.Log("testTableName=" + testTableName)

	ac := gohbase.NewAdminClient(host)
	crt := hrpc.NewCreateTable(context.Background(), []byte(testTableName), cFamilies)

	if err := ac.CreateTable(crt); err != nil {
		t.Errorf("CreateTable returned an error: %v", err)
	}

}

func TestCheck(t *testing.T) {
	// check in hbase:meta if there's a region for the table
	c := gohbase.NewClient(host)
	metaKey := "test1_20180419190721"
	keyFilter := filter.NewPrefixFilter([]byte(metaKey))
	scan, err := hrpc.NewScanStr(context.Background(), metaTableName, hrpc.Filters(keyFilter))
	if err != nil {
		t.Fatalf("Failed to create Scan request: %s", err)
	}

	var rsp []*hrpc.Result
	scanner := c.Scan(scan)
	for {
		res, err := scanner.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatal(err)
		}
		rsp = append(rsp, res)
	}
	if len(rsp) != 1 {
		t.Errorf("Meta returned %d rows for prefix '%s' , want 1", len(rsp), metaKey)
	}
}

func Test_put(t *testing.T) {
	tableName := "test1_20180419190721"
	c := gohbase.NewClient(host)
	//Insert a cell Values maps a ColumnFamily -> Qualifiers -> Values.
	values := map[string]map[string][]byte{"cf": map[string][]byte{"a": []byte("1234")}}

	// x := 1234
	// buf := bytes.NewBuffer([]byte{})
	// binary.Write(buf, binary.BigEndian, &x)
	// fmt.Printf("after write, buf is (%+v)", buf.Bytes())
	// values["cf"]["a"] = buf.Bytes()

	putRequest, err := hrpc.NewPutStr(context.Background(), tableName, "key", values)
	if err != nil {
		fmt.Printf("hrpc.NewPutStr error(%v)", err)
		return
	}
	rsp, err := c.Put(putRequest)
	if err != nil {
		fmt.Printf("c.Put error(%v)", err)
		return
	}
	fmt.Println(rsp)
}
func Test_get(t *testing.T) {
	s := []byte("中华人民共和国")
	r := bytes.Runes(s)
	fmt.Println(string(s), len(s)) //字节切片的长度
	fmt.Println(string(r), len(r)) // rune 切片的长度
	c := gohbase.NewClient(host)
	tableName := "test1_20180419190721"
	family := map[string][]string{"cf": []string{"a"}}
	getRequest, err := hrpc.NewGetStr(context.Background(), tableName, "key", hrpc.Families(family))
	if err != nil {
		fmt.Printf("hrpc.NewGetStr error(%v)", err)
		return
	}
	getRsp, err := c.Get(getRequest)
	if err != nil {
		fmt.Printf("c.Get error(%v)", err)
		return
	}
	for _, c := range getRsp.Cells {
		if c == nil {
			continue
		}
		if string(c.Family) == "cf" {
			println(len(c.Value))
			// v := binary.BigEndian.Uint32(c.Value)
			// println(v)
			fmt.Println(c.Value, string(c.Value))
			println(string(c.Qualifier[:]))
		}
	}
}

func Test_Scan(t *testing.T) {
	c := gohbase.NewClient(host)
	tableName := "test1_20180419190721"
	ctx := context.Background()
	start := "0"
	stop := "100"
	scanRequest, err := hrpc.NewScanRangeStr(ctx, tableName, start, stop)
	if err != nil {
		fmt.Printf("hrpc.NewScanRangeStr error(%v)", err)
		return
	}
	scanRsp := c.Scan(scanRequest)
	fmt.Println(scanRsp)
}
