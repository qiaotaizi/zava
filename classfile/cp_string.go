package classfile

//string常量不存放字符串数据
//而是mutf8索引

type ConstantStringInfo struct {
	cp          ConstantPool
	stringIndex uint16
}

func (c *ConstantStringInfo) readInfo(reader *ClassReader) {
	c.stringIndex = reader.readUint16()
}

//按索引从常量池中查找字符串
func (c *ConstantStringInfo) String() string {
	return c.cp.getUtf8(c.stringIndex)
}
