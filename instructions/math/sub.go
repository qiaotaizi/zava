package math

import (
	"github.com/qiaotaizi/zava/instructions/base"
	"github.com/qiaotaizi/zava/rtda"
)

//减法指令

type DSUB struct {
	base.NoOperandsInstruction
}

func (*DSUB) Execute(frame *rtda.Frame) {
	s:=frame.OperandStack()
	v2:=s.PopDouble()
	v1:=s.PopDouble()
	s.PushDouble(v1-v2)
}

type FSUB struct {
	base.NoOperandsInstruction
}
func (*FSUB) Execute(frame *rtda.Frame) {
	s:=frame.OperandStack()
	v2:=s.PopFloat()
	v1:=s.PopFloat()
	s.PushFloat(v1-v2)
}
type ISUB struct {
	base.NoOperandsInstruction
}
func (*ISUB) Execute(frame *rtda.Frame) {
	s:=frame.OperandStack()
	v2:=s.PopInt()
	v1:=s.PopInt()
	s.PushInt(v1-v2)
}
type LSUB struct {
	base.NoOperandsInstruction
}
func (*LSUB) Execute(frame *rtda.Frame) {
	s:=frame.OperandStack()
	v2:=s.PopLong()
	v1:=s.PopLong()
	s.PushLong(v1-v2)
}
