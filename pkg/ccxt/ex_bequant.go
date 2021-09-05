package ccxt

import (
)

type Bequant struct {
	*ExchangeBase
}

var _ Exchange = (*Bequant)(nil)

func init() {
	exchange := &Bequant{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Bequant) Describe(goArgs ...*Variant) *Variant{
  return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
            "id":MkString("bequant") ,
            "name":MkString("Bequant") ,
            "countries":MkArray(&VarArray{
                MkString("MT") ,
                }),
            "pro":MkBool(true) ,
            "urls":MkMap(&VarMap{
                "logo":MkString("https://user-images.githubusercontent.com/1294454/55248342-a75dfe00-525a-11e9-8aa2-05e9dca943c6.jpg") ,
                "api":MkMap(&VarMap{
                    "public":MkString("https://api.bequant.io") ,
                    "private":MkString("https://api.bequant.io") ,
                    }),
                "www":MkString("https://bequant.io") ,
                "doc":MkArray(&VarArray{
                    MkString("https://api.bequant.io/") ,
                    }),
                "fees":MkArray(&VarArray{
                    MkString("https://bequant.io/fees-and-limits") ,
                    }),
                "referral":MkString("https://bequant.io") ,
                }),
            }));
}

