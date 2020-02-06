package references

import (
	"github.com/qiaotaizi/zava/instructions/base"
	"github.com/qiaotaizi/zava/rtda"
	"github.com/qiaotaizi/zava/rtda/heap"
)

//新建引用数组
//需要两个操作数
//第一个操作数从字节码读取一个uint16，是类型引用的常量池索引
//第二个操作数从栈顶弹出，表示长度
type ANEW_ARRAY struct {
	base.Index16Instruction
}

func (i *ANEW_ARRAY) Execute(frame *rtda.Frame) {
	cp:=frame.Method().Class().ConstantPool()
	classRef:=cp.GetConstant(i.Index).(*heap.ClassRef)
	componentClass:=classRef.ResolvedClass()

	stack:=frame.OperandStack()
	count:=stack.PopInt()
	if count<0{
		panic("java.lang.NegativeArraySizeError")
	}

	arrClass:=componentClass.ArrayClass()
	arr:=arrClass.NewArray(uint(count))
	stack.PushRef(arr)

}


