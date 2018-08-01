package main

import (
	"errors"
	"flag"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang/groupcache"
)

var localStorage map[string]string

func init() {
	localStorage = make(map[string]string)
	localStorage["hello"] = "world"
	localStorage["info"] = "This is an example"
}

func main() {
	port := flag.Int("port", 4100, "Listen port")
	flag.Parse()

	// Name have to starts with http://
	self := "http://localhost:" + strconv.Itoa(*port)
	pool := groupcache.NewHTTPPool(self)
	pool.Set(self, "http://localhost:4101")

	var helloworld = groupcache.NewGroup("helloworld", 10<<20, groupcache.GetterFunc(
		func(ctx groupcache.Context, key string, dest groupcache.Sink) error {
			log.Printf("groupcache get key: %v", key)
			value, exists := localStorage[key]
			if !exists {
				dest.SetString(key + " NotExist")
				return errors.New(key + " NotExist")
			} else {
				dest.SetString(value)
				return nil
			}
		}))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		key := strings.TrimPrefix(r.RequestURI, "/")
		log.Printf("Request(%v) key(%v)", r.RemoteAddr, key)
		if key == "" {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		var data []byte
		err := helloworld.Get(nil, key, groupcache.AllocatingByteSliceSink(&data))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		log.Printf("cache data: %v", data)
		w.Write(data)
		log.Println("Gets: ", helloworld.Stats.Gets.String())
		log.Println("CacheHits: ", helloworld.Stats.CacheHits.String())
		log.Println("Loads: ", helloworld.Stats.Loads.String())
		log.Println("LocalLoads: ", helloworld.Stats.LocalLoads.String())
		log.Println("PeerErrors: ", helloworld.Stats.PeerErrors.String())
		log.Println("PeerLoads: ", helloworld.Stats.PeerLoads.String())
	})
	http.ListenAndServe(":"+strconv.Itoa(*port), nil)
}
