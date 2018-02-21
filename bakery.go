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
	"bufio"
//	"io"
//	"io/ioutil"
	"os"
	"os/exec"
)
func readStdin(out chan string, in chan bool) {
	flour.Toast_Logger("readStdin")
        //no buffering
        exec.Command("stty","-F", "/dev/tty", "cbreak", "min", "1").Run()
        //no visible output
        exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
        // restore the echoing state when exiting
        defer exec.Command("stty", "-F", "/dev/tty", "echo").Run()

        var b []byte = make([]byte, 1)
        for {
                os.Stdin.Read(b)
                out <- string(b)

        }
}




func Spatter(xvar int, yvar int, test_toast []flour.Bread) {
	flour.Toast_Logger("Spatter")
        rand.Seed(12)
        welcome := "WELCOME_TO_THE_MARCHELL"
        wel := strings.Split(welcome, "")
        fmt.Println(wel[0])
        for i := 0;i < len(test_toast)/24;i++ {
                x := rand.Intn(xvar - 1)
                y := rand.Intn(yvar - 1)
                slice := flour.Bread_Getter(x, y, test_toast)
                slice.Label = string(wel[rand.Intn(len(wel)-1)])
                test_toast = flour.Bread_Setter(x, y, test_toast, slice)
                flour.Toast(test_toast)
                fmt.Printf("0\n<:o.o:>")
                time.Sleep(75*time.Millisecond)
                }
}

func Welcome(test_toast []flour.Bread) {
	flour.Toast_Logger("Welcome")
        welcome := "WELCOME_TO_DEEGEE"
        wel := strings.Split(welcome, "")
        for i := 0;i < len(welcome);i++ {
//                                        DO STUFF HERE
                slice := flour.Bread_Getter(30+i, 11, test_toast)
                slice.Label = string(wel[i])
                test_toast = flour.Bread_Setter(30+i, 11, test_toast, slice)
        }
	flour.Toast(test_toast)
	fmt.Printf("\n<:o.o:>")
}

func Copy_Toast(welcome string, xvar int, yvar int, yend int, test_toast []flour.Bread) {
	flour.Toast_Logger("Copy_Toast")
        wel := strings.Split(welcome, "")
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
}



func Flat(label string, test_toast []flour.Bread) {
	flour.Toast_Logger("Flat")
        for i := range test_toast {
                test_toast[i].Label = label
        }
        flour.Toast(test_toast)
	fmt.Printf("\n<:o.o:>")
}

func Spawn_Button(label string, xvar int, yvar int, test_toast []flour.Bread){
	flour.Toast_Logger("Spawn_Button")
	Copy_Toast("=====", xvar, yvar, 1, test_toast)
	Copy_Toast(("| "+label+ " |"), xvar, yvar+1, 1, test_toast)
	Copy_Toast("=====", xvar, yvar+2, 1, test_toast)

}

func Spawn_Contents(path string, xvar int, yvar int, test_toast []flour.Bread, xlen int, yhei int) []flour.Bread {
        flour.Toast_Logger("Spawn_Contents")
        slice := flour.Dough(xvar+1, yvar+1)
        slice = flour.Oven(slice, "=", xvar, yvar)
        filo, err := os.Open(path)
        filscan := bufio.NewScanner(filo)
        for filscan.Scan() {
                yvar++
                Copy_Toast(filscan.Text(), xvar, yvar, 1, test_toast)

        }

        if err != nil {
                fmt.Println("Something went wrong!")
        }
        return slice
}

func Spawn_Index(path string, xvar int, yvar int, test_toast []flour.Bread, xlen int, yhei int) []flour.Bread {
	flour.Toast_Logger("Spawn_Index")
	slice := flour.Dough(xvar+1, yvar+1)
	slice = flour.Oven(slice, "=", xvar, yvar)
	filo, err := os.Open(path)
	filscan := bufio.NewScanner(filo)
	for filscan.Scan() {
		yvar++
		Copy_Toast(filscan.Text(), xvar, yvar, 1, test_toast)

	}

	if err != nil {
		fmt.Println("Something went wrong!")
	}
	return slice

}




