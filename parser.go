package main

import (
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
	var list []any;
	size := len(input)
	// input = input[1 : size-1];
	// size = len(input);

	for i:= 1 ; i < size -1 ; {
		 if input[i] == 'i' || input[i] == 'l' || input[i] == 'd' {
			j := i;
			for input[j] != 'e' {
				j++;
			}

			currEle:= input[i:j+1];
			list = append(list, decode(currEle));
			 i = j + 1;
		 }else{
			j := i;
			for input[j] != ':' {
				j++;
			}

			stringSize, _ := strconv.Atoi(input[i:j]);
			
			list = append(list, decode(input[i:j+stringSize + 1]));
			i = j + stringSize + 1;
		 }
	}
	return  list;
};  

func decodeDictionary(input string)map[any]any{
   dict := make(map[any]any)
	length := len(input);

	for i:=1 ; i < length -1; {
		//get key 
		 j := i ;

		//  fmt.Println("key starting from : ", i);
		 for input[j] != ':' {
			j++;
		 }

		 stringSize,_ := strconv.Atoi(input[i:j]);

		 
		 key := input[j+1: j+stringSize+1];

		//  fmt.Println(key);
		//  fmt.Println(stringSize);


		 j = j + stringSize + 1;

		var value any;

		// fmt.Println("value of j : ", j);

		 if input[j] == 'i' || input[j] == 'l' || input[j] == 'd' {
			start := j;
			end := j;
			
			for input[end] != 'e' {
				end++;
			}

			value = decode(input[start : end+1]);

			// fmt.Println("value of end : ", end);
			i = end + 1;

			// fmt.Println("value : %v",  value);
		 }else {
			start := j;
			end := j;
			for input[end] != ':' {
				end++;
			}

			stringSize,_ := strconv.Atoi(input[start:end]);
			
			value = input[end + 1 : end + 1 + stringSize];

			i = end + 1 + stringSize;
		 }

		 dict[key] = value;
	}
   return dict;
};




