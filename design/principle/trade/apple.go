package trade

import "fmt"

var apple = &Apple{}

type Apple struct {
}

func (t Apple) order() {
	fmt.Println("apple order")
}
