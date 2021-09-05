package ccxt

import ()

type Cdax struct {
	*ExchangeBase
}

var _ Exchange = (*Cdax)(nil)

func init() {
	exchange := &Cdax{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Cdax) Describe(goArgs ...*Variant) *Variant {
	return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
		"id":   MkString("cdax"),
		"name": MkString("CDAX"),
		"countries": MkArray(&VarArray{
			MkString("RU"),
		}),
		"hostname": MkString("cdax.io"),
		"pro":      MkBool(false),
		"urls": MkMap(&VarMap{
			"logo": MkString("https://user-images.githubusercontent.com/1294454/102157692-fd406280-3e90-11eb-8d46-4511b617cd17.jpg"),
			"api": MkMap(&VarMap{
				"market":    MkString("https://{hostname}/api"),
				"public":    MkString("https://{hostname}/api"),
				"private":   MkString("https://{hostname}/api"),
				"v2Public":  MkString("https://{hostname}/api"),
				"v2Private": MkString("https://{hostname}/api"),
			}),
			"www":      MkString("https://cdax.io"),
			"referral": MkString("https://cdax.io/invite?invite_code=esc74"),
			"doc":      MkString("https://github.com/cloudapidoc/API_Docs"),
			"fees":     MkString("https://cdax.io/about/fee"),
		}),
	}))
}
