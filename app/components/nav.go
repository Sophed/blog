package components

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func NavBar() Node {
	return Ul(Class("nav"),
		Li(Class("left"),
			A(Href("/"), Text("home")),
		),
		Li(Class("left"),
			A(Href("/resume"), Text("resume")),
		),
		Li(Class("right"),
			A(Href("/posts"), Text("blog")),
		),
		Li(Class("right"),
			A(Href("/photography"), Text("photography")),
		),
	)
}
