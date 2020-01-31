package math

import (
	"github.com/qiaotaizi/zava/instructions/base"
	"github.com/qiaotaizi/zava/rtda"
)

//读取局部变量表中的变量，加上值，放回局部变量表

type IINC struct {
	Index uint
	Const int32
}

func (i *IINC) FetchOperands(reader *base.BytecodeReader) {
	i.Index=uint(reader.ReadUint8())
	i.Const=int32(reader.ReadInt8())
}

func (i *IINC) Execute(frame *rtda.Frame) {
	localVars:=frame.LocalVars()
	val:=localVars.GetInt(i.Index)
	val+=i.Const
	localVars.SetInt(i.Index,val)
}



