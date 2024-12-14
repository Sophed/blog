package generation

import (
	"errors"
	"site/app/views"
	"strings"

	"github.com/sblinch/kdl-go"
)

var errFormat = errors.New("invalid meta format")

func getMeta(data []byte) (*views.BlogPost, error) {
	post := views.BlogPost{}
	sections := strings.SplitN(string(data), "---\n", 2)
	if len(sections) != 2 {
		return nil, errFormat
	}
	err := kdl.Unmarshal([]byte(sections[0]), &post)
	if err != nil {
		return nil, err
	}
	post.Body = []byte(sections[1])
	return &post, nil
}
