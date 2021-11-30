package transpiler

import (
	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/prompt-cash/ccxt-go/pkg/turing_parser"
)

type GoArgument struct {
	Name string
}

type GoFunction struct {
	Name      string
	Arguments arraylist.List

	Body *turing_parser.Expressions
}

type GoFile struct {
	Package string
	Imports []string

	ThisName     string
	OwnFunctions map[string]bool
}
