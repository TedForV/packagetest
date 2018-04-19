package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func mg() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn)
		log.Println("done")
		done <- struct{}{}
	}()

	mustCopy(conn, os.Stdout)
	<-done
}

func mpl() {
	naturals := make(chan int)
	squares := make(chan int)

	//counter
	go func() {
		for x := 0; x < 100; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	//squarer
	go func() {
		//for {
		//	x, ok := <-naturals
		//	if !ok {
		//		break // channel was closed and drained
		//	}
		//	squares <- x * x
		//}
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	//for {
	//	fmt.Println(<-squares)
	//}
	for x := range squares {
		fmt.Println(x)
	}
}

func counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

func printer(int <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func mc() {
	naturals := make(chan int)
	squarers := make(chan int)
	go counter(naturals)
	go squarer(squarers, naturals)
	printer(squarers)
}

func mirroredQuery() string {
	responses := make(chan string, 3)
	go func() { responses <- request("asia.gopl.io") }()
	go func() { responses <- request("europe.gopl.io") }()
	go func() { responses <- request("americas.gopl.io") }()
	return <-responses // return the quickest response
}

func request(hostname string) (response string) {

}
