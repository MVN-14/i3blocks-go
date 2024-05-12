package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/MVN-14/panggo"
)

func main() {
	var foreground string
	var background string

	index, err := strconv.Atoi(os.Getenv("idx"))
	if err != nil {
		fmt.Println("NO INDEX")
		return
	}

	if index%2 != 0 {
		foreground = os.Getenv("foreground")
		background = os.Getenv("background")
	} else {
		background = os.Getenv("foreground")
		foreground = os.Getenv("background")
	}

	dateString := time.Now().Format(time.DateOnly)

	fmt.Println(pango.Powerline("", dateString, foreground, background, false))
}
