package lesson_five

import "errors"

func Add(a int, b int) int {
	return a + b
}

func Multiply(a int, b int) (result int) {
	result = a * b
	return result
}

func Divide(numerator float64, denominator int) (result float64, err error) {
	if denominator == 0 {
		return 0, errors.New("division by zero")
	}
	return numerator/float64(denominator) , nil
}

func VariadicFunction ( numbers ...int) (sum int ){
	sum = 0
	for _, number := range numbers {
		sum += number
	}
	return 
}

func FirstClassCitizen (a int , b int, operation func(int,int)int) int{
	return operation(a,b)
}

func FirstClassCitizenReturningFunction( multiplier int) func ( int, int)int {
	return func (a int, b int) int {
		return a*b*multiplier
	}
}
