package heap

//数组object特有的方法

func (o *Object)Bytes()[]int8{//bytes和booleans共用[]int8类型
	return o.data.([]int8)
}

func (o *Object)Shorts()[]int16{
	return o.data.([]int16)
}

func (o *Object)Ints()[]int32{
	return o.data.([]int32)
}

func (o *Object)Longs()[]int64{
	return o.data.([]int64)
}

func (o *Object) Chars()[]uint16{
	return o.data.([]uint16)
}

func (o *Object)Floats()[]float32{
	return o.data.([]float32)
}

func (o *Object)Doubles()[]float64{
	return o.data.([]float64)
}

func (o *Object)Refs()[]*Object{
	return o.data.([]*Object)
}

func (o *Object)ArrayLength()int32{
	switch o.data.(type) {
	case []int8:
		return int32(len(o.Bytes()))
	case []int16:
		return int32(len(o.Shorts()))
	case []int32:
		return int32(len(o.Ints()))
	case []int64:
		return int32(len(o.Longs()))
	case []float32:
		return int32(len(o.Floats()))
	case []float64:
		return int32(len(o.Doubles()))
	case []uint16:
		return int32(len(o.Chars()))
	case []*Object:
		return int32(len(o.Refs()))
	default:
		panic("Not array!")
	}
}

func ArrayCopy(src,dst *Object,srcPos,destPos ,length int32){
	switch src.data.(type) {
	case []int8:
		_src:=src.data.([]int8)[srcPos:srcPos+length]
		_dest:=dst.data.([]int8)[destPos:destPos+length]
		copy(_dest,_src)
	case []int16:
		_src:=src.data.([]int16)[srcPos:srcPos+length]
		_dest:=dst.data.([]int16)[destPos:destPos+length]
		copy(_dest,_src)
	case []int32:
		_src:=src.data.([]int32)[srcPos:srcPos+length]
		_dest:=dst.data.([]int32)[destPos:destPos+length]
		copy(_dest,_src)
	case []int64:
		_src:=src.data.([]int64)[srcPos:srcPos+length]
		_dest:=dst.data.([]int64)[destPos:destPos+length]
		copy(_dest,_src)
	case []uint16:
		_src:=src.data.([]uint16)[srcPos:srcPos+length]
		_dest:=dst.data.([]uint16)[destPos:destPos+length]
		copy(_dest,_src)
	case []float32:
		_src:=src.data.([]float32)[srcPos:srcPos+length]
		_dest:=dst.data.([]float32)[destPos:destPos+length]
		copy(_dest,_src)
	case []float64:
		_src:=src.data.([]float64)[srcPos:srcPos+length]
		_dest:=dst.data.([]float64)[destPos:destPos+length]
		copy(_dest,_src)
	case []*Object:
		_src:=src.data.([]*Object)[srcPos:srcPos+length]
		_dest:=dst.data.([]*Object)[destPos:destPos+length]
		copy(_dest,_src)
	default:
		panic("Not array!")
	}
}
