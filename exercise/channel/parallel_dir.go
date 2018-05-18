package channel

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"sync"
	"testing"
)

func Test_parallelDir(t *testing.T) {
	// 计算指定目录下所有文件的MD5值，之后按照目录名排序并打印结果
	m, err := MD5All2(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	var paths []string
	for path := range m {
		paths = append(paths, path)
	}
	sort.Strings(paths)
	for _, path := range paths {
		fmt.Printf("%x  %s\n", m[path], path)
	}
}

func sumFiles(done <-chan struct{}, root string) (<-chan result, <-chan error) {
	// 对每个常规文件，启动一个Goroutine计算文件内容并发送结果到c。发送walk的结果到errc
	c := make(chan result)
	errc := make(chan error, 1)
	go func() {
		var wg sync.WaitGroup
		err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}
			wg.Add(1)
			go func() {
				data, err := ioutil.ReadFile(path)
				select {
				case c <- result{path, md5.Sum(data), err}:
				case <-done:
				}
				wg.Done()
			}()
			// 如果done被关闭了，停止walk
			select {
			case <-done:
				return errors.New("walk canceled")
			default:
				return nil
			}
		})
		// walk已经返回，所有wg.Add的工作都做完了。开启新进程，在所有发送完成后
		// 关闭c。
		go func() {
			wg.Wait()
			close(c)
		}()
		// 因为errc有缓冲区，所以这里不需要select。
		errc <- err
	}()
	return c, errc
}

// MD5All 在返回时关闭done channel；这个可能在从c和errc收到所有的值之前被调用
func MD5All2(root string) (m map[string][md5.Size]byte, err error) {
	done := make(chan struct{})
	defer close(done)
	c, errc := sumFiles(done, root)
	m = make(map[string][md5.Size]byte)
	for r := range c {
		if r.err != nil {
			return nil, r.err
		}
		m[r.path] = r.sum
	}
	if err = <-errc; err != nil {
		return nil, err
	}
	return m, nil
}
