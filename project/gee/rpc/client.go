package rpc

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"gowhole/project/gee/rpc/codec"
	"io"
	"log"
	"net"
	"sync"
	"time"
)

//func (t *T) MethodName(argType1 T1, replyType *T2)  error {
//
//}

type Call struct {
	Seq uint64
	ServiceMethod string
	Args interface{}
	Reply interface{}
	Error error
	Done chan *Call
}

func (call *Call) done()  {
	call.Done<-call
}

// Client represents an RPC Client.
// There may be multiple outstanding Calls associated
// with a single Client, and a Client may be used by
// multiple goroutines simultaneously.

type Client struct {
	cc codec.Codec
	opt *Option
	sending sync.Mutex
	header codec.Header
	mux sync.Mutex
	seq uint64
	pending map[uint64]*Call
	closing bool
	shutdown bool
}

var ErrShutdown = errors.New("connection is shut down")

var _ io.Closer = (*Client)(nil)

func (c Client) Close() error {
	c.mux.Lock()
	defer c.mux.Unlock()

	if c.closing{
		return ErrShutdown
	}

	c.closing = true
	return c.cc.Close()
}

// IsAvailable return true if the client does work
func (c *Client) IsAvailable() bool {
	c.mux.Lock()
	defer c.mux.Unlock()
	return !c.shutdown && !c.closing
}

func (c *Client) registerCall(call *Call) (uint64, error) {
	c.mux.Lock()
	defer c.mux.Unlock()
	if c.closing||c.shutdown{
		return 0,ErrShutdown
	}

	call.Seq =  c.seq
	c.pending[call.Seq] = call
	c.seq++
	return call.Seq,nil
}

func (c *Client) removeCall(seq uint64) *Call {
	c.mux.Lock()
	defer c.mux.Unlock()

	call := c.pending[seq]
	delete(c.pending, seq)

	return call
}

func (c *Client) terminateCalls(err error) {
	c.sending.Lock()
	defer c.sending.Unlock()

	c.mux.Lock()
	defer c.mux.Unlock()

	c.shutdown = true
	for _, call := range c.pending {
		call.Error = err
		call.done()
	}
}

func (c *Client) receive()  {
	var err error
	for  {
	   var h codec.Header
	   if err=c.cc.ReadHeader(&h);err!=nil{
	   	 break
	   }
	   
	   call:=c.removeCall(h.Seq)

		switch  {
		case call==nil:
			// it usually means that Write partially failed
			// and call was already removed.
			err = c.cc.ReadBody(nil)
		case h.Error != "":
			call.Error = fmt.Errorf(h.Error)
			err = c.cc.ReadBody(nil)
			call.done()
		default:
			err = c.cc.ReadBody(call.Reply) //读入响应
			if err != nil {
				call.Error = errors.New("reading body " + err.Error())
			}
			call.done()
		}
		// error occurs, so terminateCalls pending calls
		c.terminateCalls(err)
	}
}

func NewClient(conn net.Conn, opt *Option) (*Client, error) {
	f:=codec.NewCodecFuncMap[opt.CodecType]
	if f==nil{
		err := fmt.Errorf("invalid codec type %s", opt.CodecType)
		log.Println("rpc client: codec error:", err)
		return nil, err
	}

	// send options with server
	if err := json.NewEncoder(conn).Encode(opt); err != nil {
		log.Println("rpc client: options error: ", err)
		_ = conn.Close()
		return nil, err
	}
	return newClientCodec(f(conn), opt), nil
}

func newClientCodec(cc codec.Codec, opt *Option) *Client {
	client := &Client{
		seq:     1, // seq starts with 1, 0 means invalid call
		cc:      cc,
		opt:     opt,
		pending: make(map[uint64]*Call),
	}
	go client.receive()
	return client
}

func parseOptions(opts ...*Option) (*Option, error) {
	// if opts is nil or pass nil as parameter
	if len(opts) == 0 || opts[0] == nil {
		return DefaultOption, nil
	}
	if len(opts) != 1 {
		return nil, errors.New("number of options is more than 1")
	}

	opt := opts[0]
	opt.MagicNumber = DefaultOption.MagicNumber
	if opt.CodecType == "" {
		opt.CodecType = DefaultOption.CodecType
	}
	return opt, nil
}

