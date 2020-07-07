package main

import (
	"bufio"
	"errhandling/fib"
	"errors"
	"fmt"
	"os"
)

func writeFile(filename string) {
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	err = errors.New("this is a custom error") // 自定义一个error
	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Printf("%s, %s, %s\n", pathError.Op,
				pathError.Path,
				pathError.Err)
		}
		return
	} else {
		defer file.Close()
		writer := bufio.NewWriter(file)
		defer writer.Flush()

		f := fib.Fib()
		for i := 0; i < 30; i++ {
			fmt.Fprintln(writer, f())
		}
	}
}
func main() {
	writeFile("fibonacci.txt")
}
