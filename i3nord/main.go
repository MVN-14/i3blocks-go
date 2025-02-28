package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

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

	cmd := exec.Command("nordvpn", "status")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return
	}

	var symbol string
	if strings.Contains(string(output), "Connected") {
		symbol = "󰒘 "
	} else {
		symbol = " "
	}

	fmt.Println(pango.Powerline("", symbol, foreground, background, false))
}
