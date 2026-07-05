package main

var port int =  6800;
func test() {
	 content, _ := readFile("./test.torrent")
	decodedContent := decodeBenCode(content).(map[string]any)
	getTracker(decodedContent);

}    