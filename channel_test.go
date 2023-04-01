package golang_goroutines

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

// ----------------------------------------------------------------------------------------------------------------------
// contoh pembuatan chanel awal -- konsep
func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel) //pastikan selalu close channel agar tidak memory leak

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Willi Buli"
		fmt.Println("Selesai mengirim data ke channel")
	}()

	data := <-channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)

	// pastikan saat membuat channel ada yang mengirim dan menerima agar tidak deadlock
}

// ----------------------------------------------------------------------------------------------------------------------
// contoh pembuatan chanel as parameter
func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}

// pada pembuatan channel as parameter, itu tidak membutuhkan pointer. Karena dia sudah auto pass valuenya by reference ke chanel aslinya
func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Willi Buli"
	fmt.Println("Selesai mengirim data ke channel")
}

// ----------------------------------------------------------------------------------------------------------------------
// contoh pembuatan chanel in out
func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Willi Buli"
	fmt.Println("Selesai mengirim data ke channel")
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

// ----------------------------------------------------------------------------------------------------------------------
// contoh pembuatan buffered chanel
func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3) // penambahan buffer menghindari deadlock karena data akan masuk ke buffer/antrian. deadlock akan terjadi jika antrean full
	defer close(channel)

	go func() {
		channel <- "Willi"
		channel <- "Buli"
	}()

	fmt.Println("Kapasitas Buffer", cap(channel)) // malihat kapasitas buffer dari channel
	fmt.Println("Length Buffer", len(channel))    // melihat jumlah data antri atau data dalam buffer dari channel

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(5 * time.Second)
	fmt.Println("Selesai")
}

// ----------------------------------------------------------------------------------------------------------------------
// contoh range channel
func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan Ke-" + strconv.Itoa(i)
		}

		close(channel) // tidak bisa menggunakan defer - kena deadlock
	}()

	for data := range channel {
		fmt.Println(data)
	}

	fmt.Println("Selesai")
}

// ----------------------------------------------------------------------------------------------------------------------
// contoh select channel dan default channel
func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data channel 2", data)
			counter++
		default: // ini contoh default channel
			fmt.Println("Menunggu data")
		}

		// proses pengecekan jika data channel udah diambil semua maka perulangan for di stop
		if counter == 2 {
			break
		}
	}
}

// ----------------------------------------------------------------------------------------------------------------------
// contoh race condition pada channel
func TestRaceCondition(t *testing.T) {
	x := 0

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				x = x + 1
			}
		}()
	}

	time.Sleep(10 * time.Second)
	fmt.Println("Counter =", x)
}
