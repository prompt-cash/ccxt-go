package ccxt

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/prompt-cash/ccxt-go/pkg/turing_parser"
	"sort"
	"strconv"
	"strings"
)

//
// Go is a strongly typed language while the other CCXT languages are not
// in order to make the transpilation work we use an own universal data type
// in order to make this work we need to convert all equations to chains of
// Op... calls that implement all the operators, it would be so convenient if
// go would allow for operator overloading, like c/c++ does, but it does not :(
//

type VariantType uint8

const (
	Undefined VariantType = iota
	Boolean
	String
	Number
	Integer
	Map
	Array
	Error
	// Precise
	VEnd
)

type Variant struct {
	Type   VariantType
	Value  interface{}
	Length *Variant // used for map, array, and string, for other types its nil
}

///////////////////////////////////////////////////////////////////////////////////////////
// variant makers

func MkUndefined() *Variant {
	return &Variant{
		Type:  Undefined,
		Value: nil,
	}
}

func MkBool(boolVal bool) *Variant {
	return &Variant{
		Type:  Boolean,
		Value: boolVal,
	}
}

func MkNumber(numVal float64) *Variant {
	return &Variant{
		Type:  Number,
		Value: numVal,
	}
}

func MkInteger(numVal int64) *Variant {
	return &Variant{
		Type:  Integer,
		Value: numVal,
	}
}

func MkString(strVal string) *Variant {
	return &Variant{
		Type:   String,
		Value:  strVal,
		Length: MkInteger(int64(len(strVal))),
	}
}

func MkStringRef(strVal string) **Variant {
	str := MkString(strVal)
	return &str
}

type VarMap map[string]*Variant
type VVarMap map[string]**Variant

func MkMap(varMap *VarMap) *Variant {

	vMap := make(VVarMap, len(*varMap))

	for key, value := range *varMap {
		if value == nil {
			panic("invalid value for mkMap")
		}

		var valuePtr *Variant
		valuePtr = value
		vMap[key] = &valuePtr
	}

	return &Variant{
		Type:   Map,
		Value:  &vMap,
		Length: MkInteger(int64(len(vMap))),
	}
}

type VarArray []*Variant
type VVarArray []**Variant

func MkArray(varArray *VarArray) *Variant {

	vArr := make(VVarArray, len(*varArray))

	i := 0
	for _, value := range *varArray {
		if value == nil {
			panic("invalid value for mkArray")
		}

		var valuePtr *Variant
		valuePtr = value
		vArr[i] = &valuePtr
		i++
	}

	return &Variant{
		Type:   Array,
		Value:  &vArr,
		Length: MkInteger(int64(len(vArr))),
	}
}

///////////////////////////////////////////////////////////////////////////////////////////
// variant to native

func (v *Variant) ToStr() string {

	switch v.Type {
	case Undefined:
		return "undefined"
	case Boolean:
		if v.Value.(bool) {
			return "true"
		}
		return "false"
	case String:
		return v.Value.(string)
	case Number:
		return fmt.Sprintf("%f", v.Value.(float64))
	case Integer:
		return fmt.Sprintf("%d", v.Value.(int64))
	//case Precise:
	//	break;
	case Map:
		return "[object Object]"
	case Array:
		return "[" + v.Join(MkString(", ")).ToStr() + "]"
	case Error:
		if len(v.Value.(*VError).Message) > 0 {
			return v.Value.(*VError).Message
		}
		return fmt.Sprintf("Exception: %s", v.Value.(*VError).Type)
	}
	panic("Unknown variant type")
}

func (v *Variant) ToInt() int64 {
	i, _ := v.ToSafeInt()
	return i
}

func (v *Variant) ToSafeInt() (int64, error) {
	if v.Type == Number {
		return int64(v.Value.(float64)), nil
	}
	if v.Type == Integer {
		return v.Value.(int64), nil
	}
	/*if v.Type == Precise {
		return
	}*/
	if v.Type == String {
		str := v.ToStr()
		i, err := strconv.ParseInt(str, 10, 64)
		return i, err
	}
	return 0, errors.New("NaN")
}

