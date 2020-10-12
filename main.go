package main

import (
	"fmt"
	"strings"
)

func main() {
	var s int
	for {
		fmt.Print("Enter number of square :")
		// request number
		fmt.Scan(&s)
		switch {
		case s <= 2: // if the input is less than 2 print "to low/invalid"
			fmt.Println(s, "to low/invalid")
			continue
		case s%2 == 0: // if the input is not an odd number print "is even number"
			fmt.Println(s, "is Even number")
			continue
		default:
			fmt.Println("")
		}
		break
	}
	var p1 = make([]string, s)
	var p2 = make([]string, s)
	var pm = make([]string, s)
	var pMid = s - ((s - 1) / 2) - 1
	for i := range p1 {
		if i == 0 || i == (s-1) {
			p1[i] = "*"
			continue
		}
		p1[i] = "="
	}
	for i := range p2 {
		p2[i] = "*"
	}
	var p1String = strings.Join(p1[:], " ")
	var p2String = strings.Join(p2[:], " ")
	for i := range pm {
		if pMid == i {
			pm[i] = p2String
			continue
		}
		pm[i] = p1String
	}
	fmt.Println(strings.Join(pm[:], "\n"))
}
