package classfile

//存储类属性和方法的名称和描述符

type ConstantNameAndTypeInfo struct {
	nameIndex       uint16
	descriptorIndex uint16
}

func (c *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	c.nameIndex = reader.readUint16()       //属性、方法名称在常量池的索引
	c.descriptorIndex = reader.readUint16() //描述符，对于字段，表示字段类型，对于方法，表示参数类型和返回值类型，类型使用简短的字符描述
}
