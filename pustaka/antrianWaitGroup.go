package pustaka

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func JalanAntrianWG(){
	runtime.GOMAXPROCS(2) //memanipulasi/atur core processor jika goroutine yang dipake banyak
	var wg sync.WaitGroup

	wg.Add(3) //3 waitgroup

	go printMessageWG("GoRoutine 1", &wg)
	go printMessageWG("GoRoutine 2", &wg)
	go printMessageWG("GoRoutine 3", &wg)

	wg.Wait()
	fmt.Println("Semua GoRoutine selesai!")
}

func printMessageWG(message string, wg *sync.WaitGroup){
	defer wg.Done() //akan jalan terakhir
	for i := 1; i<=5; i++ {
		fmt.Println(message, i)
		time.Sleep(1*time.Second)
	}
}