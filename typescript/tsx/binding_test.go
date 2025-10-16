package tsx_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	sitter "github.com/tonkite/go-tree-sitter"
	"github.com/tonkite/go-tree-sitter/typescript/tsx"
)

func TestGrammar(t *testing.T) {
	assert := assert.New(t)

	n, err := sitter.ParseCtx(context.Background(), []byte("<foo />"), tsx.GetLanguage())
	assert.NoError(err)
	assert.Equal(
		"(program (expression_statement (jsx_self_closing_element name: (identifier))))",
		n.String(),
	)
}
