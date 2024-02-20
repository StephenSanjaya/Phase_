package main

import (
	"fmt"
	"strings"
)

func convertToByteAndChangeToAlay(mergeStr string)(string) {
	newStr := []byte(mergeStr)

	for i := 0; i < len(mergeStr); i++ {
		if(string(newStr[i]) == "a"){
			newStr[i] = byte('4')
		}else if(string(newStr[i]) == "e"){
			newStr[i] = byte('3')
		}else if(string(newStr[i]) == "i"){
			newStr[i] = byte('!')
		}else if(string(newStr[i]) == "l"){
			newStr[i] = byte('1')
		}else if(string(newStr[i]) == "n"){
			newStr[i] = byte('N')
		}else if(string(newStr[i]) == "s"){
			newStr[i] = byte('5')
		}else if(string(newStr[i]) == "x"){
			newStr[i] = byte('*')
		}
	}
	return string(newStr)
}

func alayGen(str ...string)(string) {
	mergeStr := strings.Join(str, " ")
	return convertToByteAndChangeToAlay(mergeStr);
}

func fibonacci(n int)(int)  {
	if(n == 1 || n == 2){
		return 1
	}
	return fibonacci(n-2) + fibonacci(n-1)
}

func main() {

	// functionNG Challenge 4 : function 2
	str := alayGen("hello", "world", "zzz", "The quick brown fox jumps over the lazy dog")
	fmt.Println(str)

	// NG Challenge 4 : function 3
	fib := fibonacci(14)
	fmt.Println(fib)
}