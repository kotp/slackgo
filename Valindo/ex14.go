package main

import (
	"io/ioutil"
	"log"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	//Read the contents of the file to "file" variable
	file, err := ioutil.ReadFile("open.txt")
	//Check for errors
	check(err)
	//variable to store position of the start of the "word"
	var position int
	// Word to be inserted
	word := " inserted"
	// convert the []byte File type to string
	str := string(file)
	//Iterate over string and save the index of "W" -1 in position
	for key, value := range str {
		// fmt.Println(key, " ", string(value)) //to debug
		variable := string(value)
		if variable == "w" {
			position = key - 1
		}
	}
	//add the word in between
	str1 := str[:position] + word + str[position:]
	//convert string to byte array
	byteConv := []byte(str1)
	//Write the modafuka
	err = ioutil.WriteFile("open.txt", byteConv, 0644)
}
