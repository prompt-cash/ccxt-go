package ccxt

import (
)

type Binanceus struct {
	*ExchangeBase
}

var _ Exchange = (*Binanceus)(nil)

func init() {
	exchange := &Binanceus{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Binanceus) Describe(goArgs ...*Variant) *Variant{
  return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
            "id":MkString("binanceus") ,
            "name":MkString("Binance US") ,
            "countries":MkArray(&VarArray{
                MkString("US") ,
                }),
            "certified":MkBool(false) ,
            "pro":MkBool(true) ,
            "urls":MkMap(&VarMap{
                "logo":MkString("https://user-images.githubusercontent.com/1294454/65177307-217b7c80-da5f-11e9-876e-0b748ba0a358.jpg") ,
                "api":MkMap(&VarMap{
                    "web":MkString("https://www.binance.us") ,
                    "sapi":MkString("https://api.binance.us/sapi/v1") ,
                    "wapi":MkString("https://api.binance.us/wapi/v3") ,
                    "public":MkString("https://api.binance.us/api/v1") ,
                    "private":MkString("https://api.binance.us/api/v3") ,
                    "v3":MkString("https://api.binance.us/api/v3") ,
                    "v1":MkString("https://api.binance.us/api/v1") ,
                    }),
                "www":MkString("https://www.binance.us") ,
                "referral":MkString("https://www.binance.us/?ref=35005074") ,
                "doc":MkString("https://github.com/binance-us/binance-official-api-docs") ,
                "fees":MkString("https://www.binance.us/en/fee/schedule") ,
                }),
            "fees":MkMap(&VarMap{"trading":MkMap(&VarMap{
                    "tierBased":MkBool(true) ,
                    "percentage":MkBool(true) ,
                    "taker":this.ParseNumber(MkString("0.001") ),
                    "maker":this.ParseNumber(MkString("0.001") ),
                    })}),
            "options":MkMap(&VarMap{"quoteOrderQty":MkBool(false) }),
            }));
}

func (this *Binanceus) FetchCurrencies(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  return MkUndefined() ;
}

