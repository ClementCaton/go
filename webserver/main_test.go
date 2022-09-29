package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"strconv"
)

var expectedScores = map[string]int{
	"Pepper":  20,
	"Jhon":    20,
	"Brendon": 10,
	"Louis":   30}

func testGETPlayers(t *testing.T) {
	for name, score := range expectedScores {
	{
		request, _ := http.NewRequest(http.MethodGet, "/players/"+name+"/wins", nil)
		response := httptest.NewRecorder()

		PlayerServer(response, request)
		
		got := response.Body.String()
		want := strconv.Itoa(score)

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}		
}
