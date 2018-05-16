package main

import "fmt"

type laptopSize float64

func (this laptopSize) getSizeOfLaptop() laptopSize {
	return this
}

func main() {
	var l laptopSize = 3.3
	fmt.Println(l.getSizeOfLaptop())
}
