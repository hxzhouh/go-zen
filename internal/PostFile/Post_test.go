package PostFile

import (
	"fmt"
	"testing"
)

func TestReadAllFile(t *testing.T) {
	posts := ReadAllFile("/Users/hxzhouh/workspace/github/go-zen/post")
	for _, post := range posts {
		fmt.Println(post.Properties.Title)
	}
}
