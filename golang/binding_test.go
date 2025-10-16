package golang_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	sitter "github.com/tonkite/go-tree-sitter"
	"github.com/tonkite/go-tree-sitter/golang"
)

func TestGrammar(t *testing.T) {
	assert := assert.New(t)

	n, err := sitter.ParseCtx(context.Background(), []byte("package main"), golang.GetLanguage())
	assert.NoError(err)
	assert.Equal(
		"(source_file (package_clause (package_identifier)))",
		n.String(),
	)
}
