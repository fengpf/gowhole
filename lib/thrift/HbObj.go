package hbase2go

import (
	"fmt"
	"net"
	"strconv"

	"github.com/lightstep/lightstep-tracer-go/thrift_0_9_2/lib/go/thrift"
)

type HbObj struct {
	Host   string
	Port   int
	Trans  *thrift.TSocket
	Client *THBaseServiceClient
}

func NewHbObj(host string, port int) HbObj {
	return HbObj{
		Host: host,
		Port: port,
	}
}

func (h *HbObj) Close() {
	h.Trans.Close()
}
func (h *HbObj) Connect() error {
	var err error
	h.Trans, err = thrift.NewTSocket(net.JoinHostPort(h.Host, strconv.Itoa(h.Port)))
	if err != nil {
		return fmt.Errorf("error resolving address:%s", err)

	}
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	h.Client = NewTHBaseServiceClientFactory(h.Trans, protocolFactory)
	return h.Trans.Open()
}

func (h *HbObj) GetRow(table string, tget *TGet) (*TResult_, error) {
	return h.Client.Get([]byte(table), tget)
}

func (h *HbObj) Put(table string, tput *TPut) (err error) {
	return h.Client.Put([]byte(table), tput)
}
