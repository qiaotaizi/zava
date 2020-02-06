package stores

import (
	"github.com/qiaotaizi/zava/instructions/base"
	"github.com/qiaotaizi/zava/rtda"
	"github.com/qiaotaizi/zava/rtda/heap"
)

//数组元素赋值
//需要三个操作数
//元素值，数组索引，数组引用
//均从操作数栈顶弹出

type AASTORE struct {
	base.NoOperandsInstruction
}

func (*AASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopRef()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	refs := arrRef.Refs()
	checkIndex(len(refs), index)
	refs[index] = val
}

type BASTORE struct {
	base.NoOperandsInstruction
}

func (*BASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	bytes := arrRef.Bytes()
	checkIndex(len(bytes), index)
	bytes[index] = int8(val)
}

type CASTORE struct {
	base.NoOperandsInstruction
}

func (*CASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	chars := arrRef.Chars()
	checkIndex(len(chars), index)
	chars[index] = uint16(val)
}

type DASTORE struct {
	base.NoOperandsInstruction
}

func (*DASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopDouble()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	doubles := arrRef.Doubles()
	checkIndex(len(doubles), index)
	doubles[index] = val
}

type FASTORE struct {
	base.NoOperandsInstruction
}

func (*FASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopFloat()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	floats := arrRef.Floats()
	checkIndex(len(floats), index)
	floats[index] = val
}

type LASTORE struct {
	base.NoOperandsInstruction
}

func (*LASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopLong()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	longs := arrRef.Longs()
	checkIndex(len(longs), index)
	longs[index] = val
}

type IASTORE struct {
	base.NoOperandsInstruction
}

func (*IASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	ints := arrRef.Ints()
	checkIndex(len(ints), index)
	ints[index] = val

}

type SASTORE struct {
	base.NoOperandsInstruction
}

func (*SASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	ints := arrRef.Shorts()
	checkIndex(len(ints), index)
	ints[index] = int16(val)
}

func checkIndex(arrLen int, index int32) {
	if index < 0 || index >= int32(arrLen) {
		panic("ArrayOutOfBoundsException")
	}
}

func checkNotNil(ref *heap.Object) {
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
}
