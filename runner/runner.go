package runner

import (
	"io"
	"monkey/evaluator"
	"monkey/lexer"
	"monkey/object"
	"monkey/parser"
	"os"
)

func Start(filePath string, out io.Writer) {
	// Read entire file
	content, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	// Parse the content
	l := lexer.New(string(content))
	p := parser.New(l)

	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		parser.PrintParserErrors(out, p.Errors())
		os.Exit(1)
	}

	// Evaluate the program
	env := object.NewEnvironment()
	evaluated := evaluator.Eval(program, env)

	if evaluated != nil {
		// Print the final return
		io.WriteString(out, evaluated.Inspect())
		io.WriteString(out, "\n")
	}
}
