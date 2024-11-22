package generation

import (
	"os"
	"strings"

	"github.com/sophed/lg"
)

const POSTS_DIR = "posts"
const BUILD_DIR = "build"

func GeneratePosts() error {
	InitDir(BUILD_DIR)
	err := clearBuildDir()
	if err != nil {
		return err
	}
	entries, err := os.ReadDir(POSTS_DIR)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		println("-", entry.Name())
		data, err := os.ReadFile(POSTS_DIR + "/" + entry.Name())
		if err != nil {
			return err
		}
		content := convertMDtoHTML(data)
		title := strings.Split(entry.Name(), ".")[0]
		err = os.WriteFile(BUILD_DIR+"/posts/"+title+".html", content, 0644)
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
