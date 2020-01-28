package classfile

//MUTF-8字符串常量
type ConstantUtf8Info struct {
	str string
}

func (c *ConstantUtf8Info) readInfo(reader *ClassReader) {
	length := uint32(reader.readUint16()) //string的tag之后有两个字节表示字符串的长度。字符串长度最大65535？
	bytes := reader.readBytes(length)
	c.str = decodeMUTF8(bytes)
}

func decodeMUTF8(bytes []byte) string {
	return string(bytes) //TODO 可以看一下复杂实现
}
