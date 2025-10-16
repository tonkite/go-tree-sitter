package php_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	sitter "github.com/tonkite/go-tree-sitter"
	"github.com/tonkite/go-tree-sitter/php"
)

func TestGrammar(t *testing.T) {
	assert := assert.New(t)

	n, err := sitter.ParseCtx(context.Background(), []byte("<?php print(1);"), php.GetLanguage())
	assert.NoError(err)
	assert.Equal(
		"(program (php_tag) (expression_statement (print_intrinsic (parenthesized_expression (integer)))))",
		n.String(),
	)
}
