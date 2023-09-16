package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model for Player
type Player struct {
	PlayerId string `json:"courseid"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Club     *Club  `json:"club"`
}

type Club struct {
	ClubName string `json:"fullName"`
	Website  string `json:"website"`
}

// db
var players []Player

// middleware - helpers
func (p *Player) IsEmpty() bool {
	// return p.PlayerId == "" && p.Name == ""
	return p.Name == ""

}

func main() {
	players = append(players, Player{
		PlayerId: "1",
		Name:     "Ronaldo",
		Age:      37,
		Club:     &Club{ClubName: "Soccer Club", Website: "http://soccerclub.com"},
	})

	// Print the players
}

//controllers - file

// serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello Go-lang</h1>"))
}

// Get All Players
func GetAllPlayers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses has been called")
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(players)
}

func GetOnePlayer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	//grab id from req
	params := mux.Vars(r)

	//loop
	for _, player := range players {
		if player.PlayerId == params["id"] {
			json.NewEncoder(w).Encode(player)
			return
		}
	}

	message := fmt.Sprintf("No Player is found: %v", params["id"])

	json.NewEncoder(w).Encode(message)
}

func createAPlayer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create A Course")
	w.Header().Set("Content-Type", "application/json")

	//what-if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Input is Empty")
		return
	}

	//what-if: wrong-data
	var player Player
	_ = json.NewDecoder(r.Body).Decode(&player)

	if player.IsEmpty() {
		json.NewEncoder(w).Encode("Not a valid Data")
		return
	}

	// generate unique id, string
	// append course into courses

	rand.Seed(time.Now().UnixNano())
	player.PlayerId = strconv.Itoa(rand.Intn(100))

	players = append(players, player)
	json.NewEncoder(w).Encode(player)
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update A Course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	//loop, id ,remove, add my ID
	for index, player := range players {
		if player.PlayerId == params["id"] {
			players = append(players[:index], players[index+1:]...)
			player.PlayerId = params["id"]
			_ = json.NewDecoder(r.Body).Decode(&l)
			players = append(players, player)
			json.NewEncoder(w).Encode(player)
		}
	}

}
