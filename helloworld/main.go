package learn_helloworld

import "fmt"

// define a struct, there is no class to define
type User struct {
	// variable in struct must be capital

	Name string

	// memory pointer to string value, can be int value or boolean value
	Nickname *string

	Age     int
	Enabled bool

	// String array variable
	Interest []string
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
	user1.Interest = []string{"Jogging", "Gaming"}

	// memory pointer can be null
	user1.Nickname = nil

	// second way of initializing struct within method scope
	user2 := User{}
	user2.Name = "James"
	user2.Age = 30
	user2.Enabled = true
	user2.Interest = []string{}

	// point to memory that hold the value J, something like map where a key hold the value to the element
	var nickname = "J"
	user2.Nickname = &nickname
}
