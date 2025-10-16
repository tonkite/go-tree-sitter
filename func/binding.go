package fun_c

//#include "parser.h"
//TSLanguage *tree_sitter_func();
import "C"
import (
	sitter "github.com/smacker/go-tree-sitter"
	"unsafe"
)

func GetLanguage() *sitter.Language {
	ptr := unsafe.Pointer(C.tree_sitter_func())
	return sitter.NewLanguage(ptr)
}
