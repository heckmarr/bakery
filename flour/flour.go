package flour

import "fmt"

type Button struct{
	x,y int
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

func Oven(butt []Button, label string, xvar int, yvar int) []Button {
	x := 0
	y := 0
	for index := range butt{

		//X values
		if index % xvar == 0{
			x = 0
		}
		butt[index].x = x
		x++
		//Y values
		if index % yvar == 0{
			y++
		}
		butt[index].y = y
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
			//the nines can be changed
			butt = make([]Button, 27, 27)
		default:
	}
	return butt

}
