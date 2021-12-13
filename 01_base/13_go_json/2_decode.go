package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
)

func ExampleUnmarshal() {
	var jsonBlob = []byte(`[
	{"Name": "Platypus", "Order": "Monotremata"},
	{"Name": "Quoll",    "Order": "Dasyuromorphia"}]`)
	type Animal struct {
		Name  string
		Order string
	}
	var animals []Animal
	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", animals)
	// Output:
	// [{Name:Platypus Order:Monotremata} {Name:Quoll Order:Dasyuromorphia}]
}

func ExampleDecoder() {
	const jsonStream = `
	{"Name": "Ed", "Text": "Knock knock."}
	{"Name": "Sam", "Text": "Who's there?"}
	{"Name": "Ed", "Text": "Go fmt."}
	{"Name": "Sam", "Text": "Go fmt who?"}
	{"Name": "Ed", "Text": "Go fmt yourself!"}`
	type Message struct {
		Name, Text string
	}
	dec := json.NewDecoder(strings.NewReader(jsonStream))
	for {
		var m Message
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s: %s\n", m.Name, m.Text)
	}
	// Output:
	// Ed: Knock knock.
	// Sam: Who's there?
	// Ed: Go fmt.
	// Sam: Go fmt who?
	// Ed: Go fmt yourself!
}

func main() {
	ExampleUnmarshal()
	fmt.Printf("\n\n")
	ExampleDecoder()
}
