package stores

import (
	"github.com/qiaotaizi/zava/instructions/base"
	"github.com/qiaotaizi/zava/rtda"
)

//存储指令将变量从操作数栈中弹出
//存入局部变量表

type FSTORE struct {
	base.Index8Instruction
}

func (i *FSTORE) Execute(frame *rtda.Frame) {
	_fstore(frame, i.Index)
}

type FSTORE_0 struct {
	base.NoOperandsInstruction
}

func (i *FSTORE_0) Execute(frame *rtda.Frame) {
	_fstore(frame, 0)
}

type FSTORE_1 struct {
	base.NoOperandsInstruction
}

func (i *FSTORE_1) Execute(frame *rtda.Frame) {
	_fstore(frame, 1)
}

type FSTORE_2 struct {
	base.NoOperandsInstruction
}

func (i *FSTORE_2) Execute(frame *rtda.Frame) {
	_fstore(frame, 2)
}

type FSTORE_3 struct {
	base.NoOperandsInstruction
}

func (i *FSTORE_3) Execute(frame *rtda.Frame) {
	_fstore(frame, 3)
}
func _fstore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopFloat()
	frame.LocalVars().SetFloat(index, val)
}