func (v *Variant) ToFloat() float64 {
	f, _ := v.ToSafeFloat()
	return f
}

func (v *Variant) ToSafeFloat() (float64, error) {
	if v.Type == Number {
		return v.Value.(float64), nil
	}
	if v.Type == Integer {
		return float64(v.Value.(int64)), nil
	}
	/*if v.Type == Precise {
		return
	}*/
	if v.Type == String {
		str := v.ToStr()
		f, err := strconv.ParseFloat(str, 64)
		return f, err
	}
	return 0, errors.New("NaN")
}

func (v *Variant) ToMap() *VVarMap {
	if v.Type != Map {
		panic("Variant is not a map")
	}
	return v.Value.(*VVarMap)
}

func (v *Variant) ToArray() *VVarArray {
	if v.Type != Array {
		panic("Variant is not a array")
	}
	return v.Value.(*VVarArray)
}

///////////////////////////////////////////////////////////////////////////////////////////
// variant functions

func (v *Variant) Call(vName *Variant, args ...*Variant) *Variant {

	if vName.Type != String {
		panic("Call name must be a string")
	}
	name := vName.Value.(string)

	if name == "toString" {
		return v.ToString()
	} else if name == "toFixed" {
		return v.ToFixed(args...)
	} else if name == "substring" {
		return v.Substring(args...)
	} else if name == "indexOf" {
		return v.IndexOf(args...)
	} else if name == "replace" {
		return v.Replace(args[0], args[1])
	} else if name == "capitalize" {
		return v.Capitalize()
	} else if name == "toUpperCase" {
		return v.ToUpperCase()
	} else if name == "toLowerCase" {
		return v.ToLowerCase()
	} else if name == "split" {
		return v.Split(args[0])
	} else if name == "join" {
		return v.Join(args[0])
	} else if name == "push" {
		return v.Push(args[0])
	} else if name == "concat" {
		return v.Concat(args[0])
	} else if name == "slice" {
		return v.Slice(args...)
	} else if name == "sort" {
		return v.Sort()
	} else {
		panic("Calling undefined variant function")
	}

	// Note: we only implement functions actually used much more could be add like pop, splice, etc....

	return nil
}

func (v *Variant) ToString() *Variant {
	return MkString(v.ToStr())
}

func (v *Variant) ToInteger() *Variant {
	return MkInteger(v.ToInt())
}

func (v *Variant) ToNumber() *Variant {
	return MkNumber(v.ToFloat())
}

func (v *Variant) ToFixed(args ...*Variant) *Variant {
	dec := 0
	if len(args) >= 1 {
		dec = int(args[0].ToInt())
	}
	return MkString(fmt.Sprintf("%."+strconv.Itoa(dec)+"f", v.ToFloat()))
}

func (v *Variant) Substring(args ...*Variant) *Variant {
	start := int(args[0].ToInt())
	length := -1
	if len(args) >= 2 {
		length = int(args[1].ToInt())
	}

	str := v.ToStr()

	var end = len(str)
	if length >= 0 && start+length <= end {
		end = start + length
	}

	return MkString(str[start:end])
}

func (v *Variant) Capitalize() *Variant {
	if v.Type != String {
		panic("Only strings can be Capitalizeed")
	}
	str := v.Value.(string)
	if len(str) == 0 {
		return v
	}
	return MkString(strings.ToUpper(string(str[0])) + str[1:])
}

func (v *Variant) ToUpperCase() *Variant {
	if v.Type != String {
		panic("Only strings can be Uppercased")
	}
	return MkString(strings.ToUpper(v.Value.(string)))
}

func (v *Variant) ToLowerCase() *Variant {
	if v.Type != String {
		panic("Only strings can be Lowercased")
	}
	return MkString(strings.ToLower(v.Value.(string)))
}

