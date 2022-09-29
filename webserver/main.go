// basic go api server
package main

import (
	"net/http"
	"strconv"
	"strings"
)

var players = map[string]int{
	"Pepper":      20,
	"Jhon":        20,
	"Brendon":     10,
	"Louis":       30,
	"Harry":       40,
	"Zayn":        50,
	"Niall":       60,
	"Mia":         70, //Théo's sister
	"Taylor":      30,
	"Mary":        10,
	"Maris":       20,
	"Maryoumaris": 30,
	"Jhonny":      30,
	"Sins":        40,
}

func main() {
	http.HandleFunc("/players/", PlayerServer)
	http.ListenAndServe(":8080", nil)
}

// returns the wins of a player on GET /players/{name}/wins
func PlayerServer(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/players/")
	name := strings.TrimSuffix(path, "/wins")

	_, playerExists := players[name]

	switch r.Method {
	case http.MethodGet:
		w.WriteHeader(http.StatusOK)
		if playerExists {
			w.Write([]byte(strconv.Itoa(players[name])))
		} else {
			w.Write([]byte("This player does not exist"))
		}
	case http.MethodPost:
		w.WriteHeader(http.StatusCreated)
		//increase the player's score

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
