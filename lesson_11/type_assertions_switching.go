package lesson_11

import "fmt"

type Person struct{
   Name string;
   Age int;
}

type School struct{
	Name string;
	Address string;
}
func TypeAssertionsSwitching() {


	var person Person = Person{Name: "Mophat", Age: 30}

	var School School = School{Name: "New York", Address: "New York"}

	var name interface{} = "Bezyl Mophat";

	name,ok := name.(string)

	if ok {
		fmt.Println("Name is an integer", name)
	}else {
		fmt.Println("Name is not an integer")
	}

	checkType(person)

	checkType(School)

}


func checkType(data interface{}) {

	switch data.(type) {
		case int:
			fmt.Println("Type is int")
		case string:
			fmt.Println("Type is string")
		case Person:
			fmt.Println("Type is Person")
		case School:
			fmt.Println("Type is School")
		default:
			fmt.Println("Type is unknown")
	}

}