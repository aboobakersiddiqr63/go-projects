package mod

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("hello Mod")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome)
	log.Fatal(http.ListenAndServe(":4000", r))
}

func greeter() {
	fmt.Println("Hello MOD detailed lesson")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Golang server</h1>"))
}
