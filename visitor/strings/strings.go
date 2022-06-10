package strings

import (
	"fmt"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"

	"nano-go/visitor/pointer"
)

func Constant(in string) *constant.CharArray {
	return constant.NewCharArray(append([]byte(in), 0))
}

func Toi8Ptr(block *ir.Block, src value.Value) *ir.InstGetElementPtr {
	return block.NewGetElementPtr(pointer.ElemType(src), src, constant.NewInt(types.I32, 0), constant.NewInt(types.I32, 0))
}

func Len(block *ir.Block, src value.Value) value.Value {
	if _, ok := src.Type().(*types.PointerType); ok {
		l := block.NewGetElementPtr(pointer.ElemType(src), src, constant.NewInt(types.I32, 0), constant.NewInt(types.I32, 0))
		return block.NewLoad(pointer.ElemType(l), l)
	}
	return block.NewExtractValue(src, 0)
}

func TreToI8Ptr(block *ir.Block, src value.Value) value.Value {
	if _, ok := src.Type().(*types.PointerType); ok {
		l := block.NewGetElementPtr(pointer.ElemType(src), src, constant.NewInt(types.I32, 0), constant.NewInt(types.I32, 1))
		return block.NewLoad(pointer.ElemType(l), l)
	}
	return block.NewExtractValue(src, 1)
}

var globalStringCounter uint

func NextStringName() string {
	name := fmt.Sprintf("str.%d", globalStringCounter)
	globalStringCounter++
	return name
}

func String() *types.StructType {
	return types.NewStruct(
		types.I64,                  // String length
		types.NewPointer(types.I8), // Content
	)
}

func StringLen(stringType types.Type) *ir.Func {
	param := ir.NewParam("input", stringType)
	res := ir.NewFunc("string_len", types.I64, param)
	block := res.NewBlock("entry")
	block.NewRet(block.NewExtractValue(param, 0))
	return res
}