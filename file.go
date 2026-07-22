package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func readFile(filePath string) (string, error){
	f, err := os.ReadFile(filePath)
	if err != nil {
		//to add a logger statement
	}
	content := string(f);
	return content,err;
}

func createSparseFile(fileDir string, size int64){
	dir := filepath.Dir(fileDir);

	err := os.MkdirAll(dir, 0755);

	if err != nil {
		fmt.Println("Error While Creating Directory Folders.");
		panic(err);
	}
	
	file,err := os.Create(fileDir);

	if err != nil {
		fmt.Println("Error While Creating Sparse File.");
	}
	defer file.Close()

	//assuming size to be in bytes

	if err := file.Truncate(size);
	 err != nil {
		log.Fatalf("Failed to truncate file: %v", err)
		panic(err);
	}

}