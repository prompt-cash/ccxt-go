package ccxt

import ()

type Okex5 struct {
	*ExchangeBase
}

var _ Exchange = (*Okex5)(nil)

func init() {
	exchange := &Okex5{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Okex5) Describe(goArgs ...*Variant) *Variant {
	return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{"id": MkString("okex5")}))
}
