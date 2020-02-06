package references

import (
	"github.com/qiaotaizi/zava/instructions/base"
	"github.com/qiaotaizi/zava/rtda"
	"github.com/qiaotaizi/zava/rtda/heap"
)

//创建基本类型数组

const (
	AT_BOOLEAN=4
	AT_CHAR=5
	AT_FLOAT=6
	AT_DOUBLE=7
	AT_BYTE=8
	AT_SHORT=9
	AT_INT=10
	AT_LONG=11
)


//newarray指令需要两个操作数
//第一个操作数紧跟在字节码指令后，表示元素类型
//第二个操作数从操作数栈弹出，表示数组长度
type NEW_ARRAY struct {
	atype uint8
}

func (i *NEW_ARRAY) FetchOperands(reader *base.BytecodeReader) {
	i.atype=reader.ReadUint8()
}

func (i *NEW_ARRAY) Execute(frame *rtda.Frame) {
	stack:=frame.OperandStack()
	count:=stack.PopInt()
	if count<0{
		panic("java.lang.NegativeArraySizeError")
	}

	classLoader:=frame.Method().Class().Loader()

	arrClass:=getPrimitiveArrayClass(classLoader,i.atype)
	arr:=arrClass.NewArray(uint(count))
	stack.PushRef(arr)
}

func getPrimitiveArrayClass(loader *heap.ClassLoader, atype uint8) *heap.Class {
	switch atype {
	case AT_BOOLEAN:
		return loader.LoadClass("[Z")
	case AT_BYTE:
		return loader.LoadClass("[B")
	case AT_CHAR:
		return loader.LoadClass("[C")
	case AT_SHORT:
		return loader.LoadClass("[S")
	case AT_INT:
		return loader.LoadClass("[I")
	case AT_LONG:
		return loader.LoadClass("[J")
	case AT_FLOAT:
		return loader.LoadClass("[F")
	case AT_DOUBLE:
		return loader.LoadClass("[D")
	default:
		panic("Invalid atype!")
	}
}



