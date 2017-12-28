package main

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

var markup = `
  <!DOCTYPE html>
  <html lang="en">
  <head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <meta http-equiv="X-UA-Compatible" content="ie=edge">
  <title>Document</title>
  </head>
  <body>

  </body>
  </html>
`

func main() {
	r := strings.NewReader(markup)
	z := html.NewTokenizer(r)

	/* 1. */
	ListAllTagTypes(z)
}

// Example of printing all types of tags in an HTML string
func ListAllTagTypes(z *html.Tokenizer) {
	for {
		tt := z.Next()

		if tt == html.ErrorToken {
			return
		}

		switch tt {
		case html.StartTagToken:
			tag, _ := z.TagName()
			fmt.Println(string(tag))
		}
	}
}
