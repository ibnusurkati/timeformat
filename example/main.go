package main

import (
	"fmt"

	"github.com/ibnusurkati/timeformat"
)

func main() {
	// Get date  now
	DateNow()
	// Get date time now
	DateTimeNow()

	fmt.Println("===============================")

	// Parse date time
	ParseDateTime()

}

func DateNow() {
	date := timeformat.Now("YYYY-MM-DD")
	fmt.Println("DATE :", date)
}

func DateTimeNow() {
	datetime := timeformat.Now("YYYY-MM-DD hh:mm:ss")
	fmt.Println("DATETIME :", datetime)
}

func ParseDateTime() {
	datetime, err := timeformat.Parse("2021-09-13", "YYYY-MM-DD", "DD MMMM YYYY")

	if err != nil {
		fmt.Println("ERROR :", err)
		return
	}
	fmt.Println("DATETIME :", datetime)
}
