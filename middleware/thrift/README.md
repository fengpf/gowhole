# gohbase


Looking forward your pr or issue

## hbase
thrift file gen by hbase v1.2.0 thrift2 file

hbase thrift2 need to be started by cmd `hbase-daemon.sh start thrift2`

## demo
#### get

    package main

    import (
	    "fmt"

        h "github.com/ianwoolf/gohbase"
    )

    func main() {
    	hbObj := h.NewHbObj("192.168.99.100", 9090)
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

#### put

	if err := hbObj.Put("test", h.GenTPut("row1", "c2", "a", []byte("2-test"))); err != nil {
		fmt.Println(err.Error())
	}
