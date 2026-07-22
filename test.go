package main

import (
	"fmt"
	"path/filepath"
	"strcnv"
	"sync"
)

var port int = 6881;
//decodeContent is the file info
func test() {
    content, _ := readFile("./test.torrent")
    decodedContent := decodeBenCode(content)
    resp := getTracker(decodedContent.(map[string]any));
    info := decodedContent.(map[string]any)["info"].(map[string]any)
    
    //## let's create the files

    basePath := "./downloads"
    folderName := info["name"].(string);

    targetPath := filepath.Join(basePath, folderName);

    piece_length := strcnv.Atoi(info["piece length"].(string)).(int);
     
    files := info["files"].([]any);
    totalBytes := 0;


    //make a prefix array, which will map the peice to the file path for the writing

    var peiceFilePath []string;
    var bytePrefix []int;

    for _, file := range files {
        bytePrefix = append(bytePrefix, totalBytes);
        fileLength :=strcnv.Atoi(file.(map[string]any)["length"].(string)); //size of file
        totalBytes+=fileLength;
        currPath := file.(map[string]any)["path"].([]string);

        curr := targetPath;
        for _, name := range currPath{
            curr = filepath.Join(curr, name);
        }
        peiceFilePath = append(peiceFilePath, curr);
        createSparseFile(curr, fileLength);
    }



    //artifacts: we have the above peicePrefix of bytes, also we have the peice path's

     


    // 1. Assert peers as []map[string]any
    peers := resp["peers"].([]any)
    
    
    var wg sync.WaitGroup;
    for _, peer := range peers {
        // 2. Access values and type-assert immediately to string/int
        // We use type assertion to convert 'any' into a usable type
		peerObj := peer.(map[string]any)
        ip := peerObj["ip"].(string)
        port := peerObj["port"].(int) // Ports are usually integers in Bencode
        id := peerObj["peer id"].(string)
        
        // 3. Construct the address string
        // Use fmt.Sprintf to join the IP and the converted port
        peerAddr := fmt.Sprintf("%s:%d", ip, port)
        
        fmt.Printf("Attempting Connection for %s, for address: %s\n", id, peerAddr)
        
        // 4. Pass the cleaned variables
        wg.Add(1);
        go connectPeer(id, peerAddr, &wg)
    }
//connect krlo , phir send the handshakes ? 
    wg.Wait()
    handshake := buildHandshake(info_hash);
    fmt.Println("PeerMap : ", peerMap);

    
    for key, _ := range peerMap {
        wg.Add(1);
        go sendMessage(key, handshake, &wg);
    }

    fmt.Println("handshakes sent ");

    wg.Wait();


    //created the sparse file, sent handshakes, now we will 




	for {
		
	}
}