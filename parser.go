package main

import (
	"fmt"
	"strconv"
)

/**
This file is related to parsing bencode to ASCII & vice-versa.
**/

func decode(input string)any{
	if len(input) == 0 {
		return nil;
	}

	inputType := input[0];
	if inputType == 'i' {
		return decodeInteger(input)
	}else if inputType == 'd'{
		return decodeDictionary(input)
	}else if inputType == 'l' {
		return decodeList(input)
	}else {
		return decodeString(input)
	}

		return nil;
}

func decodeString(input string) string{
	inputLen := len(input);
	index := -1;
	for i := 0; i < inputLen; i++ {
		if input[i] == ':' {
			index = i;
			break;
		}
	}

	if index == -1 {
		return "";
	}

	stringLen,_ := strconv.Atoi(input[:index]);

	if stringLen < 0 {
		//add a log here
		return  "";
	}

	start:= index + 1;
	end := index+stringLen+1;

	if end > len(input) {
		//add a log here

		return "";
	}

	outputString := input[start : end];

	return outputString;
}

func decodeInteger(input string) int{
	length := len(input)
	integer := input[1:length - 1];
	val,_ := strconv.Atoi(integer);
	//maybe better  error handling here
	return val;
};

func decodeList(input string)[] any{

};  

func decodeDictionary(input string)map[any]any{
 
};




