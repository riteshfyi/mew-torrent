package main

import "encoding/binary"

func buildKeepAlive() []byte {
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, 0);
	return buf
}

func buildHandshake(infoHash []byte) []byte {
	pstr := "BitTorrent protocol";
 pstrlen := len(pstr);
   buf := make([]byte, 49 + pstrlen)
   buf[0] = byte(pstrlen)
   copy(buf[1:], []byte(pstr))
   binary.BigEndian.PutUint64(buf[pstrlen+1:], 0)
   copy(buf[28:], infoHash)
   copy(buf[48:], []byte(peerId))
   return buf
}

func buildChoke() []byte {
  buf := make([]byte, 5)
  binary.BigEndian.PutUint32(buf[0:], 1)
  buf[4] = byte(0);
  return buf;
}

func buildUnChoke() []byte {
  buf := make([]byte, 5)
  binary.BigEndian.PutUint32(buf[0:], 1)
  buf[4] = byte(1);
  return buf;
}

func buildInterested() []byte {
	  buf := make([]byte, 5)
  binary.BigEndian.PutUint32(buf[0:], 1)
  buf[4] = byte(2);
  return buf;
}

func buildNotinterested() []byte {
buf := make([]byte, 5)
  binary.BigEndian.PutUint32(buf[0:], 1)
  buf[4] = byte(3);
  return buf;
}

func buildHave(pieceIndex int) []byte {
  buf := make([]byte, 9);
  binary.BigEndian.PutUint32(buf[0:], 5)
  buf[4] = byte(4);
  index := uint32(pieceIndex);
  binary.BigEndian.PutUint32(buf[5:9], index);  
  return buf
}

func buildBitfield(bitfield []byte) []byte {
 // assume the clinet itself will pass the buffer
 // little endian format 
 //we will stop after the last 1. assume rest as zero
 x := uint32(len(bitfield));
 buf := make([]byte, 4 + 1 + x);

  
binary.BigEndian.PutUint32(buf[0:],1 + x);
buf[4] = 5;
copy(buf[5:], bitfield)
return  buf;
}

/*
Clients should drop the connection if they receive bitfields that are not of the correct size, or if the bitfield has any of the spare bits set.
*/

func _()