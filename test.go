package main

import "fmt"

func test() {
	input := "l4:texti100eld1:ai1ee";
	decoded := decode(input)
	fmt.Printf("%+v\n", decoded);
}