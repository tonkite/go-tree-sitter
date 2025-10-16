package html_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	sitter "github.com/tonkite/go-tree-sitter"
	"github.com/tonkite/go-tree-sitter/html"
)

func TestGrammar(t *testing.T) {
	assert := assert.New(t)

	n, err := sitter.ParseCtx(context.Background(), []byte(`<HTML><BODY><P>Hello World</P></BODY></HTML>`), html.GetLanguage())
	assert.NoError(err)
	assert.Equal(
		"(document (element (start_tag (tag_name)) (element (start_tag (tag_name)) (element (start_tag (tag_name)) (text) (end_tag (tag_name))) (end_tag (tag_name))) (end_tag (tag_name))))",
		n.String(),
	)
}
