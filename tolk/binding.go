package tolk

//#include "parser.h"
//TSLanguage *tree_sitter_tolk();
import "C"
import (
	sitter "github.com/tonkite/go-tree-sitter"
	"unsafe"
)

func GetLanguage() *sitter.Language {
	ptr := unsafe.Pointer(C.tree_sitter_tolk())
	return sitter.NewLanguage(ptr)
}
