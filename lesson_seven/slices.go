package lesson_seven

import (
	"fmt"
)


func Slice_Operations (){
	slice:=[]int { 1,2,3,4,5,6,7,8,9,10}
	fmt.Println("Slice", slice)
	fmt.Println("Length", len(slice))
	fmt.Println("Capacity", cap(slice))
	schools:= [3]string {"Allidina", "Bahria", "Beaconhouse"}
	fmt.Println("Array", schools)
	schools_slice:= schools[0:2]
	fmt.Println("Slice", schools_slice)
	schools_slice[0] = "AiK"
	fmt.Println("Slice", schools_slice)
	// Increase the schools
}

func Create_Slice_Using_Make (){
	numbers := [5]int {1,2,3,4,5}
	slice := make([]string , 3,5)
	fmt.Println("Slice", slice)
	fmt.Println("Length", len(slice))
	fmt.Println("Capacity", cap(slice))
	newSlice := make([]int, len(numbers))
	fmt.Println("New Slice", newSlice)
	fmt.Println("New Length", len(slice))
	fmt.Println("New Capacity", cap(slice))
}

func Appending_Slices (){
	slice := []string {"Allidina", "Bahria", "Beaconhouse", "Jomvu"}
	slice = append(slice, "AiK", "Karachi Grammar School")
	fmt.Println("Slice", slice)
	fmt.Println("Length", len(slice))
	fmt.Println("Capacity", cap(slice))

}

func CopyingSlices (){
	slice1:= []string {"Allidina", "Bahria", "Beaconhouse", "Jomvu"}
	sliece2 := make([]string, 4)
	fmt.Println("Slice 1", slice1)
	fmt.Println("Slice 2", sliece2)
	copy(sliece2, slice1)
	fmt.Println("Slice 1", slice1)
	fmt.Println("Slice 2", sliece2)
}