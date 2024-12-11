package generation

import (
	"bytes"
	"os"

	"github.com/sophed/lg"
	"maragu.dev/gomponents"
)

func Static(view gomponents.Node, name string) {
	data := new(bytes.Buffer)
	view.Render(data)
	err := os.WriteFile(BUILD_DIR+"/"+name+".html", data.Bytes(), 0644)
	if err != nil {
		lg.Fatl(err)
	}
	println("-", name)
}
