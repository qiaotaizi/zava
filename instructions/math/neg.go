package math

import (
	"github.com/qiaotaizi/zava/instructions/base"
	"github.com/qiaotaizi/zava/rtda"
)

//取相反数

type DNEG struct {
	base.NoOperandsInstruction
}

func (*DNEG) Execute(frame *rtda.Frame) {
	stack:=frame.OperandStack()
	v:=stack.PopDouble()
	stack.PushDouble(-v)
}

type FNEG struct {
	base.NoOperandsInstruction
}

func (*FNEG) Execute(frame *rtda.Frame) {
	stack:=frame.OperandStack()
	v:=stack.PopFloat()
	stack.PushFloat(-v)
}

type INEG struct {
	base.NoOperandsInstruction
}

func (*INEG) Execute(frame *rtda.Frame) {
	stack:=frame.OperandStack()
	v:=stack.PopInt()
	stack.PushInt(-v)
}

type LNEG struct {
	base.NoOperandsInstruction
}

func (*LNEG) Execute(frame *rtda.Frame) {
	stack:=frame.OperandStack()
	v:=stack.PopLong()
	stack.PushLong(-v)
}

