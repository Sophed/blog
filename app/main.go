package main

import (
	"site/generation"

	"github.com/sophed/lg"
)

func main() {
	init := generation.InitDir(generation.POSTS_DIR)
	if init {
		lg.Info("add markdown files to ./posts/ to get started!")
		return
	}

	lg.Info("generating posts...")
	err := generation.GeneratePosts()
	if err != nil {
		lg.Fatl(err)
	}
	lg.Info("generated posts!")
}
