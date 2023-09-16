package main

import (
	"fmt"
	"net/url"
)

// func main() {
// 	fmt.Println("Hello Go!")
// 	// variables()
// 	// convertVarType()
// 	// conditions(9)

// 	totalSum := getTotalValueOfTwoNumber(15, 20)
// 	println("Total value is:", totalSum)

// 	x, _ := getPoint(5, 6)
// 	fmt.Println(x)
// }

// func variables() {
// 	//hard coded
// 	var age int
// 	var markPercentage float32
// 	var hasPassed bool
// 	var studentName string

// 	//short variable
// 	username := "aboobakersiddiq"
// 	password := "helloWorld123@"
// 	isLoggedIn := true
// 	sampleFloat := 0.5

// 	employeeName, employeeId := "Kamar", 100

// 	fmt.Printf("%v %f %v %q\n", age, markPercentage, hasPassed, studentName)
// 	fmt.Printf("%v %T %v %T\n", username, password, isLoggedIn, sampleFloat)
// 	fmt.Printf("%v %v\n", employeeName, employeeId)

// }

// func convertVarType() {
// 	temparature := 78.66
// 	tempWholeNumber := int(temparature)

// 	fmt.Printf("The temperature is %v\n", tempWholeNumber)
// }

// func conditions(height int) {
// 	if height < 10 {
// 		fmt.Println("You are short")
// 	} else {
// 		fmt.Println("You are tall")
// 	}
// }

// func getTotalValueOfTwoNumber(val1 int, val2 int) int {
// 	return val1 + val2
// }

// func getPoint(x int, y int) (int, int) {
// 	return x, y
// }

// Stick to this variable types mostly
// bool
// string
// int
// uint32
// byte
// rune
// float64
// complex128

//user-input
// func main(){
// 	reader:= bufio.NewReader(os.Stdin)

// 	fmt.Println("Rate our Go code between 1 and 10")

// 	input, _ := reader.ReadString('\n')
// 	fmt.Println("Thanks for rating us:", input)
// 	fmt.Printf("Type of the rating is: %T\n", input)

// 	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

// 	if err != nil {
// 		fmt.Println("Error in conversion of input")
// 	} else{
// 		fmt.Printf("The converted Value is: %v\n", numRating)
// 	}
// }

//time-handling
// func main(){
// 	fmt.Println("Hello Time-Handling in Go")
// 	presentTime := time.Now()

// 	fmt.Printf("Current time is %v\n", presentTime)
// 	fmt.Println("Formatted Date 1:", presentTime.Format("01-02-2006"))
// 	fmt.Println("Formatted Date 2:", presentTime.Format("01-02-2006 Monday"))
// 	fmt.Println("Formatted Date 3:", presentTime.Format("01-02-2006 15:04:05 Monday"))

// 	createdDate := time.Date(2000, time.October, 19, 20, 0, 0, 0, time.Local)
// 	fmt.Println("My birthday is on: ", createdDate)
// }

// pointers
// func main() {
// 	fmt.Println("Hello Pointers")

// 	var ptr *int
// 	value := 2
// 	ptr = &value

// 	fmt.Println("The ptr address of memory location", ptr)
// 	fmt.Println("The ptr value:", *ptr)

// }

// Arrays
// func main() {
// 	fmt.Println("Hello Array")

// 	var players [2]string

// 	players[0] = "Ronaldo"
// 	players[1] = "salah"

// 	fmt.Println("Players:", players)
// 	fmt.Println("length of Players:", len(players))

// 	var fighters = [3]string{"Khabib", "khamzat", "Jon Jones"}
// 	fmt.Println("fighters list:", fighters)
// }

// Slices
// func main() {

// 	fmt.Println("Slices in Go ! Hello")

// 	var fighters = []string{"khabib", "khamzat Chimaev"}
// 	fmt.Printf("fighters: %v\n", fighters)

// 	fighters = append(fighters, "Zabit MogamedSharipov", "Islam Makhachev", "zubaira tukughov")
// 	fmt.Printf("fighters after addition: %v\n", fighters)

// 	// slicing the slice
// 	fighters = append(fighters[:1])
// 	fmt.Printf("fighters after slicing 1: %v\n", fighters)

// 	// fighters = append(fighters[1:3])
// 	// fmt.Printf("fighters after slicing 2: %v\n", fighters)

