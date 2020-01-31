package stores

import (
	"github.com/qiaotaizi/zava/instructions/base"
	"github.com/qiaotaizi/zava/rtda"
)

//存储指令将变量从操作数栈中弹出
//存入局部变量表

type LSTORE struct {
	base.Index8Instruction
}

func (i *LSTORE) Execute(frame *rtda.Frame) {
	_lstore(frame, i.Index)
}

type LSTORE_0 struct {
	base.NoOperandsInstruction
}

func (i *LSTORE_0) Execute(frame *rtda.Frame) {
	_lstore(frame, 0)
}

type LSTORE_1 struct {
	base.NoOperandsInstruction
}

func (i *LSTORE_1) Execute(frame *rtda.Frame) {
	_lstore(frame, 1)
}

type LSTORE_2 struct {
	base.NoOperandsInstruction
}

func (i *LSTORE_2) Execute(frame *rtda.Frame) {
	_lstore(frame, 2)
}

type LSTORE_3 struct {
	base.NoOperandsInstruction
}

func (i *LSTORE_3) Execute(frame *rtda.Frame) {
	_lstore(frame, 3)
}
func _lstore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopLong()
	frame.LocalVars().SetLong(index, val)
}
