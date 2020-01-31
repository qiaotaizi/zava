package classfile

//暂不实现
//代码参考
//https://github.com/Demon-Cloud/jvmgo/blob/a5218c4ca40ae6a184b780114f41989d50df186c/src/jvmgo/ch03/classfile/attr_local_variable_table.go
type LocalVariableTableAttribute struct {
	localVariableTable []*LocalVariableTableEntry
}

type LocalVariableTableEntry struct {
	startPc uint16
	length uint16
	nameIndex uint16
	descriptor uint16
	index uint16
}

func (l *LocalVariableTableAttribute) readInfo(reader *ClassReader) {
	length:=reader.readUint16()
	l.localVariableTable=make([]*LocalVariableTableEntry,length)
	for i:=range l.localVariableTable{
		l.localVariableTable[i]=&LocalVariableTableEntry{
			startPc:    reader.readUint16(),
			length:     reader.readUint16(),
			nameIndex:  reader.readUint16(),
			descriptor: reader.readUint16(),
			index:      reader.readUint16(),
		}
	}
}
