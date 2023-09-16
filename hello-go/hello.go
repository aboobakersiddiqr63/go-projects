package main

import (
	"fmt"
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

func main() {
	fmt.Println("Welcome to loops in golang")

	days := []string{"Sunday", "Monday", "Tuesday"}
	fmt.Println(days)

	//Normal for loops
	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	//range loop
	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	//foreach type
	// for index, day := range days {
	// 	fmt.Printf("The index is %v and value is is %v\n", index, day)
	// }

	// do while type
	rogueValue := 1
	for rogueValue < 10 {
		if rogueValue == 5 {
			goto gotos
		}
		fmt.Println("Value is:", rogueValue)
		rogueValue++
	}

	// go-to
gotos:
	fmt.Println("Hello from go-to")
}
