package visitor

import (
	"nano-go/parser"
	"nano-go/visitor/name"
)

func (v *Visitor) visitForStmt(ctx *parser.ForStmtContext) {
	if ctx.Expression() != nil {
		v.forExpression(ctx)
		return
	}

	if ctx.ForClause() != nil {
		v.fullFor(ctx)
		return
	}

	panic("some strange for cycle")
}

func (v *Visitor) forExpression(ctx *parser.ForStmtContext) {
	condBlock := v.curBlock.Parent.NewBlock(name.BlockName() + "-cond")
	bodyBlock := v.curBlock.Parent.NewBlock(name.BlockName() + "-body")
	afterBlock := v.curBlock.Parent.NewBlock(name.BlockName() + "-after-for")

	v.curBlock.NewBr(condBlock)

	// FOR cond
	v.curBlock = condBlock
	cond := v.visitExpression(ctx.Expression().(*parser.ExpressionContext))
	v.curBlock.NewCondBr(cond, bodyBlock, afterBlock)

	// body
	v.curBlock = bodyBlock
	v.visitBlock(ctx.Block().(*parser.BlockContext))
	v.curBlock.NewBr(condBlock)

	v.curBlock = afterBlock
}

func (v *Visitor) fullFor(ctx *parser.ForStmtContext) {
	bodyBlock := v.curBlock.Parent.NewBlock(name.BlockName() + "-body")
	condBlock := bodyBlock
	afterBlock := v.curBlock.Parent.NewBlock(name.BlockName() + "-after-for")

	forClause := ctx.ForClause().(*parser.ForClauseContext)
	statementNum := 0

	// init step
	if forClause.GetInitStmt() != nil {
		initStmt := forClause.SimpleStmt(statementNum)
		statementNum++
		v.visitSimpleStmt(initStmt.(*parser.SimpleStmtContext))
	}

	if forClause.Expression() != nil {
		condBlock = v.curBlock.Parent.NewBlock(name.BlockName() + "-cond")
		v.curBlock.NewBr(condBlock)

		// cond step
		v.curBlock = condBlock
		cond := v.visitExpression(forClause.Expression().(*parser.ExpressionContext))
		v.curBlock.NewCondBr(cond, bodyBlock, afterBlock)
	} else {
		v.curBlock.NewBr(bodyBlock)
	}

	// body
	v.curBlock = bodyBlock
	v.visitBlock(ctx.Block().(*parser.BlockContext))

	if forClause.GetPostStmt() != nil {
		afterBodyBlock := v.curBlock.Parent.NewBlock(name.BlockName() + "-after-body")
		postStmt := forClause.SimpleStmt(statementNum)

		v.curBlock.NewBr(afterBodyBlock)
		v.curBlock = afterBodyBlock
		v.visitSimpleStmt(postStmt.(*parser.SimpleStmtContext))
		v.curBlock.NewBr(condBlock)
	} else {
		v.curBlock.NewBr(condBlock)
	}

	v.curBlock = afterBlock
}