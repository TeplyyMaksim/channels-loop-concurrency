package main

import (
	"fmt"
	"sync"
)

func main() {
	channel := make(chan string)

	go launchGoRoutines(channel)

	func () {
		for value := range channel {
			fmt.Println(value)
		}
	}()

	fmt.Println("Coding finished")
}

func launchGoRoutines(channel chan <- string)  {
	var ws sync.WaitGroup
	routinesAmount := 10

	for i := 0; i < routinesAmount; i++ {
		ws.Add(1)
		go func(i int) {
			for j := 0; j < 10; j++ {
				routineAndIteration := fmt.Sprintf("%v-%v", i, j)
				channel <- routineAndIteration
			}
			ws.Done()
		}(i)
	}

	ws.Wait()
	close(channel)
}
