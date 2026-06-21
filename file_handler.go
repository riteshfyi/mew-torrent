package main

import (
	"os"
)

func OpenFile(filePath string) (*os.File, error){
	f, err := os.Open(filePath)
	if err != nil {
		//to add a logger statement
	}
	return f,err;
}