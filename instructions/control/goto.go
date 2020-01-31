package control

import (
	"github.com/qiaotaizi/zava/instructions/base"
	"github.com/qiaotaizi/zava/rtda"
)

//无条件跳转
type GOTO struct {
	base.BranchInstruction
}

func (i *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame,i.Offset)
}

