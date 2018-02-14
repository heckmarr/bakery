package main

import (
	"gitlab.com/localtoast/bakery/flour"
//	"gitlab.com/localtoast/bakery/oven"
//	"gitlab.com/localtoast/bakery/loaf"
//	"gitlab.com/localtoast/bakery/dough"
	"fmt"
	"strings"
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
	new_button := flour.Dough(xvar, yvar)
	new_button = flour.Oven(new_button, "BUTOOON", xvar, yvar)
	//just toasting something
	bread := "rye"
	for _, v := range bread {
		fmt.Printf(string(v))
	}
	fmt.Println(new_button)
	for {
		for i := range new_button {
			fmt.Printf(new_button[i].Label)
			if new_button[i].Nl {
				fmt.Println("")
			}
		}
		fmt.Printf("0DG:>")
		fmt.Scan(&input)
		switch input {
			case "spawn":
				x := 0
				done := false
				y := 0
				welcome := "WELCOME_TO_DEEGEE"
				wel := strings.Split(welcome, "")
				fmt.Println(wel[0])
				for {
					for i := range new_button {
						if new_button[i].Y == y {
							if new_button[i].X > 30 {
								if done == false{
									new_button[i].Label = wel[x]
									x++
									if x == 17{
										x = 0
										y++
										if y == 11 {
										done = true
										x = 0
										}
									}
								}
							}
						
						}
						fmt.Printf(new_button[i].Label)
						if new_button[i].X % 80 == 0 {
							fmt.Println("")
							time.Sleep(250*time.Millisecond)
						}
					}
					if done {
//						time.Sleep(250*time.Millisecond)
					//	fmt.Scan(&input)
						for i := range new_button {
							if new_button[i].Y == 11{
								break
							}
							fmt.Printf(new_button[0].Label)
							time.Sleep(5*time.Millisecond)
						}
						fmt.Scan(&input)
					}
					
				}
		}
	}
//	for _, v := range new_button {
//		fmt.Println("X = ",v.X)
//		fmt.Println("Y = ", v.Y)
//	}
}
