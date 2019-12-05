package trade

func (t *Trade) init(pay int) {
	switch pay {
	case 1:
		t.biller = wechat
		t.payer = wechat
	case 2:
		t.biller = alipay
		t.payer = alipay
	case 3:
		t.payer = apple
	}
}

func (t *Trade) Download() {
	b := t.biller
	if b == nil {
		return
	}
	b.download()
}

func (t *Trade) Order() {
	p := t.payer
	if p == nil {
		return
	}
	p.order()
}

type ITrader interface {
	init(p int)
	Download()
	Order()
}

type Trade struct {
	biller Biller
	payer  Payer
}

var trader ITrader

type Biller interface {
	download()
}

type Payer interface {
	order()
}

func NewITrader(pay int) ITrader {
	trader = new(Trade)
	trader.init(pay)
	return trader
}
