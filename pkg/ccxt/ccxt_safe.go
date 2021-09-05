package ccxt

///////////////////////////////////////////////////////////////////////////////////////////////
// safe functions

func (this *ExchangeBase) SafeString(v *Variant, key *Variant, a ...*Variant) *Variant {
	return this.SafeString2(v, key, nil, a...)
}

func (this *ExchangeBase) SafeString2(v *Variant, key1 *Variant, key2 *Variant, a ...*Variant) *Variant {
	val := v.GetRef(key1)
	if val == nil || (*val).Type == Undefined && key2 != nil {
		val = v.GetRef(key2)
	}
	if val != nil && (*val).Type != Undefined {
		return (*val).ToString()
	}
	if len(a) >= 1 {
		return OpCopy(a[0])
	}
	return MkUndefined()
}

func (this *ExchangeBase) SafeStringLower(v *Variant, key *Variant, a ...*Variant) *Variant {
	return this.SafeStringLower2(v, key, nil, a...)
}

func (this *ExchangeBase) SafeStringLower2(v *Variant, key1 *Variant, key2 *Variant, a ...*Variant) *Variant {
	val := v.GetRef(key1)
	if val == nil || (*val).Type == Undefined && key2 != nil {
		val = v.GetRef(key2)
	}
	if val != nil && (*val).Type != Undefined {
		return (*val).ToLowerCase()
	}
	if len(a) >= 1 {
		return OpCopy(a[0])
	}
	return MkUndefined()
}

func (this *ExchangeBase) SafeStringUpper(v *Variant, key *Variant, a ...*Variant) *Variant {
	return this.SafeStringUpper2(v, key, nil, a...)
}

func (this *ExchangeBase) SafeStringUpper2(v *Variant, key1 *Variant, key2 *Variant, a ...*Variant) *Variant {
	val := v.GetRef(key1)
	if val == nil || (*val).Type == Undefined && key2 != nil {
		val = v.GetRef(key2)
	}
	if val != nil && (*val).Type != Undefined {
		return (*val).ToUpperCase()
	}
	if len(a) >= 1 {
		return OpCopy(a[0])
	}
	return MkUndefined()
}

func (this *ExchangeBase) SafeInteger(v *Variant, key *Variant, a ...*Variant) *Variant {
	return this.SafeInteger2(v, key, nil, a...)
}

func (this *ExchangeBase) SafeInteger2(v *Variant, key1 *Variant, key2 *Variant, a ...*Variant) *Variant {
	val := v.GetRef(key1)
	if val == nil || (*val).Type == Undefined && key2 != nil {
		val = v.GetRef(key2)
	}
	if val != nil && (*val).Type != Undefined {
		i, err := (*val).ToSafeInt()
		if err == nil {
			return MkInteger(i)
		}
	}
	if len(a) >= 1 {
		return OpCopy(a[0])
	}
	return MkUndefined()
}

func (this *ExchangeBase) SafeNumber(v *Variant, key *Variant, a ...*Variant) *Variant {
	return this.SafeNumber2(v, key, nil, a...)
}

func (this *ExchangeBase) SafeNumber2(v *Variant, key1 *Variant, key2 *Variant, a ...*Variant) *Variant {
	val := v.GetRef(key1)
	if val == nil || (*val).Type == Undefined && key2 != nil {
		val = v.GetRef(key2)
	}
	if val != nil && (*val).Type != Undefined {
		f, err := (*val).ToSafeFloat()
		if err == nil {
			return MkNumber(f)
		}
	}
	if len(a) >= 1 {
		return OpCopy(a[0])
	}
	return MkUndefined()
}

func (this *ExchangeBase) ParseNumber(v *Variant, a ...*Variant) *Variant {
	f, err := (*v).ToSafeFloat()
	if err == nil {
		return MkNumber(f)
	}
	if len(a) >= 1 {
		return OpCopy(a[0])
	}
	return MkUndefined()
}

func (this *ExchangeBase) ParseSafeNumber(goArgs ...*Variant) *Variant {
	// transpiled function
	value := GoGetArg(goArgs, 0, MkUndefined())
	if IsTrue(OpEq2(value, MkUndefined())) {
		return value
	}
	value = value.Replace(MkString(","), MkString(""))
	parts := value.Split(MkString(" "))
	return this.SafeNumber(parts, MkInteger(0))
}

func (this *ExchangeBase) SafeFloat(v *Variant, key *Variant, a ...*Variant) *Variant {
	return this.SafeNumber(v, key, a...)
}

func (this *ExchangeBase) SafeFloat2(v *Variant, key1 *Variant, key2 *Variant, a ...*Variant) *Variant {
	return this.SafeNumber2(v, key1, key2, a...)
}

func (this *ExchangeBase) SafeValue(v *Variant, key *Variant, a ...*Variant) *Variant {
	return this.SafeValue2(v, key, nil, a...)
}

func (this *ExchangeBase) SafeValue2(v *Variant, key1 *Variant, key2 *Variant, a ...*Variant) *Variant {
	val := v.GetRef(key1)
	if val == nil || (*val).Type == Undefined && key2 != nil {
		val = v.GetRef(key2)
	}
	if val != nil && (*val).Type != Undefined {
		return *val
	}
	if len(a) >= 1 {
		return OpCopy(a[0])
	}
	return MkUndefined()
}

func (this *ExchangeBase) SafeTimestamp(v *Variant, key *Variant, a ...*Variant) *Variant {
	return this.SafeTimestamp2(v, key, nil, a...)
}

func (this *ExchangeBase) SafeTimestamp2(v *Variant, key1 *Variant, key2 *Variant, a ...*Variant) *Variant {
	return this.SafeIntegerProduct2(v, key1, key2, MkInteger(1000), a...)
}

func (this *ExchangeBase) SafeIntegerProduct(v *Variant, key *Variant, factor *Variant, a ...*Variant) *Variant {
	return this.SafeIntegerProduct2(v, key, nil, factor, a...)
}

func (this *ExchangeBase) SafeIntegerProduct2(v *Variant, key1 *Variant, key2 *Variant, factor *Variant, a ...*Variant) *Variant {
	val := v.GetRef(key1)
	if val == nil && key2 != nil {
		val = v.GetRef(key2)
	}
	if val == nil || (*val).Type == Undefined {
		i, err := v.ToSafeInt()
		if err == nil {
			return MkInteger(i * factor.ToInt())
		}
	}
	if len(a) >= 1 {
		return OpCopy(a[0])
	}
	return MkUndefined()
}
