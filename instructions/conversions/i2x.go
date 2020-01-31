package conversions

import (
	"github.com/qiaotaizi/zava/instructions/base"
	"github.com/qiaotaizi/zava/rtda"
)

//转double
type I2D struct {
	base.NoOperandsInstruction
}

func (*I2D) Execute(frame *rtda.Frame) {
	stack:=frame.OperandStack()
	d:=stack.PopInt()
	f:=float64(d)
	stack.PushDouble(f)
}

//转float
type I2F struct {
	base.NoOperandsInstruction
}

func (*I2F) Execute(frame *rtda.Frame) {
	stack:=frame.OperandStack()
	d:=stack.PopInt()
	f:=float32(d)
	stack.PushFloat(f)
}

//转long
type I2L struct {
	base.NoOperandsInstruction
}

func (*I2L) Execute(frame *rtda.Frame) {
	stack:=frame.OperandStack()
	d:=stack.PopInt()
	l:=int64(d)
	stack.PushLong(l)
}

//转char
type I2C struct {
	base.NoOperandsInstruction
}

func (*I2C) Execute(frame *rtda.Frame) {
	stack:=frame.OperandStack()
	i:=stack.PopInt()
	c:=int32(uint16(i))
	stack.PushInt(c)
}

//转short
type I2S struct {
	base.NoOperandsInstruction
}

func (*I2S) Execute(frame *rtda.Frame) {
	stack:=frame.OperandStack()
	i:=stack.PopInt()
	s:=int32(uint16(i))
	stack.PushInt(s)
}

//转byte
type I2B struct {
	base.NoOperandsInstruction
}

func (*I2B) Execute(frame *rtda.Frame) {
	stack:=frame.OperandStack()
	i:=stack.PopInt()
	b:=int32(int8(i))
	stack.PushInt(b)
}


