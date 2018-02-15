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
	bread := "rye"
	for _, v := range bread {
		fmt.Printf(string(v))
	}
	fmt.Println(test_toast)
	for {
//		for i := range test_toast {
//			fmt.Printf(test_toast[i].Label)
//			if test_toast[i].Nl {
//				fmt.Println("")
//			}
//		}
		flour.Toast(test_toast)
		fmt.Printf("0DG:>")
		fmt.Scan(&input)
		switch input {
			case "spawn":
				welcome := "WELCOME_TO_DEEGEE"
				wel := strings.Split(welcome, "")
				fmt.Println(wel[0])
				for i := 0;i < len(welcome);i++ {
//        				  DO STUFF HERE
					slice := flour.Bread_Getter(30+i, 11, test_toast)
					slice.Label = string(wel[i])
					test_toast = flour.Bread_Setter(30+i, 11, test_toast, slice)
				}
				flour.Toast(test_toast)
			case "matrici":
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
					fmt.Printf("0DG:>")
					time.Sleep(100*time.Millisecond)
				}
			case "flat":
				for i := range test_toast {
					test_toast[i].Label = "_"
				}
				flour.Toast(test_toast)
			case "exit":
				break
			default:
				flour.Toast(test_toast)
		}
	}
//	for _, v := range new_button {
//		fmt.Println("X = ",v.X)
//		fmt.Println("Y = ", v.Y)
//	}
}
