package main

import (
	"fmt"
	"time"
)

func threadPrintMessage(message string, done chan bool) {
	for i := 1; i <= 5; i++ {
		fmt.Println(message)
		time.Sleep(time.Millisecond) // Introduce a small delay to see simultaneous execution
	}
	done <- true // Signal that the goroutine has finished its work
}

func main() {
	done := make(chan bool)

	go threadPrintMessage("Hello from Goroutine 1", done)
	go threadPrintMessage("Hello from Goroutine 2", done)

	// Wait for both goroutines to finish
	// <-done
	<-done
}

// The time.Sleep function above was added with a small delay (in this case, time.Millisecond) to introduce a pause between each message printed by the goroutines. This is done to demonstrate simultaneous execution more clearly, as the actual execution of goroutines is extremely fast and may result in messages being printed too quickly to be observed distinctly.
// Without the delay, the output messages might appear to be printed simultaneously, but the terminal might not be able to display them all at once due to the speed of execution. By adding a small delay, you can see each message being printed in sequence, making it easier to perceive the concurrent behavior.
// The actual timing and output behavior might vary depending on the system and terminal used, but adding a delay helps to visualize the concurrent execution of goroutines more clearly in this specific example. In real-world scenarios, you may not need such a delay and can remove it to get the full performance benefit of concurrent execution.