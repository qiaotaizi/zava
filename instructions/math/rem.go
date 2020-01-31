package math

import (
	"github.com/qiaotaizi/zava/instructions/base"
	"github.com/qiaotaizi/zava/rtda"
	"math"
)

//取模指令
//参数是操作数栈的前两个操作数
//弹出栈顶前两个操作数，取模
//将结果压入栈顶

//double取模
type DREM struct {
	base.NoOperandsInstruction
}

func (i *DREM) Execute(frame *rtda.Frame) {
	//double与float浮点数因为定义了无穷大值，所以不需要做零除判断
	stack:=frame.OperandStack()
	v2:=stack.PopDouble()
	v1:=stack.PopDouble()
	result:=math.Mod(v1,v2)
	stack.PushDouble(result)
}

//float取模
type FREM struct {
	base.NoOperandsInstruction
}

func (i *FREM) Execute(frame *rtda.Frame) {
	stack:=frame.OperandStack()
	v2:=stack.PopFloat()
	v1:=stack.PopFloat()
	result:=float32(math.Mod(float64(v1),float64(v2)))
	stack.PushFloat(result)
}

//int取模
type IREM struct {
	base.NoOperandsInstruction
}

func (i *IREM) Execute(frame *rtda.Frame) {
	stack:=frame.OperandStack()
	v2:=stack.PopInt()//被除数
	v1:=stack.PopInt()//除数
	if v2==0{
		panic("java.lang.ArithmeticException: / by zero")
	}
	result:=v1%v2
	stack.PushInt(result)
}

//long取模
type LREM struct {
	base.NoOperandsInstruction
}

func (i *LREM) Execute(frame *rtda.Frame) {
	stack:=frame.OperandStack()
	v2:=stack.PopLong()
	v1:=stack.PopLong()
	if v2==0{
		panic("java.lang.ArithmeticException: / by zero")
	}
	result:=v1%v2
	stack.PushLong(result)
}



