package math

import (
	"github.com/qiaotaizi/zava/instructions/base"
	"github.com/qiaotaizi/zava/rtda"
)

const (
	limitInt=0x1f
	limitLong=0x3f
)

//位移指令
//其中，右移分为算术右移和逻辑右移
//算术右移使用最高位补位
//逻辑右移用0补位

//整形左移
type ISHL struct {
	base.NoOperandsInstruction
}

func (i *ISHL) Execute(frame *rtda.Frame) {
	stack:=frame.OperandStack()
	v2:=stack.PopInt()//位移次数
	v1:=stack.PopInt()//被位移数
	s:=uint32(v2)&limitInt//0x1f二进制表示为11111，十进制表示为31，限制v2造成的位移长度，最多位移31位
	result:=v1<<s
	stack.PushInt(result)
}

//整形算术右移
type ISHR struct {
	base.NoOperandsInstruction
}

func (i *ISHR) Execute(frame *rtda.Frame) {
	stack:=frame.OperandStack()
	v2:=stack.PopInt()
	v1:=stack.PopInt()
	s:=uint32(v2)&limitInt
	result:=v1>>s
	stack.PushInt(result)
}

//整形逻辑右移
type IUSHR struct {
	base.NoOperandsInstruction
}

func (i *IUSHR) Execute(frame *rtda.Frame) {
	stack:=frame.OperandStack()
	v2:=stack.PopInt()
	v1:=stack.PopInt()
	s:=uint32(v2)&limitInt
	result:=int32(uint32(v1)>>s)//go的>>符号对于无符号整数表示逻辑右移，对于有符号整数标识算术右移
	stack.PushInt(result)
}

type LSHL struct {
	base.NoOperandsInstruction
}

func (i *LSHL) Execute(frame *rtda.Frame) {
	stack:=frame.OperandStack()
	v2:=stack.PopInt()
	v1:=stack.PopLong()
	s:=uint32(v2)&limitLong
	result:=v1<<s
	stack.PushLong(result)
}

type LSHR struct {
	base.NoOperandsInstruction
}

func (i *LSHR) Execute(frame *rtda.Frame) {
	stack:=frame.OperandStack()
	v2:=stack.PopInt()//位移次数
	v1:=stack.PopLong()
	s:=uint32(v2)&limitLong//最多左移63位
	result:=v1>>s
	stack.PushLong(result)

}

type LUSHR struct {
	base.NoOperandsInstruction
}

func (i *LUSHR) Execute(frame *rtda.Frame) {
	stack:=frame.OperandStack()
	v2:=stack.PopInt()
	v1:=stack.PopLong()
	s:=uint32(v2)&limitLong
	result:=int64(uint64(v1)>>s)
	stack.PushLong(result)
}


