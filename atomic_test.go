package golang_goroutines

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

// ----------------------------------------------------------------------------------------------------------------------
// contoh atomic pada channel
func TestAtomic(t *testing.T) {
	var x int64 = 0
	group := sync.WaitGroup{}

	for i := 1; i <= 1000; i++ {
		go func() {
			group.Add(1)
			for j := 1; j <= 100; j++ {
				//x = x + 1
				atomic.AddInt64(&x, 1) // proses ini untuk mengganti proses diatas tanpa terkena race condition. Hanya berlaku pada tipe data primitif
			}
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Counter =", x)
}
