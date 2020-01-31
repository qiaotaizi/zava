package comparisons

import (
	"github.com/qiaotaizi/zava/instructions/base"
	"github.com/qiaotaizi/zava/rtda"
)

type DCMPG struct {
	base.NoOperandsInstruction
}

func (*DCMPG) Execute(frame *rtda.Frame) {
	_dcmp(frame,true)
}

type DCMPL struct {
	base.NoOperandsInstruction
}

func (*DCMPL) Execute(frame *rtda.Frame) {
	_dcmp(frame,false)
}

func _dcmp(frame *rtda.Frame,gFlag bool){
	s:=frame.OperandStack()
	v2:=s.PopDouble()
	v1:=s.PopDouble()
	if v1>v2{
		s.PushInt(1)
	}else if v1==v2 {
		s.PushInt(0)
	}else if v1<v2{
		s.PushInt(-1)
	}else if gFlag{
		s.PushInt(1)
	}else{
		s.PushInt(-1)
	}
}

