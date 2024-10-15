package lesson_12

import "fmt"


type IArea interface {
	Area() float64
}

type Rectangle struct {
	length float64
	width float64
}

type Circle struct {

	radius float64
}

type Square struct {
	length float64
}

func generateArea (object IArea) float64 {
    return object.Area()
}

var mysquare = Square{length: 10.0}

var mycirle = Circle{radius: 7.0}

var myrectangle = Rectangle{length: 10.0, width: 7.0}


func (r Rectangle) Area() float64 {
	return r.length * r.width
}

func (c Circle ) Area () float64{
   return 3.142 * c.radius * c.radius
}

func (s *Square) Area() float64 {
	s.length = s.length * 2
	return s.length * s.length
}


func Main (){
	fmt.Println("Square area is", generateArea(&mysquare))
	fmt.Println("Modification of Square area is", mysquare.length)
	fmt.Println("Circle area is", generateArea(mycirle))	
}