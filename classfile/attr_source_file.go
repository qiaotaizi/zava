package classfile

//该属性的length总是2
//指定源文件名在常量池中的索引
type SourceFileAttribute struct {
	cp              ConstantPool
	sourceFileIndex uint16
}

func (s *SourceFileAttribute) readInfo(reader *ClassReader) {
	s.sourceFileIndex = reader.readUint16()
}

func (s *SourceFileAttribute) fileName() string {
	return s.cp.getUtf8(s.sourceFileIndex)
}
