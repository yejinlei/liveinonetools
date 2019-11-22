// An example simple server using the server and sipnet libraries.

package main

import (
	"fmt"
	"github.com/1lann/go-sip/server"
	"github.com/1lann/go-sip/sipnet"
)

func main() {
	listener, err := sipnet.Listen("0.0.0.0:5080")
	if err != nil {
		panic(err)
	}

	defer listener.Close()
	
	count := 1
	for {
		req, conn, err := listener.AcceptRequest()
		if err != nil {
			fmt.Println("accept request error:", err)
			continue
		}

		fmt.Println(count)
		count = count + 1

		switch req.Method {
		case sipnet.MethodRegister:
			server.HandleRegister(req, conn)
		case sipnet.MethodInvite:
			server.HandleInvite(req, conn)
		default:
			fmt.Println("unknown method:", req.Method)
		}
	}
}
