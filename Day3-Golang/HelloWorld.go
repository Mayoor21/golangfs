package main
import (
	"fmt"
)
type student struct {
	Name string
	USN float64
	Depart string
}

func main() {
	//message := "Hello World!"
	//a := 1
	//b := 2
	//add := a+b
	//fmt.Println(message)
	//fmt.Println("addition =",add)
	//ifelseDemo()
	//Greater()
	//forThreeDemo()
	//CondiDemo()
	//forPythonStyle()
	st := student{Name: "Student1", USN: 12.2, Depart: "CSE-B"}
	fmt.Println("Name:", st.Name, "\nUSN:", st.USN,"\nDepartment:",st.Depart)
}

func ifelseDemo() {
	var age int
	fmt.Println("Enter the age:")
	fmt.Scanln(&age)
	if age > 18 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Minor")
	}
}

func Greater(){
	var a,b int
	fmt.Println("Enter the numbers:");
	fmt.Scanln(&a,&b)
	if a>b {
		fmt.Println(a,"is greater.")
	} else if a==b {
		fmt.Println(a,"is eqaul to",b)
	} else {
		fmt.Println(b,"is greater.")
	}
}

func forThreeDemo() {
	sum := 0
	for i := 1; i < 5; i++ {
		sum+= i
	}

}

func CondiDemo() {
	n := 1
	for n < 5 {
		n *=2
	}
	fmt.Println(n)
}

func forPythonStyle() {
	strings := []string{"Hello","World","Golang","NIE"}
	for _, s := range strings {
		fmt.Println(s)
	}
}