package main

import (
	"fmt"
	"strings"
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

	for i := 0; i < 30; i++ {
		renderNumber(i)
	}
}

func renderNumber(number int) {
	if -1 < number && number < 10 {
		fmt.Println(numbers[number])
	} else if 9 < number && number < 100 {
		second := number % 10
		first := (number - second) / 10
		fmt.Println("First :", first, "Second :", second)
		firstArr := strings.Split(numbers[first], "\n")
		secondArr := strings.Split(numbers[second], "\n")

		for i := 0; i < 13; i++ {
			strFirst := firstArr[i]
			strSecond := secondArr[i]

			fmt.Print(strFirst, " ")
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
