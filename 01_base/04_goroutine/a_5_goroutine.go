package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

var Wait sync.WaitGroup
var Counter int = 0

func main() {
	for routine := 1; routine <= 2; routine++ {
		Wait.Add(1)
		time.Sleep(time.Second)
		log.Println("111111111")
		go Routine(routine)
	}
	Wait.Wait()
	fmt.Printf("======: %d\n", Counter)
}

func Routine(id int) {
	for count := 0; count < 2; count++ {
		value := Counter
		time.Sleep(time.Second)
		value++
		Counter = value
		time.Sleep(time.Second)
		log.Println("222222222 ", Counter)
	}
	time.Sleep(time.Second)
	Wait.Done()
}
