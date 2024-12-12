package components

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	. "maragu.dev/gomponents/html"
)

func View(title string, elements ...Node) Node {
	return HTML5(HTML5Props{
		Title:    title,
		Language: "en",
		Head: []Node{
			Link(Rel("icon"), Type("image/x-icon"), Href("/static/favicon.png")),
			Link(Rel("stylesheet"), Href("https://sophed.github.io/iosevka-webfont/3.4.1/iosevka.css")),
			Link(Rel("stylesheet"), Href("/css/global.css")),
			Link(Rel("stylesheet"), Href("/css/nav.css")),
		},
		Body: []Node{
			Div(Class("wrapper"),
				Titles(),
				NavBar(),
				Hr(),
				Div(Class("contents"),
					Group(elements),
				),
			),
		},
	})
}
