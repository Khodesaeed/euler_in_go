package main

import "fmt"

func main() {
	sum := 0
	fibSecFst := 1
	fibSecSec := 1
	for fibSecSec < 4000000 {
		if fibSecSec%2 == 0 {
			sum += fibSecSec
		}
		temp := fibSecFst + fibSecSec
		fibSecFst = fibSecSec
		fibSecSec = temp
	}
	fmt.Println(sum)
}