func (v *Variant) IndexOf(args ...*Variant) *Variant {

	what := args[0]
	from := 0
	if len(args) >= 2 {
		from = int(args[1].ToInt())
	}

	if v.Type == String {
		// add from support - not currently used
		str := what.ToStr()
		return MkInteger(int64(strings.Index(v.Value.(string), str)))
	}
	if v.Type != Array {
		panic("Only strings and arrays can do indexOf")
	}

	// find first occurence in array
	pos := -1
	arr := *v.Value.(*VVarArray)
	for i := from; i < len(arr); i++ {
		if IsTrue(OpEq(*arr[i], what)) {
			pos = i
			break
		}
	}
	return MkInteger(int64(pos))
}

func (v *Variant) Replace(what *Variant, with *Variant) *Variant {
	if v.Type != String {
		panic("Only strings can do replace")
	}

	str := what.ToStr()
	return MkString(strings.Replace(str, what.ToStr(), with.ToStr(), 1))
}

func (v *Variant) Push(val *Variant) *Variant {
	if v.Type != Array {
		panic("Only arrays can do push")
	}

	arr := *v.Value.(*VVarArray)
	tmp := OpCopy(val)
	arr = append(arr, &tmp)
	v.Value = &arr

	return val
}

func (v *Variant) Concat(val *Variant) *Variant {
	if v.Type != Array {
		panic("Only arrays can do concat")
	}

	ret := v.Slice()

	arr := *v.Value.(*VVarArray)

	for _, v := range *val.Value.(*VVarArray) {
		arr = append(arr, v)
	}
	v.Value = arr

	return ret
}

func (v *Variant) Join(sep *Variant) *Variant {
	if v.Type != Array {
		panic("Only arrays can do join")
	}

	sepStr := sep.ToStr()
	str := ""

	arr := *v.Value.(*VVarArray)
	for i := 0; i < len(arr); i++ {
		if i > 0 {
			str += sepStr
		}
		str += (*arr[i]).ToStr()
	}

	return MkString(str)
}

func (v *Variant) Split(sep *Variant) *Variant {
	if v.Type != String {
		panic("Only strings can do split")
	}

	sepStr := sep.ToStr()
	str := v.Value.(string)

	arr := strings.Split(str, sepStr)
	vArr := make(VarArray, len(arr))
	for i := 0; i < len(arr); i++ {
		vArr[i] = MkString(arr[i])
	}

	return MkArray(&vArr)
}

func (v *Variant) Slice(args ...*Variant) *Variant {

	start := 0
	end := 0

	if len(args) >= 1 {
		start = int(args[0].ToInt())
	}
	if len(args) >= 2 {
		end = int(args[1].ToInt())
	}

	if v.Type == Array {
		arr := *v.Value.(*VVarArray)
		var subArr VVarArray
		if end != 0 {
			subArr = arr[start:end]
		} else if start < 0 {
			subArr = arr[len(arr)+start:]
		} else {
			subArr = arr[start:]
		}
		return &Variant{
			Type:   Array,
			Value:  &subArr,
			Length: MkInteger(int64(len(subArr))),
		}
	} else {
		arr := v.ToStr()
		var subArr string
		if end != 0 {
			subArr = arr[start:end]
		} else if start < 0 {
			subArr = arr[len(arr)+start:]
		} else {
			subArr = arr[start:]
		}
		return MkString(subArr)
	}
}

func (v *Variant) Sort() *Variant {
	if v.Type != Array {
		panic("Only arrays can do sort")
	}

	ret := v.Slice(MkInteger(0)) // duplicate

	vArr := ret.Value.(*VVarArray)

	sort.SliceStable(*vArr, func(i, j int) bool {
		return IsTrue(OpLw((*(*vArr)[i]), (*(*vArr)[j])))
	})

	return ret
}

func GoGetKeys(v *Variant) *Variant {
	if v.Type != Map {
		panic("Only maps can do getKeys")
	}

	vMap := *v.Value.(*VVarMap)

	vArr := make(VarArray, len(vMap))
	i := 0
	for key := range vMap {
		vArr[i] = MkString(key)
		i++
	}
	return MkArray(&vArr)
}

