// Generated from GoParser.g4 by ANTLR 4.7.

package parser // GoParser

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"nano-go/ast"
)

// GoListener is a complete listener for a parse tree produced by GoParser.
type GoListener struct{}

var Listener = &GoListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *GoListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *GoListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *GoListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *GoListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterSourceFile is called when production sourceFile is entered.
func (s *GoListener) EnterSourceFile(ctx *SourceFileContext) {}

// ExitSourceFile is called when production sourceFile is exited.
func (s *GoListener) ExitSourceFile(ctx *SourceFileContext) {}

// EnterPackageClause is called when production packageClause is entered.
func (s *GoListener) EnterPackageClause(ctx *PackageClauseContext) {}

// ExitPackageClause is called when production packageClause is exited.
func (s *GoListener) ExitPackageClause(ctx *PackageClauseContext) {}

// EnterImportDecl is called when production importDecl is entered.
func (s *GoListener) EnterImportDecl(ctx *ImportDeclContext) {}

// ExitImportDecl is called when production importDecl is exited.
func (s *GoListener) ExitImportDecl(ctx *ImportDeclContext) {}

// EnterImportSpec is called when production importSpec is entered.
func (s *GoListener) EnterImportSpec(ctx *ImportSpecContext) {}

// ExitImportSpec is called when production importSpec is exited.
func (s *GoListener) ExitImportSpec(ctx *ImportSpecContext) {}

// EnterImportPath is called when production importPath is entered.
func (s *GoListener) EnterImportPath(ctx *ImportPathContext) {}

