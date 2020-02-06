package heap

//数组类特有的方法

func (c *Class)NewArray(count uint)*Object{
	if !c.IsArray(){
		panic("Not Array Class: "+c.Name())
	}
	switch c.Name() {
	case "[Z"://booleans
		return &Object{data:make([]int8,count),class:c}
	case "[B"://bytes
	return &Object{data:make([]int8,count),class:c}
	case "[C"://Chars:
		return &Object{data:make([]uint16,count),class:c}
	case "[S"://shorts
		return &Object{data:make([]int16,count),class:c}
	case "[I"://ints
		return &Object{data:make([]int32,count),class:c}
	case "[J"://longs
		return &Object{data:make([]int64,count),class:c}
	case "[F"://floats
		return &Object{data:make([]float32,count),class:c}
	case "[D"://doubles
		return &Object{data:make([]float64,count),class:c}
	default:
		return &Object{data:make([]*Object,count),class:c}
	}
}

func (c *Class) IsArray() bool {
	 return c.name[0]=='['
}

func (c *Class) ComponentClass() *Class {
	componentClassName:=getComponentClassName(c.name)
	return c.loader.LoadClass(componentClassName)
}
