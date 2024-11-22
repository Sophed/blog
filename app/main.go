package main

import (
	"log"
	"site/generation"
)

func main() {
	init := generation.InitDir(generation.POSTS_DIR)
	if init {
		println("add markdown files to ./posts/ to get started!")
		return
	}

	println("generating posts...")
	err := generation.GeneratePosts()
	if err != nil {
		log.Fatal(err)
	}
	println("generated posts!")
}
