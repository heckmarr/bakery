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
	//init
	new_button := flour.Dough("button")
	new_button = flour.Oven(new_button)
	//just toasting something
	bread := "rye"
	for _, v := range bread {
		fmt.Printf(string(v))
	}
	fmt.Println(new_button)
//	for _, v := range new_button {
//		fmt.Println("X = ",v.X)
//		fmt.Println("Y = ", v.Y)
//	}
}
