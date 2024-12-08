package internal

import (
	"bytes"
	"html/template"
	"io"

	formatters_html "github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

func MdToHTML(md []byte) template.HTML {
	// 创建 markdown 解析器
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)

	// 创建 HTML 渲染器
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{
		Flags: htmlFlags,
		RenderNodeHook: func(w io.Writer, node ast.Node, entering bool) (ast.WalkStatus, bool) {
			if code, ok := node.(*ast.CodeBlock); ok && entering {
				// 获取代码块的语言
				language := string(code.Info)
				// 获取代码内容
				content := string(code.Literal)

				// 使用 Chroma 进行语法高亮
				highlighted := highlightCode(content, language)

				// 写入高亮后的 HTML
				w.Write([]byte(highlighted))
				return ast.GoToNext, true
			}
			return ast.GoToNext, false
		},
	}
	renderer := html.NewRenderer(opts)

	// 渲染 HTML
	html := markdown.ToHTML(md, p, renderer)
	return template.HTML(html)
}

func highlightCode(code, language string) string {
	// 如果没有指定语言，默认为 plaintext
	if language == "" {
		language = "plaintext"
	}

	// 获取对应语言的词法分析器
	lexer := lexers.Get(language)
	if lexer == nil {
		lexer = lexers.Get("plaintext")
	}

	// 使用 monokai 主题
	style := styles.Get("monokai")
	if style == nil {
		style = styles.Fallback
	}

	// 创建 HTML 格式化器
	formatter := formatters_html.New(
		formatters_html.WithClasses(true),
		formatters_html.TabWidth(4),
		formatters_html.WithLineNumbers(true),
	)

	// 创建输出缓冲区
	var buf bytes.Buffer

	// 进行语法分析
	iterator, err := lexer.Tokenise(nil, code)
	if err != nil {
		return code
	}

	// 格式化为 HTML
	err = formatter.Format(&buf, style, iterator)
	if err != nil {
		return code
	}

	return buf.String()
}
