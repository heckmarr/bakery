package flour

import (
	"fmt"
	//	"io/ioutil"
	"os"

	"github.com/wayneashleyberry/terminal-dimensions"
)

type Bread struct {
	X, Y  int
	Label string
	Nl    bool
	Dirty bool
}

type Flour interface {
	Dough()
	Toast()
	Oven()
	Bread()
}

func CleanFlecks(loaf []Bread) []Bread {
	for i := range loaf {
		if loaf[i].Label != "_" {
			loaf[i].Dirty = false
		}
	}
	return loaf
}
func MakeCleanFlecks(loaf []Bread) []Bread {
	for i := range loaf {
		if loaf[i].Dirty == true {
			loaf[i].Label = "_"
			loaf[i].Dirty = false
		}
	}
	return loaf
}
func Toast(loaf []Bread) {
	ToastLogger("Toast")
	var displaytoast string
	for i := range loaf {
		if loaf[i].Dirty {
			displaytoast += Fleck(i, loaf)
		}
		if loaf[i].Dirty != true && loaf[i].Label != "_" {
			loaf[i].Label = "_"
			displaytoast += Fleck(i, loaf)
		} else {
			//do nothing
		}
		//displaytoast += loaf[i].Label
		//fmt.Printf(loaf[i].Label)
		//if loaf[i].Nl {
		//	displaytoast += "\n"
		//	fmt.Println("")
		//}
	}
	displaytoast += "_\n\x1b[93;41m\x1b[3;6H<:o.o:>\x1b[0m"
	fmt.Printf(displaytoast)
}

func Fleck(index int, loaf []Bread) string {
	text := fmt.Sprint("\x1b[", loaf[index].Y, ";", loaf[index].X, "H", loaf[index].Label, "\x1b[0m")
	return text
}
func PrintFleck(index int, loaf []Bread) {
	text := fmt.Sprint("\x1b[", loaf[index].Y, ";", loaf[index].X, "H", loaf[index].Label, "\x1b[0m")
	fmt.Printf(text)
}
func Oven(butt []Bread, label string, xvar int, yvar int) []Bread {
	ToastLogger("Oven")
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

func BreadGetter(x int, y int, loaf []Bread) Bread {
	ToastLogger("BreadGetter")
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

func BreadSetter(x int, y int, loaf []Bread, val Bread) []Bread {
	ToastLogger("BreadSetter")
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
	ToastLogger("Dough")
	var butt []Bread
	butt = make([]Bread, width*height)

	return butt
}
func DoughMax() (int, int, []Bread) {
	ToastLogger("DoughMax")
	var butt []Bread
	height, err := terminaldimensions.Height()
	width, err := terminaldimensions.Width()
	if err != nil {
		fmt.Println("Dimensional error")
		//fmt.Println(strconv.Atoi(string(height)))
		//fmt.Println(strconv.Atoi(string(width)))
	}

	//fmt.Println("Dough all mooshy!")
	//the nines can be changed
	heightInt := int(height)
	widthInt := int(width)
	butt = make([]Bread, widthInt*heightInt)

	return widthInt, heightInt, butt

}

func ToastLogger(logger string) {
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
