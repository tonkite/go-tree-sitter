package fun_c_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	sitter "github.com/tonkite/go-tree-sitter"
	funC "github.com/tonkite/go-tree-sitter/func"
	"testing"
)

func TestGrammar(t *testing.T) {
	assert := assert.New(t)

	n, err := sitter.ParseCtx(context.Background(), []byte(`(int) random() { return 4; }`), funC.GetLanguage())

	assert.NoError(err)
	assert.Equal(
		"(source_file (function_declaration return_type: (primitive_type) name: (identifier) parameters: (parameter_list) body: (block_statement (return_statement (number_literal)))))",
		n.String(),
	)
}
