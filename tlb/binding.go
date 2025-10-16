package tlb

//#include "parser.h"
//TSLanguage *tree_sitter_tlb();
import "C"
import (
	sitter "github.com/smacker/go-tree-sitter"
	"unsafe"
)

func GetLanguage() *sitter.Language {
	ptr := unsafe.Pointer(C.tree_sitter_tlb())
	return sitter.NewLanguage(ptr)
}
