package comparisons

import (
	"github.com/qiaotaizi/zava/instructions/base"
	"github.com/qiaotaizi/zava/rtda"
)

//长整形比较

type LCMP struct {
	base.NoOperandsInstruction
}

func (*LCMP) Execute(frame *rtda.Frame) {
	s:=frame.OperandStack()
	v2:=s.PopLong()
	v1:=s.PopLong()
	if v1>v2{
		s.PushInt(1)
	}else if v1==v2 {
		s.PushInt(0)
	}else{
		s.PushInt(-1)
	}
}

