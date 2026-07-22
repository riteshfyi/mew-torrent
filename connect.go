package main

import (
	"fmt"
	"net"
	"sync"
)

var peerMap = make(map[string]*net.TCPConn);

func connectPeer(id string, addr string, wg *sync.WaitGroup){
	//check if already exists ? 
	defer wg.Done()
    _, ok := peerMap[id];

	if ok {
		fmt.Println("Connection Already Exists with Peer:", id)
		return;
	}
	
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)

	// check for some error 

	if err != nil {
		fmt.Println("Unable to Resolve TCP Address:", err)
		return
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	
	if err != nil {
		fmt.Println("Error connecting:", err)
		return;
	}
	fmt.Printf("Connected With Peer %s", id);
	peerMap[id] = conn
}

func disconnectPeer(id string) {
   conn, ok := peerMap[id];
   if !ok {
		fmt.Println("Peer doesn't exists : Unable to disconnect:")
		return 
   }
	conn.Close()
	delete(peerMap, id)
}

func sendMessage(id string, message []byte, wg *sync.WaitGroup) {
	defer wg.Done()
	conn, ok := peerMap[id];
	if !ok {
		fmt.Printf("Peer not connected : unable to send message, Peer Id : %s", id)
		return
	}
	fmt.Println("Sending Message to %s", id);
	conn.Write(message)
}





func handlePeer(id string) {
    // conn, ok := peerMap[id];
	// if !ok {
	// 	fmt.Println("No Such Peer Exists : Skipping its handler");
	// 	return;
	// }
	// reader := bufio.NewReader(conn)
	// buffer := make([]byte, 64*1024)   //assume data being max to max as 64KB
	// for {



	// }
}

func closeAll(){
   for _,conn := range peerMap{
	 conn.Close()
   }
}