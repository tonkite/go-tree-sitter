package lua_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	sitter "github.com/tonkite/go-tree-sitter"
	"github.com/tonkite/go-tree-sitter/lua"
)

func TestGrammar(t *testing.T) {
	assert := assert.New(t)

	n, err := sitter.ParseCtx(context.Background(), []byte(`print("Hello World!")`), lua.GetLanguage())
	assert.NoError(err)
	assert.Equal(
		"(program (function_call prefix: (identifier) (function_call_paren) args: (function_arguments (string)) (function_call_paren)))",
		n.String(),
	)
}
