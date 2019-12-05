package trade

import "fmt"

var wechat = &WeChat{}

type WeChat struct {
}

func (t WeChat) download() {
	fmt.Println("wechat download")
}
func (t WeChat) order() {
	fmt.Println("wechat order")
}
