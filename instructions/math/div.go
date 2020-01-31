package math

import (
	"github.com/qiaotaizi/zava/instructions/base"
	"github.com/qiaotaizi/zava/rtda"
)

type DDIV struct {
	base.NoOperandsInstruction
}

func (*DDIV) Execute(frame *rtda.Frame) {
	stack:=frame.OperandStack()
	v2:=stack.PopDouble()
	v1:=stack.PopDouble()
	stack.PushDouble(v1/v2)
}

type FDIV struct {
	base.NoOperandsInstruction
}

func (*FDIV) Execute(frame *rtda.Frame) {
	stack:=frame.OperandStack()
	v2:=stack.PopFloat()
	v1:=stack.PopFloat()
	stack.PushFloat(v1/v2)
}

type IDIV struct {
	base.NoOperandsInstruction
}

func (*IDIV) Execute(frame *rtda.Frame) {
	stack:=frame.OperandStack()
	v2:=stack.PopInt()
	if v2==0{
		panic("java.lang.ArithmeticException: / by zero")
	}
	v1:=stack.PopInt()
	stack.PushInt(v1/v2)
}

type LDIV struct {
	base.NoOperandsInstruction
}

func (*LDIV) Execute(frame *rtda.Frame) {
	stack:=frame.OperandStack()
	v2:=stack.PopLong()
	if v2==0{
		panic("java.lang.ArithmeticException: / by zero")
	}
	v1:=stack.PopLong()
	stack.PushLong(v1/v2)
}

