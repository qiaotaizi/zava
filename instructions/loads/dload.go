package loads

import (
	"github.com/qiaotaizi/zava/instructions/base"
	"github.com/qiaotaizi/zava/rtda"
)

//加载指令
//从局部变量表获取变量
//然后推入操作数栈顶

type DLOAD struct {
	base.Index8Instruction
}

func (i *DLOAD) Execute(frame *rtda.Frame) {
	_dload(frame, i.Index) //索引来自操作数
}

type DLOAD_0 struct {
	base.NoOperandsInstruction
}

func (i *DLOAD_0) Execute(frame *rtda.Frame) {
	_dload(frame, 0) //DLOAD_n索引隐含在指令自身
}

type DLOAD_1 struct {
	base.NoOperandsInstruction
}

func (i *DLOAD_1) Execute(frame *rtda.Frame) {
	_dload(frame, 1)
}

type DLOAD_2 struct {
	base.NoOperandsInstruction
}

func (i *DLOAD_2) Execute(frame *rtda.Frame) {
	_dload(frame, 2)
}

type DLOAD_3 struct {
	base.NoOperandsInstruction
}

func (i *DLOAD_3) Execute(frame *rtda.Frame) {
	_dload(frame, 3)
}
func _dload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetDouble(index)
	frame.OperandStack().PushDouble(val)
}

