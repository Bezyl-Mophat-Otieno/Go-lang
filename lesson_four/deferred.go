package lesson_four

import (
	"fmt"
	"os"
);



func DefferedOperations(){
	defer fmt.Println("Deffered function executed as soon the function completed")
	fmt.Println("Risky operation started")
	panic("Panic thrown")
	fmt.Println("Operation Wouldn't be reached")
}

func ReadFile(filename string) {
    file, err := os.Open(filename)
    if err != nil {
        fmt.Println("Error opening file:", err)
        return 
    }
    defer func() {
    file.Close()
    fmt.Println("Releasing resources after the file was successfully opened")
     }()
     
     fmt.Println("File opened successfully")
    
     // Read the file contents.

     fmt.Println("Reading the file contents")

     // Something happens and an exception is raised 
     fmt.Println("Something is about to happen")

     panic("Reading failed to complete")

     // The file is closed automatically due to the deferred function     
	
}