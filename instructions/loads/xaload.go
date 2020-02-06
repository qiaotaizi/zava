package loads

import (
	"github.com/qiaotaizi/zava/instructions/base"
	"github.com/qiaotaizi/zava/rtda"
	"github.com/qiaotaizi/zava/rtda/heap"
)

//数组元素读取指令
//操作数来自栈顶前两个Slot

type AALOAD struct {
	base.NoOperandsInstruction
}

//ref
func (*AALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()  //元素在数组中的索引
	arrRef := stack.PopRef() //数组引用

	checkNotNil(arrRef)
	refs := arrRef.Refs()
	checkIndex(len(refs), index)
	stack.PushRef(refs[index])
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

type BALOAD struct {
	base.NoOperandsInstruction
}

//bytes,bool
func (*BALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()  //元素在数组中的索引
	arrRef := stack.PopRef() //数组引用

	checkNotNil(arrRef)
	bytes := arrRef. Bytes()
	checkIndex(len(bytes), index)
	stack.PushInt(int32(bytes[index]))
}

//chars
type CALOAD struct {
	base.NoOperandsInstruction
}

func (*CALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()  //元素在数组中的索引
	arrRef := stack.PopRef() //数组引用

	checkNotNil(arrRef)
	chars := arrRef.Chars()
	checkIndex(len(chars), index)
	stack.PushInt(int32(chars[index]))
}

type DALOAD struct {
	base.NoOperandsInstruction
}

func (*DALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()  //元素在数组中的索引
	arrRef := stack.PopRef() //数组引用

	checkNotNil(arrRef)
	doubles := arrRef.Doubles()
	checkIndex(len(doubles), index)
	stack.PushDouble(doubles[index])
}

type FALOAD struct {
	base.NoOperandsInstruction
}

func (*FALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()  //元素在数组中的索引
	arrRef := stack.PopRef() //数组引用

	checkNotNil(arrRef)
	floats := arrRef.Floats()
	checkIndex(len(floats), index)
	stack.PushFloat(floats[index])
}

type IALOAD struct {
	base.NoOperandsInstruction
}

func (*IALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()  //元素在数组中的索引
	arrRef := stack.PopRef() //数组引用

	checkNotNil(arrRef)
	ints := arrRef.Ints()
	checkIndex(len(ints), index)
	stack.PushInt(ints[index])
}

type LALOAD struct {
	base.NoOperandsInstruction
}

func (*LALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()  //元素在数组中的索引
	arrRef := stack.PopRef() //数组引用

	checkNotNil(arrRef)
	longs := arrRef.Longs()
	checkIndex(len(longs), index)
	stack.PushLong(longs[index])
}

type SALOAD struct {
	base.NoOperandsInstruction
}

func (*SALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()  //元素在数组中的索引
	arrRef := stack.PopRef() //数组引用

	checkNotNil(arrRef)
	refs := arrRef.Shorts()
	checkIndex(len(refs), index)
	stack.PushInt(int32(refs[index]))
}

