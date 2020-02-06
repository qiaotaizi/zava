package references

import (
	"github.com/qiaotaizi/zava/instructions/base"
	"github.com/qiaotaizi/zava/rtda"
)

//获取数组长度
type ARRAY_LENGTH struct {
	base.NoOperandsInstruction
}

//从操作数栈弹出数组引用，将数组长度推入栈顶
func (*ARRAY_LENGTH) Execute(frame *rtda.Frame) {
	stack:=frame.OperandStack()
	arrRef:=stack.PopRef()
	if arrRef==nil{
		panic("java.lang.NullPointerException")
	}

	arrLength:=arrRef.ArrayLength()
	stack.PushInt(arrLength)
}



