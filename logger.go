package main

import "fmt"

var levels = map[int]string{
	0: "info",
	1: "warning",
	2: "error",
}



func Log(level int, statement string) {
	fmt.Printf("[%s]:%s", levels[level], statement);
}