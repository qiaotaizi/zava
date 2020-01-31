package base

type BytecodeReader struct {
	code []byte //存放字节码
	pc   int    //记录读取到了哪个字节
}

func (r *BytecodeReader)PC()int{
	return r.pc
}

//重置为初始化状态
func (r *BytecodeReader) Reset(code []byte, pc int) {
	r.code = code
	r.pc = pc
}

//读取一个字节的code
func (r *BytecodeReader) ReadUint8() uint8 {
	i := r.code[r.pc]
	r.pc++
	return i
}

func (r *BytecodeReader) ReadInt8() int8 {
	return int8(r.ReadUint8())
}

//连续读取两字节
func (r *BytecodeReader) ReadUInt16() uint16 {
	byte1 := uint16(r.ReadUint8())
	byte2 := uint16(r.ReadUint8())
	return (byte1 << 8) | byte2
}

func (r *BytecodeReader) ReadInt16() int16 {
	return int16(r.ReadUInt16())
}

//连续读取4字节
func (r *BytecodeReader) ReadInt32() int32 {
	byte1 := int32(r.ReadUint8())
	byte2 := int32(r.ReadUint8())
	byte3 := int32(r.ReadUint8())
	byte4 := int32(r.ReadUint8())
	return (byte1 << 24) | (byte2 << 16) | (byte3 << 8) | byte4
}

//tableswitch指令操作码后有0~3字节的padding
//用以保证defaultOffset在字节码中的地址是4的倍数
func (r *BytecodeReader) SkipPadding() {
	for r.pc%4!=0{
		r.ReadUint8()
	}
}

func (r *BytecodeReader) ReadInt32s(count int32) []int32 {
	ints:=make([]int32 ,count)
	for i:=range ints{
		ints[i]=r.ReadInt32()
	}
	return ints
}
