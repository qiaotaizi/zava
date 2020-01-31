package comparisons

import (
	"github.com/qiaotaizi/zava/instructions/base"
	"github.com/qiaotaizi/zava/rtda"
)

//浮点数比较
//浮点数计算可能产生NaN
//当两个float变量中至少有一个是NaN时，用fcmpg指令比较的结果是1，而用fcmpl指令比较的结果是-1。

type FCMPG struct {
	base.NoOperandsInstruction
}

func (*FCMPG) Execute(frame *rtda.Frame) {
	_fcmp(frame,true)
}

type FCMPL struct {
	base.NoOperandsInstruction
}

func (*FCMPL) Execute(frame *rtda.Frame) {
	_fcmp(frame,false)
}

func _fcmp(frame *rtda.Frame,gFlag bool){
	s:=frame.OperandStack()
	v2:=s.PopFloat()
	v1:=s.PopFloat()
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

