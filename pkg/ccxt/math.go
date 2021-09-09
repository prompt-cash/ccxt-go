package ccxt

import "math"

//
// basic port of javascripts Math object
//

type MathStruct struct {
}

var Math MathStruct

func (m *MathStruct) Min(v1 *Variant, v2 *Variant) *Variant {
	if IsTrue(OpLw(v1, v2)) {
		return v1
	}
	return v2
}

func (m *MathStruct) Max(v1 *Variant, v2 *Variant) *Variant {
	if IsTrue(OpGt(v1, v2)) {
		return v1
	}
	return v2
}

func (m *MathStruct) Floor(v *Variant) *Variant {
	return MkNumber(math.Floor(v.ToFloat()))
}

func (m *MathStruct) Abs(v *Variant) *Variant {
	return MkNumber(math.Abs(v.ToFloat()))
}

func (m *MathStruct) Pow(v *Variant, to *Variant) *Variant {
	return MkNumber(math.Pow(v.ToFloat(), to.ToFloat()))
}

func (m *MathStruct) Log10(v *Variant) *Variant {
	return MkNumber(math.Log10(v.ToFloat()))
}

func (m *MathStruct) Round(v *Variant) *Variant {
	return MkNumber(math.Round(v.ToFloat()))
}

// todo: add everythign else although its not yet used
