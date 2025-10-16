package hcl_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	sitter "github.com/tonkite/go-tree-sitter"
	"github.com/tonkite/go-tree-sitter/hcl"
)

func TestGrammar(t *testing.T) {
	assert := assert.New(t)

	n, err := sitter.ParseCtx(context.Background(), []byte(`image_id = "abc123"`), hcl.GetLanguage())
	assert.NoError(err)
	assert.Equal(
		"(config_file (body (attribute (identifier) (expression (literal_value (string_lit (quoted_template_start) (template_literal) (quoted_template_end)))))))",
		n.String(),
	)
}
