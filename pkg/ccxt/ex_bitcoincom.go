package ccxt

import ()

type Bitcoincom struct {
	*ExchangeBase
}

var _ Exchange = (*Bitcoincom)(nil)

func init() {
	exchange := &Bitcoincom{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Bitcoincom) Describe(goArgs ...*Variant) *Variant {
	return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
		"id":   MkString("bitcoincom"),
		"name": MkString("bitcoin.com"),
		"countries": MkArray(&VarArray{
			MkString("KN"),
		}),
		"urls": MkMap(&VarMap{
			"logo": MkString("https://user-images.githubusercontent.com/1294454/97296144-514fa300-1861-11eb-952b-3d55d492200b.jpg"),
			"api": MkMap(&VarMap{
				"public":  MkString("https://api.exchange.bitcoin.com"),
				"private": MkString("https://api.exchange.bitcoin.com"),
			}),
			"www":      MkString("https://exchange.bitcoin.com"),
			"doc":      MkString("https://api.exchange.bitcoin.com/api/2/explore"),
			"fees":     MkString("https://exchange.bitcoin.com/fees-and-limits"),
			"referral": MkString("https://exchange.bitcoin.com/referral/da948b21d6c92d69"),
		}),
		"fees": MkMap(&VarMap{"trading": MkMap(&VarMap{
			"maker": this.ParseNumber(MkString("0.0015")),
			"taker": this.ParseNumber(MkString("0.002")),
		})}),
	}))
}
