package views

import (
	"site/app/components"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Resume() Node {
	return components.View("soph.systems ~ resume",
		H2(Text("Experience")),
		P(Text("Most of my experience is in development with Go and Java, alongside online community management. I've worked with a wide range of modern technologies, chosen carefully for the best performance and developer experience.")),
		H2(Text("Stray.gg - Administrator & Back-End Developer (July 2022 - August 2024)")),
		P(Text("My role at Stray involved frequent community interactions and consistent gathering of feedback to improve the service we provided and ensure player safety. I was involved in the development of several projects such as back-end APIs, Discord integrations and contributing to core server plugins.")),
		H2(Text("RoyalKits - Game Developer & SysAdmin (October 2023)")),
		P(Text("During the time RoyalKits was active, I developed core game functionality and managed cloud infrastructure to allow for changes to be made easily by other project maintainers and community managers.")),
		H2(Text("HydraClient - Website Developer & Designer (June 2023)")),
		P(Text("Using design tools like Figma, I created mockups for the landing page and produced a functional version using modern web technologies.")),
		H2(Text("OmniFlow - Graphics Designer (July 2024)")),
		P(Text("I created UI mockups and branding for the project, taking into consideration feedback from project managers and adjusting accordingly.")),
	)
}
