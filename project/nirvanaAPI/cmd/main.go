package main

import "github.com/caicloud/nirvana"

func main() {
	// cmd := config.NewDefaultNirvanaCommand()
	// cmd.EnablePlugin(&metrics.Option{Path: "/metrics"})
	// if err := cmd.Execute(nirvanaAPI.EchoDesc); err != nil {
	// 	log.Fatal(err)
	// }
	conf := nirvana.NewDefaultConfig()
	s := nirvana.NewServer(conf)
	s.Serve()
}
