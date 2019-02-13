package main

import (
        "fmt"
        "net"
        "runtime"
)

func listen(connection *net.UDPConn, quit chan struct{}) {
        buffer := make([]byte, 1024)
        n, remoteAddr, err := 0, new(net.UDPAddr), error(nil)
        for err == nil {
                n, remoteAddr, err = connection.ReadFromUDP(buffer)
                // you might copy out the contents of the packet here, to
                // `var r myapp.Request`, say, and `go handleRequest(r)` (or
                // send it down a channel) to free up the listening
                // goroutine. you do *need* to copy then, though,
                // because you've only made one buffer per listen().
                fmt.Println("from", remoteAddr, "-", buffer[:n])
        }
        fmt.Println("listener failed - ", err)
        quit <- struct{}{}
}

func main() {
        addr := net.UDPAddr{
                Port: 12345,
                IP:   net.IP{127, 0, 0, 1},
        }
        connection, err := net.ListenUDP("udp", &addr)
        if err != nil {
                panic(err)
        }
        quit := make(chan struct{})
        for i := 0; i < runtime.NumCPU(); i++ {
                go listen(connection, quit)
        }
        <-quit // hang until an error
}

