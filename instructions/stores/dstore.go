package stores

import (
	"github.com/qiaotaizi/zava/instructions/base"
	"github.com/qiaotaizi/zava/rtda"
)

//存储指令将变量从操作数栈中弹出
//存入局部变量表

type DSTORE struct {
	base.Index8Instruction
}

func (i *DSTORE) Execute(frame *rtda.Frame) {
	_dstore(frame, i.Index)
}

type DSTORE_0 struct {
	base.NoOperandsInstruction
}

func (i *DSTORE_0) Execute(frame *rtda.Frame) {
	_dstore(frame, 0)
}

type DSTORE_1 struct {
	base.NoOperandsInstruction
}

func (i *DSTORE_1) Execute(frame *rtda.Frame) {
	_dstore(frame, 1)
}

type DSTORE_2 struct {
	base.NoOperandsInstruction
}

func (i *DSTORE_2) Execute(frame *rtda.Frame) {
	_dstore(frame, 2)
}

type DSTORE_3 struct {
	base.NoOperandsInstruction
}

func (i *DSTORE_3) Execute(frame *rtda.Frame) {
	_dstore(frame, 3)
}
func _dstore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopDouble()
	frame.LocalVars().SetDouble(index, val)
}

