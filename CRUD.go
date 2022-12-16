package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type PlayerDetails struct {
	ID       int     `json:"ID"`
	Name     string  `json:"Name"`
	Height   float64 `json:"Height"`
	Position string  `json:"Position"`
	MVP      int     `json:"MVP"`
}

var Player []PlayerDetails

func getPlayers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/JSON")
	json.NewEncoder(w).Encode(Player)
}

func getPlayer(w http.ResponseWriter, r *http.Request) {

}

func deletePlayers(w http.ResponseWriter, r *http.Request) {

}

func updatePlayers(w http.ResponseWriter, r *http.Request) {

}

func createPlayers(w http.ResponseWriter, r *http.Request) {

}
func main() {

	Player = append(Player, PlayerDetails{
		ID:       1,
		Name:     "LeBron James",
		Height:   6.8,
		Position: "Small Forward",
		MVP:      4,
	})

	Player = append(Player, PlayerDetails{
		ID:       2,
		Name:     "Michael Jordan",
		Height:   6.6,
		Position: "Shooting Guard",
		MVP:      5,
	})

	Player = append(Player, PlayerDetails{
		ID:       3,
		Name:     "Kareem Abdul-Jabbar",
		Height:   7.2,
		Position: "Center",
		MVP:      6,
	})

	Player = append(Player, PlayerDetails{
		ID:       4,
		Name:     "Giannis Antetokounmpo",
		Height:   6.11,
		Position: "Power Forward",
		MVP:      2,
	})

	Player = append(Player, PlayerDetails{
		ID:       5,
		Name:     "Kevin Durant",
		Height:   6.9,
		Position: "Small Forward",
		MVP:      2,
	})

	Player = append(Player, PlayerDetails{
		ID:       6,
		Name:     "Stephen Curry",
		Height:   6.3,
		Position: "Point Guard",
		MVP:      3,
	})

	Player = append(Player, PlayerDetails{
		ID:       7,
		Name:     "Kobe Bryant",
		Height:   6.6,
		Position: "Shooting Guard",
		MVP:      5,
	})

	Player = append(Player, PlayerDetails{
		ID:       8,
		Name:     "Luka Doncic",
		Height:   6.7,
		Position: "Point Guard",
		MVP:      0,
	})

	r := mux.NewRouter()
	r.HandleFunc("/Players", getPlayers).Methods("GET")
	r.HandleFunc("/Players/{ID}", getPlayer).Methods("GET")
	r.HandleFunc("/Players", updatePlayers).Methods("PUT")
	r.HandleFunc("/Players", createPlayers).Methods("POST")
	r.HandleFunc("/Players", deletePlayers).Methods("DELETE")
	fmt.Printf("Starting server at port: 8040")
	log.Fatal(http.ListenAndServe(":8040", r))
}
