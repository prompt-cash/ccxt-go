package ccxt

//
// variouse other Js stuff
//

type JSONStruct struct {
}

var JSON JSONStruct

func (m *JSONStruct) Parse(data *Variant) *Variant {
	v := VariantFromJson([]byte(data.ToStr()))
	if v == nil {
		return MkMap(&VarMap{})
	}
	return v
}
