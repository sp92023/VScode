/*package main

import "fmt"

type laptopSize float64

func (this laptopSize) getSizeOfLaptop() laptopSize {
	return this
}

func main() {
	var l laptopSize = 3.3
	fmt.Println(l.getSizeOfLaptop())
}
*/

package main

import "fmt"

func main() {
	name := "Bill"

	namePointer := &name

	fmt.Println(&namePointer)
	fmt.Println(*namePointer)
	printPointer(namePointer)
}

func printPointer(namePointer *string) {
	fmt.Println(&namePointer)
	fmt.Println(*namePointer)
}