func main() {
	input := ""
	//init can be changed
	xvar := 81
	yvar := 23
	test_toast := flour.Dough(xvar, yvar)
	test_toast = flour.Oven(test_toast, "_", xvar, yvar)
	Flat("_", test_toast)
	//just toasting something
	for {
		flour.Toast(test_toast)
		fmt.Printf("\n<:o.o:>")
		fmt.Scan(&input)
		switch input {
			case "$":
			        stdin := make(chan string, 1)
        			kill := make(chan bool, 1)
				xpos := 0
        			go readStdin(stdin, kill)
        			for {
					flour.Toast(test_toast)
					fmt.Printf("_")
        			        str := <-stdin

        			        if str == "0" {
                			        kill <- true
                			        close(stdin)
                        			break
                			} else {
                				Copy_Toast(str, 35+xpos, 5, 1, test_toast)
						xpos++
// this is a good place to grab the rune printed
                			}

       				}


			case "#":
                                stdin := make(chan string, 1)
                                kill := make(chan bool, 1)
                                xpos := 0
				thread := "0"
//WIP
				for x := 0; x <  14; x++ {
					if x < 10 {
						thread = "00" + string(x) 
					}else {
						thread = "0" + string(x)
					}
					fmt.Println("AT THREAD POSITION "+thread)
				}
//WIP thread position
                                go readStdin(stdin, kill)
                                for {
                                        flour.Toast(test_toast)
                                        fmt.Printf("_")
                                        str := <-stdin

                                        if str == "0" {
                                                kill <- true
                                                close(stdin)
                                                break
                                        }
					if str == "k" {
	                                        Copy_Toast("_", 4, 5+xpos+1, 1, test_toast)
        	                                Copy_Toast("_", 30, 5+xpos+1, 1, test_toast)
//WIP, do we really need another redraw event?	Flat("_", test_toast)
                	                        //pre and post title hash
//WIP works on thread UP, not down
						if xpos < 10{
							thread = fmt.Sprint("00",xpos)
							fmt.Println("THREADPOSITION IS = " + string(thread))
						}
						if xpos >= 100{
							thread = fmt.Sprint(xpos)
							fmt.Println("THREADPOSITION IS = " + string(thread))
						}
						if xpos >= 10{
							thread = fmt.Sprint("0",xpos)
							fmt.Println("THREADPOSITION IS = " + string(thread))
						}
                        	                Spawn_Contents(fmt.Sprint("breadbox/"+thread+".1"), 35, 4, test_toast, 39, 14)
						Copy_Toast("#", 4, 5+xpos, 1, test_toast)
                                	        Copy_Toast("#", 30, 5+xpos, 1, test_toast)
						xpos--
//END WIP, don't forget to do it upwards too
					}
					if str == "j" {
						if xpos == 0 {
							Copy_Toast("#", 4, 5, 1, test_toast)
							Copy_Toast("#", 30, 5, 1, test_toast)
						} else{
							Copy_Toast("_", 4, 5+xpos-1, 1, test_toast)
							Copy_Toast("_", 30, 5+xpos-1, 1, test_toast)
							//pre and post title hash
                                                	Copy_Toast("#", 4, 5+xpos, 1, test_toast)
							Copy_Toast("#", 30, 5+xpos, 1, test_toast)
						}
						xpos++
// this is a good place to grab the rune printed
                                        }
				}

			case "spawn":
				Flat("_", test_toast)
				time.Sleep(1*time.Second)
				Copy_Toast("DEEGEE", 35, 11, 1, test_toast)
				flour.Toast(test_toast)
				fmt.Printf("_<:o.o:>")
				time.Sleep(1*time.Second)
				Spatter(xvar, yvar, test_toast)
				Flat("_", test_toast)
				Spawn_Button("$",30, 2, test_toast)
				Spawn_Button("2",30, 19, test_toast)
				Spawn_Button("#",1, 2, test_toast)
				Spawn_Button("4",1, 19, test_toast)
				Spawn_Button("5",74, 2, test_toast)
				Spawn_Button("6",74, 19, test_toast)
				//turn this into spawn_content
				Spawn_Index("breadbox/000.1", 35, 4, test_toast, 39, 14)
				Spawn_Index("breadbox/000", 5, 4, test_toast, 25, 14)
				Spawn_Index("breadbox/001", 5, 5, test_toast, 25, 14)
				//do things with them
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
