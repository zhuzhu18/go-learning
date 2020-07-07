package main

import (
	"bufio"
	"fmt"
	"os"
	"io"
	"strconv"
	"strings"
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

func printFile(filename string){
	file, err := os.Open(filename)
	if err != nil{
		panic(err)
	}else{
		printFileContents(file)
	}
}

func printFileContents(reader io.Reader){
	scanner := bufio.NewScanner(reader)
		for scanner.Scan(){
			fmt.Println(scanner.Text())
		}
}
func main(){
	fmt.Println(convertToBin(5))
	fmt.Println(convertToBin(13))
	fmt.Println(convertToBin(100))
	fmt.Println(convertToBin(0))
	fmt.Println(convertToBin(1))
	printFile("abc.txt")
	s := `abc"d"
		kkkk
		123
		
		p`
	printFileContents(strings.NewReader(s))
}