func GoGetValues(v *Variant) *Variant {
	if v.Type == Array {
		return v.Slice() // return a copy
	}
	if v.Type != Map {
		panic("Only maps and arrays can do getValues")
	}

	vMap := *v.Value.(*VVarMap)

	vArr := make(VarArray, len(vMap))
	i := 0
	for _, value := range vMap {
		vArr[i] = *value
		i++
	}
	return MkArray(&vArr)
}

type GoVarFunc func(v *Variant) *Variant

func (v *Variant) Map(do GoVarFunc) *Variant {
	vArr := v.ToArray()
	vArray := make(VarArray, len(*vArr))
	for i, val := range *vArr {
		vArray[i] = do(*val)
	}
	return MkArray(&vArray)
}

func (v *Variant) Filter(do GoVarFunc) *Variant {
	vArr := v.ToArray()
	vArray := VarArray{}
	for _, val := range *vArr {
		if IsTrue(do(*val)) {
			vArray = append(vArray, *val)
		}
	}
	return MkArray(&vArray)
}

func GoIsArray(v *Variant) *Variant {
	if v.Type == Array {
		return MkBool(true)
	}
	return MkBool(false)
}

func (v *Variant) GetRef(at *Variant) **Variant {
	if v.Type == Array {
		index := int(at.ToInt())
		vArr := v.Value.(*VVarArray)
		for index >= len(*vArr) {
			new := MkUndefined()
			*vArr = append(*vArr, &new)
			OpIncr(&v.Length)
		}
		return (*vArr)[index]
	}
	if v.Type == Map {
		name := at.ToStr()
		vMap := v.Value.(*VVarMap)
		val, ok := (*vMap)[name]
		if !ok {
			new := MkUndefined()
			val = &new
			(*vMap)[name] = val
			OpIncr(&v.Length)
		}
		return val
	}
	return nil
}

func (v *Variant) At(at *Variant) **Variant {
	val := v.GetRef(at)
	if val == nil {
		tmp := MkUndefined()
		val = &tmp
	}
	return val
}

func GoGetArg(v []*Variant, index int, def *Variant) *Variant {
	if len(v) <= index || v[index].Type == VEnd {
		return def
	}
	return v[index]
}

///////////////////////////////////////////////////////////////////////////////////////////
// variant operators

func OpNot(v1 *Variant) *Variant {
	return MkBool(!IsTrue(v1))
}

func OpIncr(v1 **Variant) *Variant { // i++
	old := OpCopy(*v1)
	*v1 = MkInteger(old.ToInt() + 1)
	return old
}

func OpDecr(v1 **Variant) *Variant { // i--
	old := OpCopy(*v1)
	*v1 = MkInteger(old.ToInt() - 1)
	return old
}

func OpIncr2(v1 **Variant) *Variant { // ++i
	*v1 = MkInteger((*v1).ToInt() + 1)
	return *v1
}

func OpDecr2(v1 **Variant) *Variant { // --i
	*v1 = MkInteger((*v1).ToInt() - 1)
	return *v1
}

func OpNeg(v1 *Variant) *Variant {
	if v1.Type == Integer {
		return MkInteger(-v1.Value.(int64))
	}
	if v1.Type == Number {
		return MkNumber(-v1.Value.(float64))
	}
	return MkNumber(v1.ToFloat())
}

func OpType(v1 *Variant) *Variant {
	return MkString(v1.GetType())
}

func OpIsType(v1 *Variant, v2 *Variant) *Variant {
	return OpEq(OpType(v2), v2)
}

func OpTrinary(v1 *Variant, v2 *Variant, v3 *Variant) *Variant {
	// todo: make this right replace arg 2 and 3 with lambdas and cal them
	if IsTrue(v1) {
		return v2
	}
	return v3
}

