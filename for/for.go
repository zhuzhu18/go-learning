package main

import (
		"fmt"
		"strconv"
		)

func convertToBin(num int)string{
	result:=""
	if num==0{
		result = "0"
	}else{
		for ;num>0;num/=2 {
			result = strconv.Itoa(num%2)+result
		}
	}
	return result
}
func main(){
	fmt.Println(convertToBin(5))
	fmt.Println(convertToBin(13))
	fmt.Println(convertToBin(100))
	fmt.Println(convertToBin(0))
	fmt.Println(convertToBin(1))
}