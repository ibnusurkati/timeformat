# Time Format

[![GoDoc](https://godoc.org/github.com/ibnusurkati/timeformat?status.svg)](https://pkg.go.dev/github.com/ibnusurkati/timeformat?tab=doc)
[![License](https://img.shields.io/github/license/ibnusurkati/timeformat?style=plastic)](https://github.com/ibnusurkati/timeformat/blob/master/LICENSE)
[![Release](https://img.shields.io/github/v/release/ibnusurkati/timeformat.svg?style=flat-square)](https://github.com/ibnusurkati/timeformat/releases)
a simple plugin to change date and time format

## Features
- parsing datetime to format global

## Installing

go mod:

```bash
go get github.com/ibnusurkati/timeformat
```

## Example

```go
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

```

### Format Time
|Format|Format Go-Language|Description                                    |
|------|------------------|-----------------------------------------------|
|hh    |15                |24-hour format of an hour with leading zeros   |
|h     |3                 |12-hour format of an hour without leading zeros|
|mm    |04                |Minutes with leading zeros                     |
|m     |4                 |Minutes without leading zeros                  |
|sss   |05                |Seconds with milliseconds                      |
|ss    |05                |Seconds with leading zeros                     |
|s     |5                 |without with leading zeros                     |
|A     |PM                |Uppercase Ante meridiem and Post meridiem      |
|a     |pm                |Lowercase Ante meridiem and Post meridiem      |

### Format Date
|Format|Format Go-Language|Description                                                       |
|------|------------------|------------------------------------------------------------------|
|YYYY  |2006              |A full numeric representation of a year, 4 digits                 |
|YY    |06                |A two digit representation of a year                              |
|MMMM  |January           |A full textual representation of a month, such as January or March|
|MMM   |Jan               |A short textual representation of a month, three letters          |
|MM    |01                |Numeric representation of a month, with leading zeros             |
|M     |1                 |Numeric representation of a month, without leading zeros          |
|DDDD  |Monday            |A full textual representation of the day of the week              |
|DDD   |Mon               |A textual representation of a day, three letters                  |
|DD    |02                |Day of the month, 2 digits with leading zeros                     |
|D     |2                 |Day of the month without leading zeros                            |