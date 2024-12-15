package components

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func ChatBox(messages int) Node {
	list := []int{}
	for i := range messages {
		list = append(list, i)
	}
	return Div(Class("chatbox"),
		Ul(ID("lines"),
			Map(list, func(i int) Node {
				return Li(Text("-"))
			}),
		),
		Input(
			ID("msg-input"),
			Placeholder("Enter to send message..."),
		),
		Script(Src("/static/chat.js")),
	)
}
