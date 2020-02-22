package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/user/{id}", getUser).Methods("GET")

	err:=http.ListenAndServe(":8080",
		csrf.Protect([]byte("32-byte-long-auth-key"))(r))

	if err!=nil{
		panic(err)
	}

	log.Println("server start")
}

func getUser(w http.ResponseWriter, r *http.Request) {
	// Authenticate the request, get the id from the route params,
	// and fetch the user from the DB, etc.
	// Get the token and pass it in the CSRF header. Our JSON-speaking client
	// or JavaScript framework can now read the header and return the token in
	// in its own "X-CSRF-Token" request header on the subsequent POST.
	params := r.URL.Query()
	fmt.Println(params["id"])

	idStr := params["id"][0]
	fmt.Println(idStr)
	user := map[int]map[string]string{
		1: {
			"name": "tom",
		},
		2: {
			"name": "jack",
		},
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println(err)
		return
	}
	u, ok := user[id]
	if !ok {
		log.Println(err)
		return
	}

	w.Header().Set("X-CSRF-Token", csrf.Token(r))
	b, err := json.Marshal(u)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write(b)
}
