package flour

import (
	"fmt"
	//	"io/ioutil"
	"os"
)

type Bread struct {
	X, Y  int
	Label string
	Nl    bool
}

type Flour interface {
	Dough()
	Toast()
	Oven()
	Bread()
}

func Toast(loaf []Bread) {
	Toast_Logger("Toast")
	var display_toast string
	for i := range loaf {

		display_toast += loaf[i].Label
		//fmt.Printf(loaf[i].Label)
		if loaf[i].Nl {
			display_toast += "\n"
			//	fmt.Println("")
		}
	}
	display_toast += "_\n\x1b[93;41m\x1b[5;5H<:o.o:>\x1b[0m"
	fmt.Printf(display_toast)
}

func Oven(butt []Bread, label string, xvar int, yvar int) []Bread {
	Toast_Logger("Oven")
	x := 0
	y := 0
	for index := range butt {

		//X values
		if x+1 == xvar {
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
	var butter []Bread
	//we get some extraneous values
	for yandex := range butt {

		if label == "_" {
			if yandex%2 == 0 {
				butt[yandex].Label = "0"
			} else {
				butt[yandex].Label = "1"
			}
		} else {
			butt[yandex].Label = label
		}
		if butt[yandex].Y >= yvar {
			butter = butt[:yandex]
			//fmt.Println("GAME OVER MAN, GAME OVER")
			break
		}
	}
	return butter
}

func Bread_Getter(x int, y int, loaf []Bread) Bread {
	Toast_Logger("Bread_Getter")
	//Gets the bread at position x, y
	var val Bread

	for i := range loaf {
		if loaf[i].Y == y {
			if loaf[i].X == x {
				val = loaf[i]
				break
			}
		}
	}
	return val
}

func Bread_Setter(x int, y int, loaf []Bread, val Bread) []Bread {
	Toast_Logger("Bread_Setter")
	//sets the Bread at position x, y
	for i := range loaf {
		if loaf[i].Y == y {
			if loaf[i].X == x {
				loaf[i] = val
				break
			}
		}
	}

	return loaf

}

func Dough(width int, height int) []Bread {
	Toast_Logger("Dough")
	var butt []Bread
	//fmt.Println("Dough all mooshy!")
	//the nines can be changed
	butt = make([]Bread, width*height)

	return butt

}

func Toast_Logger(logger string) {
	blab := 0
	loggo, err := os.OpenFile("toast.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer loggo.Close()
	_, err = os.Lstat("toast.log")
	if err != nil {
		fmt.Println("Fatal error")
	}
	if blab == 1 {
		loggo.WriteString("===" + logger + "===\n")
	}
}
