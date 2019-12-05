package trade

import "testing"

func TestBill(t *testing.T) {

	t1 := NewITrader(1)
	trader.init(1)
	trader.Download()

	t1 = NewITrader(2)
	t1.Download()
	t1.Order()

	t1 = NewITrader(3)
	t1.Order()
	t1.Download()
}
