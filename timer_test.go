package golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)
	fmt.Println("Now,", time.Now())

	time := <-timer.C
	fmt.Println("After,", time)
}

func TestAfter(t *testing.T) {
	channel := time.After(5 * time.Second)
	fmt.Println("Now,", time.Now())

	time := <-channel
	fmt.Println("After,", time)
}

func TestAfterFunc(t *testing.T) {
	group := sync.WaitGroup{}
	group.Add(1)

	time.AfterFunc(5*time.Second, func() {
		fmt.Println("After,", time.Now())
		group.Done()
	})

	fmt.Println("Now,", time.Now())

	group.Wait()
}
