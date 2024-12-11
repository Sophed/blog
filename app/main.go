package main

import (
	"site/app/views"
	"site/generation"

	"github.com/sophed/lg"
)

func main() {
	lg.Info("generating posts...")
	err := generation.GeneratePosts()
	if err != nil {
		lg.Fatl(err)
	}
	lg.Info("generating static pages...")
	generation.Static(views.Index(), "index")
	generation.Static(views.Resume(), "resume")
	generation.Static(views.Socials(), "socials")
	generation.Static(views.Photography(), "photography")
	generation.Static(views.Blog(), "blog")
	lg.Info("done!")
}
