package ccxt

import (
	"fmt"
	"reflect"
)

// exchange base

//
// This file implements the most basic aspects of our exchange base, it implements the call facilities for
// "virtual functions" and the means of retreiving properties form the values member, all this.something
// accesses are refactored to *this.At("something") and mapped ontu out values Variant
//

type Exchange interface {
}

type ExchangeBase struct {
	values *Variant

	// our "virtual" functions
	itf interface{}
}

var Exchanges = make([]Exchange, 0)

var _ Exchange = (*ExchangeBase)(nil)

func (this *ExchangeBase) At(at *Variant) **Variant {
	return this.values.At(at)
}

func RCall(name string, itf interface{}, args ...*Variant) *Variant {
	baseType := reflect.TypeOf(itf)
	for i := 0; i < baseType.NumMethod(); i++ {
		method := baseType.Method(i)
		if name == method.Name {
			in := make([]reflect.Value, len(args))
			for k, param := range args {
				in[k] = reflect.ValueOf(param)
			}
			var res []reflect.Value
			/*temp := reflect.ValueOf(itf).MethodByName(name)
			x1 := reflect.ValueOf(temp).FieldByName("flag").Uint()*/
			res = reflect.ValueOf(itf).MethodByName(name).Call(in)
			return res[0].Interface().(*Variant)
		}
	}
	return nil
}

// Note: we implement virtual functions by adding a VCall to the begin of every function we make virtual
// if the VCall finds a function on the "derived" interface it calls it instead and returns the return value
// the virtual function terminates when a return value is passed, else it executes its own implementation instead

func (this *ExchangeBase) VCall(name string, args ...*Variant) *Variant {

	// Problem: we can not distinguish functions on this from functions on this.itf
	// we want howeever only to call functions on this.itf but not on this
	// hence as a workaround we mark a call with the VEnd such that we will not end up in an endless loop
	// a call hence wil look like for example this.SetHeaders -> VCall -> this.SetHeaders -> VCall returns nill
	// and the VCall to SetHeaders does its job
	// a call to a function which is present on interface which does not invoke VCall looks like
	// this.Sign -> VCall -> this.itf.Sign and thats it
	if len(args) >= 1 && args[len(args)-1].Type == VEnd {
		return nil
	}
	args = append(args, &Variant{Type: VEnd})

	ret := RCall(name, this.itf, args...)
	//if ret == nil {
	//	ret = RCall(name + "Impl", this, args...)
	//}
	return ret
}

func (this *ExchangeBase) Call(vName *Variant, args ...*Variant) *Variant {

	if vName.Type != String {
		panic("Call name must be a string")
	}
	name := vName.Value.(string)

	ret := this.TryCallAPI(name, args...)

	//if ret == nil {
	//	ret = this.VCall(name, args...)
	//}

	if ret == nil {
		panic(fmt.Sprintf("Function %s is not implemented", name))
	}

	return ret
}

///////////////////////////////////////////////////////////////////////////////////////////////
// helper functions

func (this *ExchangeBase) Id() string {
	return (*this.At(MkString("id"))).ToStr()
}

///////////////////////////////////////////////////////////////////////////////////////////////
// misc functions

func (this *ExchangeBase) Extend(v ...*Variant) *Variant {
	vMap := VarMap{}
	for i := 0; i < len(v); i++ {
		if v[i].Type != Map {
			continue
		}
		vvMap := v[i].ToMap()
		for key, value := range *vvMap {
			vMap[key] = *value
		}
	}
	return MkMap(&vMap)
}

func (this *ExchangeBase) DeepExtend(v ...*Variant) *Variant {
	vMap := VarMap{}
	for i := 0; i < len(v); i++ {
		if v[i].Type != Map {
			continue
		}
		vvMap := v[i].ToMap()
		for key, value := range *vvMap {
			vCur, ok := vMap[key]
			if !ok || (*value).Type != Map || vCur.Type != Map { // if the object does not yet contain this key add it, of if the new value is not an object overwrite it
				vMap[key] = *value
			} else { // if it does, merge the keys
				vMap[key] = this.DeepExtend(vCur, *value)
			}
		}
	}
	return MkMap(&vMap)
}
