package cli

import (
	"io"
	"monkey/evaluator"
	"monkey/lexer"
	"monkey/object"
	"monkey/parser"
	"os"
)

func Start(filename string, out io.Writer) {
	env := object.NewEnvironment()
	macroEnv := object.NewEnvironment()

	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	l := lexer.New(string(data))
	p := parser.New(l)

	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		printParserErrors(out, p.Errors())
		return
	}

	evaluator.DefineMacros(program, macroEnv)
	expanded := evaluator.ExpandMacros(program, macroEnv)

	evaluated := evaluator.Eval(expanded, env)
	if evaluated != nil {
		io.WriteString(out, evaluated.Inspect())
		io.WriteString(out, "\n")
	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "Woops! We ran into some monkey business here!\n")
		io.WriteString(out, " parser errors:\n")
		io.WriteString(out, "\t"+msg+"\n")
	}
}
