package constants

import (
	"github.com/qiaotaizi/zava/instructions/base"
	"github.com/qiaotaizi/zava/rtda"
)

//从操作数中读取一个byte整数，推入栈顶
type BIPUSH struct {
	val int8
}

func (i *BIPUSH) FetchOperands(reader *base.BytecodeReader) {
	i.val = reader.ReadInt8()
}

func (i *BIPUSH) Execute(frame *rtda.Frame) {
	val := int32(i.val)
	frame.OperandStack().PushInt(val)
}

//从操作数中读取一个short整数，推入栈顶
type SIPUSH struct {
	val int16
}

func (i *SIPUSH) FetchOperands(reader *base.BytecodeReader) {
	i.val = reader.ReadInt16()
}

func (i *SIPUSH) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(int32(i.val))
}
