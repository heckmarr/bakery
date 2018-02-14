package main

import (
	"gitlab.com/localtoast/bakery/flour"
//	"gitlab.com/localtoast/bakery/oven"
//	"gitlab.com/localtoast/bakery/loaf"
//	"gitlab.com/localtoast/bakery/dough"
	"fmt"
)


func main() {
//	toast.Toast()
//	oven.Oven()
//	loaf.Loaf()
	//init can be changed
	xvar := 81
	yvar := 23
	new_button := flour.Dough(xvar, yvar)
	new_button = flour.Oven(new_button, "BUTOOON", xvar, yvar)
	//just toasting something
	bread := "rye"
	for _, v := range bread {
		fmt.Printf(string(v))
	}
	fmt.Println(new_button)
	for i := range new_button {
		fmt.Printf(new_button[i].Label)
		if new_button[i].Nl {
			fmt.Println("")
		}
	}
//	for _, v := range new_button {
//		fmt.Println("X = ",v.X)
//		fmt.Println("Y = ", v.Y)
//	}
}
