package fift_test

import (
	"context"
	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/fift"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGrammar(t *testing.T) {
	assert := assert.New(t)

	n, err := sitter.ParseCtx(context.Background(), []byte(`PROGRAM{ recv_internal PROC:<{ DROP }> }END>c`), fift.GetLanguage())

	assert.NoError(err)
	assert.Equal(
		"(source_file (program (definition (proc_definition name: (identifier) (instruction (identifier)))) (ERROR)))",
		n.String(),
	)
}
