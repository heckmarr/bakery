package flour

import (
	"fmt"
	"math"
	"strings"
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
type Loaf struct {
	Height, Width int
}

type Flour interface {
	Dough()
	Toast()
	Oven()
	Bread()
}

func CleanFlecks(loaf []Bread) []Bread {
	//ToastLogger("CleanFlecks")
	for i := range loaf {
		if loaf[i].Label != "_" {
			loaf[i].Dirty = false
		}
	}
	return loaf
}
func MakeCleanFlecks(loaf []Bread) []Bread {
	//ToastLogger("MakeCleanFlecks")
	for i := range loaf {
		if loaf[i].Dirty == true {
			loaf[i].Label = "_"
			loaf[i].Dirty = false
		}
	}
	return loaf
}
func SpawnWin(xvar int, yvar int) ([]Bread, Loaf) {
	//ToastLogger("SpawnWin")
	var winLoaf Loaf
	winLoaf.Height = yvar
	winLoaf.Width = xvar
	win := Dough(xvar, yvar)
	win = Oven(win, "*", xvar, yvar)
	return win, winLoaf
}

//RelWin Copies a window with size and height relative to the size of the toast
//passed to the Loaf passed, or if the last value is false, the relative size of
//the terminal in which it is called.
func RelWin(widthP float64, heightP float64, heightPadding float64, widthPadding float64, win []Bread, testToast []Bread, winLoaf Loaf, relativeToParent bool) ([]Bread, Loaf) {

	var tHeight64 float64
	var tWidth64 float64
	if !relativeToParent {
		tHeight, err := terminaldimensions.Height()
		tWidth, err := terminaldimensions.Width()
		if err != nil {
			fmt.Println("terminal sizing error!")

		}
		tHeight64 = float64(tHeight)
		tWidth64 = float64(tWidth)

	} else {
		tHeight := winLoaf.Height
		tWidth := winLoaf.Width
		tHeight64 = float64(tHeight)
		tWidth64 = float64(tWidth)

	}

	xbeg := math.Floor((tWidth64 * widthP))
	xend := math.Floor(xbeg + widthPadding)

	ybeg := math.Floor((tHeight64 * heightP))
	yend := math.Floor(ybeg + heightPadding)

	xendI := int(xend)
	yendI := int(yend)
	ybegI := int(ybeg)
	xbegI := int(xbeg)
	for i := range win {
		CopySubToast(win[i].Label, win[i].X, win[i].Y, xbegI, xendI, ybegI, yendI, testToast)
	}
	return testToast, winLoaf
}

//CopySubToast copies the string passed into the values of a []Bread given
func CopySubToast(welcome string, xvar int, yvar int, xbeg int, xend int, ybeg int, yend int, testToast []Bread) []Bread {
	//ToastLogger("CopyToast")
	//wel := strings.Split(welcome, "")
	if yend != 0 {
		for x := ybeg - int(float64(yvar)*0.5); x < yend-int(float64(yvar)*0.5); x++ {
			for i := xbeg - int(float64(xvar)*0.5); i < xend-int(float64(xvar)*0.5); i++ {
				//                           DO STUFF HERE
				//			if x < yend && x > yend {
				//slice := BreadGetter(0, 0, testToast)
				slice := BreadGetter(xvar+i, yvar+x, testToast)
				slice.Label = welcome
				slice.Dirty = true
				testToast = BreadSetter(xvar+i, yvar+x, testToast, slice)
				//			}
			}
		}
	}
	return testToast
}

//CopyToast copies the string passed into the values of a []Bread given
func CopyToast(welcome string, xvar int, yvar int, yend int, testToast []Bread) []Bread {
	//ToastLogger("CopyToast")
	wel := strings.Split(welcome, "")
	if yend != 0 {
		for x := yend; x > 0; x-- {
			for i := 0; i < len(welcome); i++ {
				//                                        DO STUFF HERE
				slice := BreadGetter(xvar+i, yvar-x, testToast)
				slice.Label = string(wel[i])
				slice.Dirty = true
				testToast = BreadSetter(xvar+i, yvar-x, testToast, slice)
			}
		}
	}
	return testToast
}
func Toast(loaf []Bread) {
	//ToastLogger("Toast")
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
	}
	displaytoast += "_\n\x1b[93;41m\x1b[3;6H<:o.o:>\x1b[0m"
	fmt.Printf(displaytoast)
}

func Fleck(index int, loaf []Bread) string {
	//ToastLogger("Fleck")
	text := fmt.Sprint("\x1b[", loaf[index].Y, ";", loaf[index].X, "H", loaf[index].Label, "\x1b[0m")
	return text
}
func PrintFleck(index int, loaf []Bread) {
	//ToastLogger("PrintFleck")
	text := fmt.Sprint("\x1b[", loaf[index].Y, ";", loaf[index].X, "H", loaf[index].Label, "\x1b[0m")
	fmt.Printf(text)
}
func Oven(butt []Bread, label string, xvar int, yvar int) []Bread {
	//ToastLogger("Oven")
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
	//ToastLogger("BreadGetter")
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
	//ToastLogger("BreadSetter")
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
	//ToastLogger("Dough")
	var butt []Bread
	butt = make([]Bread, width*height)

	return butt
}
func DoughMax() (int, int, []Bread, Loaf) {
	//ToastLogger("DoughMax")
	var butt []Bread
	var loaf Loaf
	height, err := terminaldimensions.Height()
	width, err := terminaldimensions.Width()
	if err != nil {
		fmt.Println("Dimensional error")
		//fmt.Println(strconv.Atoi(string(height)))
		//fmt.Println(strconv.Atoi(string(width)))
	}
	loaf.Height = int(height)
	loaf.Width = int(width)
	//fmt.Println("Dough all mooshy!")
	//the nines can be changed
	heightInt := int(height)
	widthInt := int(width)
	butt = make([]Bread, widthInt*heightInt)

	return widthInt, heightInt, butt, loaf

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
