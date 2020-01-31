package base

import "github.com/qiaotaizi/zava/rtda"

//分支逻辑跳转函数
func Branch(frame *rtda.Frame,offset int){
	pc:=frame.Thread().PC()
	nextPC:=pc+offset
	frame.SetNextPC(nextPC)
}
