package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

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

	cmd := exec.Command("amixer", "get", "Master")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error reading volume from amixer")
		return
	}

	lines := strings.Split(string(output), "\n")
	var text string
	for _, line := range lines {
		if strings.Contains(line, "%") {
			text = line
			break
		}
	}

	openBraceIdx := strings.LastIndex(text, "[")
	closeBraceIdx := strings.LastIndex(text, "]")

	mutedStatus := text[openBraceIdx+1 : closeBraceIdx]
	muted := mutedStatus != "on"

	button := os.Getenv(("button"))
	switch button {
	case "1":
		if muted {
			err = exec.Command("amixer", "set", "Master", "unmute").Run()
		} else {
			err = exec.Command("amixer", "set", "Master", "mute").Run()
		}

		if err != nil {
			fmt.Println(err)
		}

	case "3":
		exec.Command("pavucontrol").Run()
	}

	if len(text) <= 0 {
		fmt.Println("No volume found")
	}

	openBraceIdx = strings.Index(text, "[")
	closeBraceIdx = strings.Index(text, "]")

	var icon string
	if muted {
		icon = "󰖁"
	} else {
		icon = " "
	}

	text = icon + text[openBraceIdx+1:closeBraceIdx]

	fmt.Println(pango.Powerline("", text, foreground, background, false))
}
