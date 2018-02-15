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
        fmt.Println(wel[0])
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
        fmt.Println(wel[0])
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
        flour.Toast(test_toast)
        fmt.Printf("\nDG:>")
}



func Flat(test_toast []flour.Bread) {
        for i := range test_toast {
                test_toast[i].Label = "_"
        }
        flour.Toast(test_toast)
	fmt.Printf("\nDG:>")
}

func Spawn_Button(xvar int, yvar int, test_toast []flour.Bread){
	Copy_Toast("=====", xvar, yvar, 1, test_toast)
	Copy_Toast("| X |", xvar, yvar+1, 1, test_toast)
	Copy_Toast("=====", xvar, yvar+2, 1, test_toast)

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
				Flat(test_toast)
				time.Sleep(1*time.Second)
				Copy_Toast("DEEGEE", 35, 11, 1, test_toast)
				time.Sleep(1*time.Second)
				Spatter(xvar, yvar, test_toast)
				Flat(test_toast)
				Spawn_Button(30, 2, test_toast)
				Spawn_Button(30, 19, test_toast)
				Spawn_Button(1, 2, test_toast)
				Spawn_Button(1, 19, test_toast)
				Spawn_Button(74, 2, test_toast)
				Spawn_Button(74, 19, test_toast)
			case "spatter":
				Spatter(xvar, yvar, test_toast)
			case "welcome":
				Welcome(test_toast)
			case "flat":
				Flat(test_toast)
			case "exit":
				break
			default:
				flour.Toast(test_toast)
		}
	}
}
