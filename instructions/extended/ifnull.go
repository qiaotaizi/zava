package extended

import (
	"github.com/qiaotaizi/zava/instructions/base"
	"github.com/qiaotaizi/zava/rtda"
)
//根据引用是否为null进行跳转
type IFNULL struct {
	base.BranchInstruction
}

func (i *IFNULL) Execute(frame *rtda.Frame) {
	ref:=frame.OperandStack().PopRef()
	if ref==nil{
		base.Branch(frame,i.Offset)
	}
}

type IFNONNULL struct {
	base.BranchInstruction
}

func (i *IFNONNULL) Execute(frame *rtda.Frame) {
	ref:=frame.OperandStack().PopRef()
	if ref!=nil{
		base.Branch(frame,i.Offset)
	}
}

