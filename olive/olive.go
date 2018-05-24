package olive

import (
	"bufio"
	"fmt"
	"os"

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

	request, err := zmq4.NewSocket(zmq4.REQ)
	if err != nil {
		fmt.Println("Error creating request socket.")
	}
	request.Connect("tcp://192.168.0.101:5555")
	//request.Connect("tcp://127.0.0.1:5555")
	file, err := os.Create("poptart/101/serve0.txt")
	defer file.Close()
	if err != nil {
		fmt.Println("Error creating message file.")
	}
	writer := bufio.NewWriter(file)
	for i := 0; i < 1; i++ {
		//message := olive.PrepareMsg(testToast)
		colourMessage := flour.PrepareToast(testToast, "green", "blue")
		//fmt.Println(message)
		request.SendMessage(colourMessage)
		//fmt.Println("Message sent")
		mess, err := request.RecvMessage(zmq4.SNDMORE)
		//mess, err := reply.RecvMessage(zmq4.SNDMORE)
		writer.WriteString(mess[0])
		writer.Flush()
		//fmt.Println(mess[0])
		if err != nil {
			fmt.Println("Timeout error.")
		}
	}
}

func CreateClient(testToast []flour.Bread) {

	request, err := zmq4.NewSocket(zmq4.REQ)
	if err != nil {
		fmt.Println("Error creating request socket.")
	}
	request.Connect("tcp://192.168.0.101:5555")
	//request.Connect("tcp://127.0.0.1:5555")

	file, err := os.Create("poptart/101/client.txt")
	defer file.Close()
	if err != nil {
		fmt.Println("Error creating message file.")
	}
	writer := bufio.NewWriter(file)
	for {
		mess, err := request.RecvMessage(zmq4.SNDMORE)
		//mess, err := reply.RecvMessage(zmq4.SNDMORE)
		writer.WriteString(mess[0])
		writer.Flush()
		//message := olive.PrepareMsg(testToast)
		colourMessage := flour.PrepareToast(testToast, "red", "blue")
		//fmt.Println(message)
		request.SendMessage(colourMessage)
		fmt.Println("Message sent")
		//fmt.Println(mess[0])
		if err != nil {
			fmt.Println("Timeout error.")
		}
	}
}
