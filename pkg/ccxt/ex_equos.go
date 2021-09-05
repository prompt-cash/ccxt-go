package ccxt

import (
)

type Equos struct {
	*ExchangeBase
}

var _ Exchange = (*Equos)(nil)

func init() {
	exchange := &Equos{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Equos) Describe(goArgs ...*Variant) *Variant{
  return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
            "id":MkString("equos") ,
            "name":MkString("EQUOS") ,
            "urls":MkMap(&VarMap{"logo":MkString("https://user-images.githubusercontent.com/1294454/107758499-05edd180-6d38-11eb-9e09-0b69602a7a15.jpg") }),
            }));
}

