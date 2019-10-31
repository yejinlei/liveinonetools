package main

import (
	"fmt"
	rtmp "github.com/zhangpeihao/gortmp"
	"os"
	"time"
)

var createStreamChan chan rtmp.OutboundStream

type TestOutboundConnHandler struct {
}

func (handler *TestOutboundConnHandler) OnStatus(conn rtmp.OutboundConn) {
	var err error
	status, err := conn.Status()
	fmt.Printf("@@@@@@@@@@@@@status: %d, err: %v\n", status, err)

}

func (handler *TestOutboundConnHandler) OnClosed(conn rtmp.Conn) {
	fmt.Printf("@@@@@@@@@@@@@Closed\n")

}

func (handler *TestOutboundConnHandler) OnReceived(conn rtmp.Conn, message *rtmp.Message) {
	fmt.Printf("@@@@@@@@@@@@@Received\n")

}

func (handler *TestOutboundConnHandler) OnReceivedRtmpCommand(conn rtmp.Conn, command *rtmp.Command) {
	fmt.Printf("ReceviedCommand: %+v\n", command)

}

func (handler *TestOutboundConnHandler) OnStreamCreated(conn rtmp.OutboundConn, stream rtmp.OutboundStream) {
	fmt.Printf("Stream created: %s\n", stream.ID())
	createStreamChan <- stream

}

func main() {
	var obConn rtmp.OutboundConn
	createStreamChan = make(chan rtmp.OutboundStream)
	testHandler := &TestOutboundConnHandler{}
	obConn, err := rtmp.Dial("rtmp://202.69.69.180:443/webcast/bshdlive-pc", testHandler, 1)
	defer obConn.Close()
	fmt.Println("Hello, World!")
	err = obConn.Connect()
	if err != nil {
		fmt.Printf("Connect error: %s", err.Error())
		os.Exit(-1)

	}

	for i := 1; i < 10; i++ {
		select {
		case stream := <-createStreamChan:
			err = stream.Play("test", nil, nil, nil)
			if err != nil {
				fmt.Printf("Play error: %s", err.Error())
				os.Exit(-1)
			} else {
				fmt.Println("Playing...")
			}

		case <-time.After(time.Second):
		}
	}

	time.Sleep(10 * 1000)
}
