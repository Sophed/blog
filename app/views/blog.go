package views

import (
	"site/app/components"
	"strconv"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

type BlogPost struct {
	ID    string `kdl:"-"`
	Title string `kdl:"title"`
	Date  string `kdl:"date"`
	Body  []byte `kdl:"-"`
}

var BLOG_POSTS []*BlogPost

func Blog() Node {
	return components.View(false, "soph.systems ~ posts",
		Map(BLOG_POSTS, func(post *BlogPost) Node {
			return blogLink(post)
		}),
	)
}

func blogLink(post *BlogPost) Node {
	return Div(
		A(Href("/posts/"+post.ID),
			Raw(post.Title),
		),
		P(Text(
			"Published: "+post.Date+
				" ~ "+strconv.Itoa(readingTime(string(post.Body)))+"m read",
		)),
		Br(),
	)
}
