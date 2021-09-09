package codec

import (
	"bufio"
	"encoding/gob"
	"io"
	"log"
)

//ServiceMethod 是服务名和方法名，通常与 Go 语言中的结构体和方法相映射。
//Seq 是请求的序号，也可以认为是某个请求的 ID，用来区分不同的请求。
//Error 是错误信息，客户端置为空，服务端如果如果发生错误，将错误信息置于 Error 中。
type Header struct {
	ServiceMethod string
	Seq uint64
	Error string
}

type Codec interface {
	io.Closer
	ReadHeader(*Header) error
	ReadBody(interface{}) error
	Write(*Header, interface{}) error
}

type NewCodecFunc func(io.ReadWriteCloser) Codec

type Type string

const (
	GobType  Type = "application/gob"
	JsonType Type = "application/json" // not implemented
)

var NewCodecFuncMap map[Type]NewCodecFunc

func init()  {
	NewCodecFuncMap=make(map[Type]NewCodecFunc)
	NewCodecFuncMap[GobType] = NewGobCodec
}

type GobCodec struct {
	conn io.ReadWriteCloser
	buf *bufio.Writer
	dec *gob.Decoder
	enc *gob.Encoder
}

var _ Codec = (*GobCodec)(nil)

func (g *GobCodec) Close() error {
	return g.conn.Close()
}

func (g *GobCodec) ReadHeader(header *Header) error {
	return g.dec.Decode(header)
}

func (g *GobCodec) ReadBody(body interface{}) error {
	return g.dec.Decode(body)
}

func (g *GobCodec) Write(header *Header, body interface{}) (err error) {
	defer func() {
		_ = g.buf.Flush()
		if err!=nil{
			_ = g.Close()
		}
	}()

	if err=g.enc.Encode(header); err!=nil{
		log.Println("rpc codec: gob error encoding header:", err)
		return err
	}

	if err := g.enc.Encode(body); err != nil {
		log.Println("rpc codec: gob error encoding body:", err)
		return err
	}
	return nil
}

func NewGobCodec(conn io.ReadWriteCloser)  Codec{
	buf:=bufio.NewWriter(conn)

	return &GobCodec{
		conn: conn,
		buf:buf,
		dec:gob.NewDecoder(conn),
		enc:gob.NewEncoder(buf),
	}
}