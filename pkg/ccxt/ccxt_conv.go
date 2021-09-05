package ccxt

import (
	b64 "encoding/base64"
	"encoding/hex"
	"net/url"
)

///////////////////////////////////////////////////////////////////////////////////////////////
// binary encoding functions

func (this *ExchangeBase) Urlencode(v *Variant) *Variant {
	if v.Type == Map {
		encoded := url.Values{}
		for k, v := range *v.ToMap() {
			encoded.Set(k, (*v).ToStr())
		}
		return MkString(encoded.Encode())
	}
	return MkString(url.QueryEscape(v.ToStr()))
}

func (this *ExchangeBase) UrlencodeWithArrayRepeat(v *Variant) *Variant {
	return this.Urlencode(v) // todo is that right?
}

func (this *ExchangeBase) Rawencode(v *Variant) *Variant {
	str, _ := url.QueryUnescape(this.Urlencode(v).ToStr())
	return MkString(str)
}

func (this *ExchangeBase) Encode(v *Variant) *Variant {
	return v // string to bytearray latin1
}

func (this *ExchangeBase) Decode(v *Variant) *Variant {
	return v // bytearray latin1 to string
}

func (this *ExchangeBase) StringToBinary(v *Variant) *Variant {
	return v // in go a string and a byte array is teh same
}

func (this *ExchangeBase) StringToBase64(v *Variant) *Variant {
	return this.BinaryToBase64(v)
}

func (this *ExchangeBase) Base64ToBinary(v *Variant) *Variant {
	uDec, _ := b64.URLEncoding.DecodeString(v.ToStr())
	return MkString(string(uDec))
}

func (this *ExchangeBase) BinaryToBase64(v *Variant) *Variant {
	uEnc := b64.StdEncoding.EncodeToString([]byte(v.ToStr()))
	return MkString(uEnc)
}

func (this *ExchangeBase) BinaryToBase16(v *Variant) *Variant {
	return MkString(hex.EncodeToString([]byte(v.ToStr())))
}

func (this *ExchangeBase) Base16ToBinary(v *Variant) *Variant {
	decodedByteArray, _ := hex.DecodeString(v.ToStr())
	return MkString(string(decodedByteArray))
}

func (this *ExchangeBase) Omit(obj *Variant, what *Variant) *Variant {
	if obj.Type != Map {
		return OpCopy(obj)
	}
	if what.Type != Array {
		what = MkArray(&VarArray{what})
	}

	// duplicate map dropping all entries which are to be omited
	vMap := VarMap{}
	for k, v := range *(*obj).ToMap() {
		if what.IndexOf(MkString(k)).ToInt() == -1 {
			vMap[k] = *v
		}
	}
	return MkMap(&vMap)
}

func (this *ExchangeBase) OmitZero(v *Variant) *Variant {
	if v.ToFloat() == 0 {
		return MkUndefined()
	}
	return v
}

func (this *ExchangeBase) Base58ToBinary(goArgs ...*Variant) *Variant {
	panic("not implemented function") // todo implement
}

///////////////////////////////////////////////////////////////////////////////////////////////
// otehr encoding functions

func (this *ExchangeBase) Json(v ...*Variant) *Variant {
	if len(v) != 1 {
		panic("todo: To Json options are not implemented")
	}
	return MkString(string(VariantToJson(v[0])))
}

func (this *ExchangeBase) IsJsonEncodedObject(v ...*Variant) *Variant {
	str := v[0].ToStr()
	return MkBool((len(str) >= 2) && ((str[0] == '{') || (str[0] == '[')))
}

func (this *ExchangeBase) ParseJson(v ...*Variant) *Variant {
	return VariantFromJson([]byte(v[0].ToStr()))
}

func (this *ExchangeBase) NumberToBE(goArgs ...*Variant) *Variant {
	panic("not implemented function") // todo implement
}

func (this *ExchangeBase) NumberToLE(goArgs ...*Variant) *Variant {
	panic("not implemented function") // todo implement
}
