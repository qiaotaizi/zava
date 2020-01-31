package stores

import (
	"github.com/qiaotaizi/zava/instructions/base"
	"github.com/qiaotaizi/zava/rtda"
)

//存储指令将变量从操作数栈中弹出
//存入局部变量表

type ISTORE struct {
	base.Index8Instruction
}

func (i *ISTORE) Execute(frame *rtda.Frame) {
	_istore(frame, i.Index)
}

type ISTORE_0 struct {
	base.NoOperandsInstruction
}

func (i *ISTORE_0) Execute(frame *rtda.Frame) {
	_istore(frame, 0)
}

type ISTORE_1 struct {
	base.NoOperandsInstruction
}

func (i *ISTORE_1) Execute(frame *rtda.Frame) {
	_istore(frame, 1)
}

type ISTORE_2 struct {
	base.NoOperandsInstruction
}

func (i *ISTORE_2) Execute(frame *rtda.Frame) {
	_istore(frame, 2)
}

type ISTORE_3 struct {
	base.NoOperandsInstruction
}

func (i *ISTORE_3) Execute(frame *rtda.Frame) {
	_istore(frame, 3)
}
func _istore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopInt()
	frame.LocalVars().SetInt(index, val)
}