// 	// marks := make([]int, 1)
// 	// marks = append(marks, 100, 200, 300, 121, 453, 213, 321)
// 	// fmt.Println("Marks:", marks)

// 	//sort
// 	// sort.Ints(marks)
// 	// fmt.Println("Sorted:", marks)

// 	//remove a value from slices based on index
// 	var languages = []string{"C#", "Go", "C", "javaScript", "TypeScript", "Ruby"}

// 	index := 2
// 	languages = append(languages[:index], languages[index+1:]...)
// 	fmt.Println("After string removed by passing index:", languages)
// }

// maps
// func main() {

// 	fmt.Println("Hello Maps in Go !")

// 	languages := make(map[string]string)

// 	languages["JS"] = "JavaScript"
// 	languages["GO"] = "Golang"
// 	languages["PY"] = "Python"

// 	fmt.Println("Maped Value: ", languages)

// 	delete(languages, "PY")
// 	fmt.Println("Maped Value After deletion: ", languages)

// 	for key, value := range languages {
// 		fmt.Printf("key: %v, value: %v\n", key, value)
// 	}

// }

// Structs
// func main() {
// 	fmt.Println("Hello Structs Go")

// 	Saif := Student{"Saif", 1, true, 12}
// 	fmt.Printf("Student Detail: %+v", Saif)
// 	//no inheritance
// }

// type Student struct {
// 	Name           string
// 	RollNo         int
// 	IsPresentToday bool
// 	Class          int
// }

// if-else
// func main() {

// 	fmt.Println("Hello If-else in Go")

// 	var result string
// 	PlayerCount := 10

// 	if PlayerCount < 10 {
// 		result = "All players are present"
// 		fmt.Println(result)
// 	} else {
// 		result = "Some Players are not present"
// 		fmt.Println(result)
// 	}

// 	if PlayerCount%2 == 0 {
// 		fmt.Println("Even numebr")
// 	} else {
// 		fmt.Println("Odd number")
// 	}

// 	if num := 4; num < 10 {
// 		fmt.Println("Number is less than 10")
// 	} else {
// 		fmt.Println("Number is greater than 10")
// 	}

// }

// Switch-Case
// func main() {
// 	fmt.Println("Hello Switch-Case Go")

// 	rand.Seed(time.Now().UnixNano())
// 	diceNumber := rand.Intn(6) + 1
// 	fmt.Println("The Value of the dice is", diceNumber)

// 	switch diceNumber {
// 	case 1:
// 		fmt.Println("Dice value is 1 and you can open")
// 	case 2:
// 		fmt.Println("You can move 2 spot")
// 	case 3:
// 		fmt.Println("You can move 3 spot")
// 	case 4:
// 		fmt.Println("You can move 4 spot")
// 		fallthrough //move to the next value also
// 	case 5:
// 		fmt.Println("You can move 5 spot")
// 	case 6:
// 		fmt.Println("You can move 6 spot")
// 	default:
// 		fmt.Println("What was that!")
// 	}
// }

// loops
// func main() {
// 	fmt.Println("Welcome to loops in golang")

// 	days := []string{"Sunday", "Monday", "Tuesday"}
// 	fmt.Println(days)

// 	//Normal for loops
// 	// for i := 0; i < len(days); i++ {
// 	// 	fmt.Println(days[i])
// 	// }

// 	//range loop
// 	// for i := range days {
// 	// 	fmt.Println(days[i])
// 	// }

// 	//foreach type
// 	// for index, day := range days {
// 	// 	fmt.Printf("The index is %v and value is is %v\n", index, day)
// 	// }

// 	// do while type
// 	rogueValue := 1
// 	for rogueValue < 10 {
// 		if rogueValue == 5 {
// 			goto gotos
// 		}
// 		fmt.Println("Value is:", rogueValue)
// 		rogueValue++
// 	}

// 	// go-to
// gotos:
// 	fmt.Println("Hello from go-to")
// }

// functions
// func main() {

// 	fmt.Println("Hello Function in Go")

// 	resp := greeter()

// 	calcResult := SimpleCalculator(2, 4)

// 	proCalcultor(1, 2, 3, 4, 5)

// 	response, respStr := proCalcultorReturn(1, 2, 3, 4, 5)

// 	fmt.Println(respStr, ":", response)

// 	fmt.Println("Response from greeter method:", resp)

// 	fmt.Println("Response from SimpleCalculator method:", calcResult)

