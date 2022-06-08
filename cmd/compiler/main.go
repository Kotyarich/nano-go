package main

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/llvm-project/llvm/bindings/go/llvm"
	"nano-go/parser"
	"nano-go/visitor"
)

func main() {
	llvm.InitializeAllTargetInfos()
	llvm.InitializeAllTargets()
	llvm.InitializeAllTargetMCs()
	llvm.InitializeAllAsmParsers()
	llvm.InitializeAllAsmPrinters()

	is := antlr.NewInputStream("package main\nfunc main() {\nreturn\n}")

	lexer := parser.NewGoLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewGoParser(stream)

	sourceFile := p.SourceFile()

	goVisitor := visitor.Visitor{}
	goVisitor.VisitSourceFile(sourceFile.(*parser.SourceFileContext))

	fmt.Println(goVisitor.Module.String())

	_, err := llvm.NewMCJITCompiler(goVisitor.Module, llvm.MCJITCompilerOptions{})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
