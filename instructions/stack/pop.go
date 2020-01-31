package stack

import (
	"github.com/qiaotaizi/zava/instructions/base"
	"github.com/qiaotaizi/zava/rtda"
)

//操作数栈指令
//槽弹出

//弹出一个槽
type POP struct {
	base.NoOperandsInstruction
}

func (i *POP) Execute(frame *rtda.Frame) {
	frame.OperandStack().PopSlot()
}

//弹出两个槽
//对应long、double等类型
type POP2 struct {
	base.NoOperandsInstruction
}

func (i *POP2) Execute(frame *rtda.Frame) {
	frame.OperandStack().PopSlot()
	frame.OperandStack().PopSlot()
}


