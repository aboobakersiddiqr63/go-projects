package main

import "fmt"

func main() {
	fmt.Println("Hello Go!")
	// variables()
	// convertVarType()
	// conditions(9)

	totalSum := getTotalValueOfTwoNumber(15, 20)
	println("Total value is:", totalSum)

	x, _ := getPoint(5, 6)
	fmt.Println(x)
}

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

func getTotalValueOfTwoNumber(val1 int, val2 int) int {
	return val1 + val2
}

func getPoint(x int, y int) (int, int) {
	return x, y
}

// Stick to this variable types mostly
// bool
// string
// int
// uint32
// byte
// rune
// float64
// complex128
