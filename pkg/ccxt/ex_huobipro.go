package ccxt

import ()

type Huobipro struct {
	*ExchangeBase
}

var _ Exchange = (*Huobipro)(nil)

func init() {
	exchange := &Huobipro{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Huobipro) Describe(goArgs ...*Variant) *Variant {
	return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{"id": MkString("huobipro")}))
}
