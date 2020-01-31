package stack

import (
	"github.com/qiaotaizi/zava/instructions/base"
	"github.com/qiaotaizi/zava/rtda"
)

//复制操作数栈顶变量并加入栈顶

type DUP struct {
	base.NoOperandsInstruction
}

func (i *DUP) Execute(frame *rtda.Frame) {
	stack:=frame.OperandStack()
	slot:=stack.PopSlot()
	stack.PushSlot(slot)
	stack.PushSlot(slot)
}

//复制栈顶变量
//并插入至上一个值之前
//bottom -> top
//[...][c][b][a]
//			__/
//			|
//			V
//[...][c][a][b][a]
type DUP_X1 struct {
	base.NoOperandsInstruction
}

func (i *DUP_X1) Execute(frame *rtda.Frame) {
	stack:=frame.OperandStack()
	slot1:=stack.PopSlot()
	slot2:=stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)

}

//复制栈顶变量
//并插入至上上一个值之前
//示意图参考DUP_X1
type DUP_X2 struct {
	base.NoOperandsInstruction
}

func (i *DUP_X2) Execute(frame *rtda.Frame) {
	stack:=frame.OperandStack()
	slot1:=stack.PopSlot()
	slot2:=stack.PopSlot()
	slot3:=stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)

}

//DUP针对long、double类型的操作
//类比DUP指令
type DUP2 struct {
	base.NoOperandsInstruction
}

func (i *DUP2) Execute(frame *rtda.Frame) {
	stack:=frame.OperandStack()
	s1:=stack.PopSlot()
	s2:=stack.PopSlot()
	stack.PushSlot(s2)
	stack.PushSlot(s1)
	stack.PushSlot(s2)
	stack.PushSlot(s1)
}

type DUP2_X1 struct {
	base.NoOperandsInstruction
}

func (i *DUP2_X1) Execute(frame *rtda.Frame) {
	stack:=frame.OperandStack()
	s1:=stack.PopSlot()
	s2:=stack.PopSlot()
	s3:=stack.PopSlot()
	stack.PushSlot(s2)
	stack.PushSlot(s1)
	stack.PushSlot(s3)
	stack.PushSlot(s2)
	stack.PushSlot(s1)
}

type DUP2_X2 struct {
	base.NoOperandsInstruction
}

func (i *DUP2_X2) Execute(frame *rtda.Frame) {
	stack:=frame.OperandStack()
	s1:=stack.PopSlot()
	s2:=stack.PopSlot()
	s3:=stack.PopSlot()
	s4:=stack.PopSlot()
	stack.PushSlot(s2)
	stack.PushSlot(s1)
	stack.PushSlot(s4)
	stack.PushSlot(s3)
	stack.PushSlot(s2)
	stack.PushSlot(s1)
}
