package heap

type Object struct {
	class *Class //对象的类型
	data interface{}
}

func (o *Object) Fields() Slots {
	return o.data.(Slots)
}

func (o *Object) IsInstanceOf(class *Class) bool {
	return class.isAssignableFrom(o.class)
}

func (o *Object)Class()*Class{
	return o.class
}

func (o *Object) SetRefVar(name string, descriptor string, ref *Object) {
	field:=o.class.GetField(name,descriptor,false)
	slots:=o.data.(Slots)
	slots.SetRef(field.slotId,ref)
}

func (o *Object) GetRefVar(name string, descriptor string) *Object {
	field:=o.class.GetField(name,descriptor,false)
	return o.data.(Slots).GetRef(field.slotId)
}

func newObject(c *Class) *Object {
	return &Object{
		data:newSlots(c.instanceSlotCount),
		class:c,
	}
}
