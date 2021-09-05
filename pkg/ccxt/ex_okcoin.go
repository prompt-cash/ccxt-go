package ccxt

import (
)

type Okcoin struct {
	*ExchangeBase
}

var _ Exchange = (*Okcoin)(nil)

func init() {
	exchange := &Okcoin{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Okcoin) Describe(goArgs ...*Variant) *Variant{
  return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
            "id":MkString("okcoin") ,
            "name":MkString("OKCoin") ,
            "countries":MkArray(&VarArray{
                MkString("CN") ,
                MkString("US") ,
                }),
            "hostname":MkString("okcoin.com") ,
            "pro":MkBool(true) ,
            "urls":MkMap(&VarMap{
                "logo":MkString("https://user-images.githubusercontent.com/51840849/87295551-102fbf00-c50e-11ea-90a9-462eebba5829.jpg") ,
                "www":MkString("https://www.okcoin.com") ,
                "doc":MkString("https://www.okcoin.com/docs/en/") ,
                "fees":MkString("https://www.okcoin.com/coin-fees") ,
                "referral":MkString("https://www.okcoin.com/account/register?flag=activity&channelId=600001513") ,
                }),
            "fees":MkMap(&VarMap{"trading":MkMap(&VarMap{
                    "taker":MkNumber(0.002) ,
                    "maker":MkNumber(0.001) ,
                    })}),
            "options":MkMap(&VarMap{"fetchMarkets":MkArray(&VarArray{
                    MkString("spot") ,
                    })}),
            }));
}

