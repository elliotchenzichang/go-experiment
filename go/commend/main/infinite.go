package main

import (
	"fmt"
	inf "github.com/fzdwx/infinite"
	"github.com/fzdwx/infinite/components/selection/confirm"
)

func main() {

	val, _ := inf.NewConfirmWithSelection(
		confirm.WithDefaultYes(),
	).Display()

	if val {
		fmt.Println("yes, you are.")
	} else {
		fmt.Println("no,you are not.")
	}
}
