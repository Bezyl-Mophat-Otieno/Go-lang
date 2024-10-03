package lesson_three

import (
	"fmt"
	"errors"
 )


func DefferedOperations (){
	defer fmt.Println("Deffered Execution 1, that acts as a cleanup");
	fmt.Println("Start Execution");
	fmt.Println("FinishExecution");
}

func OperationClenups (){
	age:= 20

	years:= []string{"1st", "2nd", "3rd"}

	for i := range years {age = age + (i) }

	fmt.Println("Age after", len(years), "year is", age)

	defer func (){
		age = 20;
		println("deffered execution used for an age reset: Age = ", age)
	}()

}

func ExceptionHandlingUsingPanicandRecover(){
	fmt.Println("Execution started...")

	
	defer func (){
		if r:= recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
		}()
	CausePanic()
		
	fmt.Println("Execution continues ...");
}

func CausePanic(){
	defer fmt.Println("Exception thrown")
	fmt.Println("About to cause a panic")
	panic("Something went wrong")
}

func ComputationWithErrorHandling( numerator int, denominator int ) (int, error) {
	defer fmt.Println("Deffered Execution 2, that acts as a cleanup");
	if denominator == 0 {
		return 0, errors.New("cannot divide by zero" )
	}
	return numerator / denominator, nil
}

