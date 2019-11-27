package phone

import "testing"

func TestPhone(t *testing.T) {
	var iConnectionManager IConnectionManager
	iConnectionManager = NewService()

	p := "110"
	d := iConnectionManager.Dial(p)
	t.Log(d)

	var iDataTransfer IDataTransfer
	iDataTransfer = NewService()

	iDataTransfer.DataTransfer()

	iConnectionManager.Hangup()
}

func TestIPhone(t *testing.T) {
	var iPhone IPhone
	iPhone = NewService()

	p := "110"
	d := iPhone.Dial(p)
	t.Log(d)

	iPhone.DataTransfer()
	iPhone.Hangup()
}
