package main

import "fmt"

func longestNonrepeatedSubstring(s string)int{
	start:=0
	lastOccured:=make(map[byte]int)
	maxLen := 0
	for i, ch := range []byte(s) {
		lastI, ok := lastOccured[ch]
		if ok && lastI >= start{
			start = lastOccured[ch] + 1
		}
		if i - start + 1 > maxLen{
			maxLen = i - start + 1
		}
		lastOccured[ch] = i
	}
	return maxLen
}
func main(){
	fmt.Println(longestNonrepeatedSubstring("abcabcbb"))
	fmt.Println(longestNonrepeatedSubstring("bbbbb"))
	fmt.Println(longestNonrepeatedSubstring("pwwkew"))
}