// ExitImportPath is called when production importPath is exited.
func (s *GoListener) ExitImportPath(ctx *ImportPathContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *GoListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *GoListener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterConstDecl is called when production constDecl is entered.
func (s *GoListener) EnterConstDecl(ctx *ConstDeclContext) {}

// ExitConstDecl is called when production constDecl is exited.
func (s *GoListener) ExitConstDecl(ctx *ConstDeclContext) {}

// EnterConstSpec is called when production constSpec is entered.
func (s *GoListener) EnterConstSpec(ctx *ConstSpecContext) {}

// ExitConstSpec is called when production constSpec is exited.
func (s *GoListener) ExitConstSpec(ctx *ConstSpecContext) {}

// EnterIdentifierList is called when production identifierList is entered.
func (s *GoListener) EnterIdentifierList(ctx *IdentifierListContext) {}

// ExitIdentifierList is called when production identifierList is exited.
func (s *GoListener) ExitIdentifierList(ctx *IdentifierListContext) {}

// EnterExpressionList is called when production expressionList is entered.
func (s *GoListener) EnterExpressionList(ctx *ExpressionListContext) {}

// ExitExpressionList is called when production expressionList is exited.
func (s *GoListener) ExitExpressionList(ctx *ExpressionListContext) {}

// EnterTypeDecl is called when production typeDecl is entered.
func (s *GoListener) EnterTypeDecl(ctx *TypeDeclContext) {}

// ExitTypeDecl is called when production typeDecl is exited.
func (s *GoListener) ExitTypeDecl(ctx *TypeDeclContext) {}

// EnterTypeSpec is called when production typeSpec is entered.
func (s *GoListener) EnterTypeSpec(ctx *TypeSpecContext) {}

// ExitTypeSpec is called when production typeSpec is exited.
func (s *GoListener) ExitTypeSpec(ctx *TypeSpecContext) {}

// EnterFunctionDecl is called when production functionDecl is entered.
func (s *GoListener) EnterFunctionDecl(ctx *FunctionDeclContext) {

}

// ExitFunctionDecl is called when production functionDecl is exited.
func (s *GoListener) ExitFunctionDecl(ctx *FunctionDeclContext) {
	fmt.Println(ctx.Signature().GetText())
	fmt.Println(ctx.GetText())
	block := ctx.Block()

	fmt.Println(block)
	fNode := ast.FuncLit{
		Type: &ast.FuncType{

		},
		Body: &ast.BlockStmt{

		},
	}
	fmt.Println(fNode)
	fmt.Println(ctx.IDENTIFIER())
	fmt.Println(ctx.FUNC())
}

// EnterMethodDecl is called when production methodDecl is entered.
func (s *GoListener) EnterMethodDecl(ctx *MethodDeclContext) {}

// ExitMethodDecl is called when production methodDecl is exited.
func (s *GoListener) ExitMethodDecl(ctx *MethodDeclContext) {}

// EnterReceiver is called when production receiver is entered.
func (s *GoListener) EnterReceiver(ctx *ReceiverContext) {}

// ExitReceiver is called when production receiver is exited.
func (s *GoListener) ExitReceiver(ctx *ReceiverContext) {}

// EnterVarDecl is called when production varDecl is entered.
func (s *GoListener) EnterVarDecl(ctx *VarDeclContext) {}

// ExitVarDecl is called when production varDecl is exited.
func (s *GoListener) ExitVarDecl(ctx *VarDeclContext) {}

// EnterVarSpec is called when production varSpec is entered.
func (s *GoListener) EnterVarSpec(ctx *VarSpecContext) {}

// ExitVarSpec is called when production varSpec is exited.
func (s *GoListener) ExitVarSpec(ctx *VarSpecContext) {}

// EnterBlock is called when production block is entered.
func (s *GoListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *GoListener) ExitBlock(ctx *BlockContext) {}

// EnterStatementList is called when production statementList is entered.
func (s *GoListener) EnterStatementList(ctx *StatementListContext) {}

// ExitStatementList is called when production statementList is exited.
func (s *GoListener) ExitStatementList(ctx *StatementListContext) {}

// EnterStatement is called when production statement is entered.
func (s *GoListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *GoListener) ExitStatement(ctx *StatementContext) {}

// EnterSimpleStmt is called when production simpleStmt is entered.
func (s *GoListener) EnterSimpleStmt(ctx *SimpleStmtContext) {}

// ExitSimpleStmt is called when production simpleStmt is exited.
func (s *GoListener) ExitSimpleStmt(ctx *SimpleStmtContext) {}

// EnterExpressionStmt is called when production expressionStmt is entered.
func (s *GoListener) EnterExpressionStmt(ctx *ExpressionStmtContext) {}

// ExitExpressionStmt is called when production expressionStmt is exited.
func (s *GoListener) ExitExpressionStmt(ctx *ExpressionStmtContext) {}

// EnterSendStmt is called when production sendStmt is entered.
func (s *GoListener) EnterSendStmt(ctx *SendStmtContext) {}

// ExitSendStmt is called when production sendStmt is exited.
func (s *GoListener) ExitSendStmt(ctx *SendStmtContext) {}

// EnterIncDecStmt is called when production incDecStmt is entered.
func (s *GoListener) EnterIncDecStmt(ctx *IncDecStmtContext) {}

// ExitIncDecStmt is called when production incDecStmt is exited.
func (s *GoListener) ExitIncDecStmt(ctx *IncDecStmtContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *GoListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *GoListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterAssign_op is called when production assign_op is entered.
func (s *GoListener) EnterAssign_op(ctx *Assign_opContext) {}

// ExitAssign_op is called when production assign_op is exited.
func (s *GoListener) ExitAssign_op(ctx *Assign_opContext) {}

// EnterShortVarDecl is called when production shortVarDecl is entered.
func (s *GoListener) EnterShortVarDecl(ctx *ShortVarDeclContext) {}

// ExitShortVarDecl is called when production shortVarDecl is exited.
func (s *GoListener) ExitShortVarDecl(ctx *ShortVarDeclContext) {}

// EnterEmptyStmt is called when production emptyStmt is entered.
func (s *GoListener) EnterEmptyStmt(ctx *EmptyStmtContext) {}

// ExitEmptyStmt is called when production emptyStmt is exited.
func (s *GoListener) ExitEmptyStmt(ctx *EmptyStmtContext) {}

// EnterLabeledStmt is called when production labeledStmt is entered.
func (s *GoListener) EnterLabeledStmt(ctx *LabeledStmtContext) {}

// ExitLabeledStmt is called when production labeledStmt is exited.
func (s *GoListener) ExitLabeledStmt(ctx *LabeledStmtContext) {}

// EnterReturnStmt is called when production returnStmt is entered.
func (s *GoListener) EnterReturnStmt(ctx *ReturnStmtContext) {}

// ExitReturnStmt is called when production returnStmt is exited.
func (s *GoListener) ExitReturnStmt(ctx *ReturnStmtContext) {}

// EnterBreakStmt is called when production breakStmt is entered.
func (s *GoListener) EnterBreakStmt(ctx *BreakStmtContext) {}

// ExitBreakStmt is called when production breakStmt is exited.
func (s *GoListener) ExitBreakStmt(ctx *BreakStmtContext) {}

// EnterContinueStmt is called when production continueStmt is entered.
func (s *GoListener) EnterContinueStmt(ctx *ContinueStmtContext) {}

// ExitContinueStmt is called when production continueStmt is exited.
func (s *GoListener) ExitContinueStmt(ctx *ContinueStmtContext) {}

// EnterGotoStmt is called when production gotoStmt is entered.
func (s *GoListener) EnterGotoStmt(ctx *GotoStmtContext) {}

// ExitGotoStmt is called when production gotoStmt is exited.
func (s *GoListener) ExitGotoStmt(ctx *GotoStmtContext) {}

// EnterFallthroughStmt is called when production fallthroughStmt is entered.
func (s *GoListener) EnterFallthroughStmt(ctx *FallthroughStmtContext) {}

// ExitFallthroughStmt is called when production fallthroughStmt is exited.
func (s *GoListener) ExitFallthroughStmt(ctx *FallthroughStmtContext) {}

// EnterDeferStmt is called when production deferStmt is entered.
func (s *GoListener) EnterDeferStmt(ctx *DeferStmtContext) {}

// ExitDeferStmt is called when production deferStmt is exited.
func (s *GoListener) ExitDeferStmt(ctx *DeferStmtContext) {}

// EnterIfStmt is called when production ifStmt is entered.
func (s *GoListener) EnterIfStmt(ctx *IfStmtContext) {}

// ExitIfStmt is called when production ifStmt is exited.
func (s *GoListener) ExitIfStmt(ctx *IfStmtContext) {}

// EnterSwitchStmt is called when production switchStmt is entered.
func (s *GoListener) EnterSwitchStmt(ctx *SwitchStmtContext) {}

// ExitSwitchStmt is called when production switchStmt is exited.
func (s *GoListener) ExitSwitchStmt(ctx *SwitchStmtContext) {}

// EnterExprSwitchStmt is called when production exprSwitchStmt is entered.
func (s *GoListener) EnterExprSwitchStmt(ctx *ExprSwitchStmtContext) {}

// ExitExprSwitchStmt is called when production exprSwitchStmt is exited.
func (s *GoListener) ExitExprSwitchStmt(ctx *ExprSwitchStmtContext) {}

// EnterExprCaseClause is called when production exprCaseClause is entered.
func (s *GoListener) EnterExprCaseClause(ctx *ExprCaseClauseContext) {}

// ExitExprCaseClause is called when production exprCaseClause is exited.
func (s *GoListener) ExitExprCaseClause(ctx *ExprCaseClauseContext) {}

// EnterExprSwitchCase is called when production exprSwitchCase is entered.
func (s *GoListener) EnterExprSwitchCase(ctx *ExprSwitchCaseContext) {}

// ExitExprSwitchCase is called when production exprSwitchCase is exited.
func (s *GoListener) ExitExprSwitchCase(ctx *ExprSwitchCaseContext) {}

// EnterTypeSwitchStmt is called when production typeSwitchStmt is entered.
func (s *GoListener) EnterTypeSwitchStmt(ctx *TypeSwitchStmtContext) {}

// ExitTypeSwitchStmt is called when production typeSwitchStmt is exited.
func (s *GoListener) ExitTypeSwitchStmt(ctx *TypeSwitchStmtContext) {}

// EnterTypeSwitchGuard is called when production typeSwitchGuard is entered.
func (s *GoListener) EnterTypeSwitchGuard(ctx *TypeSwitchGuardContext) {}

// ExitTypeSwitchGuard is called when production typeSwitchGuard is exited.
func (s *GoListener) ExitTypeSwitchGuard(ctx *TypeSwitchGuardContext) {}

// EnterTypeCaseClause is called when production typeCaseClause is entered.
func (s *GoListener) EnterTypeCaseClause(ctx *TypeCaseClauseContext) {}

// ExitTypeCaseClause is called when production typeCaseClause is exited.
func (s *GoListener) ExitTypeCaseClause(ctx *TypeCaseClauseContext) {}

// EnterTypeSwitchCase is called when production typeSwitchCase is entered.
func (s *GoListener) EnterTypeSwitchCase(ctx *TypeSwitchCaseContext) {}

// ExitTypeSwitchCase is called when production typeSwitchCase is exited.
func (s *GoListener) ExitTypeSwitchCase(ctx *TypeSwitchCaseContext) {}

// EnterTypeList is called when production typeList is entered.
func (s *GoListener) EnterTypeList(ctx *TypeListContext) {}

// ExitTypeList is called when production typeList is exited.
func (s *GoListener) ExitTypeList(ctx *TypeListContext) {}

// EnterSelectStmt is called when production selectStmt is entered.
func (s *GoListener) EnterSelectStmt(ctx *SelectStmtContext) {}

// ExitSelectStmt is called when production selectStmt is exited.
func (s *GoListener) ExitSelectStmt(ctx *SelectStmtContext) {}

// EnterCommClause is called when production commClause is entered.
func (s *GoListener) EnterCommClause(ctx *CommClauseContext) {}

// ExitCommClause is called when production commClause is exited.
func (s *GoListener) ExitCommClause(ctx *CommClauseContext) {}

// EnterCommCase is called when production commCase is entered.
func (s *GoListener) EnterCommCase(ctx *CommCaseContext) {}

// ExitCommCase is called when production commCase is exited.
func (s *GoListener) ExitCommCase(ctx *CommCaseContext) {}

// EnterRecvStmt is called when production recvStmt is entered.
func (s *GoListener) EnterRecvStmt(ctx *RecvStmtContext) {}

// ExitRecvStmt is called when production recvStmt is exited.
func (s *GoListener) ExitRecvStmt(ctx *RecvStmtContext) {}

// EnterForStmt is called when production forStmt is entered.
func (s *GoListener) EnterForStmt(ctx *ForStmtContext) {}

// ExitForStmt is called when production forStmt is exited.
func (s *GoListener) ExitForStmt(ctx *ForStmtContext) {}

// EnterForClause is called when production forClause is entered.
func (s *GoListener) EnterForClause(ctx *ForClauseContext) {}

// ExitForClause is called when production forClause is exited.
func (s *GoListener) ExitForClause(ctx *ForClauseContext) {}

// EnterRangeClause is called when production rangeClause is entered.
func (s *GoListener) EnterRangeClause(ctx *RangeClauseContext) {}

// ExitRangeClause is called when production rangeClause is exited.
func (s *GoListener) ExitRangeClause(ctx *RangeClauseContext) {}

// EnterGoStmt is called when production goStmt is entered.
func (s *GoListener) EnterGoStmt(ctx *GoStmtContext) {}

// ExitGoStmt is called when production goStmt is exited.
func (s *GoListener) ExitGoStmt(ctx *GoStmtContext) {}

// EnterType_ is called when production type_ is entered.
func (s *GoListener) EnterType_(ctx *Type_Context) {}

// ExitType_ is called when production type_ is exited.
func (s *GoListener) ExitType_(ctx *Type_Context) {}

// EnterTypeName is called when production typeName is entered.
func (s *GoListener) EnterTypeName(ctx *TypeNameContext) {}

// ExitTypeName is called when production typeName is exited.
func (s *GoListener) ExitTypeName(ctx *TypeNameContext) {}

// EnterTypeLit is called when production typeLit is entered.
func (s *GoListener) EnterTypeLit(ctx *TypeLitContext) {}

// ExitTypeLit is called when production typeLit is exited.
func (s *GoListener) ExitTypeLit(ctx *TypeLitContext) {}

// EnterArrayType is called when production arrayType is entered.
func (s *GoListener) EnterArrayType(ctx *ArrayTypeContext) {}

// ExitArrayType is called when production arrayType is exited.
func (s *GoListener) ExitArrayType(ctx *ArrayTypeContext) {}

// EnterArrayLength is called when production arrayLength is entered.
func (s *GoListener) EnterArrayLength(ctx *ArrayLengthContext) {}

// ExitArrayLength is called when production arrayLength is exited.
func (s *GoListener) ExitArrayLength(ctx *ArrayLengthContext) {}

// EnterElementType is called when production elementType is entered.
func (s *GoListener) EnterElementType(ctx *ElementTypeContext) {}

// ExitElementType is called when production elementType is exited.
func (s *GoListener) ExitElementType(ctx *ElementTypeContext) {}

// EnterPointerType is called when production pointerType is entered.
func (s *GoListener) EnterPointerType(ctx *PointerTypeContext) {}

// ExitPointerType is called when production pointerType is exited.
func (s *GoListener) ExitPointerType(ctx *PointerTypeContext) {}

// EnterInterfaceType is called when production interfaceType is entered.
func (s *GoListener) EnterInterfaceType(ctx *InterfaceTypeContext) {}

// ExitInterfaceType is called when production interfaceType is exited.
func (s *GoListener) ExitInterfaceType(ctx *InterfaceTypeContext) {}

// EnterSliceType is called when production sliceType is entered.
func (s *GoListener) EnterSliceType(ctx *SliceTypeContext) {}

// ExitSliceType is called when production sliceType is exited.
func (s *GoListener) ExitSliceType(ctx *SliceTypeContext) {}

// EnterMapType is called when production mapType is entered.
func (s *GoListener) EnterMapType(ctx *MapTypeContext) {}

// ExitMapType is called when production mapType is exited.
func (s *GoListener) ExitMapType(ctx *MapTypeContext) {}

// EnterChannelType is called when production channelType is entered.
func (s *GoListener) EnterChannelType(ctx *ChannelTypeContext) {}

// ExitChannelType is called when production channelType is exited.
func (s *GoListener) ExitChannelType(ctx *ChannelTypeContext) {}

// EnterMethodSpec is called when production methodSpec is entered.
func (s *GoListener) EnterMethodSpec(ctx *MethodSpecContext) {}

// ExitMethodSpec is called when production methodSpec is exited.
func (s *GoListener) ExitMethodSpec(ctx *MethodSpecContext) {}

// EnterFunctionType is called when production functionType is entered.
func (s *GoListener) EnterFunctionType(ctx *FunctionTypeContext) {}

// ExitFunctionType is called when production functionType is exited.
func (s *GoListener) ExitFunctionType(ctx *FunctionTypeContext) {}

// EnterSignature is called when production signature is entered.
func (s *GoListener) EnterSignature(ctx *SignatureContext) {}

// ExitSignature is called when production signature is exited.
func (s *GoListener) ExitSignature(ctx *SignatureContext) {}

// EnterResult is called when production result is entered.
func (s *GoListener) EnterResult(ctx *ResultContext) {}

// ExitResult is called when production result is exited.
func (s *GoListener) ExitResult(ctx *ResultContext) {}

// EnterParameters is called when production parameters is entered.
func (s *GoListener) EnterParameters(ctx *ParametersContext) {}

// ExitParameters is called when production parameters is exited.
func (s *GoListener) ExitParameters(ctx *ParametersContext) {}

// EnterParameterDecl is called when production parameterDecl is entered.
func (s *GoListener) EnterParameterDecl(ctx *ParameterDeclContext) {}

// ExitParameterDecl is called when production parameterDecl is exited.
func (s *GoListener) ExitParameterDecl(ctx *ParameterDeclContext) {}

// EnterExpression is called when production expression is entered.
func (s *GoListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *GoListener) ExitExpression(ctx *ExpressionContext) {}

// EnterPrimaryExpr is called when production primaryExpr is entered.
func (s *GoListener) EnterPrimaryExpr(ctx *PrimaryExprContext) {}

// ExitPrimaryExpr is called when production primaryExpr is exited.
func (s *GoListener) ExitPrimaryExpr(ctx *PrimaryExprContext) {}

// EnterConversion is called when production conversion is entered.
func (s *GoListener) EnterConversion(ctx *ConversionContext) {}

// ExitConversion is called when production conversion is exited.
func (s *GoListener) ExitConversion(ctx *ConversionContext) {}

// EnterNonNamedType is called when production nonNamedType is entered.
func (s *GoListener) EnterNonNamedType(ctx *NonNamedTypeContext) {}

// ExitNonNamedType is called when production nonNamedType is exited.
func (s *GoListener) ExitNonNamedType(ctx *NonNamedTypeContext) {}

// EnterOperand is called when production operand is entered.
func (s *GoListener) EnterOperand(ctx *OperandContext) {}

// ExitOperand is called when production operand is exited.
func (s *GoListener) ExitOperand(ctx *OperandContext) {}

// EnterLiteral is called when production literal is entered.
func (s *GoListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *GoListener) ExitLiteral(ctx *LiteralContext) {}

// EnterBasicLit is called when production basicLit is entered.
func (s *GoListener) EnterBasicLit(ctx *BasicLitContext) {}

// ExitBasicLit is called when production basicLit is exited.
func (s *GoListener) ExitBasicLit(ctx *BasicLitContext) {}

// EnterInteger is called when production integer is entered.
func (s *GoListener) EnterInteger(ctx *IntegerContext) {}

// ExitInteger is called when production integer is exited.
func (s *GoListener) ExitInteger(ctx *IntegerContext) {}

// EnterOperandName is called when production operandName is entered.
func (s *GoListener) EnterOperandName(ctx *OperandNameContext) {}

// ExitOperandName is called when production operandName is exited.
func (s *GoListener) ExitOperandName(ctx *OperandNameContext) {}

// EnterQualifiedIdent is called when production qualifiedIdent is entered.
func (s *GoListener) EnterQualifiedIdent(ctx *QualifiedIdentContext) {}

// ExitQualifiedIdent is called when production qualifiedIdent is exited.
func (s *GoListener) ExitQualifiedIdent(ctx *QualifiedIdentContext) {}

// EnterCompositeLit is called when production compositeLit is entered.
func (s *GoListener) EnterCompositeLit(ctx *CompositeLitContext) {}

// ExitCompositeLit is called when production compositeLit is exited.
func (s *GoListener) ExitCompositeLit(ctx *CompositeLitContext) {}

// EnterLiteralType is called when production literalType is entered.
func (s *GoListener) EnterLiteralType(ctx *LiteralTypeContext) {}

// ExitLiteralType is called when production literalType is exited.
func (s *GoListener) ExitLiteralType(ctx *LiteralTypeContext) {}

// EnterLiteralValue is called when production literalValue is entered.
func (s *GoListener) EnterLiteralValue(ctx *LiteralValueContext) {}

// ExitLiteralValue is called when production literalValue is exited.
func (s *GoListener) ExitLiteralValue(ctx *LiteralValueContext) {}

// EnterElementList is called when production elementList is entered.
func (s *GoListener) EnterElementList(ctx *ElementListContext) {}

// ExitElementList is called when production elementList is exited.
func (s *GoListener) ExitElementList(ctx *ElementListContext) {}

// EnterKeyedElement is called when production keyedElement is entered.
func (s *GoListener) EnterKeyedElement(ctx *KeyedElementContext) {}

// ExitKeyedElement is called when production keyedElement is exited.
func (s *GoListener) ExitKeyedElement(ctx *KeyedElementContext) {}

// EnterKey is called when production key is entered.
func (s *GoListener) EnterKey(ctx *KeyContext) {}

// ExitKey is called when production key is exited.
func (s *GoListener) ExitKey(ctx *KeyContext) {}

// EnterElement is called when production element is entered.
func (s *GoListener) EnterElement(ctx *ElementContext) {}

// ExitElement is called when production element is exited.
func (s *GoListener) ExitElement(ctx *ElementContext) {}

// EnterStructType is called when production structType is entered.
func (s *GoListener) EnterStructType(ctx *StructTypeContext) {}

// ExitStructType is called when production structType is exited.
func (s *GoListener) ExitStructType(ctx *StructTypeContext) {}

// EnterFieldDecl is called when production fieldDecl is entered.
func (s *GoListener) EnterFieldDecl(ctx *FieldDeclContext) {}

// ExitFieldDecl is called when production fieldDecl is exited.
func (s *GoListener) ExitFieldDecl(ctx *FieldDeclContext) {}

// EnterString_ is called when production string_ is entered.
func (s *GoListener) EnterString_(ctx *String_Context) {}

// ExitString_ is called when production string_ is exited.
func (s *GoListener) ExitString_(ctx *String_Context) {}

// EnterEmbeddedField is called when production embeddedField is entered.
func (s *GoListener) EnterEmbeddedField(ctx *EmbeddedFieldContext) {}

// ExitEmbeddedField is called when production embeddedField is exited.
func (s *GoListener) ExitEmbeddedField(ctx *EmbeddedFieldContext) {}

// EnterFunctionLit is called when production functionLit is entered.
func (s *GoListener) EnterFunctionLit(ctx *FunctionLitContext) {}

// ExitFunctionLit is called when production functionLit is exited.
func (s *GoListener) ExitFunctionLit(ctx *FunctionLitContext) {}

// EnterIndex is called when production index is entered.
func (s *GoListener) EnterIndex(ctx *IndexContext) {}

// ExitIndex is called when production index is exited.
func (s *GoListener) ExitIndex(ctx *IndexContext) {}

// EnterSlice_ is called when production slice_ is entered.
func (s *GoListener) EnterSlice_(ctx *Slice_Context) {}

// ExitSlice_ is called when production slice_ is exited.
func (s *GoListener) ExitSlice_(ctx *Slice_Context) {}

// EnterTypeAssertion is called when production typeAssertion is entered.
func (s *GoListener) EnterTypeAssertion(ctx *TypeAssertionContext) {}

// ExitTypeAssertion is called when production typeAssertion is exited.
func (s *GoListener) ExitTypeAssertion(ctx *TypeAssertionContext) {}

// EnterArguments is called when production arguments is entered.
func (s *GoListener) EnterArguments(ctx *ArgumentsContext) {}

// ExitArguments is called when production arguments is exited.
func (s *GoListener) ExitArguments(ctx *ArgumentsContext) {}

// EnterMethodExpr is called when production methodExpr is entered.
func (s *GoListener) EnterMethodExpr(ctx *MethodExprContext) {}

// ExitMethodExpr is called when production methodExpr is exited.
func (s *GoListener) ExitMethodExpr(ctx *MethodExprContext) {}

// EnterReceiverType is called when production receiverType is entered.
func (s *GoListener) EnterReceiverType(ctx *ReceiverTypeContext) {}

// ExitReceiverType is called when production receiverType is exited.
func (s *GoListener) ExitReceiverType(ctx *ReceiverTypeContext) {}

// EnterEos is called when production eos is entered.
func (s *GoListener) EnterEos(ctx *EosContext) {}

// ExitEos is called when production eos is exited.
func (s *GoListener) ExitEos(ctx *EosContext) {}
