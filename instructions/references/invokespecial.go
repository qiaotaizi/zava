package references

import (
	"github.com/qiaotaizi/zava/instructions/base"
	"github.com/qiaotaizi/zava/rtda"
)

type INVOKE_SPECIAL struct {
	base.Index16Instruction
}

func (i *INVOKE_SPECIAL) Execute(frame *rtda.Frame) {
	// TODO 一个临时实现：new之后调用构造函数初始化
	frame.OperandStack().PopRef()
}

