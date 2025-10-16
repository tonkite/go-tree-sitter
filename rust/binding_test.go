package rust_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	sitter "github.com/tonkite/go-tree-sitter"
	"github.com/tonkite/go-tree-sitter/rust"
)

func TestGrammar(t *testing.T) {
	assert := assert.New(t)

	n, err := sitter.ParseCtx(context.Background(), []byte("mod one;"), rust.GetLanguage())
	assert.NoError(err)
	assert.Equal(
		"(source_file (mod_item name: (identifier)))",
		n.String(),
	)
}
