package bash_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	sitter "github.com/tonkite/go-tree-sitter"
	"github.com/tonkite/go-tree-sitter/bash"
)

func TestGrammar(t *testing.T) {
	assert := assert.New(t)

	n, err := sitter.ParseCtx(context.Background(), []byte("echo 1"), bash.GetLanguage())
	assert.NoError(err)
	assert.Equal(
		"(program (command name: (command_name (word)) argument: (number)))",
		n.String(),
	)
}
