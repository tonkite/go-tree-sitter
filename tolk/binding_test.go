package tolk_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	sitter "github.com/tonkite/go-tree-sitter"
	"github.com/tonkite/go-tree-sitter/tolk"
	"testing"
)

func TestGrammar(t *testing.T) {
	assert := assert.New(t)

	n, err := sitter.ParseCtx(context.Background(), []byte(`fun onInternalMessage() {}`), tolk.GetLanguage())

	assert.NoError(err)
	assert.Equal(
		"(source_file (function_declaration name: (identifier) parameters: (parameter_list) body: (block_statement)))",
		n.String(),
	)
}
