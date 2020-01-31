package loads

import (
	"github.com/qiaotaizi/zava/instructions/base"
	"github.com/qiaotaizi/zava/rtda"
)

//加载指令
//从局部变量表获取变量
//然后推入操作数栈顶

type ALOAD struct {
	base.Index8Instruction
}

func (i *ALOAD) Execute(frame *rtda.Frame) {
	_aload(frame, i.Index) //索引来自操作数
}

type ALOAD_0 struct {
	base.NoOperandsInstruction
}

func (i *ALOAD_0) Execute(frame *rtda.Frame) {
	_aload(frame, 0) //ALOAD_n索引隐含在指令自身
}

type ALOAD_1 struct {
	base.NoOperandsInstruction
}

func (i *ALOAD_1) Execute(frame *rtda.Frame) {
	_aload(frame, 1)
}

type ALOAD_2 struct {
	base.NoOperandsInstruction
}

func (i *ALOAD_2) Execute(frame *rtda.Frame) {
	_aload(frame, 2)
}

type ALOAD_3 struct {
	base.NoOperandsInstruction
}

func (i *ALOAD_3) Execute(frame *rtda.Frame) {
	_aload(frame, 3)
}
func _aload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetRef(index)
	frame.OperandStack().PushRef(val)
}

