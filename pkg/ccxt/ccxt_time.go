package ccxt

import (
	"fmt"
	"time"
)

func (this *ExchangeBase) Milliseconds() *Variant {
	return MkInteger(time.Now().UnixNano() / int64(time.Millisecond))
}
func (this *ExchangeBase) Microseconds() *Variant {
	return MkInteger(time.Now().UnixNano() / int64(time.Microsecond))
}
func (this *ExchangeBase) Seconds() *Variant {
	return MkInteger(time.Now().UnixNano() / int64(time.Second))
}
func (this *ExchangeBase) ToMilliseconds(goArgs ...*Variant) *Variant {
	// transpiled function
	dateString := GoGetArg(goArgs, 0, MkUndefined())
	_ = dateString
	if IsTrue(OpEq2(dateString, MkUndefined())) {
		return dateString
	}
	splits := dateString.Split(MkString("-"))
	_ = splits
	partOne := this.SafeString(splits, MkInteger(0))
	_ = partOne
	partTwo := this.SafeString(splits, MkInteger(1))
	_ = partTwo
	if IsTrue(OpOr(OpEq2(partOne, MkUndefined()), OpEq2(partTwo, MkUndefined()))) {
		return MkUndefined()
	}
	if IsTrue(OpNotEq2(partOne.Length, MkInteger(8))) {
		return MkUndefined()
	}
	date := OpAdd(partOne.Slice(MkInteger(0), MkInteger(4)), OpAdd(MkString("-"), OpAdd(partOne.Slice(MkInteger(4), MkInteger(6)), OpAdd(MkString("-"), partOne.Slice(MkInteger(6), MkInteger(8))))))
	_ = date
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
