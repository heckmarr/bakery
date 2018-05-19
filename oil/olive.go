package olive

import (
	"fmt"

	"github.com/pebbe/zmq4"
	flour "gitlab.com/localtoast/flourPower"
)

func InitOlive() zmq4.Context {
	var context zmq4.Context
	return context
}

func PrepareMsg(testToast []flour.Bread) string {
	var message string
	for _, v := range testToast {
		mess := v.Label
		message += mess
	}
	return message
}

func CreateServer(testToast []flour.Bread) {
	//	ctx, err := zmq4.NewContext()
	//	if err != nil {
	//		fmt.Println("Error creating new zmq Context.")
	//	}
	request, err := zmq4.NewSocket(zmq4.REQ)
	if err != nil {
		fmt.Println("Error creating request socket.")
	}
	reply, err := zmq4.NewSocket(zmq4.REP)
	if err != nil {
		fmt.Println("Error creating reply socket.")
	}
	request.Connect("tcp://192.168.0.101:5555")
	reply.Bind("tcp://192.168.0.103:5555")

	for i := 0; i < 1; i++ {
		//message := olive.PrepareMsg(testToast)
		colourMessage := flour.PrepareToast(testToast, "red", "blue")
		//fmt.Println(message)
		request.SendMessage(colourMessage)
		fmt.Println("Message sent")
		mess, err := request.RecvMessage(zmq4.SNDMORE)
		//mess, err := reply.RecvMessage(zmq4.SNDMORE)
		fmt.Println(mess[0])
		if err != nil {
			fmt.Println("Timeout error.")
		}
	}
}
