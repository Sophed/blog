package views

import (
	"site/app/components"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Index() Node {
	return components.View(true, "soph.systems",
		Script(Src("/static/status.js")),
		H2(Text("Software Engineer, Designer, & Linux Evangelist")),
		P(Text("Status: Loading..."), ID("status")),
		P(Strong(Text("Hey! ")), Text("I'm Sophia (she/her), A student software engineer from the UK learning systems and back-end development.")),
		P(
			A(Text("GitHub"), Href("https://github.com/sophed"), Target("_blank")),
			Text(" - "),
			A(Text("Mastodon"), Href("https://tech.lgbt/@null"), Target("_blank")),
			Text(" - "),
			A(Text("Bluesky"), Href("https://bsky.app/profile/soph.cat"), Target("_blank")),
		),
		H2(Text("Projects")),
		Ul(
			project(
				"dotfiles",
				"https://github.com/sophed/dotfiles",
				"My personal Linux desktop and terminal configuration files",
			),
			project(
				"mcl",
				"https://github.com/sophed/mcl",
				"Minecraft profile lookup tool",
			),
			project(
				"OCI-Claimer",
				"https://github.com/sophed/oci-claimer",
				"Automatically claim Oracle Cloud free-tier instances",
			),
		),
		P(Text("The best way to contact me is via my discord - "), Strong(Text("@sophed"))),
		H2(Text("Chatbox")),
		components.ChatBox(6),
		P(Text("Other online users will see your messages! Please don't make me write moderation tools!")),
	)
}

func project(title, link, description string) Node {
	return Li(P(
		A(Text(title), Href(link), Target("_blank")),
		Text(" : "+description),
	))
}
