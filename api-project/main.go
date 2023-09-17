package main

import (
	"encoding/json"
	"fmt"
	"log"
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
	fmt.Println("API - Section")

	r := mux.NewRouter()

	player1 := Player{
		PlayerId: "1",
		Name:     "Ronaldo",
		Age:      37,
		Club:     &Club{ClubName: "Soccer Club", Website: "http://soccerclub.com"},
	}

	player2 := Player{
		PlayerId: "2",
		Name:     "Salah",
		Age:      37,
		Club:     &Club{ClubName: "Liverpool", Website: "http://liverpool.com"},
	}

	players = append(players, player1, player2)

	//routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/players", GetAllPlayers).Methods("GET")
	r.HandleFunc("/player/{id}", GetOnePlayer).Methods("GET")
	r.HandleFunc("/player", createAPlayer).Methods("POST")
	r.HandleFunc("/player/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/player/{id}", DeleteOneCourse).Methods("DELETE")

	//listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))
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
			_ = json.NewDecoder(r.Body).Decode(&player)
			players = append(players, player)
			json.NewEncoder(w).Encode(player)
		}
	}
}

func DeleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete A Course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	//loop, id ,remove, add my ID
	for index, player := range players {
		if player.PlayerId == params["id"] {
			players = append(players[:index], players[index+1:]...)
			message := fmt.Sprintln("The Player has been removed:", player.PlayerId)
			json.NewEncoder(w).Encode(message)
			break
		}
	}
}
