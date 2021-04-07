// Netcat1 is a read-only TCP client.
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	var port int
	flag.IntVar(&port, "p", 1234, "端口号，默认为1234")
	// 解析命令行参数写入注册的flag里
	flag.Parse()

	conn, err := net.Dial("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	_, err := io.Copy(dst, src)
	if  err != nil {
		log.Fatal(err)
	}
}