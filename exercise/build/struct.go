package main

/*
struct Vertex {
    int X;
    int Y;
};
*/

import "C"
import "fmt"

//export getVertex
func getVertex(X, Y C.int) C.struct_Vertex {
	return C.struct_Vertex{X, Y}
}

func main() {
	fmt.Println(getVertex(1, 2))
}
