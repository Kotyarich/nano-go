package visitor

import (
	"fmt"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"

	"nano-go/parser"
	"nano-go/visitor/name"
)

const paramNamePrefix = "param_"

func (v *Visitor) visitMain(ctx *parser.FunctionDeclContext) {
	f := v.Module.NewFunc("main", types.I32)
	v.curBlock = f.NewBlock(name.BlockName())

	v.pushVariablesStack()

	block := ctx.Block()
	v.visitBlock(block.(*parser.BlockContext))

	v.curBlock.NewRet(constant.NewInt(types.I32, 0))
}

func (v *Visitor) getParamType(pType string) types.Type {
	switch pType {
	case "int64":
		return types.I64
	case "int32":
		return types.I32
	}

	panic(fmt.Errorf("unknown type: %s", pType))
}

func (v *Visitor) visitType_(ctx *parser.Type_Context) types.Type {
	t := ctx.TypeName().(*parser.TypeNameContext).IDENTIFIER().GetText()
	llvmType := v.getParamType(t)

	return llvmType
}

func (v *Visitor) getParameters(ctx *parser.SignatureContext) []*ir.Param {
	var params []*ir.Param

	parameters := ctx.Parameters().(*parser.ParametersContext)
	for _, parameterI := range parameters.AllParameterDecl() {
		parameter := parameterI.(*parser.ParameterDeclContext)

		identifierList := parameter.IdentifierList().(*parser.IdentifierListContext)
		paramName := identifierList.IDENTIFIER(0).GetText()

		paramType := parameter.Type_().(*parser.Type_Context)
		llvmType:= v.visitType_(paramType)

		params = append(params, ir.NewParam(paramNamePrefix + paramName, llvmType))
	}

	return params
}

func (v *Visitor) getReturnTypes(ctx *parser.SignatureContext) []types.Type {
	result := ctx.Result()
	if result == nil {
		return []types.Type{}
	}

	type_ := ctx.Result().(*parser.ResultContext).Type_()
	llvmType := v.visitType_(type_.(*parser.Type_Context))
	return []types.Type{llvmType}
}

func (v *Visitor) visitFunc(ctx *parser.FunctionDeclContext) {
	v.pushVariablesStack()
	fName := ctx.IDENTIFIER().GetText()
	signature := ctx.Signature().(*parser.SignatureContext)

	params := v.getParameters(signature)

	var funcRetType types.Type = types.Void
	returnTypes := v.getReturnTypes(signature)
	if len(returnTypes) == 1 {
		funcRetType = returnTypes[0]
	}

	f := v.Module.NewFunc(fName, funcRetType, params...)
	v.pkgVars[fName] = Value{
		Value: f,
		Type: f.Type(),
	}
	v.curBlock = f.NewBlock(name.BlockName())

	for _, param := range f.Params {
		alloca := v.curBlock.NewAlloca(param.Type())
		allocaName := param.Name()[6:]
		alloca.SetName(allocaName)
		v.curBlock.NewStore(param, alloca)

		//v.variables[param.Name()] = param
		v.setVar(allocaName, Value{
			Value: alloca,
			Type: param.Type(),
			IsVariable: true,
		})
	}

	block := ctx.Block()
	v.visitBlock(block.(*parser.BlockContext))

	if len(returnTypes) == 0 {
		v.curBlock.NewRet(nil)
	}
}

func (v *Visitor) visitFuncCall(name string, args []value.Value) value.Value {
	switch name {
	case "Print":
		return v.printFuncCall(args)
	case "Printf":
		args[0] = v.curBlock.NewExtractValue(args[0], 1)
	}
	fn := v.pkgVars[name]

	return v.curBlock.NewCall(fn.Value, args...)
}