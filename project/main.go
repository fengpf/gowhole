package main

import (
	"fmt"
	"net/http"
)

var (
	PlayURL = "https://upos-sz-mirrorks3.bilivideo.com/ugaxcode/m201014a21e2s59ms7jh8r3747yidagf-192k.m4a?e=ig8euxZM2rNcNbdlhoNvNC8BqJIzNbfqXBvEqxTEto8BTrNvN0GvT90W5JZMkX_YN0MvXg8gNEV4NC8xNEV4N03eN0B5tZlqNxTEto8BTrNvNeZVuJ10Kj_g2UB02J0mN0B5tZlqNCNEto8BTrNvNC7MTX502C8f2jmMQJ6mqF2fka1mqx6gqj0eN0B599M=&uipk=5&nbs=1&deadline=1602654858&gen=playurl&os=ks3bv&oi=2887174439&trid=a1b65cfeb14846d4948984ec1d523d4eB&platform=pc&upsig=b91000904be06e678517ac6a73a75f79&uparams=e,uipk,nbs,deadline,gen,os,oi,trid,platform&mid=0&orderid=0,1&logo=00000000"
)

func main() {
	resp, err := http.Get(PlayURL)

	fmt.Println(resp.Status, err)
}
