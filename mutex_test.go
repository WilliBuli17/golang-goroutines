package golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// contoh sync mutex
func TestMutex(t *testing.T) {
	x := 0
	var mutex sync.Mutex // ini contoh sync mutex untuk mengindari race condition

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock() // ini contoh sync mutex untuk mengindari race condition -- aksi kunci
				x = x + 1
				mutex.Unlock() // ini contoh sync mutex untuk mengindari race condition -- aksi buka kunci
			}
		}()
	}

	time.Sleep(10 * time.Second)
	fmt.Println("Counter =", x)
}

// ----------------------------------------------------------------------------------------------------------------------
// contoh rw mutex
type BancAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BancAccount) AddBalance(amount int) {
	account.RWMutex.Lock() // Untuk kunci proses write/merubah/menambah data
	account.Balance = account.Balance + amount
	account.RWMutex.Unlock() // Untuk buka kunci proses write/merubah/menambah data
}

func (account *BancAccount) GetBalance() int {
	account.RWMutex.RLock() // Untuk kunci proses read/baca data
	balance := account.Balance
	account.RWMutex.RUnlock() // Untuk buka kunci proses read/baca data

	return balance
}

func TestReadWriteMutex(t *testing.T) {
	account := BancAccount{}

	for i := 1; i <= 100; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}

	time.Sleep(10 * time.Second)
	fmt.Println("Finnal Balance =", account.GetBalance())
}

// ----------------------------------------------------------------------------------------------------------------------
// contoh deadlock pada mutex, dimaana akdinya tidak bisa jalan sampai selesai karena ada aksi yang saling melock data, sehingga akisnya stuck
type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock() // Untuk kunci proses data
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock() // Untuk buka kunci proses  data
}

func (user *UserBalance) Change(amount int) {
	user.Balance = user.Balance + amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("Lock", user1.Name)
	user1.Change(-amount)

	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Lock", user2.Name)
	user2.Change(amount)

	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()
}

func TestDeadlock(t *testing.T) {
	user1 := UserBalance{
		Name:    "Willi",
		Balance: 1000000,
	}

	user2 := UserBalance{
		Name:    "Buli",
		Balance: 1000000,
	}

	go Transfer(&user1, &user2, 100000)
	go Transfer(&user2, &user1, 200000)

	time.Sleep(10 * time.Second)

	fmt.Println("User", user1.Name, ", Balance ", user1.Balance)
	fmt.Println("User", user2.Name, ", Balance ", user2.Balance)
}

// ----------------------------------------------------------------------------------------------------------------------
