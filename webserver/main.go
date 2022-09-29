// basic go api server
package main

import (
	"net/http"
	"strconv"
	"strings"
)

type Person struct {
	Name string
	Wins int
}

func main() {
	http.HandleFunc("/players/", PlayerServer)
	http.ListenAndServe(":5000", nil)
}

// returns the wins of a player on GET /players/{name}/wins
func PlayerServer(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	println(player)

	players := []struct {
		name string
		wins int
	}{
		{"Pepper", 20},
		{"Chris", 10},
	}

	switch r.Method {
	case http.MethodGet:
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello World\n" + strconv.Itoa(players[0].wins)))
	case http.MethodPost:
		w.WriteHeader(http.StatusCreated)
		//increase the player's score

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
