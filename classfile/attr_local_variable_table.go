package classfile

//暂不实现
//代码参考
//https://github.com/Demon-Cloud/jvmgo/blob/a5218c4ca40ae6a184b780114f41989d50df186c/src/jvmgo/ch03/classfile/attr_local_variable_table.go
type LocalVariableTableAttribute struct {
	//localVariableTable []*LocalVariableTableEntry
}

func (l LocalVariableTableAttribute) readInfo(reader *ClassReader) {
	panic("implement me")
}



