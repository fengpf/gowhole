package main

func main() {
	var q Queue

	for i := 0; i < 3; i++ {
		q.Enqueue(i + 1)
		v, ok := q.Dequeue()
		if !ok {
			break
		}
		println(v)
	}

	// for i := 0; i < 10; i++ {
	// 	go func() {
	// 		for i := 0; i < 100; i++ {
	// 			q.Enqueue(i + 1)
	// 		}
	// 	}()

	// 	go func() {
	// 		for {
	// 			v, ok := q.Dequeue()
	// 			if !ok {
	// 				break
	// 			}
	// 			_ = v
	// 			// println(v)
	// 		}
	// 	}()
	// }
	// fmt.Printf("%#v\n", q)
}
