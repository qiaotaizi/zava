package base

import "github.com/qiaotaizi/zava/rtda"

type Instruction interface {
	FetchOperands(reader *BytecodeReader) //获取操作数
	Execute(frame *rtda.Frame)            //执行指令
}

//按照类型定义一些抽象指令
//仅实现指令接口的FetchOperands

//无操作数指令
type NoOperandsInstruction struct{}

func (n *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {
	//什么也不做
}

//跳转指令
type BranchInstruction struct {
	Offset int //跳转偏移量
}

func (b *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	b.Offset = int(reader.ReadInt16())
}

//局部变量存取指令
type Index8Instruction struct {
	Index uint //局部变量索引
}

func (i *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	i.Index = uint(reader.ReadUint8())
}

//运行时常量池读取指令
type Index16Instruction struct {
	Index uint
}

func (i *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	i.Index = uint(reader.ReadUInt16())
}
