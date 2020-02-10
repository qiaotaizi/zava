package classfile

type LineNumberTableAttribute struct {
	lineNumberTable []*LineNumberTableEntry
}

func (l *LineNumberTableAttribute) readInfo(reader *ClassReader) {
	lineNumberTableLength := reader.readUint16()
	l.lineNumberTable = make([]*LineNumberTableEntry, lineNumberTableLength)
	for i := range l.lineNumberTable {
		l.lineNumberTable[i] = &LineNumberTableEntry{
			startPc:    reader.readUint16(),
			lineNumber: reader.readUint16(),
		}
	}
}

func (l *LineNumberTableAttribute) GetLineNumber(pc int) int {
	for i:=len(l.lineNumberTable)-1;i>=0;i--{
		entry:=l.lineNumberTable[i]
		if pc>=int(entry.startPc){
			return int(entry.lineNumber)
		}
	}
	return -1
}

type LineNumberTableEntry struct {
	startPc    uint16
	lineNumber uint16
}
