package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const reqUrl string = "http://localhost:3000"
const filePath = "./response.json"

func main() {

	fmt.Println("Hello Web Server from Go")

	// GetRequest()
	// PostRequestofJsonType()
	//PostRequestOfFormData()
	// EncodeJSON()
	DecodeJSON()
}

func GetRequest() {
	getResp, err := http.Get(reqUrl)

	ExceptionHandler(err)

	defer getResp.Body.Close()

	getResponseBytes, err := io.ReadAll(getResp.Body)

	ExceptionHandler(err)

	// Normal Conversion
	fmt.Println("Response String:", string(getResponseBytes))

	CreateFile(filePath, string(getResponseBytes), true)
}

func CreateFile(filePath string, response string, enablingWritingToFile bool) {
	file, err := os.Create(filePath)

	ExceptionHandler(err)

	if enablingWritingToFile {
		WriteRespToFile(file, response)
	}

}

func WriteRespToFile(file *os.File, response string) {
	io.WriteString(file, response)
}

func PostRequestofJsonType() {

	requestBody := strings.NewReader(`
	  {
		"name" : "Zabit Mohamed Sharipov"
	  }
	`)

	postResp, err := http.Post(reqUrl, "application/json", requestBody)

	ExceptionHandler(err)

	defer postResp.Body.Close()

	getResponseBytes, err := io.ReadAll(postResp.Body)

	ExceptionHandler(err)

	fmt.Println("Post Response String:", string(getResponseBytes))
}

func PostRequestOfFormData() {

	data := url.Values{}
	data.Add("firstName", "Khabib")
	data.Add("lastName", "Numagomedov")
	data.Add("email", "khabib@eaglefc.com")

	formPostResp, err := http.PostForm(reqUrl, data)
	ExceptionHandler(err)

	defer formPostResp.Body.Close()

	content, errBody := io.ReadAll(formPostResp.Body)

	ExceptionHandler(errBody)

	fmt.Println("Content:", string(content))

}

func ExceptionHandler(err error) {
	if err != nil {
		panic(err)
	}
}

// handling-json
type Player struct {
	Name     string `json:"playerName"`
	Wage     int
	IsActive bool
	Password string   `json:"-"`
	Tag      []string `json:"tags,omitempty"`
}

func EncodeJSON() {
	ufcPlayer := []Player{
		{"khabib", 20000, true, "Abc@123", []string{"Top Player"}},
		{"Khamzat", 10000, true, "def@123", []string{"Middle Weight"}},
		{"Zabit", 5000, true, "ghi@123", []string{}},
	}

	// fmt.Printf("UFC PLAYERS: %+v\n", ufcPlayer)

	// ufcPlayerJSON, errJSON := json.Marshal(ufcPlayer)
	ufcPlayerJSON, errJSON := json.MarshalIndent(ufcPlayer, "", "\t")

	ExceptionHandler(errJSON)

	fmt.Printf("%s\n", ufcPlayerJSON)
}

func DecodeJSON() {
	jsonData := []byte(`
		{
			"playerName": "khabib",
			"Wage": 20000,
			"IsActive": true,
			"tags": [
				"Top Player"
			]
		}
	`)

	var playerList Player

	isValidData := json.Valid(jsonData)

	if isValidData {
		fmt.Println("JSON is valid")
		json.Unmarshal(jsonData, &playerList)
	} else {
		fmt.Println("JSON is not valid")
	}

	fmt.Printf("%#v\n", playerList)

	//some cases
	var myPlayerData map[string]interface{}
	json.Unmarshal(jsonData, &myPlayerData)

	fmt.Printf("%#v\n", myPlayerData["playerName"])

}
