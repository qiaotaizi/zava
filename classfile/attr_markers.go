package classfile

//@Deprecate注解出现之前
//使用/** deprecate */文档标记废弃的类、方法等
type DeprecatedAttribute struct {
	MarkerAttribute
}

//由编译期生成的代码会被添加synthetic属性
type SyntheticAttribute struct {
	MarkerAttribute
}

//这类属性本身没有任何信息
//其length总是为0
type MarkerAttribute struct {
}

func (m *MarkerAttribute) readInfo(reader *ClassReader) {

}
