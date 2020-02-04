package heap

type MethodDescriptor struct {
	parameterTypes []string
	returnType string
}

func (d *MethodDescriptor) addParameterType(t string) {
	pLen:=len(d.parameterTypes)
	if pLen==cap(d.parameterTypes){//这里没有使用go提供的数组伸展机制
		s:=make([]string,pLen,pLen+4)
		copy(s,d.parameterTypes)
		d.parameterTypes=s
	}
	d.parameterTypes= append(d.parameterTypes, t)
}
