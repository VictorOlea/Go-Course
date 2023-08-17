package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	// var name string
	// var age int

	//Scan
	// fmt.Print("Enter your name: ")
	// fmt.Scanf("%s", &name)
	// fmt.Println("Hello", name)

	// fmt.Print("Enter your age: ")
	// fmt.Scan(&age)
	// fmt.Println("Your age is ", age)

	//Scanln
	// fmt.Print("Enter your name: ")
	// fmt.Scanln(&name)
	// fmt.Println("Hello ", name)

	// fmt.Print("Enter your age: ")
	// fmt.Scanln(&age)
	// fmt.Println("Your age is ", age)

	// var anInt int = 5
	// var aFloat float64 = 42
	// sum := float64(anInt) + aFloat
	// fmt.Printf("Sum: %v, Type: %T \n", sum, sum)

	now := time.Now()
	fmt.Println("Current time is: ", now)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter some text: ")
	input, _ := reader.ReadString('\n')
	fmt.Println("You have entered: ", input)

}
