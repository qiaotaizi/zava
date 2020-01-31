package loads

import (
	"github.com/qiaotaizi/zava/instructions/base"
	"github.com/qiaotaizi/zava/rtda"
)

//加载指令
//从局部变量表获取变量
//然后推入操作数栈顶

type LLOAD struct {
	base.Index8Instruction
}

func (i *LLOAD) Execute(frame *rtda.Frame) {
	_lload(frame, i.Index) //索引来自操作数
}

type LLOAD_0 struct {
	base.NoOperandsInstruction
}

func (i *LLOAD_0) Execute(frame *rtda.Frame) {
	_lload(frame, 0) //LLOAD_n索引隐含在指令自身
}

type LLOAD_1 struct {
	base.NoOperandsInstruction
}

func (i *LLOAD_1) Execute(frame *rtda.Frame) {
	_lload(frame, 1)
}

type LLOAD_2 struct {
	base.NoOperandsInstruction
}

func (i *LLOAD_2) Execute(frame *rtda.Frame) {
	_lload(frame, 2)
}

type LLOAD_3 struct {
	base.NoOperandsInstruction
}

func (i *LLOAD_3) Execute(frame *rtda.Frame) {
	_lload(frame, 3)
}
func _lload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetLong(index)
	frame.OperandStack().PushLong(val)
}
