package main

import (
	"fmt"
	"i3blocks/time/cmd/pango"
	"strings"
	"time"
)

func main() {
	time := strings.ToLower(time.Now().Format(time.Kitchen))

	light := "#769ff0"
	dark := "#1f1f1f"

	edge := pango.Span("", pango.Args{
		Bg:   dark,
		Fg:   light,
		Size: "15pt",
	})

	text := pango.Span(time, pango.Args{
		Bg:   light,
		Fg:   dark,
		Size: "14pt",
	})

	fmt.Println(edge + text)
}
