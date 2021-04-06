package main

import "fmt"

func main() {
	fmt.Println(guess_numb([]int32{10, 15, 3, 7}, 16))
}

func guess_numb(list []int32, numb int32) bool {
	for i := 0; i < len(list); i++ {
		for j := i + 1; j < len(list); j++ {
			if j == i {
				continue
			}
			if list[i]+list[j] == numb {
				return true
			}
		}
	}
	return false
}
