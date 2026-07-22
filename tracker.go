package main

import (
	crand "crypto/rand"
	"crypto/sha1"
	"fmt"
	"io"
	"math/rand/v2"
	"net/http"
	"net/url"
	"strconv"
)

var info_hash []byte;
type Event string;

var peerId []byte

const (
	started   Event = "started"
	stopped   Event = "stopped"
	completed Event = "completed"
)

// peer_id convention == '-CCDDDD-'
func getTracker(dict map[string]any) map[string]any{
	info := dict["info"]; 
	//convert this info to bencode again
	info_hash := sha1.Sum([]byte(encode(info)))
	// ueInfoHash := url.QueryEscape(string(info_hash[:]));
	generatePeerId();
	// uePeerIdHash := url.QueryEscape(string(peerid_hash[:]))
	downloaded := 0;
	uploaded := 0;
	compact := 0;
	no_peer_id := 0;
	
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
	baseUrl := dict["announce"].(string)
    content := sendRequest(baseUrl, info_hash[:],  peerId, strconv.Itoa(port), strconv.Itoa(uploaded), strconv.Itoa(downloaded), strconv.Itoa(left), strconv.Itoa(compact), strconv.Itoa(no_peer_id), event)
	response := decodeBenCode(content).(map[string]any);

	fmt.Print(response)
	//TO-DO : check for the failure field,
	return response;
}

func generatePeerId(){
	var id string
	id += "-"

	ch1 := rand.IntN(26) + 'a';
	id += string(ch1);

	ch2 := rand.IntN(26) + 'a';
	id += string(ch2);

	d1 := rand.IntN(9);
	id += strconv.Itoa(d1)

	d2 := rand.IntN(9);
	id += strconv.Itoa(d2)

	d3:= rand.IntN(9);
	id += strconv.Itoa(d3)

	d4 := rand.IntN(9);
	id += strconv.Itoa(d4)

	id+= "-";

	prefix := []byte(id)
	rest := make([]byte,12)
	crand.Read(rest)
	prefix = append(prefix, rest...)
	peerId = prefix	
}
  

func sendRequest(baseUrl string, info_hash []byte, peer_id []byte, port string, uploaded string, downloaded string, left string, compact string, no_peer_id string, event Event) string{
     reqUrl,_ := url.Parse(baseUrl);
	 params := reqUrl.Query();
	 params.Set("info_hash", string(info_hash));
	 params.Set("peer_id", string(peer_id));
	 params.Set("uploaded", uploaded);
	 params.Set("compact", compact);
	 params.Set("no_peer_id", no_peer_id);
	 params.Set("event", string(event));
	 params.Set("downloaded", downloaded);
	 params.Set("left", left);
	 params.Set("port", port)

	 reqUrl.RawQuery = params.Encode();

	 resp, err1 := http.Get(reqUrl.String())

	 if err1 != nil {
		 fmt.Printf("error : ", err1);
		  return "";
		}
		
	databytes, _ := io.ReadAll(resp.Body);
	 defer resp.Body.Close();
		
	 content := string(databytes);

	//  fmt.Printf("respone : ", content);
		return content
}