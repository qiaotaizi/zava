package rtda

import "math"

//操作数栈

type OperandStack struct {
	size  uint   //操作数栈元素下次入栈位置
	slots []Slot //操作数栈槽
}

func newOperandStack(maxStack uint) *OperandStack {
	if maxStack > 0 {
		return &OperandStack{
			slots: make([]Slot, maxStack),
		}
	}
	return nil
}

//局部变量推入 int类型
func (o *OperandStack) PushInt(val int32) {
	o.slots[o.size].num = val
	o.size++
}

func (o *OperandStack) PopInt() int32 {
	o.size--
	return o.slots[o.size].num
}

//float32转int处理
func (o *OperandStack) PushFloat(val float32) {
	bits := math.Float32bits(val)
	o.slots[o.size].num = int32(bits)
	o.size++
}

func (o *OperandStack) PopFloat() float32 {
	o.size--
	return math.Float32frombits(uint32(o.slots[o.size].num))
}

//long拆成两个int处理
func (o *OperandStack) PushLong(val int64) {
	o.slots[o.size].num = int32(val)
	o.slots[o.size+1].num = int32(val >> 32)
	o.size += 2
}

func (o *OperandStack) PopLong() int64 {
	o.size -= 2
	low := uint32(o.slots[o.size].num)
	high := uint32(o.slots[o.size+1].num)
	return int64(high)<<32 + int64(low)
}

//double转long处理
func (o *OperandStack) PushDouble(val float64) {
	bits := math.Float64bits(val)
	o.PushLong(int64(bits))
}

func (o *OperandStack) PopDouble() float64 {
	return math.Float64frombits(uint64(o.PopLong()))
}

//引用类型
func (o *OperandStack) PushRef(val *Object) {
	o.slots[o.size].ref = val
	o.size++
}

func (o *OperandStack) PopRef() *Object {
	o.size--
	ref := o.slots[o.size].ref
	o.slots[o.size].ref = nil //帮助GC栈空间
	return ref
}

//栈操作指令并不关心数据类型
func (o *OperandStack) PushSlot(slot Slot){
	o.slots[o.size]=slot
	o.size++
}

func (o *OperandStack) PopSlot()Slot{
	o.size--
	return o.slots[o.size]
}
