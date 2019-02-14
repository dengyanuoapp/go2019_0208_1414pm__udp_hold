package main

import (
	"fmt"
	"net"
	"runtime"
)

func _Flisten(___Vidx int, ___Vconnection *net.UDPConn, ___Vquit chan struct{}) {

	fmt.Println("Trying to open ___Vidx ( ", ___Vidx, " )")

	buffer := make([]byte, 1024)
	n, remoteAddr, __Verr := 0, new(net.UDPAddr), error(nil)
	for __Verr == nil {
		n, remoteAddr, __Verr = ___Vconnection.ReadFromUDP(buffer)
		// you might copy out the contents of the packet here, to
		// `var r myapp.Request`, say, and `go handleRequest(r)` (or
		// send it down a channel) to free up the listening
		// goroutine. you do *need* to copy then, though,
		// because you've only made one buffer per _Flisten().
		fmt.Println("server ( ", ___Vidx, " ) , from ", remoteAddr, " - [", buffer[:n], "]")
	}
	fmt.Println("listener failed - (", ___Vidx, ") ", __Verr)
	___Vquit <- struct{}{}
}

func main() {
	__Vaddr := net.UDPAddr{
		Port: 5353,
		//IP:   net.IP{127, 0, 0, 1},
		//IP:   net.IP{0, 0, 0, 0},
	}
	__Vconnection, __Verr := net.ListenUDP("udp", &__Vaddr)
	if __Verr != nil {
		panic(__Verr)
	}
	__Vquit := make(chan struct{})
	for __Vi := 0; __Vi < runtime.NumCPU(); __Vi++ {
		fmt.Println("Trying to open thread ( ", __Vi, " )")
		go _Flisten(__Vi, __Vconnection, __Vquit)
	}

	fmt.Println("server : before want for hang until an error")

	<-__Vquit // hang until an error

	fmt.Println("server : before exit the main()")
}
