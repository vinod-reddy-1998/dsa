package main

import "fmt"

func main() {
	// var data = [6]int{2, 3, 4, 5, 6}
	// fmt.Println(data)
	// for i := 0; i < len(data); i++ {
	// 	fmt.Println(data[i])
	// }
	var data2 = [2][3]int{{2, 3, 4}, {5, 6, 7}}
	fmt.Println(data2)
	for i := 0; i < len(data2); i++ {
		for j := 0; j < len(data2[i]); j++ {
			fmt.Print(data2[i][j], " ")
			// to give a space we have to give " "#space...
		}
		fmt.Println()
	}
}
