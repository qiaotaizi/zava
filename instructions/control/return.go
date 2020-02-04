package control

import (
	"github.com/qiaotaizi/zava/instructions/base"
	"github.com/qiaotaizi/zava/rtda"
)

//void返回
type RETURN struct {
	base.NoOperandsInstruction
}

func (*RETURN) Execute(frame *rtda.Frame) {
	//弹出帧即可
	frame.Thread().PopFrame()
}

//引用返回
type ARETURN struct {
	base.NoOperandsInstruction
}

func (*ARETURN) Execute(frame *rtda.Frame) {
	thread:=frame.Thread()
	currentFrame:=thread.PopFrame()
	invokerFrame:=thread.TopFrame()
	retVal:=currentFrame.OperandStack().PopRef()
	invokerFrame.OperandStack().PushRef(retVal)
}

//int等返回
type IRETURN struct {
	base.NoOperandsInstruction
}

func (*IRETURN) Execute(frame *rtda.Frame) {
	thread:=frame.Thread()
	currentFrame:=thread.PopFrame()
	invokerFrame:=thread.TopFrame()
	retVal:=currentFrame.OperandStack().PopInt()
	invokerFrame.OperandStack().PushInt(retVal)
}

//long
type LRETURN struct {
	base.NoOperandsInstruction
}

func (*LRETURN) Execute(frame *rtda.Frame) {
	thread:=frame.Thread()
	currentFrame:=thread.PopFrame()
	invokerFrame:=thread.TopFrame()
	retVal:=currentFrame.OperandStack().PopLong()
	invokerFrame.OperandStack().PushLong(retVal)
}

//double
type DRETURN struct {
	base.NoOperandsInstruction
}

func (*DRETURN) Execute(frame *rtda.Frame) {
	thread:=frame.Thread()
	currentFrame:=thread.PopFrame()
	invokerFrame:=thread.TopFrame()
	retVal:=currentFrame.OperandStack().PopDouble()
	invokerFrame.OperandStack().PushDouble(retVal)
}

//float
type FRETURN struct {
	base.NoOperandsInstruction
}

func (*FRETURN) Execute(frame *rtda.Frame) {
	thread:=frame.Thread()
	currentFrame:=thread.PopFrame()
	invokerFrame:=thread.TopFrame()
	retVal:=currentFrame.OperandStack().PopFloat()
	invokerFrame.OperandStack().PushFloat(retVal)
}
