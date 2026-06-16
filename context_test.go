package belajar_golang_context

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
	// A sebagai Parent	
	contextA := context.Background()

	// B sebagai Parent
	contextB := context.WithValue(contextA, "b", "B")

	// C sebagai Parent
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")
	contextG := context.WithValue(contextC, "g", "G")

	// Memunculkan isi dari Context
	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)
	fmt.Println(contextG)

	// Mendapatkan value dari Context
	fmt.Println(contextF.Value("f"))
	fmt.Println(contextF.Value("c"))
	fmt.Println(contextF.Value("b"))
	fmt.Println(contextF.Value("d"))
}

// -- Problem Goroutine Leak -- 
// func CreateCounter() chan int {
// 	destination := make(chan int)

// 	go func ()  {
// 		defer close(destination)
// 		counter := 1
// 		for {
// 			destination  <- counter
// 			counter++
// 		}
// 	}()

// 	return destination 
// }

// func TestContextWithCancel(t *testing.T) {
// 	fmt.Println("Total Goroutine", runtime.NumGoroutine())

// 	destination := CreateCounter()
// 	for n := range destination {
// 		fmt.Println("Counter", n)
// 		if n == 10 {
// 			break
// 		}
// 	}

// 	fmt.Println("Total Goroutine", runtime.NumGoroutine())
// }

// func CreateCounter(ctx context.Context) chan int {
// 	destination := make(chan int)

// 	go func ()  {
// 		defer close(destination)
// 		counter := 1
// 		for {
// 			select {
// 			case <-ctx.Done():
// 				return
// 			default:
// 				destination  <- counter
// 				counter++
// 			}
			
// 		}
// 	}()

// 	return destination 
// }

// func TestContextWithCancel(t *testing.T) {
// 	fmt.Println("Total Goroutine", runtime.NumGoroutine())

// 	parent := context.Background()
// 	ctx, cancel := context.WithCancel(parent)

// 	destination := CreateCounter(ctx)

// 	fmt.Println("Total Goroutine", runtime.NumGoroutine())

// 	for n := range destination {
// 		fmt.Println("Counter", n)
// 		if n == 10 {
// 			break
// 		}
// 	}

// 	cancel() // Mengirim sinyal cancel ke context

// 	time.Sleep(2 * time.Second)

// 	fmt.Println("Total Goroutine", runtime.NumGoroutine())
// }

// -- Context dengan Timeout --
// func CreateCounter(ctx context.Context) chan int {
// 	destination := make(chan int)

// 	go func ()  {
// 		defer close(destination)
// 		counter := 1
// 		for {
// 			select {
// 			case <-ctx.Done():
// 				return
// 			default:
// 				destination  <- counter
// 				counter++
// 				time.Sleep(1 * time.Second) // Simulasi slow
// 			}
			
// 		}
// 	}()

// 	return destination 
// }

// func TestContextWithiTimeout(t *testing.T) {
// 	fmt.Println("Total Goroutine", runtime.NumGoroutine())

// 	parent := context.Background()
// 	ctx, cancel := context.WithTimeout(parent, 5 * time.Second)
// 	defer cancel()

// 	destination := CreateCounter(ctx)
// 	fmt.Println("Total Goroutine", runtime.NumGoroutine())

// 	for n := range destination {
// 		fmt.Println("Counter", n)
// 	}

// 	time.Sleep(2 * time.Second)

// 	fmt.Println("Total Goroutine", runtime.NumGoroutine())
// }

// -- Context dengan Deadline --
func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)

	go func ()  {
		defer close(destination)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination  <- counter
				counter++
				time.Sleep(1 * time.Second) // Simulasi slow
			}
			
		}
	}()

	return destination 
}

func TestContextWithDeadline(t *testing.T) {
	fmt.Println("Total Goroutine", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(10 * time.Second))
	defer cancel()

	destination := CreateCounter(ctx)
	fmt.Println("Total Goroutine", runtime.NumGoroutine())

	for n := range destination {
		fmt.Println("Counter", n)
	}

	time.Sleep(2 * time.Second)

	fmt.Println("Total Goroutine", runtime.NumGoroutine())
}