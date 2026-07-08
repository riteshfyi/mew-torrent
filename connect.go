package main

import (
	"fmt"
	"net"
)

var peerMap map[string]*net.TCPConn


//Todo 
// []Make map access concurrent safe ?



func connectPeer(id string, addr string){
	//check if already exists ? 
    _, ok := peerMap[id];

	if ok == false {
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
		return
	}

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

func sendMessage(id string, message []byte) {
	conn, ok := peerMap[id];
	if !ok {
		fmt.Println("Peer not connected : unable to send message")
		return
	}

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