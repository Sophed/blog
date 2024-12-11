package views

import (
	"site/app/components"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Photography() Node {
	return components.View("soph.systems ~ photography",
		P(Text("photography!!")),
	)
}
