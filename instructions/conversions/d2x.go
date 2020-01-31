package conversions

import (
	"github.com/qiaotaizi/zava/instructions/base"
	"github.com/qiaotaizi/zava/rtda"
)

//double转换为其他类型

//转float
type D2F struct {
	base.NoOperandsInstruction
}

func (*D2F) Execute(frame *rtda.Frame) {
	stack:=frame.OperandStack()
	d:=stack.PopDouble()
	f:=float32(d)
	stack.PushFloat(f)
}

//转int
type D2I struct {
	base.NoOperandsInstruction
}

func (*D2I) Execute(frame *rtda.Frame) {
	stack:=frame.OperandStack()
	d:=stack.PopDouble()
	i:=int32(d)
	stack.PushInt(i)
}

//转long
type D2L struct {
	base.NoOperandsInstruction
}

func (*D2L) Execute(frame *rtda.Frame) {
	stack:=frame.OperandStack()
	d:=stack.PopDouble()
	l:=int64(d)
	stack.PushLong(l)
}
