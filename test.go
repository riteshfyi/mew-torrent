package main

var port int = 6881

func test() {
	content, _ := readFile("./test.torrent")
	decodedContent := decodeBenCode(content)
	getTracker(decodedContent.(map[string]any))
}