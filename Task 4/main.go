package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(ctx context.Context, wg *sync.WaitGroup, input chan int) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Воркер завершил свою работу")
			return
		// Всё ломалось когда было так:
		//	default:
		//		fmt.Printf("Получено число %v\n", <-input)
 
		case n := <-input:
			fmt.Printf("Получено число %v\n", n)
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var w int
	flag.IntVar(&w, "w", 5, "number of lines to read from the file")
	flag.Parse()

	wg := &sync.WaitGroup{}

	ctx, cancel := context.WithCancel(context.Background())

	in := make(chan int, 2)
	defer close(in)

	for i := 0; i < w; i++ {
		go worker(ctx, wg, in)
	}
	wg.Add(w)

	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()

		wg.Add(1)
		defer wg.Done()

		for {
			select {
			case <-ctx.Done():
				fmt.Println("Завершение работы программы...")
				return
			case <-ticker.C:
				a := rand.Int()
				in <- a
				fmt.Printf("Отправляем данные: %v\n", a)
			}
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	fmt.Println("Получен сигнал остановки")
	cancel()
	wg.Wait()
}