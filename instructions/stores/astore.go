package stores

import (
	"github.com/qiaotaizi/zava/instructions/base"
	"github.com/qiaotaizi/zava/rtda"
)

//存储指令将变量从操作数栈中弹出
//存入局部变量表

type ASTORE struct {
	base.Index8Instruction
}

func (i *ASTORE) Execute(frame *rtda.Frame) {
	_astore(frame, i.Index)
}

type ASTORE_0 struct {
	base.NoOperandsInstruction
}

func (i *ASTORE_0) Execute(frame *rtda.Frame) {
	_astore(frame, 0)
}

type ASTORE_1 struct {
	base.NoOperandsInstruction
}

func (i *ASTORE_1) Execute(frame *rtda.Frame) {
	_astore(frame, 1)
}

type ASTORE_2 struct {
	base.NoOperandsInstruction
}

func (i *ASTORE_2) Execute(frame *rtda.Frame) {
	_astore(frame, 2)
}

type ASTORE_3 struct {
	base.NoOperandsInstruction
}

func (i *ASTORE_3) Execute(frame *rtda.Frame) {
	_astore(frame, 3)
}
func _astore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopRef()
	frame.LocalVars().SetRef(index, val)
}

