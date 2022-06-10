package visitor

import (
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"nano-go/parser"
	"nano-go/visitor/name"
)

func (v *Visitor) visitStatement(ctx *parser.StatementContext) {
	simpleCtx := ctx.SimpleStmt()
	if simpleCtx != nil {
		v.visitSimpleStmt(simpleCtx.(*parser.SimpleStmtContext))
	}

	if returnCtx := ctx.ReturnStmt(); returnCtx != nil {
		v.visitReturnStatement(returnCtx.(*parser.ReturnStmtContext))
	}

	if ifCtx := ctx.IfStmt(); ifCtx != nil {
		v.visitIfStmt(ifCtx.(*parser.IfStmtContext))
	}

	if forCtx := ctx.ForStmt(); forCtx != nil {
		v.visitForStmt(forCtx.(*parser.ForStmtContext))
	}
}

func (v *Visitor) visitIfStmt(ctx *parser.IfStmtContext) {
	exprI := ctx.Expression()
	if exprI == nil {
		panic("unknown condition")
	}

	cond := v.visitExpression(exprI.(*parser.ExpressionContext))

	afterBlock := v.curBlock.Parent.NewBlock(name.BlockName() + "-after")
	trueBlock := v.curBlock.Parent.NewBlock(name.BlockName() + "-true")
	falseBlock := afterBlock

	if ctx.ELSE() != nil {
		falseBlock = v.curBlock.Parent.NewBlock(name.BlockName() + "-false")
	}

	v.curBlock.NewCondBr(cond, trueBlock, falseBlock)
	v.curBlock = trueBlock

	v.visitBlock(ctx.Block(0).(*parser.BlockContext))

	if trueBlock.Term == nil {
		trueBlock.NewBr(afterBlock)
	}

	if ctx.ELSE() != nil {
		v.curBlock = falseBlock
		if ctx.IfStmt() != nil {
			v.visitIfStmt(ctx.IfStmt().(*parser.IfStmtContext))
			if v.curBlock.Term == nil {
				v.curBlock.NewBr(afterBlock)
			}
		}
		if ctx.Block(1) != nil {
			v.visitBlock(ctx.Block(1).(*parser.BlockContext))
		}

		if falseBlock.Term == nil {
			falseBlock.NewBr(afterBlock)
		}
	}

	v.curBlock = afterBlock
}

func (v *Visitor) visitReturnStatement(ctx *parser.ReturnStmtContext) {
	if ctx.ExpressionList() == nil {
		v.curBlock.NewRet(nil)
		return
	}

	expressions := ctx.ExpressionList().(*parser.ExpressionListContext).AllExpression()
	if len(expressions) == 1 {
		expr := v.visitExpression(expressions[0].(*parser.ExpressionContext))
		v.curBlock.NewRet(expr)
		return
	}

	panic("multiple return statement is not realised yet")
}

func (v *Visitor) visitSimpleStmt(ctx *parser.SimpleStmtContext) {
	shortDeclCtx := ctx.ShortVarDecl()
	if shortDeclCtx != nil {
		v.visitShortVarDecl(shortDeclCtx.(*parser.ShortVarDeclContext))
	}

	if assigmentCtx := ctx.Assignment(); assigmentCtx != nil {
		v.visitAssigment(assigmentCtx.(*parser.AssignmentContext))
	}

	if expressionCtx := ctx.ExpressionStmt(); expressionCtx != nil {
		v.visitExpression(expressionCtx.(*parser.ExpressionStmtContext).Expression().(*parser.ExpressionContext))
	}

	if incDecCtx := ctx.IncDecStmt(); incDecCtx != nil {
		v.visitIncDecCtx(incDecCtx.(*parser.IncDecStmtContext))
	}
}

func (v *Visitor) visitIncDecCtx(ctx *parser.IncDecStmtContext) {
	identifier := ctx.Expression().(*parser.ExpressionContext).GetText()
	value := v.contextBlockVariables[len(v.contextBlockVariables) - 1][identifier]

	one := constant.NewInt(types.I64, 1)

	if ctx.PLUS_PLUS() != nil {
		load := v.curBlock.NewLoad(value.Type, value.Value)
		inced := v.curBlock.NewAdd(load, one)
		v.curBlock.NewStore(inced, value.Value)
	}

	if ctx.MINUS_MINUS() != nil {
		load := v.curBlock.NewLoad(value.Type, value.Value)
		deced := v.curBlock.NewSub(load, one)
		v.curBlock.NewStore(deced, value.Value)
	}
}

func (v *Visitor) visitShortVarDecl(ctx *parser.ShortVarDeclContext) {
	identifier := ctx.IdentifierList().(*parser.IdentifierListContext).IDENTIFIER(0).GetText()

	expressionCtx := ctx.ExpressionList().(*parser.ExpressionListContext).Expression(0).(*parser.ExpressionContext)
	exprValue := v.visitExpression(expressionCtx)

	alloca := v.curBlock.NewAlloca(exprValue.Type())
	alloca.SetName(identifier)
	v.curBlock.NewStore(exprValue, alloca)

	v.setVar(identifier, Value{
		Value:      alloca,
		Type:       exprValue.Type(),
		IsVariable: true,
	})
}

func (v *Visitor) visitAssigment(ctx *parser.AssignmentContext) {
	leftCtx := ctx.ExpressionList(0).(*parser.ExpressionListContext).Expression(0)
	varName := leftCtx.GetText()
	variable := v.contextBlockVariables[len(v.contextBlockVariables) - 1][varName]

	rightCtx := ctx.ExpressionList(1).(*parser.ExpressionListContext).Expression(0)
	rightValue := v.visitExpression(rightCtx.(*parser.ExpressionContext))

	v.curBlock.NewStore(rightValue, variable.Value)
	v.curBlock.NewLoad(rightValue.Type(), variable.Value)
}