func OpAdd(v1 *Variant, v2 *Variant) *Variant {
	if v1.Type == String || v2.Type == String {
		return MkString(v1.ToStr() + v2.ToStr())
	}
	if v1.Type == Number && v2.Type == Number {
		return MkNumber(v1.Value.(float64) + v2.Value.(float64))
	}
	if v1.Type == Integer && v2.Type == Integer {
		return MkInteger(v1.Value.(int64) + v2.Value.(int64))
	}
	return MkNumber(v1.ToFloat() + v2.ToFloat())
}

func OpSub(v1 *Variant, v2 *Variant) *Variant {
	if v1.Type == Number && v2.Type == Number {
		return MkNumber(v1.Value.(float64) - v2.Value.(float64))
	}
	if v1.Type == Integer && v2.Type == Integer {
		return MkInteger(v1.Value.(int64) - v2.Value.(int64))
	}
	return MkNumber(v1.ToFloat() - v2.ToFloat())
}

func OpMulti(v1 *Variant, v2 *Variant) *Variant {
	if v1.Type == Number && v2.Type == Number {
		return MkNumber(v1.Value.(float64) * v2.Value.(float64))
	}
	if v1.Type == Integer && v2.Type == Integer {
		return MkInteger(v1.Value.(int64) * v2.Value.(int64))
	}
	return MkNumber(v1.ToFloat() * v2.ToFloat())
}

func OpDiv(v1 *Variant, v2 *Variant) *Variant {
	if v1.Type == Number && v2.Type == Number {
		return MkNumber(v1.Value.(float64) / v2.Value.(float64))
	}
	if v1.Type == Integer && v2.Type == Integer {
		return MkInteger(v1.Value.(int64) / v2.Value.(int64))
	}
	return MkNumber(v1.ToFloat() / v2.ToFloat())
}

func OpMod(v1 *Variant, v2 *Variant) *Variant {
	return MkInteger(v1.ToInt() % v2.ToInt()) // todo fix me java can do thts also with floats
}

func OpBitNot(v1 *Variant) *Variant {
	return MkInteger(int64(0xFFFFFFFFFFFFFFFF ^ uint64(v1.ToInt())))
}

func OpBitAnd(v1 *Variant, v2 *Variant) *Variant {
	return MkInteger(v1.ToInt() & v2.ToInt())
}

func OpBitOr(v1 *Variant, v2 *Variant) *Variant {
	return MkInteger(v1.ToInt() | v2.ToInt())
}

func OpBitXOr(v1 *Variant, v2 *Variant) *Variant {
	return MkInteger(v1.ToInt() ^ v2.ToInt())
}

func OpShiftLeft(v1 *Variant, v2 *Variant) *Variant {
	return MkInteger(v1.ToInt() << v2.ToInt())
}

func OpShiftRight(v1 *Variant, v2 *Variant) *Variant {
	return MkInteger(v1.ToInt() >> v2.ToInt())
}

func OpAnd(v1 *Variant, v2 *Variant) *Variant {
	return MkBool(IsTrue(v1) && IsTrue(v2))
}

func OpOr(v1 *Variant, v2 *Variant) *Variant {
	return MkBool(IsTrue(v1) || IsTrue(v2))
}

func OpEq(v1 *Variant, v2 *Variant) *Variant {
	if v1.Type != v2.Type {
		if (v1.Type == Integer || v1.Type == Number) && (v2.Type == Integer || v2.Type == Number) {
			return MkBool(v1.ToFloat() == v2.ToFloat())
		}
		return MkBool(v1.ToStr() == v2.ToStr())
	}

	switch v1.Type {
	case Undefined:
		return MkBool(true)
	case Boolean:
		return MkBool(v1.Value.(bool) == v2.Value.(bool))
	case String:
		return MkBool(v1.Value.(string) == v2.Value.(string))
	case Number:
		return MkBool(v1.Value.(float64) == v2.Value.(float64))
	case Integer:
		return MkBool(v1.Value.(int64) == v2.Value.(int64))
		//case Precise:
		//	break;
	}
	return MkBool(v1.Value == v2.Value)
}

