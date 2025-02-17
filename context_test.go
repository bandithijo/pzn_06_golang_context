package main

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}

func TestContextWithValue(t *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")

	fmt.Println(contextA) // => context.Background
	fmt.Println(contextB) // => context.Background.WithValue(b, B)
	fmt.Println(contextC) // => context.Background.WithValue(c, C)
	fmt.Println(contextD) // => context.Background.WithValue(b, B).WithValue(d, D)
	fmt.Println(contextE) // => context.Background.WithValue(b, B).WithValue(e, E)
	fmt.Println(contextF) // => context.Background.WithValue(c, C).WithValue(f, F)

	fmt.Println(contextF.Value("f")) // dapat => F
	fmt.Println(contextF.Value("c")) // dapat milik parent => C
	fmt.Println(contextF.Value("b")) // tidak dapat, beda parent => nil
	fmt.Println(contextA.Value("b")) // tidak bisa mengambil data child => nil
}

func CreateCounter() chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)

		counter := 1

		for {
			destination <- counter
			counter++
		}
	}()

	return destination
}

func TestContextWithLeak(t *testing.T) {
	fmt.Println("Total Goroutine:", runtime.NumGoroutine())

	destination := CreateCounter()

	fmt.Println("Total Goroutine:", runtime.NumGoroutine())

	for n := range destination {
		fmt.Println("Counter", n)

		if n == 10 {
			break
		}
	}

	fmt.Println("Total Goroutine:", runtime.NumGoroutine())
}

func CreateCounterWithContext(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)

		counter := 1

		for {
			select {
			case <- ctx.Done():
				return
			default:
				destination <- counter
				counter++
			}
		}
	}()

	return destination
}

func TestContextWithCancel(t *testing.T) {
	fmt.Println("Total Goroutine:", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	destination := CreateCounterWithContext(ctx)

	fmt.Println("Total Goroutine:", runtime.NumGoroutine())

	for n := range destination {
		fmt.Println("Counter", n)

		if n == 10 {
			break
		}
	}

	cancel() // mengirim sinyal cancel ke context

	time.Sleep(2 * time.Second)

	fmt.Println("Total Goroutine:", runtime.NumGoroutine())
}
