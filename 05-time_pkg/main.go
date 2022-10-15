package main

import (
	"fmt"
	"time"
)

func main() {

	timeNow := time.Now()
	fmt.Println(timeNow)

	fmt.Println(timeNow.Format("01-02-2006 15:04:05 Monday"))

	fmt.Println(timeNow.Format("2006/01/02"))

	createdDate := time.Date(2020, time.August, 11, 00, 00, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))

}

/*
Basic full date  "Mon Jan 2 15:04:05 MST 2006" gives "Wed Feb 25 11:06:39 PST 2015"
Basic short date "2006/01/02" gives "2015/02/25"
AM/PM            "3PM==3pm==15h" gives "11AM==11am==11h"
No fraction      "Mon Jan _2 15:04:05 MST 2006" gives "Wed Feb 25 11:06:39 PST 2015"
0s for fraction  "15:04:05.00000" gives "11:06:39.12340"
9s for fraction  "15:04:05.99999999" gives "11:06:39.1234"
*/
