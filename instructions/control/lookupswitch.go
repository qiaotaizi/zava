package control

import (
	"github.com/qiaotaizi/zava/instructions/base"
	"github.com/qiaotaizi/zava/rtda"
)

type LOOKUP_SWITCH struct {
	defaultOffset int32
	npairs int32
	matchOffsets []int32//长度为npairs的两倍，前一个值是从操作数栈顶读取的key
	//也就是switch关键字，该值后一索引上的元素是跳转偏移量
}

func (i *LOOKUP_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	reader.SkipPadding()
	i.defaultOffset=reader.ReadInt32()
	i.npairs=reader.ReadInt32()
	i.matchOffsets=reader.ReadInt32s(i.npairs*2)
}

func (i *LOOKUP_SWITCH) Execute(frame *rtda.Frame) {
	key:=frame.OperandStack().PopInt()
	for j:=int32(0);j<i.npairs*2;j+=2{
		if i.matchOffsets[j]==key{
			offset:=i.matchOffsets[j+1]
			base.Branch(frame,int(offset))
			return
		}
	}
	base.Branch(frame,int(i.defaultOffset))
}



