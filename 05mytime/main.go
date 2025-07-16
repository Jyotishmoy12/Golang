package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	presentTime :=time.Now()
	fmt.Println(presentTime)
    // formatting the time and date based on our needs 
	// but why that particular date and time?
	// its because we want to print it in a particular format
	// and its a standard and its in documentation of golang 
	fmt.Println(presentTime.Format("01-02-2006 Monday 15:04:05 PM"))

	// create a date
	// time.Date(year, month, day, hour, minute, second, nanoseconds, location)
	createdDate := time.Date(2025, time.February, 10, 23, 45, 0, 0, time.Local)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday 15:04:05 PM"))
}