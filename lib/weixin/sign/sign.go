package weixin

import (
	"bufio"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
)

//CheckSignature check sginature.
func CheckSignature(c *gin.Context) bool {
	signature := c.Query("signature")
	timestamp := c.Query("timestamp")
	nonce := c.Query("nonce")
	echostr := c.Query("echostr")
	token := token
	sigin := sign(token, timestamp, nonce)
	if sigin == signature {
		println("check signature success")
		fmt.Fprintf(c.Writer, echostr)
		return true
	}
	return false
}

//https://api.weixin.qq.com/cgi-bin/token?
//grant_type=client_credential&appid=APPID&secret=APPSECRET
func testAccessToken() (accessToken string, err error) {
	params := url.Values{}
	params.Set("grant_type", "client_credential")
	params.Set("appid", appID)
	params.Set("secret", appSecret)
	aurl := accessTokenHost + "?" + params.Encode()
	resp, err := http.Get(aurl)
	if err != nil {
		log.Printf("http.Get %v\n", err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Printf("http.Status %v\n", resp.Status)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	//{"errcode":40013,"errmsg":"invalid appid"}
	//{"access_token":"ACCESS_TOKEN","expires_in":7200}
	var result struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int64  `json:"expires_in"`
	}
	if err = json.Unmarshal(body, &result); err != nil {
		log.Printf("json.Unmarshal %v\n", resp.Status)
		return
	}
	accessToken = result.AccessToken
	return
}
func sign(token, timestamp, nonce string) (sigin string) {
	arr := []string{token, timestamp, nonce}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	str := strings.Join(arr, "")
	sh := sha1.New()
	sh.Write([]byte(str))
	bs := sh.Sum(nil)
	sigin = hex.EncodeToString(bs)
	return
}

func newSign(token, timestamp, nonce string) (signature string) {
	strs := sort.StringSlice{token, timestamp, nonce}
	strs.Sort()
	buf := make([]byte, 0, len(token)+len(timestamp)+len(nonce))
	buf = append(buf, strs[0]...)
	buf = append(buf, strs[1]...)
	buf = append(buf, strs[2]...)
	hashsum := sha1.Sum(buf)
	return hex.EncodeToString(hashsum[:])
}

func MsgSign(token, timestamp, nonce, encryptedMsg string) (signature string) {
	strs := sort.StringSlice{token, timestamp, nonce, encryptedMsg}
	strs.Sort()
	h := sha1.New()
	bufw := bufio.NewWriterSize(h, 128)
	bufw.WriteString(strs[0])
	bufw.WriteString(strs[1])
	bufw.WriteString(strs[2])
	bufw.WriteString(strs[3])
	bufw.Flush()
	hashsum := h.Sum(nil)
	return hex.EncodeToString(hashsum)
}
