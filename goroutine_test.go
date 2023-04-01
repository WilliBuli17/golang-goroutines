package golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func TestCreateGoroutine(t *testing.T) {
	/*
		go di depan nama fungsi itu artinya memanggil function yang akan kita jalankan dalam goroutine
		function kita jalankan dalam goroutine, function tersebut akan berjalan secara asynchronous,
		artinya tidak akan ditunggu sampai function tersebut selesai

		jika menggunakan go routine pada fungsi yang menggembalikan/mereturn sebuah nilai maka return valuenya tidak bisa di tangkap
		sehingga go routine ini tidak di sarankan untuk digunakan pada fungsi yang mengembalikan sebuah nilai
	*/
	go RunHelloWorld()

	fmt.Println("Ups")

	time.Sleep(1 * time.Second) // ini gunanya untuk programnya gak mati dulu sebelum si fungsi asynchronous selesai di jalankan
}

func DisplayNumber(number int) {
	fmt.Println("Display", number)
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 100000; i++ {
		// hasilnya tidak akan berurutan karena fungsinya berjalan secara async. Selain itu, fungsi ini juga berjalan secara parallel dan concurency
		go DisplayNumber(i)
	}

	time.Sleep(10 * time.Second)
}
