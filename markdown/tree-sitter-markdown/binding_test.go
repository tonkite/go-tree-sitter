package tree_sitter_markdown_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	sitter "github.com/tonkite/go-tree-sitter"
	tree_sitter_markdown "github.com/tonkite/go-tree-sitter/markdown/tree-sitter-markdown"
)

func TestGrammar(t *testing.T) {
	assert := assert.New(t)

	n, err := sitter.ParseCtx(context.Background(), []byte("# Hello world!\n- Here is a picture: ![picture](https://example.com/picture.png)"), tree_sitter_markdown.GetLanguage())
	assert.NoError(err)
	assert.Equal(
		"(document (section (atx_heading (atx_h1_marker) heading_content: (inline)) (list (list_item (list_marker_minus) (paragraph (inline))))))",
		n.String(),
	)
}
