package phone

import (
	"fmt"
	"time"
)

type IPhone interface {
	IConnectionManager
	IDataTransfer
}

type IConnectionManager interface {
	Dial(string) bool
	Hangup()
}

type IDataTransfer interface {
	DataTransfer()
}

func (s Service) Dial(p string) bool {
	fmt.Println("do")
	fmt.Printf("mobile: %s\n", p)
	return true
}

func (s Service) Hangup() {
	fmt.Println("hangup done")
	return
}

func (s Service) DataTransfer() {
	for i := 0; i < 2; i++ {
		fmt.Println("balabala...")
		time.Sleep(1 * time.Second)
	}
	return
}
