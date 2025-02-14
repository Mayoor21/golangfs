package main
import (
	"fmt"
)

func f1() {
	fmt.Println("This is beggining of f1 func")

	fmt.Println("This is end of f1 func")
}

func f2() {
	fmt.Println("This is beggining of f2 func")

	fmt.Println("This is end of f2 func")
}

func f3() {
	fmt.Println("This is beggining of f3 func")

	fmt.Println("This is end of f3 func")
}

func main() {
	fmt.Println("Start of main")
	f1()
	f2()
	f3()
	fmt.Println("End of main")
}