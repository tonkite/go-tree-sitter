package dockerfile_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	sitter "github.com/tonkite/go-tree-sitter"
	"github.com/tonkite/go-tree-sitter/dockerfile"
)

func TestGrammar(t *testing.T) {
	assert := assert.New(t)

	n, err := sitter.ParseCtx(context.Background(), []byte("FROM microsoft/nanoserver"), dockerfile.GetLanguage())
	assert.NoError(err)
	assert.Equal(
		"(source_file (from_instruction (image_spec name: (image_name))) (MISSING \"\n\"))",
		n.String(),
	)
}
