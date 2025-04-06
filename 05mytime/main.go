package main

import (
	"fmt"
	"time"
)

func main() {
	/**
	for formating the time, everytime you should use same date i.e. - 01-02-2006 15:04:05 Monday
	if you want only date -> 01-02-2006
	If you want date and time -> 01-02-2006 15:04:05
	If you want date, time and day -> 01-02-2006 15:04:05 Monday
	*/
	fmt.Println("Welcome to time study of golang")
	presentTime := time.Now()
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))
	fmt.Println(presentTime.UnixMilli())
	createdTime := time.Date(2020, time.November, 3, 00, 00, 00, 00, time.Local)
	fmt.Println(createdTime)
	fmt.Println(createdTime.Format("01-02-2006 15:04:05 Monday"))
}
