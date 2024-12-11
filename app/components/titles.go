package components

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Titles() Node {
	return Div(Class("titles"),
		H2(Style("float: left"), Raw("soph.systems")),
		H2(Style("float: right"), Raw("ᓚᘏᗢ")),
	)
}
