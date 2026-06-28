package main

import "fmt"

func test() {
	 content, _ := readFile("./test.torrent")

	// tests := []string{"i42e","i-42e","le", "d4:name1:a3:agei20ee","d1:a1:a1:a1:be"};

	// for i := 0; i < len(tests) ; i++ {
	// content := tests[i];
	decodedContent := decodeBenCode(content)
	 fmt.Println("output : ", decodedContent)
	// }
	
}