// Dial connects to an RPC server at the specified network address
//func Dial(network, address string, opts ...*Option) (client *Client, err error) {
//	opt,err:=parseOptions(opts...)
//	if err!=nil{
//		return nil, err
//	}
//
//	conn,err:=net.Dial(network, address)
//	if err!=nil{
//		return nil, err
//	}
//
//	// close the connection if client is nil
//	defer func() {
//		if client == nil {
//			_ = conn.Close()
//		}
//	}()
//	return NewClient(conn, opt)
//}Client) Call


// Dial connects to an RPC server at the specified network address
func Dial(network, address string, opts ...*Option) (*Client, error) {
	return dialTimeout(NewClient, network, address, opts...)
}

func (c *Client) send(call *Call) {
	// make sure that the client will send a complete request
   c.sending.Lock()
   defer c.sending.Unlock()

	// register this call.
	seq,err:=c.registerCall(call)
	if err!=nil{
		call.Error = err
		call.done()
		return
	}

	// prepare request header
	c.header.ServiceMethod = call.ServiceMethod
	c.header.Seq = seq
	c.header.Error = ""

	// encode and send the request
	if err:=c.cc.Write(&c.header, call.Args);err!=nil{
		call:=c.removeCall(seq)
		// call may be nil, it usually means that Write partially failed,
		// client has received the response and handled
		if call != nil {
			call.Error = err
			call.done()
		}
	}
}

// Go invokes the function asynchronously.
// It returns the Call structure representing the invocation.
func (c *Client) Go(serviceMethod string, args, reply interface{}, done chan *Call) *Call {
	if done==nil{
		done =make(chan *Call, 10)
	}else if cap(done) ==0 {
		log.Panic("rpc client: done channel is unbuffered")
	}

	call:=&Call{
		ServiceMethod: serviceMethod,
		Args: args,
		Reply: reply,
		Done: done,
	}
	c.send(call)
	return call
}

// Call invokes the named function, waits for it to complete,
// and returns its error status.
//func (c *Client) Call(serviceMethod string, args, reply interface{}) error {
//	 call:=<-c.Go(
//	 	serviceMethod,
//	 	args,
//	 	reply,
//	 	make(chan *Call, 1)).Done
//
//	 return call.Error
//}

// Call invokes the named function, waits for it to complete,
// and returns its error status.
func (client *Client) Call(ctx context.Context, serviceMethod string, args, reply interface{}) error {
	call := client.Go(serviceMethod, args, reply, make(chan *Call, 1))
	select {
	case <-ctx.Done():
		client.removeCall(call.Seq)
		return errors.New("rpc client: call failed: " + ctx.Err().Error())
	case call := <-call.Done:
		return call.Error
	}
}

//Go 和 Call 是客户端暴露给用户的两个 RPC 服务调用接口，Go 是一个异步接口，返回 call 实例。
//Call 是对 Go 的封装，阻塞 call.Done，等待响应返回，是一个同步接口。


type clientResult struct {
	client *Client
	err    error
}

type newClientFunc func(conn net.Conn, opt *Option) (client *Client, err error)

func dialTimeout(f newClientFunc, network, address string, opts ...*Option) (client *Client, err error) {
    opt,err:=parseOptions(opts...)
    if err!=nil{
    	return nil, err
	}

	conn,err:=net.DialTimeout(network, address, opt.ConnectTimeout)
	if err!=nil{
		return nil, err
	}
	// close the connection if client is nil
	defer func() {
		if err != nil {
			_ = conn.Close()
		}
	}()

	ch:=make(chan clientResult)
	go func() {
		client,err:=f(conn, opt)
		ch<-clientResult{client: client,err: err}
	}()

	if opt.ConnectTimeout == 0 {
		result := <-ch
		return result.client, result.err
	}

	select {
	  case <-time.After(opt.ConnectTimeout):
		  return nil, fmt.Errorf("rpc client: connect timeout: expect within %s", opt.ConnectTimeout)
      case result:=<-ch:
		 return result.client, result.err
	}
}
