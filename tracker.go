package main

import (
	"crypto/sha1"
	"math/rand/v2"
	"net/url"
	"strconv"
)


type Event string

var peerId string

const (
	started   Event = "started"
	stopped   Event = "stopped"
	completed Event = "completed"
)

// peer_id convention == '-CCDDDD-'
func getTracker(dict map[string]any) {
	info := dict["info"]; 
	//convert this info to bencode again
	info_hash := sha1.Sum([]byte(encode(info)))
	ueInfoHash := url.QueryEscape(string(info_hash[:]));
	generatePeerId();
	peerid_hash := sha1.Sum([]byte(peerId));
	uePeerIdHash := url.QueryEscape(string(peerid_hash[:]))
	downloaded := 0;
	uploaded := 0;
	compact := false;
	no_peer_id := false;
	
	fileInfo := dict["info"].(map[string]any);
	var size int;
	length, ok := fileInfo["length"];

	if ok {
		size = length.(int);
	}else {
		for _,file := range fileInfo["files"].([]any){
			size+=(file.(map[string]any)["length"]).(int);
		}
	}

	left := size;
	event := started;

    sendRequest([]byte(ueInfoHash),  []byte(uePeerIdHash), port, uploaded, downloaded, left, compact, no_peer_id, event)
}

func generatePeerId(){
	var id string
	id += "-"

	ch1 := rand.IntN(26) + 'a'
	id += string(ch1) 

	ch2 := rand.IntN(26) + 'a'
	id += string(ch2)

	d1 := rand.IntN(9);
	id += strconv.Itoa(d1)

	d2 := rand.IntN(9);
	id += strconv.Itoa(d2)

	id+= "-";

	peerId = id;
}
  

func sendRequest(info_hash []byte, peer_id []byte, port int, uploaded int, downloaded int, left int, compact bool, no_peer_id bool, event Event) {

}