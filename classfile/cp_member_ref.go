package classfile

//CONSTANT_Fieldref_info:字段符号引用
//CONSTANT_Methodref_info：普通方法（非接口方法）符号引用
//CONSTANT_InterfaceMethodref_info：接口方法引用
//三种类型的常量结构相同
//
//符号引用的意义就是在方法体中引用其他类型的成员、方法、接口方法等

type ConstantMemberrefInfo struct {
	cp               ConstantPool
	classIndex       uint16 //class类型常量见cp_class.go
	nameAndTypeIndex uint16 //nameAndType类型常量见cp_name_and_type.go
}

func (c *ConstantMemberrefInfo) readInfo(reader *ClassReader) {
	c.classIndex = reader.readUint16()
	c.nameAndTypeIndex = reader.readUint16()
}

func (c *ConstantMemberrefInfo) ClassName() string {
	return c.cp.getClassName(c.classIndex)
}

func (c *ConstantMemberrefInfo) NameAndDescriptor() (string, string) {
	return c.cp.getNameAndType(c.nameAndTypeIndex)
}

type ConstantFieldrefInfo struct {
	ConstantMemberrefInfo
}

type ConstantMethodrefInfo struct {
	ConstantMemberrefInfo
}

type ConstantInterfaceMethodrefInfo struct {
	ConstantMemberrefInfo
}
