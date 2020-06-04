package main

import (
	"log"
	"runtime"
	"time"
)

func a(c chan int) {
	i := 0
	for {
		log.Printf("отправляю %d", i)
		c <- i
		log.Printf("отправил %d", i)
		i++
		// time.Sleep(time.Second)
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.Println("->main")

	chan1 := make(chan int)
	defer close(chan1)

	go a(chan1)
	time.Sleep((time.Second))
	for i := 0; i < 10; i++ {
		log.Println("жду следующего значения")
		v := <-chan1
		log.Printf("забрал %d", v)
		// time.Sleep((time.Second))
	}
	log.Println("<-main")
	time.Sleep(time.Second * 5)
}
