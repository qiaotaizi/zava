package math

import (
	"github.com/qiaotaizi/zava/instructions/base"
	"github.com/qiaotaizi/zava/rtda"
)

//加法指令

//double加法
type DADD struct {
	base.NoOperandsInstruction
}

func (*DADD) Execute(frame *rtda.Frame) {
	stack:=frame.OperandStack()
	v2:=stack.PopDouble()
	v1:=stack.PopDouble()
	stack.PushDouble(v1+v2)
}

type FADD struct {
	base.NoOperandsInstruction
}

func (*FADD) Execute(frame *rtda.Frame) {
	stack:=frame.OperandStack()
	v2:=stack.PopFloat()
	v1:=stack.PopFloat()
	stack.PushFloat(v1+v2)
}

type IADD struct {
	base.NoOperandsInstruction
}

func (*IADD) Execute(frame *rtda.Frame) {
	stack:=frame.OperandStack()
	v2:=stack.PopInt()
	v1:=stack.PopInt()
	stack.PushInt(v1+v2)
}

type LADD struct {
	base.NoOperandsInstruction
}

func (*LADD) Execute(frame *rtda.Frame) {
	stack:=frame.OperandStack()
	v2:=stack.PopLong()
	v1:=stack.PopLong()
	stack.PushLong(v1+v2)
}
