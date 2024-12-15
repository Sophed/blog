package views

import (
	"site/app/components"
	"strconv"
	"strings"
	"unicode"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Post(post *BlogPost, body string) Node {
	return components.View(false, post.Title,
		P(Text(
			"Published: "+post.DateStr+
				" ~ "+strconv.Itoa(readingTime(string(post.Body)))+"m read",
		)),
		H2(Text(post.Title)),
		Raw(string(body)),
	)
}

func readingTime(body string) int {
	specials := ""
	for _, c := range body {
		if !unicode.IsLetter(c) && !unicode.IsNumber(c) && c != ' ' {
			specials += string(c)
		}
	}
	for _, c := range specials {
		body = strings.ReplaceAll(body, string(c), "")
	}
	words := 0
	for _, s := range strings.Split(body, " ") {
		if s == "" {
			continue
		}
		words++
	}
	return words / 200
}
