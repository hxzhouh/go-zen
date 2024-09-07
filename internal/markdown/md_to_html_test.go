package markdown

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

// MdToHTML converts markdown to HTML.
func TestMdToHTML(t *testing.T) {
	mds, err := os.ReadFile("读写锁和互斥锁的性能比较.md")
	assert.Nil(t, err)
	start := time.Now()
	html := MdToHTML(mds)
	fmt.Printf("cost:%d ms \n", time.Since(start)/time.Millisecond)
	fmt.Println(string(html))
}
