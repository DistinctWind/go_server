package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func handler(conn net.Conn) {
    defer conn.Close()
    reader:=bufio.NewReader(conn)
    writer:=bufio.NewWriter(conn)
    rw:=bufio.NewReadWriter(reader, writer)
    request_line, _:=rw.ReadString('\n')
    response:=fmt.Sprintf("Server got message: %s", request_line)
    rw.Write([]byte(response))
}

func main() {
    ln, err:=net.Listen("tcp", "0.0.0.0:8080")
    if err!=nil {
        panic(err)
    }
    
    for {
        conn, err:=ln.Accept()
        if err!=nil {
            log.Println("Can't have connection\n", conn)
        }

        go handler(conn)
    }
}
