package extended

import (
	"github.com/qiaotaizi/zava/instructions/base"
	"github.com/qiaotaizi/zava/rtda"
)

//与GOTO的区别在于索引从2字节变成4字节

type GOTO_W struct {
	offset int
}

func (i *GOTO_W) FetchOperands(reader *base.BytecodeReader) {
	i.offset=int(reader.ReadInt32())
}

func (i *GOTO_W) Execute(frame *rtda.Frame) {
	base.Branch(frame,i.offset)
}

