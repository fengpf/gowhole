package main

import (
	"fmt"
	"reflect"
	"unsafe"
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
	var stu student
	stu.Age = 18
	stu.Name = "xiaohan"
	stu.Score = 90
	fmt.Println(stu)

	t := reflect.TypeOf(stu)
	name, _ := t.FieldByName("Name")
	age, _ := t.FieldByName("Age")
	score, _ := t.FieldByName("Score")

	fmt.Printf("Name :%p \t 大小 :%d\t 对齐:%d \n", &stu.Name, unsafe.Sizeof(stu.Name), name.Type.FieldAlign())    //Name :0xc42009a020  sizeof :16
	fmt.Printf("Age :%p  \t 大小 :%d\t 对齐:%d\n", &stu.Age, unsafe.Sizeof(stu.Name), age.Type.FieldAlign())       //Age :0xc42009a030   sizeof :16
	fmt.Printf("Score :%p \t 大小 :%d\t 对齐:%d \n", &stu.Score, unsafe.Sizeof(stu.Name), score.Type.FieldAlign()) //Score :0xc42009a038 sizeof :16

	println("=================================")
	typ := reflect.TypeOf(stu)
	fmt.Printf("结构体大小：(%2v)，对齐系数：(%2v)\n", unsafe.Sizeof(stu), unsafe.Alignof(stu)) //aaaa|aaaa|aaaa|aaaa|bxxx|cccc
	n := typ.NumField()
	for i := 0; i < n; i++ {
		field := typ.Field(i)
		fmt.Printf(
			"字段：%2v\t 类型：%5v\t 大小：%2v\t 对齐：%2v\t 字段对齐：%2v\t 偏移：%2v\t 地址：%2v\n",
			field.Name,
			field.Type.Kind(),
			field.Type.Size(),
			field.Type.Align(),
			field.Type.FieldAlign(),
			field.Offset,
			&field.Name,
		)
	}

	println("=================================")

	f := unsafe.Pointer(uintptr(unsafe.Pointer(&stu)) + unsafe.Offsetof(stu.Name))
	fmt.Println(
		unsafe.Pointer(&stu),
		uintptr(unsafe.Pointer(&stu)),
		unsafe.Offsetof(stu.Name),
		uintptr(unsafe.Pointer(&stu))+unsafe.Offsetof(stu.Name),
		f,
		*(*string)(f),
	)

	println("=================================")

	f2 := unsafe.Pointer(uintptr(unsafe.Pointer(&stu)) + unsafe.Offsetof(stu.Age))
	fmt.Println(
		unsafe.Pointer(&stu),
		uintptr(unsafe.Pointer(&stu)),
		unsafe.Offsetof(stu.Age),
		uintptr(unsafe.Pointer(&stu))+unsafe.Offsetof(stu.Age),
		f2,
		*(*int)(f2),
	)
}
