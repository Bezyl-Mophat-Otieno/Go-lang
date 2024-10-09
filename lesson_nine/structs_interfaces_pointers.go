package lesson_nine

import "fmt"

type Person struct {
	Name string;
	Age int8;
	School string;
}

func (p Person) Speak() {
	fmt.Printf("Speak has been called for %v of age %v", p.Name, p.Age)
}

func (p *Person) IncreamentAge(){
	p.Age = p.Age + 1
}

func GenerateDuplicatePeople (Name string, Age int8, School string) []Person {

	people:= []Person{
		{
		Name: Name,
		Age: Age,
		School: School,
	},
			{
		Name: Name,
		Age: Age,
		School: School,
	},

	}

	return people

}

func StructOperations (){

	person1:= Person{Name: "John", Age: 30, School: "New York"}
    var person2 Person;
	fmt.Println(person1)
	fmt.Println(person2)
	person1.Name = "Jane"
	person2.Name = "Mike"
	person2.Age = 40
	person1.IncreamentAge()
	person2.IncreamentAge()
	fmt.Println(person1)
	fmt.Println(person2)


	person1.Speak()
	fmt.Println()
	person2.Speak()

}

type Cat struct {
	Name string;
	Age int8;
}


func ( a Cat) makeSound(sound string){
	  fmt.Printf("%v %v",a.Name, sound)
}

type Dog struct {
	Name string;
	Age int8;
	speed string;

}

func ( a Dog) makeSound(sound string){
	  fmt.Printf("%v %v",a.Name, sound)
}

type IAnimal interface {
	makeSound(sound string)
}

func animal_sound ( animal IAnimal, sound string){
	animal.makeSound(sound)
}


func InterfaceOperations (){

	var mycat Cat = Cat{
		Name: "She",
		Age: 1,
	}
	var mydog Dog = Dog{
		Name: "Bruno",
		Age: 2,
		speed: "40km/h",
		
	}

	mycat.makeSound("meows")

	mydog.makeSound("barks")

	

	println()

	animal_sound(mycat, "meows")

	println()

	animal_sound(mydog, "barks")


}