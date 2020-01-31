package math

import (
	"github.com/qiaotaizi/zava/instructions/base"
	"github.com/qiaotaizi/zava/rtda"
)

//乘法指令

type DMUL struct {
	base.NoOperandsInstruction
}

func (*DMUL) Execute(frame *rtda.Frame) {
	stack:=frame.OperandStack()
	v2:=stack.PopDouble()
	v1:=stack.PopDouble()
	stack.PushDouble(v1*v2)
}

type FMUL struct {
	base.NoOperandsInstruction
}

func (*FMUL) Execute(frame *rtda.Frame) {
	stack:=frame.OperandStack()
	v2:=stack.PopFloat()
	v1:=stack.PopFloat()
	stack.PushFloat(v1*v2)
}

type IMUL struct {
	base.NoOperandsInstruction
}

func (*IMUL) Execute(frame *rtda.Frame) {
	stack:=frame.OperandStack()
	v2:=stack.PopInt()
	v1:=stack.PopInt()
	stack.PushInt(v1*v2)
}

type LMUL struct {
	base.NoOperandsInstruction
}

func (*LMUL) Execute(frame *rtda.Frame) {
	stack:=frame.OperandStack()
	v2:=stack.PopLong()
	v1:=stack.PopLong()
	stack.PushLong(v1*v2)
}


