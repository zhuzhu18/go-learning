package main

import "fmt"

func main() {
	var s1 []int
	fmt.Println(len(s1), cap(s1))

	arr := []int{1, 2, 3, 4, 5}
	s1 = arr[:]
	s2 := make([]int, 16)
	copy(s2, s1)
	fmt.Println(s2)
	s3 := append(s2[:2], s2[3:]...)
	fmt.Println(s3)
	fmt.Println(len(s3), cap(s3))
	popElement := s3[len(s3)-1]
	s3 = s3[:len(s3)-1]
	fmt.Println(popElement)
	fmt.Println(len(s3), cap(s3))
}
