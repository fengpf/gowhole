package wrr

import (
	"fmt"
	"strconv"
)

type server struct {
	Name   string
	Weight int
}
type client struct {
	Name string
}

func test() {
	var svrs []*server
	svr := &server{Name: "server1", Weight: 2}
	svrs = append(svrs, svr)
	svr = &server{Name: "server2", Weight: 4}
	svrs = append(svrs, svr)
	svr = &server{Name: "server3", Weight: 6}
	svrs = append(svrs, svr)
	svr = &server{Name: "server4", Weight: 8}
	svrs = append(svrs, svr)
	svr = &server{Name: "server4", Weight: 5}
	svrs = append(svrs, svr)
	cs := make([]*client, 0, len(svrs))
	weights := 0
	for i, svr := range svrs {
		weights += svr.Weight
		cli := &client{"client" + strconv.Itoa(i)}
		cs = append(cs, cli)
	}
	println(weights)
	wcs := make([]*client, 0, weights)
	for i, j := 0, 0; i < weights; j++ {
		idx := j % len(svrs)
		if svr := svrs[idx]; svr.Weight > 0 {
			i++
			svr.Weight--
			wcs = append(wcs, cs[idx])
		}
	}
	for _, w := range wcs {
		fmt.Printf("%v\n", w)
	}
}
