package references

import (
	"fmt"
	"github.com/qiaotaizi/zava/instructions/base"
	"github.com/qiaotaizi/zava/rtda"
	"github.com/qiaotaizi/zava/rtda/heap"
)

type INVOKE_VIRTUAL struct {
	base.Index16Instruction
}

func (i *INVOKE_VIRTUAL) Execute(frame *rtda.Frame) {
	cp:=frame.Method().Class().ConstantPool()
	methodRef:=cp.GetConstant(i.Index).(*heap.MethodRef)
	if methodRef.Name()=="println"{//临时调用java的控制台打印
		stack:=frame.OperandStack()
		switch methodRef.Descriptor() {
		case "(Z)V":
			fmt.Printf("%v\n",stack.PopInt()!=0)
		case "(C)V":
			fmt.Printf("%v\n",stack.PopInt())
		case "(B)V":
			fmt.Printf("%v\n",stack.PopInt())
		case "(S)V":
			fmt.Printf("%v\n",stack.PopInt())
		case "(I)V":
			fmt.Printf("%v\n",stack.PopInt())
		case "(J)V":
			fmt.Printf("%v\n",stack.PopLong())
		case "(F)V":
			fmt.Printf("%v\n",stack.PopFloat())
		case "(D)V":
			fmt.Printf("%v\n",stack.PopDouble())
		default:
			panic("println: "+methodRef.Descriptor())
		}
		stack.PopRef()
	}
}