func OpNotEq(v1 *Variant, v2 *Variant) *Variant {
	return OpNot(OpEq(v1, v2))
}

func OpEq2(v1 *Variant, v2 *Variant) *Variant {
	if v1.Type != v2.Type {
		return MkBool(false)
	}
	return OpEq(v1, v2)
}

func OpNotEq2(v1 *Variant, v2 *Variant) *Variant {
	if v1.Type != v2.Type {
		return MkBool(true)
	}
	return OpNotEq(v1, v2)
}

func OpLw(v1 *Variant, v2 *Variant) *Variant {
	if v1.Type != v2.Type {
		if (v1.Type == Integer || v1.Type == Number) && (v2.Type == Integer || v2.Type == Number) {
			return MkBool(v1.ToFloat() < v2.ToFloat())
		}
		return MkBool(v1.ToStr() < v2.ToStr())
	}

	switch v1.Type {
	//case Boolean:
	//	return MkBool(*v1.Value.(*bool) < *v2.Value.(*bool))
	case String:
		return MkBool(v1.Value.(string) < v2.Value.(string))
	case Number:
		return MkBool(v1.Value.(float64) < v2.Value.(float64))
	case Integer:
		return MkBool(v1.Value.(int64) < v2.Value.(int64))
		//case Precise:
		//	break;
	}
	return MkBool(false)
}

func OpGt(v1 *Variant, v2 *Variant) *Variant {
	if v1.Type != v2.Type {
		if (v1.Type == Integer || v1.Type == Number) && (v2.Type == Integer || v2.Type == Number) {
			return MkBool(v1.ToFloat() > v2.ToFloat())
		}
		return MkBool(v1.ToStr() > v2.ToStr())
	}

	switch v1.Type {
	//case Boolean:
	//	return MkBool(*v1.Value.(*bool) > *v2.Value.(*bool))
	case String:
		return MkBool(v1.Value.(string) > v2.Value.(string))
	case Number:
		return MkBool(v1.Value.(float64) > v2.Value.(float64))
	case Integer:
		return MkBool(v1.Value.(int64) > v2.Value.(int64))
		//case Precise:
		//	break;
	}
	return MkBool(false)
}

func OpLwEq(v1 *Variant, v2 *Variant) *Variant {
	if v1.Type != v2.Type {
		if (v1.Type == Integer || v1.Type == Number) && (v2.Type == Integer || v2.Type == Number) {
			return MkBool(v1.ToFloat() <= v2.ToFloat())
		}
		return MkBool(v1.ToStr() <= v2.ToStr())
	}

	switch v1.Type {
	//case Boolean:
	//	return MkBool(*v1.Value.(*bool) <= *v2.Value.(*bool))
	case String:
		return MkBool(v1.Value.(string) <= v2.Value.(string))
	case Number:
		return MkBool(v1.Value.(float64) <= v2.Value.(float64))
	case Integer:
		return MkBool(v1.Value.(int64) <= v2.Value.(int64))
		//case Precise:
		//	break;
	}
	return OpEq(v1, v2)
}

func OpGtEq(v1 *Variant, v2 *Variant) *Variant {
	if v1.Type != v2.Type {
		if (v1.Type == Integer || v1.Type == Number) && (v2.Type == Integer || v2.Type == Number) {
			return MkBool(v1.ToFloat() >= v2.ToFloat())
		}
		return MkBool(v1.ToStr() >= v2.ToStr())
	}

	switch v1.Type {
	//case Boolean:
	//	return MkBool(*v1.Value.(*bool) >= *v2.Value.(*bool))
	case String:
		return MkBool(v1.Value.(string) >= v2.Value.(string))
	case Number:
		return MkBool(v1.Value.(float64) >= v2.Value.(float64))
	case Integer:
		return MkBool(v1.Value.(int64) >= v2.Value.(int64))
		//case Precise:
		//	break;
	}
	return OpEq(v1, v2)
}

