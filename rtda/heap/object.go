package heap

type Object struct {
	class *Class      //对象的类型
	data  interface{} //字段槽
	extra interface{} //目前用该字段表示类对象对应的Class结构体指针
}

func (o *Object) Extra() interface{} {
	return o.extra
}

func (o *Object) SetExtra(extra interface{}) {
	o.extra = extra
}

func (o *Object) Fields() Slots {
	return o.data.(Slots)
}

func (o *Object) IsInstanceOf(class *Class) bool {
	return class.isAssignableFrom(o.class)
}

func (o *Object) Class() *Class {
	return o.class
}

func (o *Object) SetRefVar(name string, descriptor string, ref *Object) {
	field := o.class.GetField(name, descriptor, false)
	slots := o.data.(Slots)
	slots.SetRef(field.slotId, ref)
}

func (o *Object) GetRefVar(name string, descriptor string) *Object {
	field := o.class.GetField(name, descriptor, false)
	return o.data.(Slots).GetRef(field.slotId)
}

func (o *Object) Clone() *Object {
	return &Object{
		class: o.class,
		data:  o.cloneData(),
	}
}

func (o *Object) cloneData() interface{} {
	switch o.data.(type) {
	case []int8:
		int8s := o.data.([]int8)
		int8s2 := make([]int8, len(int8s))
		copy(int8s2, int8s)
		return int8s2
	case []int16:
		int16s:=o.data.([]int16)
		int16s2:=make([]int16,len(int16s))
		copy(int16s2,int16s)
		return int16s2
	case []uint16:
		uint16s:=o.data.([]uint16)
		uint16s2:=make([]uint16,len(uint16s))
		copy(uint16s2,uint16s)
		return uint16s2
	case []int32:
		int32s:=o.data.([]int32)
		int32s2:=make([]int32,len(int32s))
		copy(int32s2,int32s)
		return int32s2
	case []int64:
		int64s:=o.data.([]int64)
		int64s2:=make([]int64,len(int64s))
		copy(int64s2,int64s)
		return int64s2
	case []float32:
		float32s:=o.data.([]float32)
		float32s2:=make([]float32,len(float32s))
		copy(float32s2,float32s)
		return float32s2
	case []float64:
		float64s:=o.data.([]float64)
		float64s2:=make([]float64,len(float64s))
		copy(float64s2,float64s)
		return float64s2
	case []*Object:
		elements := o.data.([]*Object)
		elements2 := make([]*Object, len(elements))
		copy(elements2, elements)
		return elements2
	default: //[]Slot
		slots := o.data.(Slots)
		slots2 := newSlots(uint(len(slots)))
		copy(slots2, slots)
		return slots2
	}

}

func newObject(c *Class) *Object {
	return &Object{
		data:  newSlots(c.instanceSlotCount),
		class: c,
	}
}
