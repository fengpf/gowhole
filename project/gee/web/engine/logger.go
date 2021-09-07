package engine

import (
	log "github.com/sirupsen/logrus"
	"time"
)

func Logger() HandlerFunc {
	return func(c *Context) {
		t := time.Now()

		c.Next()

		log.Printf("Logger [%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
