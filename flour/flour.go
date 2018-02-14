package flour

import "fmt"

type Bread struct{
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

func Toast(loaf []Bread) {
	for i := range loaf {
		fmt.Printf(loaf[i].Label)
		if loaf[i].Nl {
			fmt.Println("")
		}
	}
}

func Oven(butt []Bread, label string, xvar int, yvar int) []Bread {
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
	var butter []Bread
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

func Bread_Getter(x int, y int, loaf []Bread) Bread {
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
	var butt []Bread
	fmt.Println("Dough all mooshy!")
	//the nines can be changed
	butt = make([]Bread, width*height)

	return butt

}
