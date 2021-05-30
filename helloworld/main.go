package learn_helloworld

import "fmt"

// define a struct, there is no class to define
type User struct {
	// variable in struct must be capital
	Name    string
	Age     int
	Enabled bool
}

// define global variable
var (
	globalVar1 string = "Hello"
	globalVar2 string = "World"
)

func main() {

	// print hello world on console
	fmt.Println("Hello world")

	// print hello world on console with substitution
	fmt.Printf("%s %s\n", globalVar1, globalVar2)

	// first way of initializing struct within method scope
	var user1 User
	user1.Name = "Ben"
	user1.Age = 20
	user1.Enabled = true

	// second way of initializing struct within method scope
	user2 := User{}
	user2.Name = "James"
	user2.Age = 30
	user2.Enabled = true

}
