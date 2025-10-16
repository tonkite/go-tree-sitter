package css_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	sitter "github.com/tonkite/go-tree-sitter"
	"github.com/tonkite/go-tree-sitter/css"
)

func TestGrammar(t *testing.T) {
	assert := assert.New(t)

	n, err := sitter.ParseCtx(context.Background(), []byte(`p { color: red; }`), css.GetLanguage())
	assert.NoError(err)
	assert.Equal(
		"(stylesheet (rule_set (selectors (tag_name)) (block (declaration (property_name) (plain_value)))))",
		n.String(),
	)
}
