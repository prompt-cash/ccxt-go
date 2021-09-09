package ccxt

import (
	b64 "encoding/base64"
	"encoding/hex"
	"fmt"
	"net/url"
	"time"
)

//
// this file contains functions related to data conversion
//

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

///////////////////////////////////////////////////////////////////////////////////////////////
// time conversion functions

func (this *ExchangeBase) ToMilliseconds(goArgs ...*Variant) *Variant {
	// transpiled function
	dateString := GoGetArg(goArgs, 0, MkUndefined())
	if IsTrue(OpEq2(dateString, MkUndefined())) {
		return dateString
	}
	splits := dateString.Split(MkString("-"))
	partOne := this.SafeString(splits, MkInteger(0))
	partTwo := this.SafeString(splits, MkInteger(1))
	if IsTrue(OpOr(OpEq2(partOne, MkUndefined()), OpEq2(partTwo, MkUndefined()))) {
		return MkUndefined()
	}
	if IsTrue(OpNotEq2(partOne.Length, MkInteger(8))) {
		return MkUndefined()
	}
	date := OpAdd(partOne.Slice(MkInteger(0), MkInteger(4)), OpAdd(MkString("-"), OpAdd(partOne.Slice(MkInteger(4), MkInteger(6)), OpAdd(MkString("-"), partOne.Slice(MkInteger(6), MkInteger(8))))))
	return this.Parse8601(OpAdd(date, OpAdd(MkString(" "), partTwo)))
}

func (this *ExchangeBase) Uuidv1(goArgs ...*Variant) *Variant {
	biasSeconds := MkInteger(12219292800)
	bias := OpMulti(biasSeconds, MkInteger(10000000))
	time := OpAdd(OpMulti(this.Microseconds(), MkInteger(10)), bias)
	timeHex := MkString(fmt.Sprintf("%x", time.ToInt()))
	arranged := OpAdd(timeHex.Slice(MkInteger(7), MkInteger(15)), OpAdd(timeHex.Slice(MkInteger(3), MkInteger(7)), OpAdd(MkString("1"), timeHex.Slice(MkInteger(0), MkInteger(3)))))
	clockId := MkString("9696")
	macAddress := MkString("ffffffffffff")
	return OpAdd(arranged, OpAdd(clockId, macAddress))
}

func (this *ExchangeBase) Iso8601(timestamp *Variant) *Variant {
	t := time.Unix(timestamp.ToInt()/1000, (timestamp.ToInt()%1000)*1000000).UTC()
	//return MkString(t.Format(time.RFC3339))
	formatted := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d.%03dZ",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second(), t.Nanosecond()/1000000)
	return MkString(formatted)
}

func (this *ExchangeBase) Parse8601(timestamp *Variant) *Variant {
	t, err := time.Parse(time.RFC3339, timestamp.ToStr())
	if err != nil {
		return MkUndefined()
	}
	return MkInteger(t.UTC().UnixNano() / 1000000)
}

func (this *ExchangeBase) Nonce(goArgs ...*Variant) *Variant {
	ret := this.VCall("Nonce", goArgs...)
	if ret != nil {
		return ret
	}
	return this.Seconds()
}

func (this *ExchangeBase) Ymd(goArgs ...*Variant) *Variant {
	timestamp := GoGetArg(goArgs, 0, MkUndefined())
	infix := GoGetArg(goArgs, 1, MkString("-")).ToStr()
	t := time.Unix(timestamp.ToInt()/1000, 0).UTC()
	formatted := fmt.Sprintf("%d"+infix+"%02d"+infix+"%02d",
		t.Year(), t.Month(), t.Day())
	return MkString(formatted)
}

func (this *ExchangeBase) Ymdhms(goArgs ...*Variant) *Variant {
	timestamp := GoGetArg(goArgs, 0, MkUndefined())
	infix := GoGetArg(goArgs, 1, MkString(" ")).ToStr()
	t := time.Unix(timestamp.ToInt()/1000, 0).UTC()
	formatted := fmt.Sprintf("%d-%02d-%02d"+infix+"%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	return MkString(formatted)
}
