package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	pango "github.com/MVN-14/panggo"
)

func main() {
	var foreground string
	var background string

	const batteryPath = "/sys/class/power_supply/BAT0/"

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

	bytes, err := os.ReadFile(batteryPath + "capacity")
	if err != nil || len(bytes) <= 0 {
		fmt.Println("Error reading BAT0 capacity")
		return
	}

	percent := string(bytes[:len(bytes)-1])
	symbol := '󰁹'

	bytes, err = os.ReadFile(batteryPath + "status")
	if err != nil {
		fmt.Println("Error reading BAT0 status")
		return
	}
	if strings.Trim(string(bytes), " \n\r\t") == "Charging" {
		symbol = '󰂄'
	}
	percentStr := fmt.Sprintf("%c %s%% ", symbol, percent)
	fmt.Println(pango.Powerline("", percentStr, foreground, background, false))
}
