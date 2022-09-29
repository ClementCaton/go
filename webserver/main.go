// basic go api server
package main

import (
	"log"
	"net/http"
)

type Person struct {
	Name string
	Wins int
}

func main() {
}

players := []struct{
	name string
	wins int
}{
	{"Pepper", 20},
	{"Floyd", 10},
}


// returns the wins of a player on GET /players/{name}/wins
func PlayerServer(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello World"))
	case http.MethodPost:
		w.WriteHeader(http.StatusCreated)
		//increase the player's score
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
