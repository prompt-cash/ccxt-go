package ccxt

import ()

type Coinbaseprime struct {
	*ExchangeBase
}

var _ Exchange = (*Coinbaseprime)(nil)

func init() {
	exchange := &Coinbaseprime{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Coinbaseprime) Describe(goArgs ...*Variant) *Variant {
	return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
		"id":       MkString("coinbaseprime"),
		"name":     MkString("Coinbase Prime"),
		"pro":      MkBool(true),
		"hostname": MkString("exchange.coinbase.com"),
		"urls": MkMap(&VarMap{
			"test": MkMap(&VarMap{
				"public":  MkString("https://public.sandbox.exchange.coinbase.com"),
				"private": MkString("https://public.sandbox.exchange.coinbase.com"),
			}),
			"logo": MkString("https://user-images.githubusercontent.com/1294454/44539184-29f26e00-a70c-11e8-868f-e907fc236a7c.jpg"),
			"api": MkMap(&VarMap{
				"public":  MkString("https://api.{hostname}"),
				"private": MkString("https://api.{hostname}"),
			}),
			"www": MkString("https://exchange.coinbase.com"),
			"doc": MkString("https://docs.exchange.coinbase.com"),
		}),
	}))
}
