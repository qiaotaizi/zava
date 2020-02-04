package rtda

import (
	"github.com/qiaotaizi/zava/rtda/heap"
	"math"
)

type LocalVars []Slot

func newLocalVars(maxLocals uint) LocalVars {
	if maxLocals > 0 {
		return make([]Slot, maxLocals)
	}
	return nil
}

//负责处理int,bool,byte,short,char
func (l LocalVars) SetInt(index uint, val int32) {
	l[index].num = val
}

func (l LocalVars) GetInt(index uint) int32 {
	return l[index].num
}

//float32类型的数据转成int处理
func (l LocalVars) SetFloat(index uint, val float32) {
	bits := math.Float32bits(val)
	l[index].num = int32(bits)
}

func (l LocalVars) GetFloat(index uint) float32 {
	return math.Float32frombits(uint32(l[index].num))
}

//long拆成两个连续的slot
func (l LocalVars) SetLong(index uint, val int64) {
	l[index].num = int32(val)         //截断，保留低位
	l[index+1].num = int32(val >> 32) //保留高位
}

func (l LocalVars) GetLong(index uint) int64 {
	low := uint32(l[index].num)
	high := uint32(l[index+1].num)
	return int64(high)<<32 + int64(low) //这个二进制运算是怎么回事？
}

//double转成long，然后按照和long一样处理
func (l LocalVars) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	l.SetLong(index, int64(bits))
}

func (l LocalVars) GetDouble(index uint) float64 {
	return math.Float64frombits(uint64(l.GetLong(index)))
}

//引用值直接存取
func (l LocalVars) SetRef(index uint, val *heap.Object) {
	l[index].ref = val
}

func (l LocalVars) GetRef(index uint) *heap.Object {
	return l[index].ref
}

func (l LocalVars) SetSlot(index uint,slot Slot) {
	l[index]=slot
}
