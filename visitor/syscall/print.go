package syscall

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"nano-go/visitor/strings"
)

func Print(block *ir.Block, value value.Value, goos string) {
	asmFunc := ir.NewInlineAsm(types.NewPointer(types.NewFunc(types.I64)), "syscall", "=r,{rax},{rdi},{rsi},{rdx}")
	asmFunc.SideEffect = true

	strPtr := strings.TreToI8Ptr(block, value)
	strLen := strings.Len(block, value)

	block.NewCall(asmFunc,
		constant.NewInt(types.I64, Convert(WRITE, goos)), // rax
		constant.NewInt(types.I64, 1),                    // rdi, stdout
		strPtr,                                           // rsi
		strLen,                                           // rdx
	)
}
