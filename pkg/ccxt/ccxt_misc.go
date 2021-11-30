package ccxt

//
// variouse other ccxt functions
//

import (
	"github.com/prompt-cash/ccxt-go/utils"
	"regexp"
	"sort"
	"strings"
	"time"
)

func (this *ExchangeBase) ArrayConcat(goArgs ...*Variant) *Variant {
	vArr := VarArray{}
	j := 0
	for m := 0; m < len(goArgs); m++ {
		arr := goArgs[m].ToArray()
		for i := 0; i < len(*arr); i++ {
			vArr[j] = *(*arr)[i]
			j++
		}
	}
	return MkArray(&vArr)
}

func (this *ExchangeBase) BinaryConcat(goArgs ...*Variant) *Variant {
	return this.ArrayConcat(goArgs...) // seams to be the same
}

func (this *ExchangeBase) BinaryConcatArray(goArgs ...*Variant) *Variant {
	return this.ArrayConcat(goArgs...) // seams to be the same
}

func (this *ExchangeBase) InArray(what *Variant, obj *Variant) *Variant {
	pos := obj.IndexOf(what)
	return MkBool(pos.ToInt() != -1)
}

func (this *ExchangeBase) ToArray(obj *Variant) *Variant {
	if obj.Type == Array {
		return OpCopy(obj)
	} else if obj.Type == Map {
		return GoGetValues(obj)
	}
	return MkArray(&VarArray{obj})
}

func (this *ExchangeBase) IsEmpty(obj *Variant) *Variant {
	if obj.Type == Array {
		return MkBool(obj.Length.ToInt() == 0)
	} else if obj.Type == Map {
		return MkBool(len(*obj.ToMap()) == 0)
	}
	return MkBool(true)
}

func (this *ExchangeBase) ImplodeParams(str *Variant, params *Variant) *Variant {
	if params.Type != Map {
		return OpCopy(str)
	}
	newStr := str.ToStr()
	for k, v := range *(*params).ToMap() {
		newStr = strings.ReplaceAll(newStr, k, (*v).ToStr())
	}
	return MkString(newStr)
}

func (this *ExchangeBase) ImplodeHostname(url *Variant) *Variant {
	return this.ImplodeParams(url, MkMap(&VarMap{"hostname": *this.At(MkString("hostname"))}))
}

func (this *ExchangeBase) ExtractParams(v *Variant) *Variant {
	r := regexp.MustCompile(`{([\w-]+)}`)
	matches := r.FindAllString(v.ToStr(), -1)
	vArr := make(VarArray, len(matches))
	for i, c := range matches {
		vArr[i] = MkString(c)
	}
	return MkArray(&vArr)
}

func (this *ExchangeBase) Strip(v *Variant) *Variant {
	return MkString(strings.TrimSpace(v.ToStr()))
}

func (this *ExchangeBase) Keysort(v *Variant) *Variant {
	return v // todo - Sort a map by key in ascending order, problem go does not have sorted maps
}

func (this *ExchangeBase) Ordered(v *Variant) *Variant {
	return v // todo // a stub to keep assoc keys in order (in JS it does nothing, it's mostly for Python)
}

func (this *ExchangeBase) IndexBy(arr *Variant, key *Variant) *Variant {
	if arr.Type == Map {
		arr = GoGetValues(this.Keysort(arr))
	}

	vMap := VarMap{}
	for _, element := range *arr.ToArray() {
		if (*element).Has(key.ToStr()) {
			k := *((*element).At(key))
			vMap[k.ToStr()] = *element
		}
	}
	return MkMap(&vMap)
}

func (this *ExchangeBase) GroupBy(arr *Variant, key *Variant) *Variant {
	if arr.Type == Map {
		arr = this.ToArray(arr)
	}

	vMap := VarMap{}
	for _, element := range *arr.ToArray() {
		if (*element).Has(key.ToStr()) {
			k := *((*element).At(key))
			if vMap[k.ToStr()] == nil {
				vMap[k.ToStr()] = MkArray(&VarArray{})
			}
			vMap[k.ToStr()].Push(*element)
		}
	}
	return MkMap(&vMap)
}

func (this *ExchangeBase) FilterBy(arr *Variant, key *Variant, value *Variant) *Variant {
	if arr.Type == Map {
		arr = this.ToArray(arr)
	}

	vMap := VarMap{}
	for _, element := range *arr.ToArray() {
		if IsTrue(OpEq(*((*element).At(key)), value)) {
			k := *((*element).At(key))
			vMap[k.ToStr()] = *element
		}
	}
	return MkMap(&vMap)
}

