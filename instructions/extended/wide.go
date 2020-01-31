package extended

import (
	"github.com/qiaotaizi/zava/instructions/base"
	"github.com/qiaotaizi/zava/instructions/loads"
	"github.com/qiaotaizi/zava/instructions/math"
	"github.com/qiaotaizi/zava/instructions/stores"
	"github.com/qiaotaizi/zava/rtda"
)

//扩展指令用于应对方法的局部变量表长度大于1字节存储范围，即大于256时的情况
//仅仅针对不同的情况进行索引读取方式的变更
//Execute方法不发生变化

type WIDE struct {
	modifiedInstruction base.Instruction
}

func (w *WIDE) Execute(frame *rtda.Frame) {
	w.modifiedInstruction.Execute(frame)
}

func (w *WIDE) FetchOperands(reader *base.BytecodeReader) {
	//解码具体的指令
	opcode := reader.ReadUint8()
	switch opcode {
	case 0x15: //iload
		inst := &loads.ILOAD{}
		inst.Index = uint(reader.ReadUInt16())//局部变量表索引超过了1字节表示范围，扩展至2字节
		w.modifiedInstruction = inst
	case 0x16: //lload
		inst:=&loads.LLOAD{}
		inst.Index=uint(reader.ReadUInt16())
		w.modifiedInstruction=inst
	case 0x17: //fload
		inst:=&loads.FLOAD{}
		inst.Index=uint(reader.ReadUInt16())
		w.modifiedInstruction=inst
	case 0x18: //dload
		inst:=&loads.DLOAD{}
		inst.Index=uint(reader.ReadUInt16())
		w.modifiedInstruction=inst
	case 0x19: //aload
		inst:=&loads.ALOAD{}
		inst.Index=uint(reader.ReadUInt16())
		w.modifiedInstruction=inst
	case 0x36: //istore
		inst:=&stores.ISTORE{}
		inst.Index=uint(reader.ReadUInt16())
		w.modifiedInstruction=inst
	case 0x37: //lstore
		inst:=&stores.LSTORE{}
		inst.Index=uint(reader.ReadUInt16())
		w.modifiedInstruction=inst
	case 0x38: //fstore
		inst:=&stores.FSTORE{}
		inst.Index=uint(reader.ReadUInt16())
		w.modifiedInstruction=inst
	case 0x39: //dstore
		inst:=&stores.DSTORE{}
		inst.Index=uint(reader.ReadUInt16())
		w.modifiedInstruction=inst
	case 0x3a: //astore
		inst:=&stores.ASTORE{}
		inst.Index=uint(reader.ReadUInt16())
		w.modifiedInstruction=inst
	case 0x84: //iinc
		inst:=&math.IINC{}
		inst.Index=uint(reader.ReadUInt16())
		inst.Const=int32(reader.ReadUInt16())
		w.modifiedInstruction=inst
	case 0xa9: //ret
		panic("Unsupported opcode: 0xa9")
	}
}
