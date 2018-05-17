package olive

import (
	"fmt"

	"github.com/pebbe/zmq4"
)

func InitOlive() zmq4.Context {
	var context zmq4.Context
	return context
}

func CreateServer(ctx zmq4.Context) {

	request, err := ctx.NewSocket(zmq4.REQ)
	if err != nil {
		fmt.Println("Error creating request socket.")
	}
	reply, err := ctx.NewSocket(zmq4.REP)
	if err != nil {
		fmt.Println("Error creating reply socket.")
	}
	reply.Bind("oil/endpoint")
	request.Connect("oil/endpoint")
	request.SendBytes([]byte("hello"), zmq4.Flag(0))
	poller := zmq4.NewPoller()
	poller.Add(reply, 0)
	for {
		replyState, err := poller.Poll(100)
		if err != nil {
			fmt.Println("Timeout on polling.")
		}
		message, err := replyState[0].Socket.RecvMessage(0)
		fmt.Println(message)
	}
}
