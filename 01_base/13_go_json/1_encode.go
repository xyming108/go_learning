package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func ExampleMarshal() {
	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
	}
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)
	fmt.Println()
	fmt.Println(string(b))
	// Output:
	// {"ID":1,"Name":"Reds","Colors":["Crimson","Red","Ruby","Maroon"]}
}

func ExampleMarshalIndent() {
	/*data := map[string]int{
		"a": 1,
		"b": 2,
	}*/

	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
	}
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}

	//b, err := json.MarshalIndent(data, "<prefix>", "<indent>")
	b, err := json.MarshalIndent(group, "", "\t")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(b))
	// Output:
	// {
	// <prefix><indent>"a": 1,
	// <prefix><indent>"b": 2
	// <prefix>}

	/*
		{
		        "ID": 1,
		        "Name": "Reds",
		        "Colors": [
		                "Crimson",
		                "Red",
		                "Ruby",
		                "Maroon"
		        ]
		}
	*/
}

func main() {
	ExampleMarshal()
	fmt.Printf("\n\n")
	ExampleMarshalIndent()
}
