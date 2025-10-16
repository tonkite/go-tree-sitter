package tlb_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	sitter "github.com/tonkite/go-tree-sitter"
	"github.com/tonkite/go-tree-sitter/tlb"
	"testing"
)

func TestGrammar(t *testing.T) {
	assert := assert.New(t)

	n, err := sitter.ParseCtx(context.Background(), []byte(`simple#1234cafe query_id:uint64 = InMsgBody;`), tlb.GetLanguage())

	assert.NoError(err)
	assert.Equal(
		"(program (declaration constructor: (constructor_ name: (identifier) tag: (constructor_tag (hex))) (field (field_named name: (identifier) expr: (cond_expr (cond_type_expr (type_expr (simple_expr (ref_expr (ref_inner (type_identifier))))))))) combinator: (combinator name: (type_identifier))))",
		n.String(),
	)
}
