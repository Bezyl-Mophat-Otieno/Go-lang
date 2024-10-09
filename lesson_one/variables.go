package lesson_one

import "fmt"


func VariablesandConstants(){
	var a int = 10
	var b string = "golang"
	var c float32 = 4.17
	var d bool = true
	var schools []string = []string{"Yale", "MIT", "Stanford"};
	var (
		firstname, lastname = "John", "Doe"
	)

	// Enums

	const  (
		Admin = "admin"
		User = "user"
		Guest = "guest"
	)

	const (
		Session = iota
		Cookie
		Token

	)

	fmt.Println(Admin, User, Guest)
	fmt.Println(Session, Cookie, Token)

	for i , school := range schools {
		println(i, school)
	}

	fmt.Println("int: ", a, "string",b, "float",c, "bool",d)
	fmt.Println("firstname: ", firstname, "lastname: ", lastname)


	const CONNECTION_URL = "https://www.google.com"

}