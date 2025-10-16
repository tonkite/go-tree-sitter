package elixir_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	sitter "github.com/tonkite/go-tree-sitter"
	"github.com/tonkite/go-tree-sitter/elixir"
)

func TestGrammar(t *testing.T) {
	assert := assert.New(t)
	n, err := sitter.ParseCtx(context.Background(), []byte("40 + 2"), elixir.GetLanguage())
	assert.NoError(err)
	assert.Equal(
		"(source (binary_operator left: (integer) right: (integer)))",
		n.String(),
	)

	n, err = sitter.ParseCtx(context.Background(), []byte(`"hello" <> " world"`), elixir.GetLanguage())
	assert.NoError(err)
	assert.Equal(
		"(source (binary_operator left: (string (quoted_content)) right: (string (quoted_content))))",
		n.String(),
	)
}
