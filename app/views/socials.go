package views

import (
	"site/app/components"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Socials() Node {
	return components.View("soph.systems ~ socials",
		P(Text("socials!!")),
	)
}
