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

func buildHave(pieceIndex uint32) []byte {
  buf := make([]byte, 9);
  binary.BigEndian.PutUint32(buf[0:], 5)
  buf[4] = byte(4);
  binary.BigEndian.PutUint32(buf[5:9], pieceIndex);  
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
return buf;
}

/*
Clients should drop the connection if they receive bitfields that are not of the correct size, or if the bitfield has any of the spare bits set.
*/

func buildBlockRequest(index uint32, begin uint32, length uint32)[]byte {
  buf := make([]byte, 17);
  binary.BigEndian.PutUint32(buf[0:], 13);
  buf[4] = 6;
  binary.BigEndian.PutUint32(buf[5:], index);
  binary.BigEndian.PutUint32(buf[9:], begin);
  binary.BigEndian.PutUint32(buf[13:], length);
  return buf;
}

func buildPiece(index uint32, begin uint32, block []byte)[] byte{
  x := uint32(len(block));
  buf := make([]byte, 9+x+4);
  binary.BigEndian.PutUint32(buf[0:], 9 + x);
  buf[4] = 8;
  binary.BigEndian.PutUint32(buf[5:] , index);
  binary.BigEndian.PutUint32(buf[9:] , begin);
  copy(buf[13:], block[0:])
  return buf;
}

func buildCancel(index uint32, begin uint32, length uint32) []byte{
 buf := make([]byte, 17);
 binary.BigEndian.PutUint32(buf[0:], 13);
 buf[4] = 8;
 binary.BigEndian.PutUint32(buf[5:], index); 
 binary.BigEndian.PutUint32(buf[9:], begin);
 binary.BigEndian.PutUint32(buf[13:], length);
 return buf;
}

/*
 
All current implementations use 2^15 (32KB), and close connections which request an amount greater than 2^17 (128KB)." */