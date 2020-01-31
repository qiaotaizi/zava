package constants

import (
	"github.com/qiaotaizi/zava/instructions/base"
	"github.com/qiaotaizi/zava/rtda"
)

//空操作

type NOP struct {
	base.NoOperandsInstruction
}

func (N NOP) Execute(frame *rtda.Frame) {
	//什么也不做
}
