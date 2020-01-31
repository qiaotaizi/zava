package control

import (
	"github.com/qiaotaizi/zava/instructions/base"
	"github.com/qiaotaizi/zava/rtda"
)

//switch-case
//如果case值可以编码称索引表(case值是连续的非负整数？)
//实现为TABLE_SWITCH指令
//否则实现为LOOKUPSWOTCH指令
type TABLE_SWITCH struct {
	defaultOffset int32 //默认跳转偏移量
	low           int32
	high          int32
	jumpOffsets   []int32 //跳转偏移量数组
}

func (i *TABLE_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	reader.SkipPadding()
	i.defaultOffset = reader.ReadInt32()
	i.low = reader.ReadInt32()
	i.high = reader.ReadInt32()
	jumpOffsetsCount := i.high - i.low + 1
	i.jumpOffsets = reader.ReadInt32s(jumpOffsetsCount)
}

func (i *TABLE_SWITCH) Execute(frame *rtda.Frame) {
	index := frame.OperandStack().PopInt()
	var offset int
	if index >= i.low && index <= i.high {
		offset = int(i.jumpOffsets[index-i.low])
	} else {
		offset = int(i.defaultOffset)
	}
	base.Branch(frame, offset)
}
