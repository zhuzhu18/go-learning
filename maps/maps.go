package main

import "fmt"

func main() {
	m1 := map[string]string{"a": "1",
		"b": "2",
		"c": "3"}
	for k, v := range m1 {
		fmt.Println(k, v)
	}
	fmt.Println(m1["m"])
	if v, ok := m1["m"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("key does not exist")
	}
	delete(m1, "d") // 删除不存在的key不会报错
	m1["e"] = "4"
	fmt.Println(m1)
}
