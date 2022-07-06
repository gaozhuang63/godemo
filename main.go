package main

import "fmt"

func main() {
	fmt.Println("hello,world!")
	var stockcode = 123
	var url = "Code=%d&endDate=%s"
	var enddate = "2022-7-6"
	var target = fmt.Sprintf(url, stockcode, enddate)
	fmt.Println(target)

}
