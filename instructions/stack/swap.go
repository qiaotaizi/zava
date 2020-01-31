package stack

import (
	"github.com/qiaotaizi/zava/instructions/base"
	"github.com/qiaotaizi/zava/rtda"
)

type SWAP struct {
	base.NoOperandsInstruction
}

func (i *SWAP) Execute(frame *rtda.Frame) {
	stack:=frame.OperandStack()
	s1:=stack.PopSlot()
	s2:=stack.PopSlot()
	stack.PushSlot(s1)
	stack.PushSlot(s2)
}

