package basictests

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"testing"
	"unsafe"
)

func Test_byte(t *testing.T) {
	str2 := "aAhello"
	data2 := []byte(str2)
	fmt.Println(data2)
	str2 = string(data2[:])
	fmt.Println(str2)

	upBytes()
	upBytes2()
}

func upBytes() {
	bytes := []byte("I am byte array !")
	println(&bytes)
	str := string(bytes)
	println(&str)
	bytes[0] = 'i' //注意这一行，bytes在这里修改了数据，但是str打印出来的依然没变化，
	fmt.Println(str)
}

func upBytes2() {
	bytes := []byte("I am byte array !")
	//这样做的意义在于，在网络通信中，大多数的接受方式都是[]byte，如果[]byte的数据比较大，内存拷贝的话会影响系统的性能。
	str := (*string)(unsafe.Pointer(&bytes))
	bytes[0] = 'i' //str和bytes共用一片内存
	println(&bytes)
	println(&str)
	fmt.Println(*str)
}

type AB struct {
	One int32
	Two int32
}

func Test_byteSort(t *testing.T) {
	i := uint16(1)
	size := binary.Size(i)
	fmt.Println(i, size)

	var a AB
	a.One = int32(1)
	a.Two = int32(2)
	buf := new(bytes.Buffer)
	fmt.Println("a size is ", binary.Size(a))
	binary.Write(buf, binary.LittleEndian, a)
	fmt.Println("after write, buf is:", buf.Bytes())

	var aa AB
	buf1 := new(bytes.Buffer)
	binary.Write(buf1, binary.LittleEndian, a)
	binary.Read(buf1, binary.LittleEndian, &aa)
	fmt.Println("after aa is ", aa)

	int16buf := new(bytes.Buffer)
	binary.Write(int16buf, binary.LittleEndian, i)
	fmt.Println("write buf is:", int16buf.Bytes())

	var int16buf2 [2]byte
	binary.LittleEndian.PutUint16(int16buf2[:], uint16(1))
	fmt.Println("put buffer is :", int16buf2[:])

	ii := binary.LittleEndian.Uint16(int16buf2[:])
	fmt.Println("Get buf is :", ii)
}

//这个是一个常见的在tcp 拼包得例子。
//在例子中通过binary.BigEndian.Uint16将数据按照网络序的格式读出来，放入到head中对应的结构里面。
type Head struct {
	Cmd     byte
	Version byte
	Magic   uint16
	Reserve byte
	HeadLen byte
	BodyLen uint16
}

func NewHead(buf []byte) *Head {
	head := new(Head)
	head.Cmd = buf[0]
	head.Version = buf[1]
	head.Magic = binary.BigEndian.Uint16(buf[2:4])
	head.Reserve = buf[4]
	head.HeadLen = buf[5]
	head.BodyLen = binary.BigEndian.Uint16(buf[6:8])
	return head
}

func Test_IntByte(t *testing.T) {
	i := uint64(1)
	size := binary.Size(i)
	fmt.Println(i, size)

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, i)
	fmt.Println("after write, buf is:", buf.Bytes())

	var u uint8 = 255
	fmt.Println(u, u+1, u*u) // "255 0 1"

	var j int8 = 127
	fmt.Println(j, j+1, j*j) // "127 -128 1"

	var a int = 12
	fmt.Println(a, unsafe.Sizeof(a), binary.Size(a))
	buf2 := new(bytes.Buffer)
	binary.Write(buf2, binary.BigEndian, a)
	fmt.Println("after write, buf2 is:", buf2.Bytes())
	var b int32
	binary.Read(buf2, binary.BigEndian, &b)
	fmt.Println(b)
}

func Test_byteScope(t *testing.T) {
	i := uint32(1)
	size := binary.Size(i)
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, i)
	fmt.Println(size, len(buf.Bytes()))
}

func Test_byteToInt(t *testing.T) {
	c := make([]byte, 4, 4)
	// c := []byte{255, 255, 255, 228}
	// c := []byte{255, 255, 255, 248}
	// c := []byte{0, 0, 2, 248}
	// c = append(c, byte(1))
	c[0] = byte(0)
	c[1] = byte(0)
	c[2] = byte(2)
	c[3] = byte(248)
	fmt.Println(byteToInt32(c))

}

func byteToInt32(b []byte) int32 {
	return int32(b[3]) | int32(b[2])<<8 | int32(b[1])<<16 | int32(b[0])<<24
}
