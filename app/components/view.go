package components

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	. "maragu.dev/gomponents/html"
)

func View(ajax bool, title string, elements ...Node) Node {
	return HTML5(HTML5Props{
		Title:    title,
		Language: "en",
		Head: []Node{
			If(ajax, Script(
				Src("https://cdnjs.cloudflare.com/ajax/libs/jquery/3.6.0/jquery.min.js"),
				Integrity("sha512-894YE6QWD5I59HgZOGReFYm4dnWc1Qt5NtvYSaNcOP+u1T9qYdvdihz0PPSiiqn/+/3e7Jo4EaG7TubfWGUrMQ=="),
				CrossOrigin("anonymous"),
			)),
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
