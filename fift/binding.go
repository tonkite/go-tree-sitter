package fift

//#include "parser.h"
//TSLanguage *tree_sitter_fift();
import "C"
import (
	sitter "github.com/tonkite/go-tree-sitter"
	"unsafe"
)

func GetLanguage() *sitter.Language {
	ptr := unsafe.Pointer(C.tree_sitter_fift())
	return sitter.NewLanguage(ptr)
}