func OpCopy(v1 *Variant) *Variant {

	switch v1.Type {
	case Undefined:
		return MkUndefined()
	case Boolean:
		return MkBool(v1.Value.(bool))
	case String:
		return MkString(v1.Value.(string))
	case Number:
		return MkNumber(v1.Value.(float64))
	case Integer:
		return MkInteger(v1.Value.(int64))
		//case Precise:
		//	break;
	}
	return v1 // return the poitner to assign the pointer
}

func OpAssign(v1 **Variant, v2 *Variant) *Variant {
	*v1 = OpCopy(v2)
	return *v1
}

func OpAddAssign(v1 **Variant, v2 *Variant) *Variant {
	tmp := OpAdd(*v1, v2)
	*v1 = tmp
	return tmp
}

func OpSubAssign(v1 **Variant, v2 *Variant) *Variant {
	tmp := OpSub(*v1, v2)
	*v1 = tmp
	return tmp
}

func OpMultiAssign(v1 **Variant, v2 *Variant) *Variant {
	tmp := OpMulti(*v1, v2)
	*v1 = tmp
	return tmp
}

func OpDivAssign(v1 **Variant, v2 *Variant) *Variant {
	tmp := OpDiv(*v1, v2)
	*v1 = tmp
	return tmp
}

func OpModAssign(v1 **Variant, v2 *Variant) *Variant {
	tmp := OpMod(*v1, v2)
	*v1 = tmp
	return tmp
}

func OpOrAssign(v1 **Variant, v2 *Variant) *Variant {
	tmp := OpOr(*v1, v2)
	*v1 = tmp
	return tmp
}

func OpAndAssign(v1 **Variant, v2 *Variant) *Variant {
	tmp := OpAnd(*v1, v2)
	*v1 = tmp
	return tmp
}

func OpBitOrAssign(v1 **Variant, v2 *Variant) *Variant {
	tmp := OpBitOr(*v1, v2)
	*v1 = tmp
	return tmp
}

func OpBitAndAssign(v1 **Variant, v2 *Variant) *Variant {
	tmp := OpBitAnd(*v1, v2)
	*v1 = tmp
	return tmp
}

func OpBitXorAssign(v1 **Variant, v2 *Variant) *Variant {
	tmp := OpBitXOr(*v1, v2)
	*v1 = tmp
	return tmp
}

func OpShiftLeftAssign(v1 **Variant, v2 *Variant) *Variant {
	tmp := OpShiftLeft(*v1, v2)
	*v1 = tmp
	return tmp
}

func OpShiftRightAssign(v1 **Variant, v2 *Variant) *Variant {
	tmp := OpShiftRight(*v1, v2)
	*v1 = tmp
	return tmp
}

func OpHasMember(v1 *Variant, v2 *Variant) *Variant {
	return MkBool(v2.Has(v1.ToStr()))
}

///////////////////////////////////////////////////////////////////////////////////////////
// variant helpers

func IsTrue(v *Variant) bool {
	switch v.Type {
	case Undefined:
		return false
	case Boolean:
		return v.Value.(bool)
	case String:
		return len(v.Value.(string)) > 0
	case Number:
		return v.Value.(float64) != 0
	case Integer:
		return v.Value.(int64) != 0.0
	//case Precise:
	//	break;
	case Map:
		return true
	case Array:
		return true
	}
	return false
}

func (v *Variant) GetType() string {
	switch v.Type {
	case Undefined:
		return "undefined"
	case Boolean:
		return "boolean"
	case String:
		return "string"
	case Number:
		return "number"
	case Integer:
		return "number"
	//case Precise:
	//	break;
	case Error:
		return v.Value.(*VError).Type
	}
	return "object"
}

func (v *Variant) Has(str string) bool {
	if v.Type != Map {
		return false
	}
	vMap := *v.Value.(*VVarMap)
	_, ok := vMap[str]
	return ok
}

