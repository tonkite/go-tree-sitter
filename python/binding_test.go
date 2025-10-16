package python_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	sitter "github.com/tonkite/go-tree-sitter"
	"github.com/tonkite/go-tree-sitter/python"
)

func TestGrammar(t *testing.T) {
	assert := assert.New(t)

	n, err := sitter.ParseCtx(context.Background(), []byte("print(1)"), python.GetLanguage())
	assert.NoError(err)
	assert.Equal(
		"(module (expression_statement (call function: (identifier) arguments: (argument_list (integer)))))",
		n.String(),
	)
}
