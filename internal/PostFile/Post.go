package PostFile

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/hxzhouh/go-zen.git/domain"
	"github.com/hxzhouh/go-zen.git/internal"
	"gopkg.in/yaml.v3"
)

type Post struct {
	Properties *PostProperties
	Content    string
}

type PostProperties struct {
	Title       string   `yaml:"title"`
	Description string   `yaml:"description"`
	SubTitle    string   `yaml:"subtitle"`
	Date        string   `yaml:"date"`
	Lastmod     string   `yaml:"lastmod"`
	draft       bool     `yaml:"draft"`
	Tags        []string `yaml:"tags"`
	Categories  []string `yaml:"categories"`
	Author      string   `yaml:"author"`
	Slug        string   `yaml:"slug"`
	Image       string   `yaml:"image"`
	Keywords    []string `yaml:"keywords"`
	LongTimeUrl string   `yaml:"long_time_url"`
	Id          string   `yaml:"id"`
}

func (p *Post) ToPost() *domain.Post {

	// implement the conversion logic here

	return &domain.Post{Title: p.Properties.Title,
		SubTitle:    p.Properties.SubTitle,
		Summary:     p.Properties.Description,
		Cover:       p.Properties.Image,
		Content:     p.Content,
		TagIds:      p.Properties.Tags,
		CategoryId:  p.Properties.Categories,
		Draft:       p.Properties.draft,
		AuthorID:    p.Properties.Author,
		Md5:         "",
		PostId:      p.Properties.Id,
		ContentHtml: internal.MdToHTML([]byte(p.Content)),
		Reads:       0,
		Likes:       0,
	}

}
func (p *Post) formatByFile(path string) {
	// read md File by path
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}
	p.Properties, p.Content, err = p.getTacYaml(content)
	if err != nil {
		slog.Error("formatByFile error", path, err)
		return
	}
}

func (t *Post) getTacYaml(content []byte) (*PostProperties, string, error) {
	// Define the regular expression to find content between ---
	re := regexp.MustCompile(`(?s)---(.*?)---`)

	// Find the content between the markers
	matches := re.FindStringSubmatch(string(content))
	if len(matches) < 2 {
		return nil, "", nil
	}
	// Parse the content as YAML
	properties := &PostProperties{}
	err := yaml.Unmarshal([]byte(matches[1]), properties)
	if err != nil {
		log.Printf("Failed to unmarshal yaml: %v\n", err)
		return properties, "", err
	}
	return properties, string(content[len(matches[1]):]), nil
}

func ReadAllFile(path string) []*Post {
	posts := make([]*Post, 0)

	err := filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("Error accessing file %s: %v", filePath, err)
			return err
		}

		// 检查是否是符号链接
		if info.Mode()&os.ModeSymlink != 0 {
			realPath, err := filepath.EvalSymlinks(filePath)
			if err != nil {
				log.Printf("Failed to resolve symlink %s: %v", filePath, err)
				return err
			}
			log.Printf("Resolved symlink: %s -> %s", filePath, realPath)
			// 如果是目录符号链接，递归处理目标路径
			if resolvedInfo, err := os.Stat(realPath); err == nil && resolvedInfo.IsDir() {
				return filepath.Walk(realPath, func(subFilePath string, subInfo os.FileInfo, subErr error) error {
					if subErr != nil {
						log.Printf("Error accessing file %s: %v", subFilePath, subErr)
						return subErr
					}
					if !subInfo.IsDir() && strings.HasSuffix(subFilePath, ".md") {
						tmp := &Post{}
						fmt.Println(subFilePath)
						tmp.formatByFile(subFilePath)
						if tmp.Properties != nil {
							posts = append(posts, tmp)
						}
					}
					return nil
				})
			}
		}

		// 普通文件处理逻辑
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".md") {
			tmp := &Post{}
			tmp.formatByFile(filePath)
			if tmp.Properties != nil {
				posts = append(posts, tmp)
			}
		}
		return nil
	})

	if err != nil {
		log.Fatalf("Failed to walk through directory: %v", err)
	}

	return posts
}
