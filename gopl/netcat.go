package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan string)
	go func() {
		io.Copy(os.Stdout, conn) // NOTE: ignoring errors
		log.Println("done")
		done <- "message" // signal the main goroutine
	}()
	mustCopy(conn, os.Stdin)
	conn.Close()
	x := <-done // wait for background goroutine to finish
	fmt.Println(x)
}

func mustCopy(dst io.Writer, src io.Reader) {
	_, err := io.Copy(dst, src)
	if  err != nil {
		log.Fatal(err)
	}
}