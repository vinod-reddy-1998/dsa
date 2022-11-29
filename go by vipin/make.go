package main

import "fmt"

func main() {
	var temp = make(map[string]int)
	temp["vinod"] = 22
	temp["v"] = 23
	temp["vi"] = 24
	fmt.Println(temp)
	delete(temp, "v")
	fmt.Println(temp)
	_, check := temp["v"]
	fmt.Println(check)

}
