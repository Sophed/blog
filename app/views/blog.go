package views

import (
	"site/app/components"
	"sort"
	"strconv"
	"time"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

type BlogPost struct {
	ID      string    `kdl:"-"`
	Title   string    `kdl:"title"`
	DateStr string    `kdl:"date"`
	Date    time.Time `kdl:"-"`
	Body    []byte    `kdl:"-"`
}

var posts []*BlogPost

func AddPost(p *BlogPost) {
	posts = append(posts, p)
	sort.Slice(posts, func(i, j int) bool {
		return posts[i].Date.After(posts[j].Date) // latest...oldest
	})
}

func LatestPost() Node {
	p := posts[0]
	if p == nil {
		return P(Text("No posts available"))
	} else {
		return blogLink(p)
	}
}

func Blog() Node {
	return components.View(false, "soph.systems ~ posts",
		H2(Text("Latest - Oldest")),
		postList(),
	)
}

func postList() Node {
	if len(posts) == 0 {
		return P(Text("No posts available"))
	}
	return Div(
		Map(posts, func(post *BlogPost) Node {
			return blogLink(post)
		}),
	)
}

func blogLink(post *BlogPost) Node {
	return Div(
		A(Href("/posts/"+post.ID),
			Raw(post.Title),
		),
		P(Text(
			"Published: "+post.DateStr+
				" ~ "+strconv.Itoa(readingTime(string(post.Body)))+"m read",
		)),
		Br(),
	)
}
