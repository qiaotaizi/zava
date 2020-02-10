package classfile

//变长属性
//存储方法字节码信息

type CodeAttribute struct {
	cp             ConstantPool
	maxStack       uint16
	maxLocals      uint16
	code           []byte
	exceptionTable []*ExceptionTableEntry
	attributes     []AttributeInfo
}

func (c *CodeAttribute) MaxStack() uint {
	return uint(c.maxStack)
}

func (c *CodeAttribute) MaxLocals() uint {
	return uint(c.maxLocals)
}

func (c *CodeAttribute) Code() []byte {
	return c.code
}

func (c *CodeAttribute) readInfo(reader *ClassReader) {
	c.maxStack = reader.readUint16()
	c.maxLocals = reader.readUint16()
	codeLength := reader.readUint32() //2^32-1代码最大长度？
	c.code = reader.readBytes(codeLength)
	c.exceptionTable = readExceptionTable(reader)
	c.attributes = readAttributes(reader, c.cp)
}

type ExceptionTableEntry struct {
	startPc   uint16
	endPc     uint16
	handlerPc   uint16
	catchType uint16
}

func (e *ExceptionTableEntry)StartPc()uint16{
	return e.startPc
}

func (e *ExceptionTableEntry)EndPc()uint16{
	return e.endPc
}

func (e *ExceptionTableEntry)HandlerPc()uint16{
	return e.handlerPc
}

func (e *ExceptionTableEntry)CatchType()uint16{
	return e.catchType
}

func readExceptionTable(reader *ClassReader) []*ExceptionTableEntry {
	exceptionTableLength := reader.readUint16()
	exceptionTable := make([]*ExceptionTableEntry, exceptionTableLength)
	for i := range exceptionTable {
		exceptionTable[i] = &ExceptionTableEntry{
			startPc:   reader.readUint16(),
			endPc:     reader.readUint16(),
			handlerPc:   reader.readUint16(),
			catchType: reader.readUint16(),
		}
	}

	return exceptionTable
}
