package lesson_eight

import (
	"fmt"
	"strings"
)


func MapOperations(){
	ageMap:= make(map[string]int)
	ageMap["John"] = 30
	ageMap["Jane"] = 25
	ageMap["Mike"] = 35
	json:= map[string]string{
		"name": "John",
		"age": "25",
		"city": "New York",
	}
	fmt.Println("Age of John:", ageMap["John"])
	delete(json, "city")
	for key, value := range json {
		fmt.Printf("%s: %s\n", key, value)
		if key == "age" {
			fmt.Println("Age found in JSON:", value)
		}
	}
}

func Processing (){
	person:= map[string]string{
		"name": "Alice",
		"age": "30",
		"city": "New York",
	}
	modifyPerson(person, "age", "31")
	for key, value := range person{
		fmt.Printf("%s: %s\n", key, value)
	}

}

func modifyPerson(person map[string]string, key string,value string) {
	fmt.Println("Original person:", person)
	person[key] = value
	fmt.Println("Modified person:", person)
}


func WordCount1 (){
	text := "This is a sample text. This text contains some words."
	text = strings.ToLower(text)
	wordCount:= make(map[string]int)
	words := strings.Fields(text)
	for _, word := range words {
		wordCount[word] = wordCount[word] + 1
	}
	fmt.Println(wordCount)
}

func WordCount2() {
	text := "This is a sample text. This text contains some words."

    // Create a map to store word counts
    wordCount := make(map[string]int)

    // Split the text into words and count them
    words := strings.Fields(text)
    for _, word := range words {
        wordCount[word]++
    }

    // Print the word frequencies
    for word, count := range wordCount {
        fmt.Printf("%s: %d\n", word, count)
    }
}