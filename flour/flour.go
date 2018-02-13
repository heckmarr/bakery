package flour

import "fmt"

type Button struct{
	X,Y int
	label string
}

type Flour interface {
	Dough()
	Toast()
	Oven()
	Bread()
}

func Toast() {

}

func Oven(butt []Button) []Button {
	for i := range butt {
		butt[i].X = 1
		butt[i].Y = 1
	}
	return butt
}

func Bread() {

}

func Dough(dough string) []Button {
	var butt []Button
	fmt.Println("Dough all mooshy!")
	switch dough {
		case "button":
			butt = make([]Button, 9)
		default:
	}
	return butt

}
