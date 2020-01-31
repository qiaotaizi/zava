package constants

import (
	"github.com/qiaotaizi/zava/instructions/base"
	"github.com/qiaotaizi/zava/rtda"
)

//将null引用推入栈顶
type ACONST_NULL struct {
	base.NoOperandsInstruction
}

func (i *ACONST_NULL) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushRef(nil)
}

//将双精度浮点数0退出栈顶
type DCONST_0 struct {
	base.NoOperandsInstruction
}

func (i *DCONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushDouble(0)
}

//类比上面的注释
type DCONST_1 struct {
	base.NoOperandsInstruction
}

func (i *DCONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushDouble(1)
}

type FCONST_0 struct {
	base.NoOperandsInstruction
}

func (i *FCONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(0)
}

type FCONST_1 struct {
	base.NoOperandsInstruction
}

func (i *FCONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(1)
}

type FCONST_2 struct {
	base.NoOperandsInstruction
}

func (i *FCONST_2) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(2)
}

//整形-1推入栈顶
type ICONST_M1 struct {
	base.NoOperandsInstruction
}

func (i *ICONST_M1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(-1)
}

type ICONST_0 struct {
	base.NoOperandsInstruction
}

func (i *ICONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(0)
}

type ICONST_1 struct {
	base.NoOperandsInstruction
}

func (i *ICONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(1)
}

type ICONST_2 struct {
	base.NoOperandsInstruction
}

func (i *ICONST_2) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(2)
}

type ICONST_3 struct {
	base.NoOperandsInstruction
}

func (i *ICONST_3) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(3)
}

type ICONST_4 struct {
	base.NoOperandsInstruction
}

func (i *ICONST_4) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(4)
}

type ICONST_5 struct {
	base.NoOperandsInstruction
}

func (i *ICONST_5) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(5)
}

type LCONST_0 struct {
	base.NoOperandsInstruction
}

func (i *LCONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushLong(0)
}

type LCONST_1 struct {
	base.NoOperandsInstruction
}

func (i *LCONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushLong(1)
}
