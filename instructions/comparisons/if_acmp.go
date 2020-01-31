package comparisons

import (
	"github.com/qiaotaizi/zava/instructions/base"
	"github.com/qiaotaizi/zava/rtda"
)

//引用对象比较

type IF_ACMPEQ struct {
	base.BranchInstruction
}

func (i *IF_ACMPEQ) Execute(frame *rtda.Frame) {
	s := frame.OperandStack()
	v2 := s.PopRef()
	v1 := s.PopRef()
	if v1 == v2 {
		base.Branch(frame, i.Offset)
	}
}

type IF_ACMPNE struct {
	base.BranchInstruction
}

func (i *IF_ACMPNE) Execute(frame *rtda.Frame) {
	s := frame.OperandStack()
	v2 := s.PopRef()
	v1 := s.PopRef()
	if v1 != v2 {
		base.Branch(frame, i.Offset)
	}
}
