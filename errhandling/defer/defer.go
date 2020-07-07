package main

import (
	"bufio"
	"errhandling/fib"
	"fmt"
	"os"
)

func writeFile(filename string){
	file, err := os.Create(filename)
	if err!=nil{
		panic(err)
	}else{
		defer file.Close()
		writer:= bufio.NewWriter(file)
		defer writer.Flush()

		f:=fib.Fib()
		for i:=0; i<30;i++{
			fmt.Fprintln(writer, f())
		}
	}
}
func main() {
	writeFile("fibonacci.txt")
}
