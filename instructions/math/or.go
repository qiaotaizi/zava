package math

import (
	"github.com/qiaotaizi/zava/instructions/base"
	"github.com/qiaotaizi/zava/rtda"
)

//按位或

type IOR struct {
	base.NoOperandsInstruction
}

func (i *IOR) Execute(frame *rtda.Frame) {
	stack:=frame.OperandStack()
	v2:=stack.PopInt()
	v1:=stack.PopInt()
	stack.PushInt(v1|v2)
}

type LOR struct {
	base.NoOperandsInstruction
}

func (i *LOR) Execute(frame *rtda.Frame) {
	stack:=frame.OperandStack()
	v2:=stack.PopLong()
	v1:=stack.PopLong()
	stack.PushLong(v1|v2)
}

