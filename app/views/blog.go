package views

import (
	"site/app/components"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

var BLOG_POSTS []string

func Blog() Node {
	return components.View("soph.systems ~ posts",
		Map(BLOG_POSTS, func(title string) Node {
			return blogLink(title)
		}),
	)
}

func blogLink(title string) Node {
	return Div(
		A(Href("/posts/"+title),
			Raw(title),
		),
	)
}
