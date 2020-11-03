package main

import (
	"fmt"
	//"time"
	//"os/user"
)

// func init() {
// 	fmt.Println("Init!")
// }

func main() {
	var (
		i   int     = 1
		f64 float64 = 1.2
		s   string  = "test"
		t   bool    = true
		f   bool    = false
	)

	fmt.Println(i, f64, s, t, f)

	xi := 1
	xf64 := 1.2
	xs := "test"
	xt, xf := true, false

	fmt.Println(xi, xf64, xs, xt, xf)
	fmt.Printf("%T", xf64)
}
