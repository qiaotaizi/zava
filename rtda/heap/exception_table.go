package heap

import "github.com/qiaotaizi/zava/classfile"

type ExceptionTable []*ExceptionHandler

type ExceptionHandler struct {
	startPc int //try块开始pc
	endPc int //try块结束pc
	handlerPc int //catch块pc
	catchType *ClassRef //异常类型引用，当这个值为0时，表示catch all
}

func newExceptionTable(entries []*classfile.ExceptionTableEntry,
	cp *ConstantPool) ExceptionTable{
	table:=make([]*ExceptionHandler,len(entries))
	for i,entry:=range entries{
		table[i]=&ExceptionHandler{
			startPc:   int(entry.StartPc()),
			endPc:     int(entry.EndPc()),
			handlerPc: int(entry.HandlerPc()),
			catchType: getCatchType(uint(entry.CatchType()),cp),
		}
	}

	return table
}

func getCatchType(index uint, cp *ConstantPool) *ClassRef {
	if index==0{
		return nil
	}
	return cp.GetConstant(index).(*ClassRef)
}

func (e ExceptionTable) findExceptionHandler(exClass *Class,pc int)*ExceptionHandler{
	for _,handler:=range e{
		if pc>=handler.startPc && pc<handler.endPc{
			if handler.catchType==nil{
				return handler//catch-all
			}

			catchClass:=handler.catchType.ResolvedClass()
			if catchClass==exClass || catchClass.IsSuperClassOf(exClass){
				return handler
			}
		}
	}

	return nil
}
