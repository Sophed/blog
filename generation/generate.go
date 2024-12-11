package generation

import (
	"bytes"
	"os"
	"site/app/views"
	"strings"
)

const POSTS_DIR = "posts"
const BUILD_DIR = "build"
const STYLE_DIR = "styles"

func GeneratePosts() error {
	err := clearBuildDir()
	if err != nil {
		return err
	}
	err = copyCSS()
	if err != nil {
		return err
	}
	entries, err := os.ReadDir(POSTS_DIR)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		title := strings.Split(entry.Name(), ".")[0]
		println("-", title)
		data, err := os.ReadFile(POSTS_DIR + "/" + entry.Name())
		if err != nil {
			return err
		}
		html := convertMDtoHTML(data, title)
		post := new(bytes.Buffer)
		views.Post(title, html).Render(post)
		content := post.String()
		err = os.WriteFile(BUILD_DIR+"/posts/"+title+".html", []byte(content), 0644)
		if err != nil {
			return err
		}
		views.BLOG_POSTS = append(views.BLOG_POSTS, title)
	}
	return nil
}

func copyCSS() error {
	entries, err := os.ReadDir(STYLE_DIR)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		data, err := os.ReadFile(STYLE_DIR + "/" + entry.Name())
		if err != nil {
			return err
		}
		err = os.WriteFile(BUILD_DIR+"/css/"+entry.Name(), data, 0644)
		if err != nil {
			return err
		}
	}
	return nil
}
