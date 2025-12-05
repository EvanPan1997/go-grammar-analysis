package main

import (
	"fmt"

	"github.com/EvanPan1997/go-grammar-analysis/grammar"
	"github.com/antlr4-go/antlr/v4"
)

func main() {
	fmt.Println("Hello")

	input := antlr.NewInputStream("SELECT * FROM TABLE_NAME")

	lexer := grammar.NewUniversalLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := grammar.NewUniversalParser(stream)
	parser.BuildParseTrees = true
	statements := parser.Statements()
	statement := statements.Statement(0)
	fmt.Println(statement.GetText())
}
