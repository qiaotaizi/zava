package loads

import (
	"github.com/qiaotaizi/zava/instructions/base"
	"github.com/qiaotaizi/zava/rtda"
)

//加载指令
//从局部变量表获取变量
//然后推入操作数栈顶

type FLOAD struct {
	base.Index8Instruction
}

func (i *FLOAD) Execute(frame *rtda.Frame) {
	_fload(frame, i.Index) //索引来自操作数
}

type FLOAD_0 struct {
	base.NoOperandsInstruction
}

func (i *FLOAD_0) Execute(frame *rtda.Frame) {
	_fload(frame, 0) //FLOAD_n索引隐含在指令自身
}

type FLOAD_1 struct {
	base.NoOperandsInstruction
}

func (i *FLOAD_1) Execute(frame *rtda.Frame) {
	_fload(frame, 1)
}

type FLOAD_2 struct {
	base.NoOperandsInstruction
}

func (i *FLOAD_2) Execute(frame *rtda.Frame) {
	_fload(frame, 2)
}

type FLOAD_3 struct {
	base.NoOperandsInstruction
}

func (i *FLOAD_3) Execute(frame *rtda.Frame) {
	_fload(frame, 3)
}
func _fload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetFloat(index)
	frame.OperandStack().PushFloat(val)
}

