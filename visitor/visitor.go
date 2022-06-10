package visitor

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"

	"nano-go/parser"
	"nano-go/visitor/strings"
	"nano-go/visitor/type"
)


type Value struct {
	Type  types.Type
	Value value.Value

	IsVariable bool
	MultiValues []Value
}

type Visitor struct {
	Module   *ir.Module
	curBlock *ir.Block

	contextBlockVariables []map[string]Value
	pkgVars				  map[string]Value

}

func (v *Visitor) setVar(name string, val Value) {
	v.contextBlockVariables[len(v.contextBlockVariables)-1][name] = val
}

func (v *Visitor) pushVariablesStack() {
	v.contextBlockVariables = append(v.contextBlockVariables, make(map[string]Value))
}

func (v *Visitor) popVariablesStack() {
	v.contextBlockVariables = v.contextBlockVariables[0 : len(v.contextBlockVariables)-1]
}

func (v *Visitor) VisitSourceFile(ctx *parser.SourceFileContext) {
	v.pkgVars = make(map[string]Value)
	v.contextBlockVariables = make([]map[string]Value, 0)
	v.Module = ir.NewModule()

	_type.ModuleStringType = v.Module.NewTypeDef("string", strings.String())

	// Create empty string constant
	_type.EmptyStringConstant = v.Module.NewGlobalDef(strings.NextStringName(), strings.Constant(""))
	_type.EmptyStringConstant.Immutable = true

	setExternal := func(internalName string, fn *ir.Func, variadic bool) Value {
		fn.Sig.Variadic = variadic
		val := Value{
			Type: fn.Type(),
			Value: fn,
		}
		v.pkgVars[internalName] = val
		return val
	}

	setExternal("Printf", v.Module.NewFunc("printf",
		types.I32,
		ir.NewParam("_", types.NewPointer(types.I8)),
	), true)

	v.visitFunctionDecls(ctx.AllFunctionDecl())
}

func (v *Visitor) visitFunctionDecls(ctxs []parser.IFunctionDeclContext) {
	for _, ctx := range ctxs {
		fCtx := ctx.(*parser.FunctionDeclContext)
		if fCtx.IDENTIFIER().GetText() != "main" {
			v.visitFunc(fCtx)
		}
	}

	for _, ctx := range ctxs {
		fCtx := ctx.(*parser.FunctionDeclContext)
		if fCtx.IDENTIFIER().GetText() == "main" {
			v.visitMain(fCtx)
		}
	}
}

func (v *Visitor) visitBlock(ctx *parser.BlockContext) {
	list := ctx.StatementList()
	children := list.GetChildren()

	for _, sCtx := range children {
		switch sCtx.(type) {
		case *parser.StatementContext:
			v.visitStatement(sCtx.(*parser.StatementContext))
		}
	}
}
