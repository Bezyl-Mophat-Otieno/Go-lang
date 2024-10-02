package lesson_two

import "fmt"


func Conditionals(){
	age:= 18;
	if age >= 18{
		fmt.Println("You can vote")

	}else if age >= 16 && age < 18 {
		fmt.Println("You can drive but not vote")

	}else {
		fmt.Println("You can't vote or drive")
	}
}

func SwichStatement(){
	var dayOfTheWeek = "Saturday";

	switch dayOfTheWeek {
		case "Saturday", "Sunday":
			fmt.Println("It's the weekend")
		case "Monday":
			fmt.Println("It's Monday, the worst day of the week")

		case "Tuesday", "Wednesday", "Thursday", "Friday":
			fmt.Println("It's a weekday, get back to work!")
		default:
			fmt.Println("It's friday, the day of the week")
	}
}

func SwichStatementWithNoCondition(){
	age:=18

	switch {
	case age >= 18:
			fmt.Println("You can vote")
	case age >= 16 && age < 18:
			fmt.Println("You can drive but not vote")

	case age < 16:
			fmt.Println("You can't vote or drive")
	default:
			fmt.Println("Go home")
	}
}

func Loops (){
	universities:= []string{"Stanford", "MIT", "Harvard", "Oxford", "Cambridge", "Kisii University"}
	for i:=0; i < len(universities); i++{
		fmt.Println(universities[i])
	} 

	for i, school:= range universities{
		fmt.Println(i, school)
	}
}

func SkipEvenIncludeOdd(){
	numbers:= []int{1,2,3,4,5,6,7,8,9}

	for i:=0; i<len(numbers); i++{
		if numbers[i] % 2 ==0 {
			continue
		}

		print(numbers[i])
		
		if i == 7{
			break
		}
		
	}
}
