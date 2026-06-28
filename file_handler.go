package main

import (
	"os"
)

func readFile(filePath string) (string, error){
	f, err := os.ReadFile(filePath)
	if err != nil {
		//to add a logger statement
	}
	content := string(f);
	return content,err;
}