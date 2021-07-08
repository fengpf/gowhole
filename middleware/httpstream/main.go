package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os/exec"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	data := make(url.Values)
	data.Set("code", r.FormValue("code"))
	data.Set("ip", r.RemoteAddr)

	resp, _ := http.PostForm("http://mppp.it/exists", data)
	code, _ := ioutil.ReadAll(resp.Body)
	ipresp, _ := http.PostForm("http://mppp.it/ip", data)

	ip, _ := ioutil.ReadAll(ipresp.Body)
	if strings.TrimSpace(string(code)) == "true" || strings.TrimSpace(string(ip)) == "true" {
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s.mp3", r.FormValue("name")))

		url, _ := exec.Command("youtube-dl", "-g", "--cookies", "cookie", r.FormValue("url")).Output()

		cmd := exec.Command("mencoder", "-msglevel", "all=-1", "-cookies", "-cookies-file", "cookie", strings.TrimSpace(string(url)), "-of", "rawaudio", "-ss", r.FormValue("stime"), "-endpos", r.FormValue("etime"), "-oac", "mp3lame", "-lameopts", "cbr:mode=2:br=192", "-mc", "0", "-noskip", "-ofps", "24000/1001", "-ovc", "copy", "-o", "-")
		reader, _ := cmd.StdoutPipe()
		cmd.Start()
		io.Copy(w, reader)
		cmd.Wait()
		cmd.Process.Kill()
		http.PostForm("http://mppp.it/convert", make(map[string][]string))
		//exec.Command("kill","-9",string(cmd.Process.Pid)).Run()
	} else {
		fmt.Fprintf(w, "Either your daily limit of 3 converts was reached, or your code expired!")
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":4567", nil)
}
