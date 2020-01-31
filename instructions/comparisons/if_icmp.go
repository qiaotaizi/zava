package comparisons

import (
	"github.com/qiaotaizi/zava/instructions/base"
	"github.com/qiaotaizi/zava/rtda"
)

//if_icmp<cond>指令
//从栈顶弹出两个整数进行比较
//根据结果进行跳转

type IF_ICMPEQ struct {
	base.BranchInstruction
}

func (i *IF_ICMPEQ) Execute(frame *rtda.Frame) {
	s:=frame.OperandStack()
	v2:=s.PopInt()
	v1:=s.PopInt()
	if v1==v2{
		base.Branch(frame,i.Offset)
	}
}

type IF_ICMPNE struct {
	base.BranchInstruction
}
func (i *IF_ICMPNE) Execute(frame *rtda.Frame) {
	s:=frame.OperandStack()
	v2:=s.PopInt()
	v1:=s.PopInt()
	if v1!=v2{
		base.Branch(frame,i.Offset)
	}
}
type IF_ICMPLT struct {
	base.BranchInstruction
}
func (i *IF_ICMPLT) Execute(frame *rtda.Frame) {
	s:=frame.OperandStack()
	v2:=s.PopInt()
	v1:=s.PopInt()
	if v1<v2{
		base.Branch(frame,i.Offset)
	}
}
type IF_ICMPLE struct {
	base.BranchInstruction
}
func (i *IF_ICMPLE) Execute(frame *rtda.Frame) {
	s:=frame.OperandStack()
	v2:=s.PopInt()
	v1:=s.PopInt()
	if v1<=v2{
		base.Branch(frame,i.Offset)
	}
}
type IF_ICMPGT struct {
	base.BranchInstruction
}
func (i *IF_ICMPGT) Execute(frame *rtda.Frame) {
	s:=frame.OperandStack()
	v2:=s.PopInt()
	v1:=s.PopInt()
	if v1>v2{
		base.Branch(frame,i.Offset)
	}
}
type IF_ICMPGE struct {
	base.BranchInstruction
}
func (i *IF_ICMPGE) Execute(frame *rtda.Frame) {
	s:=frame.OperandStack()
	v2:=s.PopInt()
	v1:=s.PopInt()
	if v1>=v2{
		base.Branch(frame,i.Offset)
	}
}
