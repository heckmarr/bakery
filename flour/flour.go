package flour

import "fmt"

type Button struct{
	X,Y int
	Label string
	Nl bool
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
		if x + 1 == xvar {
			x = 0
//			butt[index].Y = y
			butt[index].Nl = true
			y++
		}
		butt[index].X = x
		x++
		//Y values
//		if index % yvar == 0{
//			y++
//		}
		butt[index].Y = y
	}
	var butter []Button
	//we get some extraneous values
	for yandex := range butt{
		if yandex % 2 == 0{
			butt[yandex].Label = "0"
		} else {
			butt[yandex].Label = "1"
		}
		if butt[yandex].Y >= yvar {
			butter = butt[:yandex]
			fmt.Println("GAME OVER MAN, GAME OVER")
			break
		}
	}
	return butter
}

func Bread() {

}

func Dough(width int, height int) []Button {
	var butt []Button
	fmt.Println("Dough all mooshy!")
	//the nines can be changed
	butt = make([]Button, width*height)

	return butt

}
