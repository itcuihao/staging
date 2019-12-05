package trade

import "fmt"

var alipay = &AliPay{}

type AliPay struct {
}

func (t AliPay) download() {
	fmt.Println("alipay download")
}

func (t AliPay) order() {
	fmt.Println("alipay order")
}
