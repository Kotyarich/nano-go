package visitor

import (
	"fmt"
	"strconv"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"

	"nano-go/parser"
	"nano-go/visitor/pointer"
	"nano-go/visitor/strings"
	"nano-go/visitor/type"
)

func (v *Visitor) visitExpression(ctx *parser.ExpressionContext) value.Value {
	unary := ctx.GetUnary_op()
	if unary != nil {
		if unary.GetTokenType() == parser.GoParserMINUS {
			exprValue := v.visitExpression(ctx.Expression(0).(*parser.ExpressionContext))
			valueType := exprValue.Type()
			switch valueType {
			case types.I32, types.I64:
				zero := constant.NewInt(types.I64, 0)
				return v.curBlock.NewSub(zero, exprValue)
			}
		}
	}

	if addOp := ctx.GetAdd_op(); addOp != nil {
		leftValue := v.visitExpression(ctx.Expression(0).(*parser.ExpressionContext))
		rightValue := v.visitExpression(ctx.Expression(1).(*parser.ExpressionContext))

		switch addOp.GetTokenType() {
		case parser.GoParserPLUS:
			return v.curBlock.NewAdd(leftValue, rightValue)
		case parser.GoParserMINUS:
			return v.curBlock.NewSub(leftValue, rightValue)
		case parser.GoParserOR:
			return v.curBlock.NewOr(leftValue, rightValue)
		case parser.GoParserCARET:
			panic("caret is not implemented yet")
		}
	}

	if mulOp := ctx.GetMul_op(); mulOp != nil {
		leftValue := v.visitExpression(ctx.Expression(0).(*parser.ExpressionContext))
		rightValue := v.visitExpression(ctx.Expression(1).(*parser.ExpressionContext))

		switch mulOp.GetTokenType() {
		case parser.GoParserSTAR:
			return v.curBlock.NewMul(leftValue, rightValue)
		case parser.GoParserDIV:
			return v.curBlock.NewSDiv(leftValue, rightValue)
		case parser.GoParserMOD:
			panic("mod is not implemented yet")
		case parser.GoParserLSHIFT:
			return v.curBlock.NewShl(leftValue, rightValue)
		case parser.GoParserRSHIFT:
			return v.curBlock.NewAShr(leftValue, rightValue)
		case parser.GoParserAMPERSAND:
			return v.curBlock.NewAnd(leftValue, rightValue)
		case parser.GoParserBIT_CLEAR:
			panic("bit clear is not implemented yet")
		}
	}

	if relOp := ctx.GetRel_op(); relOp != nil {
		leftValue := v.visitExpression(ctx.Expression(0).(*parser.ExpressionContext))
		rightValue := v.visitExpression(ctx.Expression(1).(*parser.ExpressionContext))

		switch relOp.GetTokenType() {
		case parser.GoParserLESS:
			return v.curBlock.NewICmp(enum.IPredSLT, leftValue, rightValue)
		case parser.GoParserGREATER:
			return v.curBlock.NewICmp(enum.IPredSGT, leftValue, rightValue)
		case parser.GoParserLESS_OR_EQUALS:
			return v.curBlock.NewICmp(enum.IPredSLE, leftValue, rightValue)
		case parser.GoParserGREATER_OR_EQUALS:
			return v.curBlock.NewICmp(enum.IPredSGE, leftValue, rightValue)
		case parser.GoParserEQUALS:
			return v.curBlock.NewICmp(enum.IPredEQ, leftValue, rightValue)
		case parser.GoParserNOT_EQUALS:
			return v.curBlock.NewICmp(enum.IPredNE, leftValue, rightValue)
		}
	}

	if primaryExpr := ctx.PrimaryExpr(); primaryExpr != nil {
		return v.visitPrimaryExpr(primaryExpr.(*parser.PrimaryExprContext))
	}

	panic("unknown expression")
}

func (v *Visitor) visitPrimaryExpr(ctx *parser.PrimaryExprContext) value.Value {
	if operand := ctx.Operand(); operand != nil {
		return v.visitOperand(operand.(*parser.OperandContext))
	}

	if primaryExprI := ctx.PrimaryExpr(); primaryExprI != nil {
		primaryExpr := primaryExprI.(*parser.PrimaryExprContext)
		name := primaryExpr.GetText()
		argumentsI := ctx.Arguments()
		if argumentsI != nil {
			arguments := argumentsI.(*parser.ArgumentsContext)
			var argumentsValues []value.Value
			for _, exprI := range arguments.ExpressionList().(*parser.ExpressionListContext).AllExpression() {
				expr := exprI.(*parser.ExpressionContext)
				exprValue := v.visitExpression(expr)
				argumentsValues = append(argumentsValues, exprValue)
			}
			return v.visitFuncCall(name, argumentsValues)
		}
	}

	panic("unknown primary expression")
}

func (v *Visitor) visitOperand(ctx *parser.OperandContext) value.Value {
	if name := ctx.OperandName(); name != nil {
		return v.visitOperandName(name.(*parser.OperandNameContext))
	}

	if literal := ctx.Literal(); literal != nil {
		return v.visitLiteral(literal.(*parser.LiteralContext))
	}

	panic(fmt.Errorf("unknown operand"))
}

func (v *Visitor) visitOperandName(ctx *parser.OperandNameContext) value.Value {
	name := ctx.IDENTIFIER().GetText()
	//variable, ok := v.variables[name]
	variable, ok := v.contextBlockVariables[len(v.contextBlockVariables)-1][name]
	if !ok {
		panic(fmt.Errorf("unknown variable: %s", name))
	}

	load := v.curBlock.NewLoad(variable.Type, variable.Value)

	return load
}

func (v *Visitor) visitLiteral(ctx *parser.LiteralContext) value.Value {
	if basicLit := ctx.BasicLit(); basicLit != nil {
		return v.visitBasicLit(basicLit.(*parser.BasicLitContext))
	}

	panic("unknown literal")
}

func (v *Visitor) visitBasicLit(ctx *parser.BasicLitContext) value.Value {
	if str := ctx.String_(); str != nil {

	}

	if integer := ctx.Integer(); integer != nil {
		literal := integer.(*parser.IntegerContext).DECIMAL_LIT().GetText()
		constantValue, _ := strconv.Atoi(literal)
		return constant.NewInt(types.I64, int64(constantValue))
	}

	if floatLit := ctx.FLOAT_LIT(); floatLit != nil {
		panic("float is not realised yet")
	}

	if stringLit := ctx.String_(); stringLit != nil {
		var constString *ir.Global
		valueStr := ctx.GetText()
		valueStr = valueStr[1:len(valueStr)-1]


		constString = v.Module.NewGlobalDef(strings.NextStringName(), strings.Constant(valueStr))
		constString.Immutable = true

		alloc := v.curBlock.NewAlloca(_type.String.LLVM())

		// Save length of the string
		lenItem := v.curBlock.NewGetElementPtr(pointer.ElemType(alloc), alloc, constant.NewInt(types.I32, 0), constant.NewInt(types.I32, 0))
		v.curBlock.NewStore(constant.NewInt(types.I64, int64(len(valueStr))), lenItem)

		// Save i8* version of string
		strItem := v.curBlock.NewGetElementPtr(pointer.ElemType(alloc), alloc, constant.NewInt(types.I32, 0), constant.NewInt(types.I32, 1))
		v.curBlock.NewStore(strings.Toi8Ptr(v.curBlock, constString), strItem)

		return v.curBlock.NewLoad(pointer.ElemType(alloc), alloc)
	}

	panic("unknown basic literal")
}

