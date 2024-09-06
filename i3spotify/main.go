package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	//"strings"

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
		foreground = os.Getenv("background")
		background = os.Getenv("foreground")
	}

	cmd := exec.Command("playerctl", "--player=spotify", "status")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error reading status from playerctl: ", err)
		return
	}

	//status := strings.Split(string(output), "\n")[0]

	cmd = exec.Command("playerctl", "--player=spotify", "metadata", "artist")
	output, err = cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	artist := strings.Split(string(output), "\n")[0]

	cmd = exec.Command("playerctl", "--player=spotify", "metadata", "title")
	output, err = cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	title := strings.Split(string(output), "\n")[0]

	text := fmt.Sprintf("%s %s - %s ", "󰎇", artist, title)

	fmt.Println(pango.Powerline("", text, foreground, background, true))
}
