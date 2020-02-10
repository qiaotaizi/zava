package rtda

type Stack struct {
	maxSize uint   //最大栈长
	size    uint   //当前栈长
	_top    *Frame //当前帧
}

//帧入栈
func (s *Stack) push(frame *Frame) {
	if s.size >= s.maxSize {
		panic("java.lang.StackOverflowError")
	}
	if s._top != nil {
		frame.lower = s._top
	}
	s._top = frame
	s.size++
}

//帧出栈
func (s *Stack) pop() *Frame {
	if s._top == nil {
		panic("jvm stack is empty!")
	}
	top := s._top
	s._top = top.lower
	top.lower = nil //节点出链
	s.size--
	return top
}

func (s *Stack) top() *Frame {
	if s._top == nil {
		panic("jvm stack is empty!")
	}
	return s._top
}

func (s *Stack) isEmpty() bool {
	return s._top==nil
}

func (s *Stack) clear() {
	for !s.isEmpty(){
		s.pop()
	}
}

func (s *Stack) getFrames() []*Frame {
	frames:=make([]*Frame,0,s.size)
	for frame:=s._top;frame!=nil;frame=frame.lower{
		frames=append(frames, frame)
	}
	return frames
}

func newStack(maxSize uint) *Stack {
	return &Stack{maxSize: maxSize}
}
