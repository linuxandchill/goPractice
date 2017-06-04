package main

import "fmt"

/*

Given a year, report if it is a leap year.
on every year that is evenly divisible by 4
  except every year that is evenly divisible by 100
	    unless the year is also evenly divisible by 400

*/

func isLeapYear(yr int) bool {
	if (yr%4 == 0 && yr%100 != 0) || yr%400 == 0 {
		return true
	}
	return false
}

func main() {
	fmt.Println(isLeapYear(2016))
}
