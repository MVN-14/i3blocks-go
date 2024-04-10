package pango

import "fmt"

type Args struct {
  Bg string
  Fg string
  Size string
}

func Span(s string, a Args) string {
  argString := ""
  
  if a.Bg != "" {
    argString += fmt.Sprintf("background=\"%s\"", a.Bg)
  }
if a.Fg != "" {
    argString += fmt.Sprintf("foreground=\"%s\"", a.Fg)
  }
if a.Size != "" {
    argString += fmt.Sprintf("background=\"%s\"", a.Size)
  }

  return fmt.Sprintf("<span%s>%s</span>", argString, s)
}
