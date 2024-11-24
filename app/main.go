package main

import (
	"os"
	"site/generation"

	"github.com/sophed/lg"
)

func main() {
	init := generation.InitDir(generation.POSTS_DIR)
	if init {
		lg.Info("add markdown files to ./" + generation.POSTS_DIR + "/ to get started!")
		return
	}
	templates := generation.InitDir(generation.TEMPLATES_DIR)
	if templates {
		os.OpenFile(generation.TEMPLATES_DIR+"/post.html", os.O_RDONLY|os.O_CREATE, 0666)
		lg.Info("add markdown files to ./" + generation.TEMPLATES_DIR + "/ to get started!")
		return
	}

	lg.Info("generating posts...")
	err := generation.GeneratePosts()
	if err != nil {
		lg.Fatl(err)
	}
	lg.Info("generated posts!")
}
