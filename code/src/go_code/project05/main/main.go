package main

import (
	"fmt"
)

type USB interface{
	start()
	stop()
}

type phone struct{
}

func (h phone) start()  {
	fmt.Println("手机启动")
}
func (h phone) stop()  {
	fmt.Println("手机关闭")
}

type camera struct{
}

func (c camera) start()  {
	fmt.Println("相机启动")
}
func (c camera) stop()  {
	fmt.Println("相机关闭")
}

func main(){

	var h USB=phone{}
	h.start()
	h.stop()

	var c USB=camera{}
	c.start()
	c.stop()

}