package generation

import (
	"bytes"
	"os"
	"site/app/views"
	"strings"
)

const POSTS_DIR = "posts"
const BUILD_DIR = "build"
const STYLES_DIR = "styles"
const STATIC_DIR = "static"

func GeneratePosts() error {
	err := clearBuildDir()
	if err != nil {
		return err
	}
	err = copyStatics(STYLES_DIR, "css")
	err = copyStatics(STATIC_DIR, "static")
	if err != nil {
		return err
	}
	entries, err := os.ReadDir(POSTS_DIR)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		id := strings.Split(entry.Name(), ".")[0]
		println("-", id)
		data, err := os.ReadFile(POSTS_DIR + "/" + entry.Name())
		if err != nil {
			return err
		}
		post, err := getMeta(data)
		if err != nil {
			return err
		}
		post.ID = id
		page := new(bytes.Buffer)
		views.Post(post, convertMDtoHTML(post.Body)).Render(page)
		err = os.WriteFile(BUILD_DIR+"/posts/"+post.ID+".html", []byte(page.String()), 0644)
		if err != nil {
			return err
		}
		views.BLOG_POSTS = append(views.BLOG_POSTS, post)
	}
	return nil
}

func copyStatics(target, destination string) error {
	entries, err := os.ReadDir(target)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		data, err := os.ReadFile(target + "/" + entry.Name())
		if err != nil {
			return err
		}
		err = os.WriteFile(BUILD_DIR+"/"+destination+"/"+entry.Name(), data, 0644)
		if err != nil {
			return err
		}
	}
	return nil
}
