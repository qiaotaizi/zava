package rtda

import "github.com/qiaotaizi/zava/rtda/heap"

//线程和栈

type Thread struct {
	pc    int //计数器
	stack *Stack
}

func NewThread() *Thread {
	return &Thread{stack: newStack(1024)} //最大栈帧1024
}

func (t *Thread)NewFrame(method *heap.Method)*Frame{
	return newFrame(t,method)
}

func (t *Thread) PC() int {
	return t.pc
}

func (t *Thread) SetPC(pc int) {
	t.pc = pc
}

//帧入栈
func (t *Thread) PushFrame(frame *Frame) {
	t.stack.push(frame)
}

//帧出栈
func (t *Thread) PopFrame() *Frame {
	return t.stack.pop()
}

//获取当前帧
func (t *Thread) CurrentFrame() *Frame {
	return nil
}
