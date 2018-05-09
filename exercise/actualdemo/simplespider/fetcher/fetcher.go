package fetcher

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

func Fetch(url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	// req.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 10_3 like Mac OS X) AppleWebKit/602.1.50 (KHTML, like Gecko) CriOS/56.0.2924.75 Mobile/14E5239e Safari/602.1")
	client := http.Client{
		CheckRedirect: func(request *http.Request, via []*http.Request) error {
			fmt.Println("Redirect:", request)
			return nil
		},
	}
	if err != nil {
		return nil, err
	}
	// resp, err := http.DefaultClient.Do(req)//use default client
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error status code %d\n", resp.StatusCode)
	}

	// s, err := httputil.DumpResponse(resp, true)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%s\n", s)
	buf := bufio.NewReader(resp.Body)
	e := determineEncodeing(buf)
	utf8Reader := transform.NewReader(buf, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

func determineEncodeing(buf *bufio.Reader) encoding.Encoding {
	bs, err := bufio.NewReader(buf).Peek(1024)
	if err != nil {
		log.Printf("fetch error:%v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bs, "")
	return e
}
