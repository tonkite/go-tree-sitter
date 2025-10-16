package solidity_test

import (
	"context"
	"github.com/tonkite/go-tree-sitter/solidity"
	"testing"

	"github.com/stretchr/testify/assert"
	sitter "github.com/tonkite/go-tree-sitter"
)

const code = `pragma solidity 0.4.4;
contract c {
    event e(uint[10] a, bytes7[8] indexed b, c[3] x);
}
`

const expected = `(source_file (pragma_directive (solidity_pragma_token version_constraint: (solidity_version))) (contract_declaration name: (identifier) body: (contract_body (event_definition name: (identifier) (event_parameter type: (type_name (type_name (primitive_type)) (expression (number_literal))) name: (identifier)) (event_parameter type: (type_name (type_name (primitive_type)) (expression (number_literal))) name: (identifier)) (event_parameter type: (type_name (type_name (user_defined_type (identifier))) (expression (number_literal))) name: (identifier))))))`

func TestGrammar(t *testing.T) {
	assert := assert.New(t)

	n, err := sitter.ParseCtx(context.Background(), []byte(code), solidity.GetLanguage())
	assert.NoError(err)
	assert.Equal(
		expected,
		n.String(),
	)
}
