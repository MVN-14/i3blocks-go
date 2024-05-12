package main

import (
	"fmt"
	"os"
	"strconv"

	pango "github.com/MVN-14/panggo"
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

	percent, err := os.ReadFile("/sys/class/power_supply/BAT0/capacity")
	if err != nil || len(percent) <= 0 {
		fmt.Println("Error reading BAT0 capacity")
		return
	}

	percentStr := fmt.Sprintf("%c %s%% ", '󰁹', string(percent[:len(percent)-1]))
	fmt.Println(pango.Powerline("", percentStr, foreground, background, true))
}