func ArrayUnpack(a *Variant, v ...**Variant) { // pArray, pVar1, pVar2, ... pVarN
	if a.Type != Array {
		panic("Invalid Unpack argument")
	}

	arr := *a.Value.(*VVarArray)
	for i := 0; i < len(v); i++ {
		if len(arr) >= i {
			*v[i] = MkUndefined()
		} else {
			*v[i] = *arr[i]
		}
	}
}

func ObjectUnpack(o *Variant, v ...**Variant) { // pObject, pVar1Name, pVar1, pVar2Name, pVar2, ... , pVarNName, pVarN
	if o.Type != Map {
		panic("Invalid Unpack argument")
	}

	vMap := *o.Value.(*VarMap)
	for i := 0; i < len(v); i += 2 {
		name := (*(v[i])).ToStr()
		*v[i+1] = vMap[name]
	}
}

func parseInt(v *Variant) *Variant {
	str := v.ToStr()
	i, _ := strconv.ParseInt(str, 10, 64)
	return MkInteger(i)
}

func parseFloat(v *Variant) *Variant {
	str := v.ToStr()
	f, _ := strconv.ParseFloat(str, 64)
	return MkNumber(f)
}

///////////////////////////////////////////////////////////////////////////////////////////
// variant operators

/*type ItfMap map[string]interface{}

func MkIMap(itfMap *ItfMap) *Variant {
	return &Variant{
		Type:  Map,
		Value: nil, // todo
	}
}

type ItfArray []interface{}

func MkIArray(itfArray *ItfArray) *Variant {
	return &Variant{
		Type:  Array,
		Value: nil, // todo
	}
}*/

func ItfToVariant(itf interface{}) *Variant {

	switch typ := itf.(type) {
	case int:
		return MkInteger(int64(itf.(int)))
	case float64:
		return MkNumber(itf.(float64))
	case string:
		return MkString(itf.(string))
	case bool:
		return MkBool(itf.(bool))
	case map[string]interface{}:
		vMap := make(VVarMap, len(itf.(map[string]interface{})))
		for key, value := range itf.(map[string]interface{}) {
			var valuePtr *Variant
			valuePtr = ItfToVariant(value)
			vMap[key] = &valuePtr
		}
		return &Variant{
			Type:   Map,
			Value:  &vMap,
			Length: MkInteger(int64(len(vMap))),
		}
	case []interface{}:
		vArr := make(VVarArray, len(itf.([]interface{})))
		i := 0
		for _, value := range itf.([]interface{}) {
			var valuePtr *Variant
			valuePtr = ItfToVariant(value)
			vArr[i] = &valuePtr
			i++
		}
		return &Variant{
			Type:   Array,
			Value:  &vArr,
			Length: MkInteger(int64(len(vArr))),
		}
	case nil:
		return MkUndefined()
	default:
		panic(fmt.Sprintf("Unknown type in json: %s", typ))
		return MkUndefined()
	}
}

func VariantFromJson(data []byte) *Variant {
	var result map[string]interface{}
	err := json.Unmarshal(data, &result)
	if err != nil {
		return nil
	}
	variant := ItfToVariant(result)
	return variant
}

func VariantToItf(v *Variant) interface{} {

	switch v.Type {
	case Boolean:
		return v.Value.(bool)
	case String:
		return turing_parser.EscapeString(v.Value.(string))
	case Number:
		return v.Value.(float64)
	case Integer:
		return int(v.Value.(int64))
	//case Precise:
	//	break;
	case Map:
		iMap := make(map[string]interface{}, len(*v.ToMap()))
		for key, value := range *v.ToMap() {
			iMap[key] = VariantToItf(*value)
		}
		return iMap
	case Array:
		iArr := make([]interface{}, len(*v.ToArray()))
		i := 0
		for _, value := range *v.ToArray() {
			iArr[i] = VariantToItf(*value)
			i++
		}
		return iArr
	}
	return nil
}

func VariantToJson(v *Variant) []byte {
	obj := VariantToItf(v)
	data, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}
	return data
}
