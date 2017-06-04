package main

import "fmt"

func isLeapYear(yr int) bool {
	if (yr%4 == 0 && yr%100 != 0) || yr%400 == 0 {
		return true
	}
	return false
}

func main() {
	fmt.Println(isLeapYear(2016))
}
