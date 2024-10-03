package main

import (
	"basics/lesson_one"
	"basics/lesson_two"
	"basics/lesson_three"
	"basics/lesson_four"
	"basics/lesson_five"
	"fmt"
)

func main() {
	lesson_one.VariablesandConstants()
	lesson_two.Conditionals()
	lesson_two.SwichStatement()
	lesson_two.SwichStatementWithNoCondition()
	lesson_two.Loops()
	lesson_two.SkipEvenIncludeOdd()
	lesson_three.DefferedOperations()
	lesson_three.OperationClenups()
	lesson_three.CausePanic()
	lesson_three.ExceptionHandlingUsingPanicandRecover()
	 output, error :=  lesson_three.ComputationWithErrorHandling(1,0)
	 if error != nil {
		fmt.Println("Error thrown:", error)
	 } else {
		fmt.Println("Computation perfomed successfully: Result is", output)
	 }

	lesson_four.DefferedOperations()
	lesson_four.ReadFile("nonexistent-file.txt")

	 var result = lesson_five.Add(20,20)
	 fmt.Println(result)
	 var result1 = lesson_five.Multiply(29,2)
	 fmt.Println(result1)
	 divide_comp, error := lesson_five.Divide(20.0,2)
	 fmt.Println(result)
	 if error != nil {
		fmt.Println("Error thrown:", error)
	 } else {
		fmt.Println("Computation perfomed successfully: Result is", divide_comp)
	 }

	 var sumOfArguments = lesson_five.VariadicFunction(1,2,3,4,4,5,5,6,7,8,)
	 fmt.Println(sumOfArguments, "Sum of Arguments")

	 // Adding two numbers 
	 sum:=lesson_five.FirstClassCitizen(204,20, lesson_five.Add)
	 fmt.Println("Sum", sum)
}
