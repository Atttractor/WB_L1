package main

import (
	"context"
	"fmt"
	"time"
)

func cClose1(c chan int) {
	for {
		v, ok := <-c
		if !ok {
			fmt.Println("Канал закрыт, выход из горутины")
			return
		}
		fmt.Printf(" got: %d\n", v)
	}
}

func cClose2(ch chan int) {
	for v := range ch {
		fmt.Println(v)
	}
	fmt.Println("Горутина прочитла все значения из канал и закончила работу")
}

func сClose3(ch chan int, stop chan struct{}) {
	for {
		select {
		case v := <-ch:
			fmt.Printf(" got: %d\n", v)
		case <-stop:
			fmt.Println("Получен сигнал завешения горутины через отдельный канал")
			return
		}
	}
}

func cClose4(ctx context.Context, ch <-chan int) {
	for {
		select {
		case v := <-ch:
			fmt.Printf(" got: %d\n", v)
		case <-ctx.Done():
			fmt.Println("Получен сигнал из контекста о завершении горутины")
			return
		}
	}
}

func cClose5(ctx context.Context, ch <-chan int) {
	for {
		select {
		case <-time.After(10 * time.Second):
			fmt.Println(<-ch)
		case <-ctx.Done():
			fmt.Println("Выход из горутины по таймауту")
			return
		}
	}
}

func main() {
}