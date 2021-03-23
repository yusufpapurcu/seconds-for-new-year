package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

const (
	zero = `
██████████████████
██              ██
██              ██
██              ██
██              ██
██              ██
██              ██
██              ██
██              ██
██              ██
██              ██
██████████████████
`
	one = `
      █████       
    ███████       
  █████████       
       ████       
       ████       
       ████       
       ████       
       ████       
       ████       
       ████       
       ████       
██████████████████
`

	two = `
██████████████████
                ██
                ██
                ██
                ██
██████████████████
██                
██                
██                
██                
██                
██████████████████
`
	three = `
██████████████████
                ██
                ██
                ██
                ██
██████████████████
                ██
                ██
                ██
                ██
                ██
██████████████████
`
	four = `
██              ██
██              ██
██              ██
██              ██
██              ██
██████████████████
                ██
                ██
                ██
                ██
                ██
                ██
`
	five = `
██████████████████
██                
██                
██                
██                
██████████████████
                ██
                ██
                ██
                ██
                ██
██████████████████
`
	six = `
██████████████████
██                
██                
██                
██                
██████████████████
██              ██
██              ██
██              ██
██              ██
██              ██
██████████████████
`
	seven = `
██████████████████
                ██
                ██
                ██
                ██
                ██
                ██
                ██
                ██
                ██
                ██
                ██
`
	eight = `
██████████████████
██              ██
██              ██
██              ██
██              ██
██████████████████
██              ██
██              ██
██              ██
██              ██
██              ██
██████████████████
`
	nine = `
██████████████████
██              ██
██              ██
██              ██
██              ██
██████████████████
                ██
                ██
                ██
                ██
                ██
██████████████████
`
)

var numbers = []string{zero, one, two, three, four, five, six, seven, eight, nine}

func main() {

	for i := 60; i > 0; i-- {
		renderNumber(i)
		time.Sleep(1000 * time.Millisecond)
		var clear *exec.Cmd
		if runtime.GOOS == "windows" {
			clear = exec.Command("cmd", "/c", "cls")
		} else {
			clear = exec.Command("clear")
		}
		clear.Stdout = os.Stdout
		err := clear.Run()
		if err != nil {
			fmt.Println(err)
		}
	}
}

func renderNumber(number int) {
	if -1 < number && number < 10 {
		fmt.Print(numbers[number], "\n")
	} else if 9 < number && number < 100 {
		second := number % 10
		first := (number - second) / 10
		//	fmt.Println("First :", first, "Second :", second)
		firstArr := strings.Split(numbers[first], "\n")
		secondArr := strings.Split(numbers[second], "\n")
		fmt.Print("\n\n")
		for i := 0; i < 13; i++ {
			strFirst := firstArr[i]
			strSecond := secondArr[i]

			fmt.Print("\t\t\t", strFirst, "  ")
			if len(strFirst) < 22 {
				for a := 0; a < 22-len(strFirst); a++ {
					//	fmt.Print("a")
				}
			}
			fmt.Print(strSecond)
			fmt.Print("\n")
		}

	}
}

// █ █ █ █ █ █
