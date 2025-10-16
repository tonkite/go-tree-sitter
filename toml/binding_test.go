package toml_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	sitter "github.com/tonkite/go-tree-sitter"
	"github.com/tonkite/go-tree-sitter/toml"
)

func TestGrammar(t *testing.T) {
	assert := assert.New(t)

	n, err := sitter.ParseCtx(context.Background(), []byte(`key = "value"`), toml.GetLanguage())
	assert.NoError(err)
	assert.Equal(
		"(document (pair (bare_key) (string)))",
		n.String(),
	)
}
