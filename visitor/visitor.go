package visitor

import (
	"fmt"
	"nano-go/parser"
	"strconv"

	"github.com/llvm-project/llvm/bindings/go/llvm"
)

type Visitor struct {
	Module    llvm.Module
	builder   llvm.Builder
	ctx       llvm.Context
	variables map[string]llvm.Value
}

func (v *Visitor) VisitSourceFile(ctx *parser.SourceFileContext) {
	v.ctx = llvm.NewContext()
	v.builder = llvm.NewBuilder()
	v.Module = llvm.NewModule("output")

	fType := llvm.FunctionType(llvm.Int32Type(), []llvm.Type{}, false)
	f := llvm.AddFunction(v.Module, "main", fType)

	block := llvm.AddBasicBlock(f, "body")
	v.builder.SetInsertPointAtEnd(block)

	v.visitFunctionDecls(ctx.AllFunctionDecl())

	v.builder.CreateRet(llvm.ConstInt(llvm.Int32Type(), 0, true))
}

func (v *Visitor) visitFunctionDecls(ctxs []parser.IFunctionDeclContext) {
	for _, ctx := range ctxs {
		fCtx := ctx.(*parser.FunctionDeclContext)
		if fCtx.IDENTIFIER().GetText() == "main" {
			v.visitMain(fCtx)
		}
	}
}

func (v *Visitor) visitMain(ctx *parser.FunctionDeclContext) {
	block := ctx.Block()
	v.visitBlock(block.(*parser.BlockContext))
}

func (v *Visitor) visitBlock(ctx *parser.BlockContext) {
	list := ctx.StatementList()
	children := list.GetChildren()[1 : len(list.GetChildren())-1] // skip { and }

	for _, sCtx := range children {
		v.visitStatement(sCtx.(*parser.StatementContext))
	}
}

func (v *Visitor) visitStatement(ctx *parser.StatementContext) {
	simpleCtx := ctx.SimpleStmt()
	if simpleCtx != nil {
		v.visitSimpleStmt(simpleCtx.(*parser.SimpleStmtContext))
	}
}

func (v *Visitor) visitSimpleStmt(ctx *parser.SimpleStmtContext) {
	shortDeclCtx := ctx.ShortVarDecl()
	if shortDeclCtx != nil {
		v.visitShortVarDecl(shortDeclCtx.(*parser.ShortVarDeclContext))
	}

	if assigmentCtx := ctx.Assignment(); assigmentCtx != nil {
		v.visitAssigment(assigmentCtx.(*parser.AssignmentContext))
	}
}

func (v *Visitor) visitShortVarDecl(ctx *parser.ShortVarDeclContext) {
	identifier := ctx.IdentifierList().(*parser.IdentifierListContext).IDENTIFIER(0).GetText()

	expressionCtx := ctx.ExpressionList().(*parser.ExpressionListContext).Expression(0).(*parser.ExpressionContext)
	exprValue := v.visitExpression(expressionCtx)

	alloca := v.builder.CreateAlloca(exprValue.Type(), identifier)
	v.builder.CreateStore(exprValue, alloca)

	v.variables[identifier] = alloca
}

func (v *Visitor) visitAssigment(ctx *parser.AssignmentContext) {
	/*
    this->builder.CreateStore(expression, variable);

    return this->builder.CreateLoad(variable->getType()->getPointerElementType(), variable);
	 */
	leftCtx := ctx.ExpressionList(0).(*parser.ExpressionListContext).Expression(0)
	leftValue := v.visitExpression(leftCtx.(*parser.ExpressionContext))
	name := leftValue.Name()
	variable := v.variables[name]

	rightCtx := ctx.ExpressionList(0).(*parser.ExpressionListContext).Expression(0)
	rightValue := v.visitExpression(rightCtx.(*parser.ExpressionContext))

	v.builder.CreateStore(rightValue, variable)

	v.builder.CreateLoad(variable, name)
}

func (v *Visitor) visitExpression(ctx *parser.ExpressionContext) llvm.Value {
	if unary := ctx.GetUnary_op(); unary != nil {
		if unary.GetTokenType() == parser.GoParserMINUS {
			value := v.visitExpression(ctx.Expression(0).(*parser.ExpressionContext))
			valueType := value.Type()
			switch valueType.TypeKind() {
			case llvm.IntegerTypeKind:
				zero := llvm.ConstInt(valueType, 0, true)
				return v.builder.CreateSub(zero, value, "unary_minus")
			}
		}
	}

	if primaryExpr := ctx.PrimaryExpr(); primaryExpr != nil {
		v.visitPrimaryExpr(primaryExpr.(*parser.PrimaryExprContext))
	}

	panic("unknown expression")
}

func (v *Visitor) visitPrimaryExpr(ctx *parser.PrimaryExprContext) llvm.Value {
	if operand := ctx.Operand(); operand != nil {
		return v.visitOperand(operand.(*parser.OperandContext))
	}

	panic("unknown primary expression")
}

func (v *Visitor) visitOperand(ctx *parser.OperandContext) llvm.Value {
	if name := ctx.OperandName(); name != nil {
		return v.visitOperandName(name.(*parser.OperandNameContext))
	}

	if literal := ctx.Literal(); literal != nil {
		return v.visitLiteral(literal.(*parser.LiteralContext))
	}

	panic(fmt.Errorf("unknown operand"))
}

func (v *Visitor) visitOperandName(ctx *parser.OperandNameContext) llvm.Value {
	name := ctx.IDENTIFIER().GetText()
	variable, ok := v.variables[name]
	if !ok {
		panic(fmt.Errorf("unknown variable: %s", name))
	}

	return v.builder.CreateLoad(variable, name)
}

func (v *Visitor) visitLiteral(ctx *parser.LiteralContext) llvm.Value {
	if basicLit := ctx.BasicLit(); basicLit != nil {
		return v.visitBasicLit(basicLit.(*parser.BasicLitContext))
	}

	panic("unknown literal")
}

func (v *Visitor) visitBasicLit(ctx *parser.BasicLitContext) llvm.Value {
	if str := ctx.String_(); str != nil {

	}

	if integer := ctx.Integer(); integer != nil {
		literal := integer.(*parser.IntegerContext).DECIMAL_LIT().GetText()
		value, _ := strconv.Atoi(literal)
		return llvm.ConstInt(llvm.Int64Type(), uint64(value), true)
	}

	if floatLit := ctx.FLOAT_LIT(); floatLit != nil {

	}

	panic("unknown basic literal")
}
