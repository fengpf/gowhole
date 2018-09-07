package main

import "testing"

func BenchmarkQueue(b *testing.B) {
	var q Queue
	for i := 0; i < b.N; i++ {
		for i := 0; i < 10; i++ {
			go func() {
				for i := 0; i < 100; i++ {
					q.Enqueue(i)
				}
			}()

			go func() {
				for {
					v, ok := q.Dequeue()
					if !ok {
						break
					}
					_ = v
					// println(v)
				}
			}()
		}
	}
	// fmt.Printf("%#v\n", q)
}

func TestQueue(t *testing.T) {
	tests := []struct {
		x []int
	}{
		{[]int{1, 2, 3, 4, 5, 6}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 100, 10}},
	}

	for _, tt := range tests {
		o := tt
		t.Run("TestQueue", func(t *testing.T) {
			var q Queue

			for _, i := range o.x {
				q.Enqueue(i)
			}

			for i := 0; i < len(o.x); i++ {
				v, ok := q.Dequeue()

				if !ok || v != o.x[i] {
					t.Fatal()
				}
			}
		})
	}
}
