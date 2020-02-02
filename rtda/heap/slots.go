package heap

import "math"

//代码与localVars基本上一致
//主要是为了避免循环依赖
type Slot struct{
	num int32
	ref *Object
}

type Slots []Slot


func newSlots(slotCount uint) Slots {
	if slotCount > 0 {
		return make([]Slot, slotCount)
	}
	return nil
}

//负责处理int,bool,byte,short,char
func (l Slots) SetInt(index uint, val int32) {
	l[index].num = val
}

func (l Slots) GetInt(index uint) int32 {
	return l[index].num
}

//float32类型的数据转成int处理
func (l Slots) SetFloat(index uint, val float32) {
	bits := math.Float32bits(val)
	l[index].num = int32(bits)
}

func (l Slots) GetFloat(index uint) float32 {
	return math.Float32frombits(uint32(l[index].num))
}

//long拆成两个连续的slot
func (l Slots) SetLong(index uint, val int64) {
	l[index].num = int32(val)         //截断，保留低位
	l[index+1].num = int32(val >> 32) //保留高位
}

func (l Slots) GetLong(index uint) int64 {
	low := uint32(l[index].num)
	high := uint32(l[index+1].num)
	return int64(high)<<32 + int64(low) //这个二进制运算是怎么回事？
}

//double转成long，然后按照和long一样处理
func (l Slots) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	l.SetLong(index, int64(bits))
}

func (l Slots) GetDouble(index uint) float64 {
	return math.Float64frombits(uint64(l.GetLong(index)))
}

//引用值直接存取
func (l Slots) SetRef(index uint, val *Object) {
	l[index].ref = val
}

func (l Slots) GetRef(index uint) *Object {
	return l[index].ref
}