func (this *ExchangeBase) SortBy(arr *Variant, key *Variant, v ...*Variant) *Variant {
	if arr.Type != Array {
		panic("Only arrays can do sort")
	}

	descending := GoGetArg(v, 0, MkBool(false))

	ret := arr.Slice(MkInteger(0)) // duplicate

	vArr := ret.Value.(*VVarArray)

	sort.SliceStable(*vArr, func(i, j int) bool {
		return IsTrue(OpLw((*(*((*vArr)[i])).At(key)), (*(*((*vArr)[j])).At(key)))) != IsTrue(descending)
	})

	return ret
}

func (this *ExchangeBase) Flatten(arr *Variant) *Variant {
	if arr.Type == Map {
		arr = GoGetValues(this.Keysort(arr))
	}

	ret := VarArray{}
	for _, val := range *arr.ToArray() {
		if (*val).Type == Array {
			ret = append(ret, this.Flatten(*val))
		} else {
			ret = append(ret, *val)
		}
	}
	return MkArray(&ret)
}

func (this *ExchangeBase) Sum(v ...*Variant) *Variant {
	var fVal float64 = 0
	var iVal int64 = 0
	isFloat := false
	for i := 0; i < len(v); i++ {
		if isFloat {
			fVal = +v[i].ToFloat()
		} else if v[i].Type == Number || (v[i].Type == String && strings.Index(v[i].ToStr(), ".") != -1) {
			isFloat = true
			fVal = float64(i) + v[i].ToFloat()
		} else {
			iVal += v[i].ToInt()
		}
	}
	if isFloat {
		return MkNumber(fVal)
	}
	return MkInteger(iVal)
}

func (this *ExchangeBase) NumberToString(goArgs ...*Variant) *Variant {
	// transpiled function
	x := GoGetArg(goArgs, 0, MkUndefined())
	if IsTrue(OpEq2(OpType(x), MkString("string"))) {
		return x
	}
	s := x.ToString()
	if IsTrue(OpLw(Math.Abs(x), MkNumber(1.0))) {
		n_e := s.Split(MkString("e-"))
		n := (*(n_e).At(MkInteger(0))).Replace(MkString("."), MkString(""))
		e := parseInt(*(n_e).At(MkInteger(1)))
		neg := OpEq2(*(s).At(MkInteger(0)), MkString("-"))
		if IsTrue(e) {
			x = OpAdd(OpTrinary(neg, MkString("-"), MkString("")), OpAdd(MkString("0."), OpAdd(MkString(utils.StrRepeat("0", int(e.ToInt())-1)), n.Substring(neg))))
			return x
		}
	} else {
		parts := s.Split(MkString("e"))
		if IsTrue(*(parts).At(MkInteger(1))) {
			e := parseInt(*(parts).At(MkInteger(1)))
			m := (*(parts).At(MkInteger(0))).Split(MkString("."))
			part := MkString("")
			if IsTrue(*(m).At(MkInteger(1))) {
				OpSubAssign(&e, *(*(m).At(MkInteger(1))).At(MkString("length")))
				part = *(m).At(MkInteger(1))
			}
			return OpAdd(*(m).At(MkInteger(0)), OpAdd(part, MkString(utils.StrRepeat("0", int(e.ToInt())))))
		}
	}
	return s
}

func (this *ExchangeBase) Capitalize(v *Variant) *Variant {
	return v.Capitalize()
}

func (this *ExchangeBase) FindBroadlyMatchedKey(v *Variant, str *Variant) *Variant {

	if v.Type != Map {
		panic("Only maps have keys")
	}

	sstr := str.ToStr()
	vMap := *v.Value.(*VVarMap)
	for key, value := range vMap {
		if strings.Index(key, sstr) != -1 {
			return *value
		}
	}
	return nil
}

func (this *ExchangeBase) Milliseconds() *Variant {
	return MkInteger(time.Now().UnixNano() / int64(time.Millisecond))
}
func (this *ExchangeBase) Microseconds() *Variant {
	return MkInteger(time.Now().UnixNano() / int64(time.Microsecond))
}
func (this *ExchangeBase) Seconds() *Variant {
	return MkInteger(time.Now().UnixNano() / int64(time.Second))
}
