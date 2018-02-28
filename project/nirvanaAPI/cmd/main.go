package main

import (
	"gowhole/project/nirvanaAPI"

	"github.com/caicloud/nirvana/config"
	"github.com/caicloud/nirvana/log"
	"github.com/caicloud/nirvana/plugins/metrics"
)

func main() {
	cmd := config.NewDefaultNirvanaCommand()
	cmd.EnablePlugin(&metrics.Option{Path: "/metrics"})
	if err := cmd.Execute(nirvanaAPI.EchoDesc); err != nil {
		log.Fatal(err)
	}
}
