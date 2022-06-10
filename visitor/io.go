package visitor

import (
	"github.com/llir/llvm/ir/value"
	"nano-go/visitor/syscall"
)

func (v *Visitor) printFuncCall(arguments []value.Value) value.Value {
	syscall.Print(v.curBlock, arguments[0], "linux")
	return nil
}
