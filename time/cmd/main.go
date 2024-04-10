package main

import (
	"fmt"
	"i3blocks/time/cmd/pango"
	"time"
)





func main() {
  time:= time.Now().Format(time.Kitchen)
  
  bg := "#769ff0"
  fg := "#1f1f1f"

  output := pango.Span(time, pango.Args{
    Bg: bg,
    Fg: fg,
  })

  fmt.Println(output)
}
