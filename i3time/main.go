package main

import (
	"fmt"
	"github.com/MVN-14/panggo"
	"os"
	"strconv"
	"strings"
	"time"
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

	time := strings.ToLower(time.Now().Format(time.Kitchen))

	fmt.Println(pango.Powerline("", time, foreground, background, false))
}
