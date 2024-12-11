package views

import (
	"site/app/components"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Index() Node {
	return components.View("soph.systems",
		P(Text("bwahh!!")),
	)
}
