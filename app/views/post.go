package views

import (
	"site/app/components"

	. "maragu.dev/gomponents"
)

func Post(title, content string) Node {
	return components.View(false, title, Raw(content))
}
