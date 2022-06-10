package _type

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"nano-go/visitor/pointer"
	"nano-go/visitor/strings"
)

type Type interface {
	LLVM() types.Type
	Name() string

	// Size of type in bytes
	Size() int64

	Zero(*ir.Block, value.Value)
}

type backingType struct {
}

func (backingType) Size() int64 {
	panic("Type does not have size set")
}

func (backingType) Zero(*ir.Block, value.Value) {
	// NOOP
}

type StringType struct {
	backingType
	Type types.Type
}

var String = &StringType{}

// Populated by compiler.go
var ModuleStringType types.Type
var EmptyStringConstant *ir.Global

func (StringType) LLVM() types.Type {
	return ModuleStringType
}

func (StringType) Name() string {
	return "string"
}

func (StringType) Size() int64 {
	return 16
}

func (s StringType) Zero(block *ir.Block, alloca value.Value) {
	lenPtr := block.NewGetElementPtr(pointer.ElemType(alloca), alloca, constant.NewInt(types.I32, 0), constant.NewInt(types.I32, 0))
	backingDataPtr := block.NewGetElementPtr(pointer.ElemType(alloca), alloca, constant.NewInt(types.I32, 0), constant.NewInt(types.I32, 1))
	block.NewStore(constant.NewInt(types.I64, 0), lenPtr)
	block.NewStore(strings.Toi8Ptr(block, EmptyStringConstant), backingDataPtr)
}