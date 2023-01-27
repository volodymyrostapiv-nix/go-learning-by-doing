package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func main() {
	// Do read
	input, err := os.ReadFile("persistent.txt")
	if err != nil {
		fmt.Printf("Got error: %s!", err.Error())
	}

	var dictionary map[string]string
	err = json.Unmarshal(input, &dictionary)
	if err != nil {
		fmt.Printf("Got error: %s!", err.Error())
	}
	fmt.Println(dictionary)
	dictionary["last updated at"] = time.Now().Format("01-02-2006 15:04:05")

	jsonStr, err := json.Marshal(dictionary)
	if err != nil {
		fmt.Printf("Got error: %s!", err.Error())
	}

	//Do save
	fo, err := os.Create("persistent.txt")
	if err != nil {
		panic(err)
	}
	// close
	defer func() {
		if err := fo.Close(); err != nil {
			fmt.Printf("Got error: %s!", err.Error())
		}
	}()

	//Do write
	if _, err := fo.Write(jsonStr); err != nil {
		fmt.Printf("Got error: %s!", err.Error())
	}
}
