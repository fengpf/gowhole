package parall

import "sync"

//Goroutine的主函数是没有参数的。传引用的情况利用了upvalue,而需要传值的i变量用了一个外包函数的参数来复制。
//因为每次循环都会调用这个外包函数，从而复制了一次i的数值，
//虽然里层的Goroutine主函数还是 通过 upvalue来捕获i，不过每次捕获的都是外包函数的i副本而已。
func t() {
	var wg sync.WaitGroup
	wg = sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		func(i int) {
			go func() {
				println(i)
				wg.Done()
			}()
		}(i)
	}
	wg.Wait()
}
