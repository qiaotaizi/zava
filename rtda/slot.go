package rtda

import "github.com/qiaotaizi/zava/rtda/heap"

//局部变量表大小可预知，按索引访问
//使用数组来实现
//每个元素至少容纳一个int或者引用值，
//两个元素容纳一个long或double

//一个槽要么保存数值
//要么保存引用，即内存地址
type Slot struct {
	num int32
	ref *heap.Object
}
