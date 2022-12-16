package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	//	"math/rand"
	//	"time"
)

//func generatePrice() float64 {
//	rand.Seed(time.Now().UnixNano())
//	return rand.Float64()*20 + 1
//}
//func sayHello(s *string) {
//	*s = "Bablu"
//}

// var Name string
// var Item string
// var Input int
// var myBill bill

func userInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}
func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	//	fmt.Println("Create a new bill name ")
	//	name, _ := reader.ReadString('\n')
	//	name = strings.TrimSpace(name)
	name, _ := userInput("Create a new bill Name ", reader)
	b := newbill(name)
	fmt.Println("Created bill for - ", b.name)
	return b
}
func main() {
	fmt.Println(createBill())
	//	fmt.Println("Welcome to McDonalds!!")
	//	fmt.Println("Whats your good name?")
	//	fmt.Scan(&Name)
	//	myBill = newbill(Name)
	//	for {
	//		fmt.Println("What Would you like to have")
	//		fmt.Scan(&Item)
	//		product := Item
	//
	//		myBill.addItem(product, generatePrice())
	//		fmt.Println("Press 1 to add more. Press 2 to exit.")
	//		fmt.Scan(&Input)
	//
	//		if Input == 1 {
	//			continue
	//		} else {
	//			break
	//		}
	//
	//	}
	//	myBill.updateBill(20)
	//	fmt.Println(myBill)
	//	fmt.Println(myBill.format())
}