// }

// func greeter() string {
// 	greeter := "Hello from Greeter"
// 	return greeter
// }

// func SimpleCalculator(num1 int, num2 int) int {
// 	return num1 + num2
// }

// func proCalcultor(values ...int) {
// 	total := 0

// 	for _, value := range values {
// 		total += value
// 	}
// 	fmt.Println("Total value of all the values:", total)
// }

// func proCalcultorReturn(values ...int) (int, string) {
// 	total := 0

// 	for _, value := range values {
// 		total += value
// 	}
// 	return total, "Response from proCalculator"
// }

// Methods
// func main() {

// 	fmt.Println("Hello Method from Go")
// 	carDetail := Cars{"Honda", 2000, 1000, true}
// 	Cars.GetCarDetails(Cars{"Honda", 2000, 1000, true}, 2000)
// 	CarDetailsFunc(carDetail)
// }

// type Cars struct {
// 	Model       string
// 	Year        int
// 	HP          int
// 	IsAvailable bool
// }

// func (cars Cars) GetCarDetails(taxes int) {
// 	fmt.Println("Is Car available:", cars.IsAvailable)
// 	fmt.Println("Car tax:", taxes)
// 	cars.GetCarHP()
// }

// func (cars Cars) GetCarHP() {
// 	fmt.Println("The Cars HP is:", cars.HP)
// }

// func CarDetailsFunc(cars Cars) {
// 	fmt.Println("Is Car available from CarDetailsFunc:", cars.IsAvailable)
// }

// Defer Statement
// func main() {

// 	fmt.Println("Hello from defer Statements ")
// 	DeferFunc()

// }

// func DeferFunc() {
// 	// defer fmt.Println("Hello1")
// 	// defer fmt.Println("Hello2")

// 	// fmt.Println("World")
// 	//The data will not be printed directly where the resp data will be add to the stack and then printed finally
// 	for i := 0; i <= 5; i++ {
// 		defer fmt.Println("Hello", i)
// 	}
// }

// files-Handling
// func main() {

// 	fmt.Println("Hello file handling")

// 	filePath := "./file-handler.txt"
// 	CreateFile(filePath)
// 	ReadFromFile(filePath)
// }

// func CreateFile(filePath string) {
// 	content := "Hello World: I wanted this file to go into a file"

// 	file, err := os.Create(filePath)

// 	if err != nil {
// 		panic(err)
// 	}

// 	fileLen, errFileCreation := io.WriteString(file, content)

// 	if errFileCreation != nil {
// 		panic(err)
// 	}
// 	fmt.Println("Length of the file is:", fileLen)
// 	defer file.Close()
// }

// func ReadFromFile(filePath string) {
// 	contentByte, errBytes := os.ReadFile(filePath)

// 	if errBytes != nil {
// 		panic(errBytes)
// 	}

// 	fmt.Println("Content from the file in Bytes: \n", contentByte)

// 	fmt.Println("Content from the file: \n", string(contentByte))

// }

// web-request
// const url = "https://lco.dev"

// func main() {
// 	fmt.Println("LCO Web Request")

// 	response, err := http.Get(url)

// 	ExceptionHandler(err)

// 	if response.StatusCode == 200 {
// 		//fmt.Printf("Response is of Type %T\n", response)
// 		defer response.Body.Close()

// 		contentBytes, errBytes := io.ReadAll(response.Body)

// 		ExceptionHandler(errBytes)

// 		fmt.Println("Response from the Endpoint is:\n", string(contentBytes))
// 	}
// }

// func ExceptionHandler(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }

// handling-URL

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&courseid=25"

func main() {

	fmt.Println("Hello from handling-URL")

	ExtractingDataFromURL(myurl)

	CreatingURLFromData()

}

func ExtractingDataFromURL(myurl string) {

	fmt.Println(myurl)

	response, err := url.Parse(myurl)

	ExceptionHandler(err)

	// fmt.Println(response.RawQuery)

	qParams := response.Query()

	for key, val := range qParams {
		fmt.Printf("key: %v, Value: %v\n", key, val[0])

	}
}

func CreatingURLFromData() {

	partsofURL := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/players",
		RawPath: "player=khabib",
	}

	playerURL := partsofURL.String()

	fmt.Println("The Player URL is:", playerURL)
}

func ExceptionHandler(err error) {
	if err != nil {
		panic(err)
	}
}
