package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var n int
	flag.IntVar(&n, "n", 5, "seconds timeout")
	flag.Parse()

	ctx, err := context.WithTimeout(context.Background(), time.Duration(n) * time.Second)
	if err != nil {
		fmt.Println(err)
	}

	wg := &sync.WaitGroup{}

	input := make(chan int)
	wg.Add(1)
	go takeMsg(ctx, wg, input)

	wg.Add(1)
	go func() {
		defer wg.Done()
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				fmt.Println("Выход по таймауту")
				close(input)
				return
			case <-ticker.C:
				a := rand.Int()
				input <- a
				fmt.Printf("Отправлено сообщение: %v\n", a)
			}
		}
	}()

	wg.Wait()
}

func takeMsg(ctx context.Context, wg *sync.WaitGroup, msgChan chan int) {
	defer wg.Done()
	
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Завершение работы воркера")
			return
		default:
			fmt.Printf("Получено сообщение: %v\n", <-msgChan)
		}
	}
}
// cd Desktop/wb/wb_l1/task 4