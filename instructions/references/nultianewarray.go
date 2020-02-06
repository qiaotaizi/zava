package references

import (
	"github.com/qiaotaizi/zava/instructions/base"
	"github.com/qiaotaizi/zava/rtda"
	"github.com/qiaotaizi/zava/rtda/heap"
)

//创建多维数组
//类符号引用的常量池索引和维度数均从字节码读取，紧跟在指令之后
//还需要从栈顶弹出n个整数，分别表示每个维度上数组的长度
type MULTI_ANEW_ARRAY struct {
	index      uint16
	dimensions uint8
}

func (i *MULTI_ANEW_ARRAY) FetchOperands(reader *base.BytecodeReader) {
	i.index=reader.ReadUInt16()
	i.dimensions=reader.ReadUint8()
}

func (i *MULTI_ANEW_ARRAY) Execute(frame *rtda.Frame) {
	cp:=frame.Method().Class().ConstantPool()
	classRef:=cp.GetConstant(uint(i.index)).(*heap.ClassRef)
	arrClass:=classRef.ResolvedClass()

	stack:=frame.OperandStack()
	counts:=popAndCheckCounts(stack,int(i.dimensions))

	arr:=newMultiDimensionArray(counts,arrClass)
	stack.PushRef(arr)

}

func newMultiDimensionArray(counts []int32, class *heap.Class) *heap.Object {
	count:=uint(counts[0])
	arr:=class.NewArray(count)

	if len(counts)>1{//还存在更高维度
		refs:=arr.Refs()
		for i:=range refs{
			refs[i]=newMultiDimensionArray(counts[1:],class.ComponentClass())
		}
	}
	return arr//没有更高维度，直接返回这个一维数组
}

func popAndCheckCounts(stack *rtda.OperandStack, dimensions int) []int32 {
	counts:=make([]int32,dimensions)
	for i:=dimensions-1;i>=0;i--{
		counts[i]=stack.PopInt()
		if counts[i]<0{
			panic("java.lang.NegativeArraySizeException")
		}
	}
	return counts
}
