package test

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"testing"

	"github.com/lightstep/lightstep-tracer-go/thrift_0_9_2/lib/go/thrift"
)

func TestGetMulti(t *testing.T) {
	startTime := currentTimeMillis()
	logformatstr := "----%s\n"
	logformatstr := "----%s 用时:%d-%d=%d毫秒\n\n"
	logformattitle := "建立连接"
	rowkey := "1"
	temptable := "test_idoall_org"
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	transport, err := thrift.NewTSocket(net.JoinHostPort(HOST, PORT))
	if err != nil {
		fmt.Fprintln(os.Stderr, "error resolving address:", err)
		os.Exit(1)
	}
	client := hbase2go.NewTHBaseServiceClientFactory(transport, protocolFactory)
	if err := transport.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to "+HOST+":"+PORT, " ", err)
		os.Exit(1)
	}
	tmpendTime := currentTimeMillis()
	fmt.Printf(logformatstr, logformattitle, tmpendTime, startTime, (tmpendTime - startTime))
	defer transport.Close()
	//------------------GetMultiple-----------------------------
	logformattitle = "调用GetMultiple方法获取" + strconv.Itoa(TESTRECORD) + "数据"
	fmt.Printf(logformatstr_, logformattitle)
	tmpstartTime = currentTimeMillis()
	//
	var tgets []*hbase.TGet
	for i := 0; i < TESTRECORD; i++ {
		putrowkey := strconv.Itoa(i)
		tgets = append(tgets, &hbase.TGet{
			Row: []byte(putrowkey)})
	}
	results, err := client.GetMultiple([]byte(temptable), tgets)
	if err != nil {
		fmt.Printf("GetMultiple err:%s", err)
	} else {
		fmt.Printf("GetMultiple Count:%d\n", len(results))
		for _, k := range results {
			fmt.Println("Rowkey:" + string(k.Row))
			for _, cv := range k.ColumnValues {
				printscruct(cv)
			}
		}
	}
	tmpendTime = currentTimeMillis()
	fmt.Printf(logformatstr, logformattitle, tmpendTime, tmpstartTime, (tmpendTime - tmpstartTime))
}

func Test_hbObj(t *testing.T) {
	hbObj := hbase2go.NewHbObj("192.168.99.100", 9090)
	if err := hbObj.Connect(); err != nil {
		fmt.Println(err.Error())
	}
	defer hbObj.Close()
	// get 'test','row1',{COLUMN => ['c2:a']}
	TRow, _ := hbObj.GetRow("test", h.GenTGet("row1", "col", "a"))
	for _, col := range TRow.ColumnValues {
		fmt.Println(string(col.GetFamily()), string(col.GetQualifier()), string(col.GetValue()), col.GetTags(), col.GetTimestamp())
	}
}
