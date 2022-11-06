package Engine

import (
	"sync"
)

var wg sync.WaitGroup

// Handler creates the files
func Handler(channel <-chan string) {
	for op := range channel {
		Create(op) // Create file
		Reload(op) // Delete file
	}
}

// Producer starts the list and sends the filenames
func Producer(Channel chan<- string) {
	wg.Add(1)
	go func() {
		for i := range Files {
			// Check if files exist if not then its a good thing, create them, otherwise ignore
			if !ExistQ(Files[i]) || ExistQ(Files[i]) {
				Channel <- Files[i] // Send them to the channel
			}
		}
		wg.Done()
	}()
	wg.Wait()
	close(Channel)
}
