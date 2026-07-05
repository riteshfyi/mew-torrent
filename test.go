package main

import "fmt"

var port int = 6800

func test() {
	content, _ := readFile("./test.torrent")
	decodedContent := decodeBenCode(content)
	output := encode(decodedContent)  
	//encoder is working fine too

}