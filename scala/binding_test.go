package scala_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	sitter "github.com/tonkite/go-tree-sitter"
	"github.com/tonkite/go-tree-sitter/scala"
)

const code = `package com.foo.bar`

const expected = `(compilation_unit (package_clause name: (package_identifier (identifier) (identifier) (identifier))))`

func TestGrammar(t *testing.T) {
	assert := assert.New(t)

	n, err := sitter.ParseCtx(context.Background(), []byte(code), scala.GetLanguage())
	assert.NoError(err)
	assert.Equal(
		expected,
		n.String(),
	)
}
