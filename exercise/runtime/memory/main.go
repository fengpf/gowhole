package main

import (
	"fmt"
	"reflect"
	"unsafe"

	"gowhole/exercise/convert"
)

//逃逸分析（escape analysis）
//go run -gcflags '-m -l' main.go

//查看汇编代码
//go build -gcflags '-l' -o  main main.go
//go tool objdump -s "main\.main" main

type student struct {
	Name  string
	Age   int
	Score float32
}

func main() {
	var (
		stu student
	)
	stu.Age = 18
	stu.Name = "xiaohan"
	stu.Score = 90

	fmt.Println(stu)

	fmt.Printf("Name :%p \t sizeof :%d\n", &stu.Name, unsafe.Sizeof(stu.Name))   //Name :0xc42009a020  sizeof :16
	fmt.Printf("Age :%p  \t sizeof :%d\n", &stu.Age, unsafe.Sizeof(stu.Name))    //Age :0xc42009a030   sizeof :16
	fmt.Printf("Score :%p \t sizeof :%d\n", &stu.Score, unsafe.Sizeof(stu.Name)) //Score :0xc42009a038 sizeof :16

	t := reflect.TypeOf(stu)
	fmt.Println(unsafe.Sizeof(stu), t.Align()) //32 8

	name, _ := t.FieldByName("Name")
	age, _ := t.FieldByName("Age")
	score, _ := t.FieldByName("Score")

	//Name :8	Age :8	 Score :4
	fmt.Printf("Name :%d\t Age :%d\t Score :%d\t \n", name.Type.FieldAlign(), age.Type.FieldAlign(), score.Type.FieldAlign())

	// First ask Go to give us some information about the MyData type
	typ := reflect.TypeOf(stu)
	fmt.Printf("Struct is %d bytes long\n", typ.Size())
	// We can run through the fields in the structure in order
	n := typ.NumField()
	for i := 0; i < n; i++ {
		field := typ.Field(i)
		fmt.Printf("%s at offset %v, size=%d, align=%d, addr=%p\n",
			field.Name, field.Offset, field.Type.Size(),
			field.Type.Align(), &field.Name)
	}

	fmt.Println(convert.HexBin("0xc42009a020"), convert.HexBin("0xc42009a030"))
}

//test
func foo() *int {
	var x int
	return &x
}

func bar() int {
	x := new(int)
	*x = 1
	return *x
}

// ./main.go:61:9: &x escapes to heap
// ./main.go:60:6: moved to heap: x
// ./main.go:65:10: bar new(int) does not escape
// 上面的意思是 foo() 中的 x 最后在堆上分配，而 bar() 中的 x 最后分配在了栈上

// 对齐系数,GOARCH=amd64是8，GOARCH=386是4
// So(unsafe.Sizeof(true), ShouldEqual, 1)
// So(unsafe.Sizeof(int8(0)), ShouldEqual, 1)
// So(unsafe.Sizeof(int16(0)), ShouldEqual, 2)
// So(unsafe.Sizeof(int32(0)), ShouldEqual, 4)
// So(unsafe.Sizeof(int64(0)), ShouldEqual, 8)
// So(unsafe.Sizeof(int(0)), ShouldEqual, 8)
// So(unsafe.Sizeof(float32(0)), ShouldEqual, 4)
// So(unsafe.Sizeof(float64(0)), ShouldEqual, 8)
// So(unsafe.Sizeof(""), ShouldEqual, 16)
// So(unsafe.Sizeof("hello world"), ShouldEqual, 16)
// So(unsafe.Sizeof([]int{}), ShouldEqual, 24)
// So(unsafe.Sizeof([]int{1, 2, 3}), ShouldEqual, 24)
// So(unsafe.Sizeof([3]int{1, 2, 3}), ShouldEqual, 24)
// So(unsafe.Sizeof(map[string]string{}), ShouldEqual, 8)
// So(unsafe.Sizeof(map[string]string{"1": "one", "2": "two"}), ShouldEqual, 8)
// So(unsafe.Sizeof(struct{}{}), ShouldEqual, 0)

// bool 类型虽然只有一位，但也需要占用1个字节，因为计算机是以字节为单位
// 64为的机器，一个 int 占8个字节
// string 类型占16个字节，内部包含一个指向数据的指针（8个字节）和一个 int 的长度（8个字节）
// slice 类型占24个字节，内部包含一个指向数据的指针（8个字节）和一个 int 的长度（8个字节）和一个 int 的容量（8个字节）
// map 类型占8个字节，是一个指向 map 结构的指针
// 可以用 struct{} 表示空类型，这个类型不占用任何空间，用这个作为 map 的 value，可以讲 map 当做 set 来用
