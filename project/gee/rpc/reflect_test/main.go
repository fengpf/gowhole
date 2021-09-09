package main

import (
	"fmt"
	"reflect"
	"strings"
	"sync"
)

func main()  {
	var wg sync.WaitGroup
	typ :=reflect.TypeOf(&wg)

	for i:=0;i<typ.NumMethod();i++{
		fmt.Println(111, i)

		method:=typ.Method(i)
		argv:=make([]string,0,method.Type.NumIn())
		returns:=make([]string,0,method.Type.NumOut())

		// j 从 1 开始，第 0 个入参是 wg 自己。
		for j:=1;j<method.Type.NumIn();i++{
			argv = append(argv, method.Type.In(j).Name())
		}
		for j := 0; j < method.Type.NumOut(); j++ {
			returns = append(returns, method.Type.Out(j).Name())
		}


		fmt.Printf("func (w *%s) %s(%s) %s",
			typ.Elem().Name(),
			method.Name,
			strings.Join(argv, ","),
			strings.Join(returns, ","))
	}
}
