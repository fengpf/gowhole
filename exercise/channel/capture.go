package channel

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"testing"
)

const (
	//URL 抓取目标地址.
	URL string = "http://www.mzitu.com/xinggan"
	//PATH img path
	PATH string = "meizi"
)

var (
	c1 chan string // 保存一级页面url
	c2 chan string //  保存图片url
	c3 chan int    // 记录下载数量
)

func init() {
	c1 = make(chan string, 1)
	c2 = make(chan string, 1)
	c3 = make(chan int, 1)
	go regexParentHTML()
	go regexChildHTML()
	go downloadImg()
}

func Test_capture(t *testing.T) {
	fmt.Println("start")
	for item := range c3 {
		fmt.Println("已经下载完成第" + strconv.Itoa(item) + "张图片")
	}
	fmt.Println("End")
}

// 获取HTML内容
func getHTMLContent(url string) (content string) {
	response, err := http.Get(url)
	checkErr(err)
	defer response.Body.Close()
	str, err := ioutil.ReadAll(response.Body)
	checkErr(err)
	return string(str)
}

// 解析第一级页面，获取一级url
func regexParentHTML() {
	reg, err := regexp.Compile(`<li><a href=\"(.*?)\" target=\"_blank\">`)
	checkErr(err)
	content := getHTMLContent(URL)
	lists := reg.FindAllStringSubmatch(content, 1)
	for _, item := range lists {
		println(item[1])
		c1 <- item[1]
	}
	println(c1)
	close(c1)
}

// 解析二级页面
func regexChildHTML() {
	for url := range c1 {
		content := getHTMLContent(url)
		page := getChildPage(content)
		for i := 1; i < page; i++ {
			scontent := getHTMLContent(url + "/" + strconv.Itoa(i))
			getImgURL(scontent)
		}
	}
	close(c2)
}

// 获取图片地址，存入c2
func getImgURL(content string) {
	reg, err := regexp.Compile(`<img src=\"(.*?)\" alt=(.*?)>`)
	checkErr(err)
	lists := reg.FindAllStringSubmatch(content, 1000)
	for _, item := range lists {
		c2 <- item[1]
	}
}

// 获取二级页面页数
func getChildPage(content string) (page int) {
	reg, err := regexp.Compile(`<span>([0-9]{2})<\/span>`)
	checkErr(err)
	list := reg.FindAllStringSubmatch(content, 100)
	page, _ = strconv.Atoi(list[len(list)-1][1])
	return
}

// 下载图片
func downloadImg() {
	info, err := os.Stat(PATH)
	if err != nil || info.IsDir() == false {
		err := os.Mkdir(PATH, os.ModePerm)
		checkErr(err)
	}
	fnum := 1
	for item := range c2 {
		fileName := PATH + "/" + strconv.Itoa(fnum) + ".jpg"
		fmt.Println(fileName)
		body := getImgBody(item)
		ioutil.WriteFile(fileName, body, 0644)
		fnum++
		c3 <- fnum
	}
	close(c3)
}

// 获取图片body
func getImgBody(url string) (body []byte) {
	response, err := http.Get(url)
	checkErr(err)
	defer response.Body.Close()
	str, err := ioutil.ReadAll(response.Body)
	checkErr(err)
	body = str
	return
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
