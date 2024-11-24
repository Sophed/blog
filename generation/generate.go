package generation

import (
	"os"
	"strings"

	"github.com/sophed/lg"
)

const POSTS_DIR = "posts"
const BUILD_DIR = "build"
const TEMPLATES_DIR = "templates"

func GeneratePosts() error {
	InitDir(BUILD_DIR)
	err := clearBuildDir()
	if err != nil {
		return err
	}
	template, err := os.ReadFile(TEMPLATES_DIR + "/page.html")
	if err != nil {
		return err
	}
	nav, err := os.ReadFile(TEMPLATES_DIR + "/nav.html")
	if err != nil {
		return err
	}
	styles, err := os.ReadFile(TEMPLATES_DIR + "/global.css")
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
		content := strings.ReplaceAll(string(template), "{{content}}", convertMDtoHTML(data, title))
		content = strings.ReplaceAll(content, "{{title}}", title)
		content = strings.ReplaceAll(content, "{{css}}", string(styles))
		content = strings.ReplaceAll(content, "{{nav}}", string(nav))
		err = os.WriteFile(BUILD_DIR+"/posts/"+title+".html", []byte(content), 0644)
	}
	return nil
}

func InitDir(name string) bool {
	s, err := os.Stat(name)
	if os.IsNotExist(err) {
		err = os.Mkdir(name, os.ModePerm)
		if err != nil {
			lg.Fatl(err)
		}
		return true
	}
	if !s.IsDir() {
		lg.Fatl("a file already exists with expected name of" + name + "dir")
	}
	return false
}
