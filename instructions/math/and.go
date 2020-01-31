package math

import (
	"github.com/qiaotaizi/zava/instructions/base"
	"github.com/qiaotaizi/zava/rtda"
)

//按位与运算

type IAND struct {
	base.NoOperandsInstruction
}

func (i *IAND) Execute(frame *rtda.Frame) {
	stack:=frame.OperandStack()
	v2:=stack.PopInt()
	v1:=stack.PopInt()
	stack.PushInt(v1&v2)
}

type LAND struct {
	base.NoOperandsInstruction
}

func (i *LAND) Execute(frame *rtda.Frame) {
	stack:=frame.OperandStack()
	v2:=stack.PopLong()
	v1:=stack.PopLong()
	stack.PushLong(v1&v2)
}



