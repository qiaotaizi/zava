package heap

type Object struct {
	class *Class //对象的类型
	fields Slots //实例变量
}

func (o *Object) Fields() Slots {
	return o.fields
}

func (o *Object) IsInstanceOf(class *Class) bool {
	return class.isAssignableFrom(o.class)
}
