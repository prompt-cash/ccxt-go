package ccxt

import ()

type Huobijp struct {
	*ExchangeBase
}

var _ Exchange = (*Huobijp)(nil)

func init() {
	exchange := &Huobijp{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Huobijp) Describe(goArgs ...*Variant) *Variant {
	return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
		"id":   MkString("huobijp"),
		"name": MkString("Huobi Japan"),
		"countries": MkArray(&VarArray{
			MkString("JP"),
		}),
		"hostname": MkString("api-cloud.huobi.co.jp"),
		"pro":      MkBool(true),
		"has":      MkMap(&VarMap{"fetchDepositAddress": MkBool(false)}),
		"urls": MkMap(&VarMap{
			"logo": MkString("https://user-images.githubusercontent.com/1294454/85734211-85755480-b705-11ea-8b35-0b7f1db33a2f.jpg"),
			"api": MkMap(&VarMap{
				"market":  MkString("https://{hostname}"),
				"public":  MkString("https://{hostname}"),
				"private": MkString("https://{hostname}"),
			}),
			"www":      MkString("https://www.huobi.co.jp"),
			"referral": MkString("https://www.huobi.co.jp/register/?invite_code=znnq3"),
			"doc":      MkString("https://api-doc.huobi.co.jp"),
			"fees":     MkString("https://www.huobi.co.jp/support/fee"),
		}),
	}))
}

func (this *Huobijp) FetchDepositAddress(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	panic(NewNotSupported(OpAdd(*this.At(MkString("id")), MkString(" fetchDepositAddress not supported yet"))))
	return MkUndefined()
}
