package main

import (
	"fmt"
)

func floatVsInteger(i int, f float32) string {
	var conv = float32 (i)
	if conv > f {
		return "Integer"
	} else if (conv<f) {
		return "Float"
	} else {return "Same"}
	
}

func main () {
	var answer  = IntVsFloat (8, 7.2)
	fmt.Println(answer)
}
