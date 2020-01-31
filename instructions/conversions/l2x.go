package conversions

import (
	"github.com/qiaotaizi/zava/instructions/base"
	"github.com/qiaotaizi/zava/rtda"
)

//转double
type L2D struct {
	base.NoOperandsInstruction
}

func (*L2D) Execute(frame *rtda.Frame) {
	stack:=frame.OperandStack()
	d:=stack.PopLong()
	f:=float64(d)
	stack.PushDouble(f)
}

//转int
type L2I struct {
	base.NoOperandsInstruction
}

func (*L2I) Execute(frame *rtda.Frame) {
	stack:=frame.OperandStack()
	d:=stack.PopLong()
	i:=int32(d)
	stack.PushInt(i)
}

//转float
type L2F struct {
	base.NoOperandsInstruction
}

func (*L2F) Execute(frame *rtda.Frame) {
	stack:=frame.OperandStack()
	d:=stack.PopLong()
	f:=float32(d)
	stack.PushFloat(f)
}

