package main

import (
	"gitlab.com/localtoast/bakery/flour"
//	"gitlab.com/localtoast/bakery/oven"
//	"gitlab.com/localtoast/bakery/loaf"
//	"gitlab.com/localtoast/bakery/dough"
	"fmt"
	"strings"
	"math/rand"
	"time"
)

func Spatter(xvar int, yvar int, test_toast []flour.Bread) {
        rand.Seed(12)
        welcome := "WELCOME_TO_HECK"
        wel := strings.Split(welcome, "")
        fmt.Println(wel[0])
        for i := 0;i < len(test_toast)/24;i++ {
                x := rand.Intn(xvar - 1)
                y := rand.Intn(yvar - 1)
                slice := flour.Bread_Getter(x, y, test_toast)
                slice.Label = string(wel[rand.Intn(len(wel)-1)])
                test_toast = flour.Bread_Setter(x, y, test_toast, slice)
                flour.Toast(test_toast)
                fmt.Printf("0\nDG:>")
                time.Sleep(75*time.Millisecond)
                }
}

func Welcome(test_toast []flour.Bread) {
        welcome := "WELCOME_TO_DEEGEE"
        wel := strings.Split(welcome, "")
//        fmt.Println(wel[0])
        for i := 0;i < len(welcome);i++ {
//                                        DO STUFF HERE
                slice := flour.Bread_Getter(30+i, 11, test_toast)
                slice.Label = string(wel[i])
                test_toast = flour.Bread_Setter(30+i, 11, test_toast, slice)
        }
	flour.Toast(test_toast)
	fmt.Printf("\nDG:>")
}

func Copy_Toast(welcome string, xvar int, yvar int, yend int, test_toast []flour.Bread) {
        wel := strings.Split(welcome, "")
//        fmt.Println(wel[0])
	if yend != 0{
		for x := yend; x > 0;x--{
		        for i := 0;i < len(welcome);i++ {
//                                        DO STUFF HERE
       		        slice := flour.Bread_Getter(xvar+i, yvar - x, test_toast)
                	slice.Label = string(wel[i])
                	test_toast = flour.Bread_Setter(xvar+i, yvar - x, test_toast, slice)
			}
		}
	}
//        flour.Toast(test_toast)
//        fmt.Printf("\nDG:>")
}



func Flat(label string, test_toast []flour.Bread) {
        for i := range test_toast {
                test_toast[i].Label = label
        }
        flour.Toast(test_toast)
	fmt.Printf("\nDG:>")
}

func Spawn_Button(label string, xvar int, yvar int, test_toast []flour.Bread){
	Copy_Toast("=====", xvar, yvar, 1, test_toast)
	Copy_Toast(("| "+label+ " |"), xvar, yvar+1, 1, test_toast)
	Copy_Toast("=====", xvar, yvar+2, 1, test_toast)

}

func Spawn_View(xvar int, yvar int, test_toast []flour.Bread, xlen int, yhei int) []flour.Bread {

	slice := flour.Dough(xvar+1, yvar+1)
	slice = flour.Oven(slice, "=", xvar, yvar)

	for y := 0;y < yhei;y++ {
		for x := 0;x < xlen;x++{
			Copy_Toast(slice[0].Label, xvar+x, yvar+y, 1, test_toast)
		}
	}
	return slice

}



func main() {
//	toast.Toast()
//	oven.Oven()
//	loaf.Loaf()
	input := ""
	//init can be changed
	xvar := 81
	yvar := 23
	test_toast := flour.Dough(xvar, yvar)
	test_toast = flour.Oven(test_toast, "BUTOOON", xvar, yvar)
	//just toasting something
	for {
		flour.Toast(test_toast)
		fmt.Printf("\nDG:>")
		fmt.Scan(&input)

		switch input {
			case "spawn":
				var contentview []flour.Bread
				var nodeview []flour.Bread
				Flat("_", test_toast)
				time.Sleep(1*time.Second)
				Copy_Toast("DEEGEE", 35, 11, 1, test_toast)
				flour.Toast(test_toast)
				time.Sleep(1*time.Second)
				Spatter(xvar, yvar, test_toast)
				Flat("_", test_toast)
				Spawn_Button("1",30, 2, test_toast)
				Spawn_Button("2",30, 19, test_toast)
				Spawn_Button("3",1, 2, test_toast)
				Spawn_Button("4",1, 19, test_toast)
				Spawn_Button("5",74, 2, test_toast)
				Spawn_Button("6",74, 19, test_toast)
				contentview = Spawn_View(35, 5, test_toast, 39, 14)
				nodeview = Spawn_View(5, 5, test_toast, 25, 14)
				//do things with them
				contentview[0] = contentview[0]
				nodeview[0] = nodeview[0]
			case "spatter":
				Spatter(xvar, yvar, test_toast)
			case "welcome":
				Welcome(test_toast)
			case "flat":
				Flat("_", test_toast)
			case "exit":
				break
			default:
				flour.Toast(test_toast)
		}
	}
}
