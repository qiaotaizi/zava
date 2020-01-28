package classfile

type ConstantClassInfo struct {
	cp        ConstantPool
	nameIndex uint16
}

func (c *ConstantClassInfo) readInfo(reader *ClassReader) {
	c.nameIndex = reader.readUint16()
}

//获取类或者接口的引用符号
//即类名
func (c *ConstantClassInfo) Name() string {
	return c.cp.getUtf8(c.nameIndex)
}
