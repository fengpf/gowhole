package main

import (
	"fmt"

	"github.com/let-z-go/aspector"
)

type Greeter interface {
	Hello(name string) bool
}

type GreeterImpl struct{}

func (*GreeterImpl) Hello(name string) bool {
	fmt.Println("Hello, " + name + "!")
	return true
}

func SayHello(greeter Greeter) {
	ok := greeter.Hello("Mr.Aspector")
	if ok {
		fmt.Println("- Succeeded")
	} else {
		fmt.Println("- Failed")
	}
}

func main() {
	greeterImpl := &GreeterImpl{}
	SayHello(greeterImpl)
	// output:
	// Hello, Mr.Aspector!
	// - Succeeded

	greeterWrap := &GreeterWrap{
		Origin: greeterImpl, // proxy `greeterImpl`
	}

	// Add an interceptor
	greeterWrap.AddMethodInterceptor(func(wrap aspector.Wrap, methodID int, rawArgs interface{}, methodHandler aspector.MethodHandler) interface{} {
		fmt.Printf("===== Before: %s.%s\n", wrap.GetName(), wrap.GetMethodName(methodID))

		if methodID == Greeter_Hello { // the method `Hello` is called
			args := rawArgs.(*Greeter_HelloArgs) // fetch the arguments passed
			args.V1 = "World"                    // update the argument (`V1` is the first argument, there is only one argument)
		}

		rawResults := methodHandler(rawArgs) // call the next method interceptor or the real method if no more method interceptor

		if methodID == Greeter_Hello { // the method `Hello` is called
			results := rawResults.(*Greeter_HelloResults) // fetch the results returned
			results.V1 = false                            // update the result (`V1` is the first result, there is only one result)
		}

		fmt.Printf("===== After: %s.%s\n", wrap.GetName(), wrap.GetMethodName(methodID))
		return rawResults // pass the results
	})

	SayHello(greeterWrap)
	// output:
	// ===== Before: Greeter.Hello
	// Hello, World!
	// ===== After: Greeter.Hello
	// - Failed
}
