package views

import (
	"site/app/components"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Resume() Node {
	return components.View("soph.systems ~ resume",
		P(Text("resume!!")),
	)
}
