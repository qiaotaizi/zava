package conversions

import (
	"github.com/qiaotaizi/zava/instructions/base"
	"github.com/qiaotaizi/zava/rtda"
)

//转double
type F2D struct {
	base.NoOperandsInstruction
}

func (*F2D) Execute(frame *rtda.Frame) {
	stack:=frame.OperandStack()
	d:=stack.PopFloat()
	f:=float64(d)
	stack.PushDouble(f)
}

//转int
type F2I struct {
	base.NoOperandsInstruction
}

func (*F2I) Execute(frame *rtda.Frame) {
	stack:=frame.OperandStack()
	d:=stack.PopFloat()
	i:=int32(d)
	stack.PushInt(i)
}

//转long
type F2L struct {
	base.NoOperandsInstruction
}

func (*F2L) Execute(frame *rtda.Frame) {
	stack:=frame.OperandStack()
	d:=stack.PopFloat()
	l:=int64(d)
	stack.PushLong(l)
}

