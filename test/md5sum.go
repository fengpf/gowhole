package base

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

const testPath = "/path/to/file"

// BenchmarkMd5Sum go test -test.run=none -test.bench="^BenchmarkMd5Sum
func BenchmarkMd5Sum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		md5sum1(testPath)
	}
}

func md5sum1(file string) string {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%x", md5.Sum(data))
}

func md5sum2(file string) string {
	f, err := os.Open(file)
	if err != nil {
		return ""
	}
	defer f.Close()
	h := md5.New()
	_, err = io.Copy(h, f)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}

func md5sum3(file string) string {
	f, err := os.Open(file)
	if err != nil {
		return ""
	}
	defer f.Close()
	r := bufio.NewReader(f)
	h := md5.New()
	_, err = io.Copy(h, r)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%x", h.Sum(nil))

}
