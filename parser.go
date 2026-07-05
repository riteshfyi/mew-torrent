package main

import (
	"maps"
	"slices"
	"strconv"
)

/**
This file is related to decoding & encoding bencode.
**/

func decodeBenCode(input string) any {
	stack := []any{}
	n := len(input)

	for i := 0; i < n; i++ {
		ch := input[i]

		if ch >= '0' && ch <= '9' {
			j := i
			for j < n && input[j] != ':' {
				j++
			}

			stringSize, _ := strconv.Atoi(input[i:j]) 
			i = j + 1
			j = j + 1 + stringSize
			currString := input[i:j]
			stack = append(stack, currString)
			i = j - 1

		} else if ch == 'i' {
			j := i
			for j < n && input[j] != 'e' {
				j++
			}

			num, _ := strconv.Atoi(input[i+1 : j])

			stack = append(stack, num)
			i = j

		} else if ch == 'd' || ch == 'l' {
			stack = append(stack, rune(ch))
		} else if ch == 'e' {
			//ch == 'e'
			revStack := []any{}

			for len(stack) > 0 && (stack[len(stack)-1] != 'l') && (stack[len(stack)-1] != 'd') {
				revStack = append(revStack, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}

			if len(stack) == 0 {
				//throw an error
			}

			if stack[len(stack)-1] == 'l' {
				// 'l'
				list := []any{}
				stack = stack[0 : len(stack)-1]
				for len(revStack) > 0 {
					list = append(list, revStack[len(revStack)-1])
					revStack = revStack[:len(revStack)-1]
				}

				stack = append(stack, list)
			} else {
				// 'd'
				dict := make(map[string]any)
				stack = stack[0 : len(stack)-1]
				for len(revStack) >= 2 {
					key := revStack[len(revStack)-1]
					value := revStack[len(revStack)-2]
					dict[key.(string)] = value
					revStack = revStack[:len(revStack)-2]
				}

				stack = append(stack, dict)
			}
		}

		if len(stack) == 0 {
			//thorw an error
		}

	}
	return stack[len(stack)-1]
}


func encode(input any) string{
	var output string
	switch v := input.(type) {
	case map[string]any: 
		curr :="d";
		keys := slices.Collect(maps.Keys(v));
		slices.Sort(keys);
		for _, key := range keys {
			curr+=encode(key);
			curr+=encode(v[key]);
		}
		curr+="e";
		output = curr;
	case string:
		curr := strconv.Itoa(len(input.(string)))
		curr += ":"
		curr+=input.(string)
		output = curr;
	case int:
		curr := "i"
		curr+=strconv.Itoa(input.(int));
		curr+="e"
		output = curr;
	case []any:
		curr := "l"
		for _,val := range input.([]any){
			curr+=encode(val);
		}

		curr+="e";
		output = curr;
	}
	

	return output;
}