package main

import (
	"fmt"
	"io/ioutil"
)

func grade(score int)string{
	g:=""
	switch {
	case score<0 || score>100:
		panic("invalid score")
	case score<60:
		g = "F"
	case score < 80:
		g = "D"
	case score<90:
		g="C"
	case score<100:
		g = "B"
	default:
		g = "A"
	}
	return g
}
func main() {
	const filename = "abc.txt"
	if contents, err := ioutil.ReadFile(filename); err!=nil{
	}else{
		fmt.Printf("%s", contents)
	}
	fmt.Println(grade(54),grade(68),grade(88),grade(100))
}
