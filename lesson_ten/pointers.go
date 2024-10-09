package lesson_ten

import "fmt"


type Person struct {
	Name string;
}

var mophat Person = Person{Name: "Bezyl Mophat"}


func PointerOperations () {
	var myage = 22
	intpointer:= &myage
	fmt.Println("Address of myage is", intpointer)
	fmt.Println("Address of myage is", &myage)
	var valueFromPointer int = *intpointer
	fmt.Println("Value of myage is from pointer", valueFromPointer)
	valueFromPointer = 23
	fmt.Println("Modified value from pointer",valueFromPointer)

	fmt.Printf("Value of %v before modification", mophat.Name)

	println()

	failed_modify(mophat,"John Doe")

	fmt.Printf("Value of %v after modification", mophat.Name)

	println()
	modify_name(&mophat,"John Doe")

	fmt.Printf("Successfull modification after passing the address of the struct, new value %v", mophat.Name)





	
}

func modify_name( person *Person, name string){
	person.Name = name
}

func failed_modify(person Person, name string){
	person.Name = name
}