package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)


func main() {
	if len(os.Args) < 5 {
		log.Fatal("Usage: ", os.Args[0], " port serverAddr username peername")
	}
	__Vport := fmt.Sprintf(":%s", os.Args[1])
	__VserverAddr := os.Args[2]
	__Vusername := os.Args[3]
	__Vpeer := os.Args[4]
	__Vbuf := make([]byte, 2048)

	// Prepare to register user to server.
	__Vsaddr, __Verr := net.ResolveUDPAddr("udp4", __VserverAddr)
	if __Verr != nil {
		log.Print("Resolve server address failed.")
		log.Fatal(__Verr)
	}

	// Prepare for local listening.
	__Vaddr, __Verr := net.ResolveUDPAddr("udp4", __Vport)
	if __Verr != nil {
		log.Print("Resolve local address failed.")
		log.Fatal(__Verr)
	}
	__Vconn, __Verr := net.ListenUDP("udp", __Vaddr)
	if __Verr != nil {
		log.Print("Listen UDP failed.")
		log.Fatal(__Verr)
	}

	// Send registration information to server.
	initChatRequest := _TchatRequest{
		"New",
		__Vusername,
		"",
	}
	__VjsonRequest, __Verr := json.Marshal(initChatRequest)
	if __Verr != nil {
		log.Print("Marshal Register information failed.")
		log.Fatal(__Verr)
	}
	_, __Verr = __Vconn.WriteToUDP(__VjsonRequest, __Vsaddr)
	if __Verr != nil {
		log.Fatal(__Verr)
	}

	log.Print("Waiting for server response...")
	_, _, __Verr = __Vconn.ReadFromUDP(__Vbuf)
	if __Verr != nil {
		log.Print("Register to server failed.")
		log.Fatal(__Verr)
	}

	// Send connect request to server
	__VconnectChatRequest := _TchatRequest{
		"Get",
		__Vusername,
		__Vpeer,
	}
	__VjsonRequest, __Verr = json.Marshal(__VconnectChatRequest)
	if __Verr != nil {
		log.Print("Marshal connection information failed.")
		log.Fatal(__Verr)
	}

	var __VserverResponse _TchatRequest
	for i := 0; i < 3; i++ {
		__Vconn.WriteToUDP(__VjsonRequest, __Vsaddr)
		__Vn, _, __Verr := __Vconn.ReadFromUDP(__Vbuf)
		if __Verr != nil {
			log.Print("Get peer address from server failed.")
			log.Fatal(__Verr)
		}
		__Verr = json.Unmarshal(__Vbuf[:__Vn], &__VserverResponse)
		if __Verr != nil {
			log.Print("Unmarshal server response failed.")
			log.Fatal(__Verr)
		}
		if __VserverResponse.Message != "" {
			break
		}
		time.Sleep(10 * time.Second)
	}
	if __VserverResponse.Message == "" {
		log.Fatal("Cannot get peer's address")
	}
	log.Print("Peer address: ", __VserverResponse.Message)
	__VpeerAddr, __Verr := net.ResolveUDPAddr("udp4", __VserverResponse.Message)
	if __Verr != nil {
		log.Print("Resolve peer address failed.")
		log.Fatal(__Verr)
	}

	// Start chatting.
	go _Flisten(__Vconn)
	for {
		fmt.Print("Input message: ")
		__Vmessage := make([]byte, 2048)
		fmt.Scanln(&__Vmessage)
		messageRequest := _TchatRequest{
			"Chat",
			__Vusername,
			string(__Vmessage),
		}
		__VjsonRequest, __Verr = json.Marshal(messageRequest)
		if __Verr != nil {
			log.Print("Error: ", __Verr)
			continue
		}
		__Vconn.WriteToUDP(__VjsonRequest, __VpeerAddr)
	}
} // main

func _Flisten(___Vconn *net.UDPConn) {
	for {
		__Vbuf := make([]byte, 2048)
		__Vn, _, __Verr := ___Vconn.ReadFromUDP(__Vbuf)
		if __Verr != nil {
			log.Print(__Verr)
			continue
		}
		// log.Print("Message from ", __Vaddr.IP)

		var __Vmessage _TchatRequest
		__Verr = json.Unmarshal(__Vbuf[:__Vn], &__Vmessage)
		if __Verr != nil {
			log.Print(__Verr)
			continue
		}
		fmt.Println(__Vmessage.Username, ":", __Vmessage.Message)
	}
} // _Flisten
