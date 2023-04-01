package golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
)

// contoh wait group
func RunAsynchronous(group *sync.WaitGroup, count int) {
	defer group.Done()

	group.Add(1)

	fmt.Println("Hello ", count)
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		go RunAsynchronous(group, i)
	}

	group.Wait()
	fmt.Println("Complete")
}

// ----------------------------------------------------------------------------------------------------------------------
