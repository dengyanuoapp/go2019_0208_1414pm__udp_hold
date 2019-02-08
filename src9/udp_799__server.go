package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
)

var _VuserIP map[string]string

func main() {
	_VuserIP = map[string]string{}
	__Vservice := ":9999"
	__VudpAddr, __Verr := net.ResolveUDPAddr("udp4", __Vservice)
        if __Verr != nil {
                log.Fatal(__Verr)
        }

	__Vconn, __Verr := net.ListenUDP("udp", __VudpAddr)
        if __Verr != nil {
                log.Fatal(__Verr)
        }

	for {
		_FhandleClient(__Vconn)
	}
} // main

func _FhandleClient(___Vconn *net.UDPConn) {
	var __Vbuf [2048]byte

	__Vn, addr, __Verr := ___Vconn.ReadFromUDP(__Vbuf[0:])
	if __Verr != nil {
		return
	}

	var __VchatRequest _TchatRequest
	__Verr = json.Unmarshal(__Vbuf[:__Vn], &__VchatRequest)
	if __Verr != nil {
		log.Print(__Verr)
		return
	}

	switch __VchatRequest.Action {
	case "New":
		__VremoteAddr := fmt.Sprintf("%s:%d", addr.IP, addr.Port)
		fmt.Println(__VremoteAddr, "connecting")
		_VuserIP[__VchatRequest.Username] = __VremoteAddr

		// Send message back
		__VmessageRequest := _TchatRequest{
			"Chat",
			__VchatRequest.Username,
			__VremoteAddr,
		}
		__VjsonRequest, __Verr := json.Marshal(&__VmessageRequest)
		if __Verr != nil {
			log.Print(__Verr)
			break
		}
		___Vconn.WriteToUDP(__VjsonRequest, addr)
	case "Get":
		// Send message back
                peerAddr := ""
                if _, ok := _VuserIP[__VchatRequest.Message]; ok {
                        peerAddr = _VuserIP[__VchatRequest.Message]
                }

		__VmessageRequest := _TchatRequest{
			"Chat",
			__VchatRequest.Username,
                        peerAddr,
		}
		__VjsonRequest, __Verr := json.Marshal(&__VmessageRequest)
		if __Verr != nil {
			log.Print(__Verr)
			break
		}
		_, __Verr = ___Vconn.WriteToUDP(__VjsonRequest, addr)
		if __Verr != nil {
			log.Print(__Verr)
		}
	}
	fmt.Println("User table:", _VuserIP)
} // _FhandleClient
