package golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
)

func TestPool(t *testing.T) {
	pool := sync.Pool{
		New: func() interface{} {
			return "New"
		},
	}
	group := &sync.WaitGroup{}

	pool.Put("Styephen")
	pool.Put("William")
	pool.Put("Buli")

	for i := 0; i < 10; i++ {
		go func() {
			group.Add(1)
			data := pool.Get()
			fmt.Println(data)
			//time.Sleep(1 * time.Second)
			pool.Put(data)
			group.Done()
		}()
	}

	group.Wait()
	//time.Sleep(5 * time.Second)
	fmt.Println("Selesai")
}
