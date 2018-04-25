package buffer

import (
	"math/rand"
	"time"
)

var buf = make(map[interface{}]*boxType)
var msgChan = make(chan interface{})
var returnChan = make(chan interface{})

type boxType struct {
	data  interface{}
	timer *time.Timer
	id    int
}

type setMsg struct {
	key     interface{}
	value   interface{}
	timeout time.Duration
}

type getMsg struct {
	key interface{}
}

type removeMsg struct {
	key interface{}
	id  int
}

func init() {
	go func() {
		for msg := range msgChan {
			switch m := msg.(type) {
			case *setMsg:
				box, has := buf[m.key]
				if has {
					box.timer.Stop()
				}
				box = &boxType{}
				box.data = m.value
				box.id = rand.Int()
				box.timer = time.AfterFunc(m.timeout, func() {
					msgChan <- &removeMsg{
						key: m.key,
						id:  box.id,
					}
				})
				buf[m.key] = box
			case *getMsg:
				box, ok := buf[m.key]
				if ok {
					returnChan <- box.data
				} else {
					returnChan <- nil
				}
			case *removeMsg:
				box, has := buf[m.key]
				if has {
					if m.id == box.id {
						delete(buf, m.key)
					}
				}
			}
		}
	}()
}

// Set key -> value to buffer, when timeout(nanosecond) remove it.
func Set(key, value interface{}, timeout time.Duration) {
	msgChan <- &setMsg{
		key:     key,
		value:   value,
		timeout: timeout,
	}
}

// Get value by key from buffer. If it not in buffer then return nil.
func Get(key interface{}) interface{} {
	msgChan <- &getMsg{
		key: key,
	}
	return <-returnChan
}

// Close this buffer.
func Close() {
	close(msgChan)
	close(returnChan)
}
