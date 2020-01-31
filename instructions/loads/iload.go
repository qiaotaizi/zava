package loads

import (
	"github.com/qiaotaizi/zava/instructions/base"
	"github.com/qiaotaizi/zava/rtda"
)

//加载指令
//从局部变量表获取变量
//然后推入操作数栈顶

type ILOAD struct {
	base.Index8Instruction
}

func (i *ILOAD) Execute(frame *rtda.Frame) {
	_iload(frame, i.Index) //索引来自操作数
}

type ILOAD_0 struct {
	base.NoOperandsInstruction
}

func (i *ILOAD_0) Execute(frame *rtda.Frame) {
	_iload(frame, 0) //ILOAD_n索引隐含在指令自身
}

type ILOAD_1 struct {
	base.NoOperandsInstruction
}

func (i *ILOAD_1) Execute(frame *rtda.Frame) {
	_iload(frame, 1)
}

type ILOAD_2 struct {
	base.NoOperandsInstruction
}

func (i *ILOAD_2) Execute(frame *rtda.Frame) {
	_iload(frame, 2)
}

type ILOAD_3 struct {
	base.NoOperandsInstruction
}

func (i *ILOAD_3) Execute(frame *rtda.Frame) {
	_iload(frame, 3)
}
func _iload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetInt(index)
	frame.OperandStack().PushInt(val)
}
