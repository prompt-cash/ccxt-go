package transpiler

import (
	"fmt"
	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/prompt-cash/ccxt-go/log"
	"github.com/prompt-cash/ccxt-go/pkg/turing_parser"
	"strings"
)

type GoGenerator struct {
	//StrictVariance bool
	ForceUsage bool

	BaseMethods map[string]bool
	//BaseMembers map[string]bool

	logger log.Logger
}

func NewGoGenerator(baseMethods map[string]bool, logger log.Logger) (*GoGenerator, error) {
	generator := &GoGenerator{
		ForceUsage: false,
		logger: logger.WithFields(log.Fields{
			"module": "go_gen",
		}),
	}

	generator.BaseMethods = baseMethods
	//generator.StrictVariance = true
	//if !g.StrictVariance {
	//	panic("g.StrictVariance not set -> MkIMap and MkIArray are not yet implemented")
	//}
	generator.ForceUsage = true

	return generator, nil
}

func (g *GoGenerator) Genrate(js *arraylist.List, goFile *GoFile) string {

	thisName := strings.Title(goFile.Package)
	goFile.ThisName = thisName

	goFile.OwnFunctions = map[string]bool{}

	str := "package ccxt\n\n"

	str += "import (\n"
	for i := 0; i < len(goFile.Imports); i++ {
		str += "\t\"" + goFile.Imports[i] + "\"\n"
	}
	str += ")\n\n"

	str += "type " + thisName + " struct {\n"
	str += "\t*ExchangeBase\n"
	str += "}\n\n"
	str += "var _ Exchange = (*" + thisName + ")(nil)\n\n"
	str += "func init() {\n"
	str += "\texchange := &" + thisName + "{}\n"
	str += "\tExchanges = append(Exchanges, exchange)\n"
	str += "}\n\n"

	for i := 0; i < js.Size(); i++ {
		op_, _ := js.Get(i)
		op := op_.(*turing_parser.JSOpcode)

		if op.Type == turing_parser.JsFunction {
			goFile.OwnFunctions[strings.Title(op.Value.(*turing_parser.JSFunction).Name)] = true
		}
	}

	for i := 0; i < js.Size(); i++ {
		op_, _ := js.Get(i)
		op := op_.(*turing_parser.JSOpcode)
		str += g.BuildGOBlock(op, goFile, 1)
	}

	return str
}

func (g *GoGenerator) FixGoName(name string) string {

	if name == "undefined" || name == "null" {
		return "MkUndefined()"
	}
	//if g.StrictVariance {
	if name == "false" {
		return "MkBool(false)"
	}
	if name == "true" {
		return "MkBool(true)"
	}
	//}
	if name == "Object.keys" {
		return "GoGetKeys"
	}
	if name == "Object.values" {
		return "GoGetValues"
	}
	if name == "Array.isArray" {
		return "GoIsArray"
	}

	if name == "super.describe" {
		return "this.BaseDescribe"
	}

	names := strings.Split(name, ".")

	if names[0] == "type" {
		names[0] = "type_"
	} else if names[0] == "super" {
		names[0] = "this"
	}

	for i := 1; i < len(names); i++ {
		names[i] = strings.Title(names[i])
	}

	return strings.Join(names, ".")
}

func (g *GoGenerator) FixGoString(str string) string {

	if len(str) == 0 {
		return str
	}

	raw := str[1 : len(str)-1]
	raw = strings.Replace(raw, "\"", "\\\"", -1)
	raw = strings.Replace(raw, "\\0", "\\000", -1)
	return "\"" + raw + "\""
}

func IsString(str string) bool {
	return len(str) > 0 && (str[0] == '"' || str[0] == '\'')
}

func IsNumber(str string) bool {
	return len(str) > 0 && turing_parser.IsNumber(str[0])
}

func (g *GoGenerator) GoRebuildArgument(arg *turing_parser.JSOpcode) *turing_parser.JSOpcode {
	if arg.Type != turing_parser.JsEquation {
		return arg
	}
	return g.GoRebuildEquation(arg.Value.(*turing_parser.JSEquation).Params, 0)
}

func (g *GoGenerator) GoNeedRebuildOne(oeq *arraylist.List, from int) bool {
	if from == 2 {
		prev_, _ := oeq.Get(from - 1)
		prev := prev_.(*turing_parser.JSOpcode)
		if prev.Type == turing_parser.JsOperator && prev.Value.(*turing_parser.JSOperator).Op == "=" {
			return true
		}
	}
	return false
}

func (g *GoGenerator) GoRebuildOne(oeq *arraylist.List, from int) *turing_parser.JSOpcode {

	if oeq.Size() == from+1 {
		only_, _ := oeq.Get(from)
		only := only_.(*turing_parser.JSOpcode)
		if only.Type == turing_parser.JsVarName {

			arg := &turing_parser.JSEquation{Params: &arraylist.List{}}
			arg.Params.Add(only)

			args := &arraylist.List{}
			args.Add(arg)

			fx := &turing_parser.JSOpcode{Type: turing_parser.JsVarName, Value: &turing_parser.JSVarName{Name: "OpCopy"}}
			jsCall := &turing_parser.JSFuncCall{Func: fx, Arguments: args}

			return &turing_parser.JSOpcode{Type: turing_parser.JsFuncCall, Value: jsCall}
		}
	}
	return nil
}

func (g *GoGenerator) GoRebuildEquation(oeq *arraylist.List, from int) *turing_parser.JSOpcode {

	// ensure proper variabel assignment when we leave the assign '=' not replaced with OpAssign
	if g.GoNeedRebuildOne(oeq, from) {
		tmp := g.GoRebuildOne(oeq, from)
		if tmp != nil {
			return tmp
		}
	}

	if oeq.Size() == from+1 {
		only_, _ := oeq.Get(from)
		only := only_.(*turing_parser.JSOpcode)
		return g.GoRebuildArgument(only)
	}

	// case 1: a = b + c - d * ...
	// case 2: !t, --i, ~128
	// case 3: i++
	// case 4: a ? b : c

	first_, _ := oeq.Get(from + 0)
	first := first_.(*turing_parser.JSOpcode)

	second_, _ := oeq.Get(from + 1)
	second := second_.(*turing_parser.JSOpcode)

	if first.Type == turing_parser.JsOperator {
		// case 2: !t, --i, ~128

		op := first.Value.(*turing_parser.JSOperator)
		opFunc := "Func" + op.Op
		opAssign := false
		if op.Op == "!" {
			opFunc = "OpNot"
		} else if op.Op == "--" {
			opFunc = "OpDecr2"
			opAssign = true
		} else if op.Op == "++" {
			opFunc = "OpIncr2"
			opAssign = true
		} else if op.Op == "~" {
			opFunc = "OpBitNot"
		} else if op.Op == "-" {
			opFunc = "OpNeg"
		} else if op.Op == "typeof" {
			opFunc = "OpType"
		}

		arg1 := &turing_parser.JSEquation{Params: &arraylist.List{}}
		if opAssign {
			arg1.Params.Add(&turing_parser.JSOpcode{Type: turing_parser.JsOperator, Value: &turing_parser.JSOperator{Op: "&"}})
			arg1.Params.Add(g.GoRebuildArgument(second))
		} else {
			arg1.Params.Add(g.GoRebuildArgument(second))
		}

		args := &arraylist.List{}
		args.Add(arg1)

		fx := &turing_parser.JSOpcode{Type: turing_parser.JsVarName, Value: &turing_parser.JSVarName{Name: opFunc}}
		jsCall := &turing_parser.JSFuncCall{Func: fx, Arguments: args}

		return &turing_parser.JSOpcode{Type: turing_parser.JsFuncCall, Value: jsCall}
	}

	if oeq.Size() == from+2 && second.Type == turing_parser.JsOperator {
		// case 3: i++

		op := second.Value.(*turing_parser.JSOperator)
		opFunc := "Func" + op.Op
		if op.Op == "--" {
			opFunc = "OpDecr"
		} else if op.Op == "++" {
			opFunc = "OpIncr"
		}

		arg1 := &turing_parser.JSEquation{Params: &arraylist.List{}}
		arg1.Params.Add(&turing_parser.JSOpcode{Type: turing_parser.JsOperator, Value: &turing_parser.JSOperator{Op: "&"}})
		arg1.Params.Add(g.GoRebuildArgument(first))

		args := &arraylist.List{}
		args.Add(arg1)

		fx := &turing_parser.JSOpcode{Type: turing_parser.JsVarName, Value: &turing_parser.JSVarName{Name: opFunc}}
		jsCall := &turing_parser.JSFuncCall{Func: fx, Arguments: args}

		return &turing_parser.JSOpcode{Type: turing_parser.JsFuncCall, Value: jsCall}
	}

	if oeq.Size() == from+5 && second.Type == turing_parser.JsOperator && second.Value.(*turing_parser.JSOperator).Op == "?" {
		// case 4: a ? b : c

		opFunc := "OpTrinary"

		onTrue_, _ := oeq.Get(from + 2)
		onTrue := onTrue_.(*turing_parser.JSOpcode)

		onFalse_, _ := oeq.Get(from + 4)
		onFalse := onFalse_.(*turing_parser.JSOpcode)

		arg1 := &turing_parser.JSEquation{Params: &arraylist.List{}}
		arg1.Params.Add(g.GoRebuildArgument(first))
		arg2 := &turing_parser.JSEquation{Params: &arraylist.List{}}
		arg2.Params.Add(g.GoRebuildArgument(onTrue))
		arg3 := &turing_parser.JSEquation{Params: &arraylist.List{}}
		arg3.Params.Add(g.GoRebuildArgument(onFalse))
		// todo: note this is technically not corrent as we always execute booth code paths
		// we should only execute the true path, idealy we shoudl wrap the two paths in lambda functions !!!

		args := &arraylist.List{}
		args.Add(arg1)
		args.Add(arg2)
		args.Add(arg3)

		fx := &turing_parser.JSOpcode{Type: turing_parser.JsVarName, Value: &turing_parser.JSVarName{Name: opFunc}}
		jsCall := &turing_parser.JSFuncCall{Func: fx, Arguments: args}

		return &turing_parser.JSOpcode{Type: turing_parser.JsFuncCall, Value: jsCall}
	}

	// case 1: a = b + c - d * ...

	if second.Type != turing_parser.JsOperator || oeq.Size() < from+3 {
		g.logger.Errorf("Encountered problem when rebuilding equation!")
		fx := &turing_parser.JSOpcode{Type: turing_parser.JsVarName, Value: &turing_parser.JSVarName{Name: "problem"}}
		jsCall := &turing_parser.JSFuncCall{Func: fx, Arguments: &arraylist.List{}}
		return &turing_parser.JSOpcode{Type: turing_parser.JsFuncCall, Value: jsCall}
		return nil // todo log error
	}

	op := second.Value.(*turing_parser.JSOperator)
	opFunc := "Func" + op.Op
	opAssign := false
	obFirstStr := false
	if op.Op == "+" {
		opFunc = "OpAdd"
	} else if op.Op == "-" {
		opFunc = "OpSub"
	} else if op.Op == "*" {
		opFunc = "OpMulti"
	} else if op.Op == "/" {
		opFunc = "OpDiv"
	} else if op.Op == "%" {
		opFunc = "OpMod"
	} else if op.Op == "&" {
		opFunc = "OpBitAnd"
	} else if op.Op == "|" {
		opFunc = "OpBitOr"
	} else if op.Op == "^" {
		opFunc = "OpBitXOr"
	} else if op.Op == "<<" {
		opFunc = "OpShiftLeft"
	} else if op.Op == ">>" {
		opFunc = "OpShiftRight"
	} else if op.Op == "&&" {
		opFunc = "OpAnd"
	} else if op.Op == "||" {
		opFunc = "OpOr"
	} else if op.Op == "==" {
		opFunc = "OpEq"
	} else if op.Op == "!=" {
		opFunc = "OpNotEq"
	} else if op.Op == "===" {
		opFunc = "OpEq2"
	} else if op.Op == "!==" {
		opFunc = "OpNotEq2"
	} else if op.Op == "<" {
		opFunc = "OpLw"
	} else if op.Op == ">" {
		opFunc = "OpGt"
	} else if op.Op == "<=" {
		opFunc = "OpLwEq"
	} else if op.Op == ">=" {
		opFunc = "OpGtEq"
	} else if op.Op == "=" {
		opFunc = "OpAssign"
		opAssign = true
	} else if op.Op == "+=" {
		opFunc = "OpAddAssign"
		opAssign = true
	} else if op.Op == "-=" {
		opFunc = "OpSubAssign"
		opAssign = true
	} else if op.Op == "*=" {
		opFunc = "OpMultiAssign"
		opAssign = true
	} else if op.Op == "/=" {
		opFunc = "OpDivAssign"
		opAssign = true
	} else if op.Op == "%=" {
		opFunc = "OpModAssign"
		opAssign = true
	} else if op.Op == "||=" {
		opFunc = "OpOrAssign"
		opAssign = true
	} else if op.Op == "&&=" {
		opFunc = "OpAndAssign"
		opAssign = true
	} else if op.Op == "|=" {
		opFunc = "OpBitOrAssign"
		opAssign = true
	} else if op.Op == "&=" {
		opFunc = "OpBitAndAssign"
		opAssign = true
	} else if op.Op == "^=" {
		opFunc = "OpBitXorAssign"
		opAssign = true
	} else if op.Op == "<<=" {
		opFunc = "OpShiftLeftAssign"
		opAssign = true
	} else if op.Op == ">>=" {
		opFunc = "OpShiftRightAssign"
		opAssign = true
	} else if op.Op == "in" {
		opFunc = "OpHasMember"
		//obFirstStr = true
	} else if op.Op == "instanceof" {
		opFunc = "OpIsType"
	}

	arg1 := &turing_parser.JSEquation{Params: &arraylist.List{}}
	if opAssign {
		arg1.Params.Add(&turing_parser.JSOpcode{Type: turing_parser.JsOperator, Value: &turing_parser.JSOperator{Op: "&"}})
		arg1.Params.Add(g.GoRebuildArgument(first))
	} else if obFirstStr {
		if first.Type == turing_parser.JsVarName {
			arg1.Params.Add(&turing_parser.JSOpcode{Type: turing_parser.JsConstVal, Value: &turing_parser.JSConstVal{Value: "\"" + first.Value.(*turing_parser.JSVarName).Name + "\""}})
		} else {
			arg1.Params.Add(g.GoRebuildArgument(first))
		}
	} else {
		arg1.Params.Add(g.GoRebuildArgument(first))
	}
	arg2 := &turing_parser.JSEquation{Params: &arraylist.List{}}
	arg2.Params.Add(g.GoRebuildEquation(oeq, from+2))

	args := &arraylist.List{}
	args.Add(arg1)
	args.Add(arg2)

	fx := &turing_parser.JSOpcode{Type: turing_parser.JsVarName, Value: &turing_parser.JSVarName{Name: opFunc}}
	jsCall := &turing_parser.JSFuncCall{Func: fx, Arguments: args}

	return &turing_parser.JSOpcode{Type: turing_parser.JsFuncCall, Value: jsCall}
}

func (g *GoGenerator) IsThisFunction(Func *turing_parser.JSOpcode, goFile *GoFile) bool {

	if Func.Type == turing_parser.JsVarName {
		val := Func.Value.(*turing_parser.JSVarName)
		if len(val.Name) > 5 && val.Name[:5] == "this." {

			name := strings.Title(val.Name[5:])

			if goFile.OwnFunctions[name] == true {
				return false
			}

			if g.BaseMethods[name] == false {
				return true
			}
		}
	}
	return false
}

func (g *GoGenerator) IsThisMember(Member *turing_parser.JSOpcode) bool {

	if Member.Type == turing_parser.JsVarName {
		val := Member.Value.(*turing_parser.JSVarName)
		if len(val.Name) > 5 && val.Name[:5] == "this." {

			//	names := strings.Split(val.Name, ".")
			//	name := strings.Title(names[1])
			//
			//	return g.BaseMembers[name] == false

			return true // we dont have native members all is in the values map
		}
	}
	return false
}

func (g *GoGenerator) BuildGOAtChain(names []string, from int, to int) string {
	if from == to-1 {
		return names[from]
	}
	return "*((" + g.BuildGOAtChain(names, from, to-1) + ").At(MkString(\"" + names[to-1] + "\")))"
}

func (g *GoGenerator) BuildGOBlock(op *turing_parser.JSOpcode, goFile *GoFile, indent int) string {

	//inEq := (indent & 0x80000000) != 0
	//indent &= 0x7FFFFFFF

	str := ""

	if op.Type == turing_parser.JsFunction {
		fx := op.Value.(*turing_parser.JSFunction)
		str += "func (this *" + goFile.ThisName + ") "
		str += strings.Title(fx.Name)
		str += "(goArgs ...*Variant"
		/*for i := 0; i < fx.Arguments.Size(); i++ {
			if i > 0 {
				str += ", "
			}
			arg_, _ := fx.Arguments.Get(i)
			arg := arg_.(*turing_parser.JSNameValue)
			str += arg.Name
			if arg.Value.Params.Size() > 0 {
				str += " = "
				str += BuildGOBlock(&turing_parser.JSOpcode{Type: turing_parser.JsEquation, Value: arg.Value}, indent+1)
			}
		}*/
		str += ") *Variant"
		str += "{\n"

		for i := 0; i < fx.Arguments.Size(); i++ {
			arg_, _ := fx.Arguments.Get(i)
			arg := arg_.(*turing_parser.JSNameValue)
			str += turing_parser.PrintIndent(indent)
			name := g.FixGoName(arg.Name)
			str += fmt.Sprintf("%s := GoGetArg(goArgs, %d, ", name, i) // todo: get actual default value
			if arg.Value.Params.Size() > 0 {
				str += g.BuildGOBlock(&turing_parser.JSOpcode{Type: turing_parser.JsEquation, Value: arg.Value}, goFile, indent+1)
			} else {
				str += "MkUndefined()"
			}
			str += ")"
			str += ";"
			if g.ForceUsage {
				str += " _ = " + name + ";"
			}
			str += "\n"
			//str += "_ = " + arg.Name + "\n" // sledge hammer approche to avoid unised parameter errors
		}

		var last *turing_parser.JSOpcode = nil
		for i := 0; i < fx.FuncBody.Size(); i++ {
			tmp_, _ := fx.FuncBody.Get(i)
			tmp := tmp_.(*turing_parser.JSOpcode)
			str += turing_parser.PrintIndent(indent)
			str += g.BuildGOBlock(tmp, goFile, indent+1)
			turing_parser.AddSemiIfNeeded(&str)
			last = tmp
		}
		if last == nil || last.Type != turing_parser.JsReturn {
			str += turing_parser.PrintIndent(indent)
			str += "return MkUndefined()\n"
		}
		str += turing_parser.PrintIndent(indent - 1)
		str += "}\n\n"
	} else if op.Type == turing_parser.JsVarInit {
		init := op.Value.(*turing_parser.JSVarInit)
		if len(init.Names) == 1 {
			name := g.FixGoName(init.Names[0])
			str += name
			str += " := "
			tmp := g.GoRebuildOne(init.Value.Params, 0)
			if tmp != nil {
				str += g.BuildGOBlock(tmp, goFile, indent+1)
			} else {
				str += g.BuildGOBlock(&turing_parser.JSOpcode{Type: turing_parser.JsEquation, Value: init.Value}, goFile, indent+1)
			}

			if indent > 0 {
				str += ";"
				if g.ForceUsage {
					str += " _ = " + name
				}
			}
		} else {

			for i := 0; i < len(init.Names); i++ {
				name := g.FixGoName(init.Names[i])
				if i > 0 {
					str += turing_parser.PrintIndent(indent - 1)
				}
				str += name
				str += " := MkUndefined()"
				if indent > 0 {
					str += "; "
					if g.ForceUsage {
						str += "_ = " + name
					}
				}
				str += ";\n"
			}

			str += turing_parser.PrintIndent(indent - 1)
			if init.MultiType == 1 { // Array
				str += "ArrayUnpack("
				str += g.BuildGOBlock(&turing_parser.JSOpcode{Type: turing_parser.JsEquation, Value: init.Value}, goFile, indent+1)
				for i := 0; i < len(init.Names); i++ {
					name := g.FixGoName(init.Names[i])
					str += ", &" + name
				}
				str += ")"
			} else if init.MultiType == 2 { // Object
				str += "ObjectUnpack("
				str += g.BuildGOBlock(&turing_parser.JSOpcode{Type: turing_parser.JsEquation, Value: init.Value}, goFile, indent+1)
				for i := 0; i < len(init.Names); i++ {
					name := g.FixGoName(init.Names[i])
					str += ", MkStringRef(\"" + name + "\")"
					str += ", &" + name
				}
				str += ")"
			}
		}
	} else if op.Type == turing_parser.JsVarName {
		val := op.Value.(*turing_parser.JSVarName)
		if g.IsThisMember(op) {
			names := strings.Split(val.Name, ".")
			if len(names) > 2 {
				str += "(" + g.BuildGOAtChain(names, 0, len(names)-1)
				str += ")." + strings.Title(names[len(names)-1])
			} else {
				str += "*this.At(MkString(\""
				str += names[1]
				str += "\"))"
			}
		} else {
			str += g.FixGoName(val.Name)
		}
		str += " "
	} else if op.Type == turing_parser.JsMember {
		val := op.Value.(*turing_parser.JSMember)

		//str += "*(" + g.BuildGOBlock(val.Object, (indent+1)|0x80000000) + ").At("
		str += "*(" + g.BuildGOBlock(val.Object, goFile, indent+1) + ").At("
		str += g.BuildGOBlock(&turing_parser.JSOpcode{Type: turing_parser.JsEquation, Value: val.Member}, goFile, indent+1)
		str += ")"

	} else if op.Type == turing_parser.JsConstVal {
		val := op.Value.(*turing_parser.JSConstVal)
		if IsString(val.Value) {
			//if g.StrictVariance {
			str += "MkString(" + g.FixGoString(val.Value) + ")"
			//} else {
			//	str += g.FixGoString(val.Value)
			//}
		} else if IsNumber(val.Value) {
			//if g.StrictVariance {
			if strings.Contains(val.Value, ".") {
				str += "MkNumber(" + val.Value + ")"
			} else {
				str += "MkInteger(" + val.Value + ")"
			}
			//} else {
			//	str += val.Value
			//}
		} else {
			str += val.Value
		}
		str += " "
	} else if op.Type == turing_parser.JsObject {
		obj := op.Value.(*turing_parser.JSObject)
		//if g.StrictVariance {
		str += "MkMap(&VarMap{"
		//} else {
		//	str += "MkIMap(&ItfMap{"
		//}
		if obj.Params.Size() > 1 {
			str += "\n"
		}
		for i := 0; i < obj.Params.Size(); i++ {
			//if i > 0 {
			//	str += ",\n"
			//}
			arg_, _ := obj.Params.Get(i)
			arg := arg_.(*turing_parser.JSNameValue)
			if obj.Params.Size() > 1 {
				str += turing_parser.PrintIndent(indent)
			}
			if len(arg.Name) > 0 {
				str += g.FixGoString(arg.Name)
				str += ":"
				str += g.BuildGOBlock(&turing_parser.JSOpcode{Type: turing_parser.JsEquation, Value: arg.Value}, goFile, indent+1)
			}
			if obj.Params.Size() > 1 {
				str += ",\n"
			}
		}
		if obj.Params.Size() > 1 {
			str += turing_parser.PrintIndent(indent)
		}
		str += "})"
	} else if op.Type == turing_parser.JsArray {
		arr := op.Value.(*turing_parser.JSArray)
		//if g.StrictVariance {
		str += "MkArray(&VarArray{"
		//} else {
		//	str += "MkIArray(&ItfArray{"
		//}
		if arr.Params.Size() > 0 {
			str += "\n"
		}
		for i := 0; i < arr.Params.Size(); i++ {
			//if i > 0 {
			//	str += ",\n"
			//}
			arg_, _ := arr.Params.Get(i)
			arg := arg_.(*turing_parser.JSEquation)
			str += turing_parser.PrintIndent(indent)
			str += g.BuildGOBlock(&turing_parser.JSOpcode{Type: turing_parser.JsEquation, Value: arg}, goFile, indent+1)
			str += ",\n"
		}
		if arr.Params.Size() > 0 {
			str += turing_parser.PrintIndent(indent)
		}
		str += "})"
	} else if op.Type == turing_parser.JsOperator {
		val := op.Value.(*turing_parser.JSOperator)
		str += val.Op
		str += " "
	} else if op.Type == turing_parser.JsEquation {
		eq := op.Value.(*turing_parser.JSEquation)

		/*if inEq {
			str += "("
		}
		for i := 0; i < oeq.Size(); i++ {
			cur_, _ := oeq.Get(i)
			cur := cur_.(*turing_parser.JSOpcode)
			str += g.BuildGOBlock(cur, (indent+1)|0x80000000)
		}
		if inEq {
			str += ")"
		}*/

		from := 0
		bSub := true

		if eq.Params.Size() > 2 { // special case for simple assignment
			second_, _ := eq.Params.Get(from + 1)
			second := second_.(*turing_parser.JSOpcode)

			if second.Type == turing_parser.JsOperator && second.Value.(*turing_parser.JSOperator).Op == "=" {
				from += 2
				bSub = false
			}
		}

		oeq := turing_parser.GoOrderEquation(eq.Params, bSub) // order equation a = b + c * d to a = (b + (c * d))
		//DumpEquation(oeq)

		if oeq.Size() == 2 { // keep (&VarName) unaltered
			first_, _ := oeq.Get(from)
			first := first_.(*turing_parser.JSOpcode)

			if first.Type == turing_parser.JsOperator && first.Value.(*turing_parser.JSOperator).Op == "&" {
				from += 2
			}
		}

		for i := 0; i < from; i++ {
			cur_, _ := oeq.Get(i)
			cur := cur_.(*turing_parser.JSOpcode)
			str += g.BuildGOBlock(cur, goFile, indent+1)
		}

		if from < oeq.Size() {

			veq := g.GoRebuildEquation(oeq, from) // reconstruct a = (b + (c * d)) to a = vAdd(b, vMulti(c, d))

			str += g.BuildGOBlock(veq, goFile, indent+1)
		}

	} else if op.Type == turing_parser.JsFuncCall {
		fx := op.Value.(*turing_parser.JSFuncCall)
		//str += g.FixGoName(fx.Name)

		if fx.Func.Type == turing_parser.JsMember {

			val := fx.Func.Value.(*turing_parser.JSMember)

			/*if val.Object.Type == turing_parser.JsVarName && val.Object.Value.(*turing_parser.JSVarName).Name == "this" {
			str += "this.Call("*/
			if val.Object.Type == turing_parser.JsVarName {
				val := val.Object.Value.(*turing_parser.JSVarName)
				str += g.FixGoName(val.Name) + ".Call("
			} else {
				str += "(" + g.BuildGOBlock(val.Object, goFile, indent+1) + ").Call("
			}
			str += g.BuildGOBlock(&turing_parser.JSOpcode{Type: turing_parser.JsEquation, Value: val.Member}, goFile, indent+1)
			if fx.Arguments.Size() > 0 {
				str += ", "
			}
		} else if g.IsThisFunction(fx.Func, goFile) {
			names := strings.Split(fx.Func.Value.(*turing_parser.JSVarName).Name, ".")
			if len(names) > 2 {
				str += "(" + g.BuildGOAtChain(names, 0, len(names)-1)
				str += ").Call(MkString(\"" + names[len(names)-1] + "\")"
			} else {
				//if g.StrictVariance {
				str += "this.Call(MkString(\"" + names[1] + "\")"
				//} else {
				//	str += "this.Call(\"" + fx.Func.Value.(*turing_parser.JSVarName).Name[5:] + "\""
				//}
			}
			if fx.Arguments.Size() > 0 {
				str += ", "
			}
		} else {
			//str += g.BuildGOBlock(fx.Func, indent+1)
			str += g.FixGoName(fx.Func.Value.(*turing_parser.JSVarName).Name)
			str += "("
		}

		for i := 0; i < fx.Arguments.Size(); i++ {
			if i > 0 {
				str += ", "
			}
			arg_, _ := fx.Arguments.Get(i)
			arg := arg_.(*turing_parser.JSEquation)
			str += g.BuildGOBlock(&turing_parser.JSOpcode{Type: turing_parser.JsEquation, Value: arg}, goFile, indent+1)
		}
		str += ")"
	} else if op.Type == turing_parser.JsCondition {
		cond := op.Value.(*turing_parser.JSCondition)
		//str += PrintIndent(indent)
		str += "if IsTrue("
		str += g.BuildGOBlock(&turing_parser.JSOpcode{Type: turing_parser.JsEquation, Value: cond.Condition}, goFile, indent+1)
		str += ") {\n"
		for i := 0; i < cond.CodeBlock.Opcodes.Size(); i++ {
			tmp_, _ := cond.CodeBlock.Opcodes.Get(i)
			tmp := tmp_.(*turing_parser.JSOpcode)
			str += turing_parser.PrintIndent(indent)
			str += g.BuildGOBlock(tmp, goFile, indent+1)
			turing_parser.AddSemiIfNeeded(&str)
		}
		str += turing_parser.PrintIndent(indent - 1)
		str += "}"
		if cond.ElseBlock != nil {
			str += " else "
			str += "{\n"
			for i := 0; i < cond.ElseBlock.Opcodes.Size(); i++ {
				tmp_, _ := cond.ElseBlock.Opcodes.Get(i)
				tmp := tmp_.(*turing_parser.JSOpcode)
				str += turing_parser.PrintIndent(indent)
				str += g.BuildGOBlock(tmp, goFile, indent+1)
				turing_parser.AddSemiIfNeeded(&str)
			}
			str += turing_parser.PrintIndent(indent - 1)
			str += "}\n"
		} else {
			str += "\n"
		}
	} else if op.Type == turing_parser.JsFor {
		loop := op.Value.(*turing_parser.JSFor)
		//str += PrintIndent(indent)
		str += "for "
		str += g.BuildGOBlock(loop.Init, goFile, -100)
		str += "; IsTrue("
		str += g.BuildGOBlock(&turing_parser.JSOpcode{Type: turing_parser.JsEquation, Value: loop.Condition}, goFile, -100)
		str += "); "
		str += g.BuildGOBlock(&turing_parser.JSOpcode{Type: turing_parser.JsEquation, Value: loop.Next}, goFile, -100)
		str += "{\n"
		for i := 0; i < loop.CodeBlock.Opcodes.Size(); i++ {
			tmp_, _ := loop.CodeBlock.Opcodes.Get(i)
			tmp := tmp_.(*turing_parser.JSOpcode)
			str += turing_parser.PrintIndent(indent)
			str += g.BuildGOBlock(tmp, goFile, indent+1)
			turing_parser.AddSemiIfNeeded(&str)
		}
		str += turing_parser.PrintIndent(indent - 1)
		str += "}\n"
	} else if op.Type == turing_parser.JsForEach {
		loop := op.Value.(*turing_parser.JSForEach)
		//str += PrintIndent(indent)
		str += "for ("
		str += loop.Name
		str += " in "
		str += g.BuildGOBlock(&turing_parser.JSOpcode{Type: turing_parser.JsEquation, Value: loop.In}, goFile, -100)
		str += ")"
		str += "{\n"
		for i := 0; i < loop.CodeBlock.Opcodes.Size(); i++ {
			tmp_, _ := loop.CodeBlock.Opcodes.Get(i)
			tmp := tmp_.(*turing_parser.JSOpcode)
			str += turing_parser.PrintIndent(indent)
			str += g.BuildGOBlock(tmp, goFile, indent+1)
			turing_parser.AddSemiIfNeeded(&str)
		}
		str += turing_parser.PrintIndent(indent - 1)
		str += "}\n"
	} else if op.Type == turing_parser.JsReturn {
		ret := op.Value.(*turing_parser.JSReturn)
		str += "return "
		if ret.Value != nil {
			str += g.BuildGOBlock(&turing_parser.JSOpcode{Type: turing_parser.JsEquation, Value: ret.Value}, goFile, indent+1)
		} else {
			str += "MkUndefined()"
		}
		turing_parser.AddSemiIfNeeded(&str)
	} else if op.Type == turing_parser.JsThrow {
		throw := op.Value.(*turing_parser.JSThrow)
		str += "panic("

		// fix for throw'ign and forgting to new
		tmp_, _ := throw.Value.Params.Get(0)
		tmp := tmp_.(*turing_parser.JSOpcode)
		if tmp.Type == turing_parser.JsFuncCall {
			str += "New"
		}

		str += g.BuildGOBlock(&turing_parser.JSOpcode{Type: turing_parser.JsEquation, Value: throw.Value}, goFile, indent+1)
		str += ")"
		turing_parser.AddSemiIfNeeded(&str)
	} else if op.Type == turing_parser.JsTryCatch {
		try := op.Value.(*turing_parser.JSTryCatch)

		str += "{\n"
		str += turing_parser.PrintIndent(indent - 1)
		str += "ret__ := func(this *" + goFile.ThisName + ") (ret_ *Variant) {\n"

		str += turing_parser.PrintIndent(indent)
		str += "defer func() {\n"

		str += turing_parser.PrintIndent(indent + 1)
		str += "if " + try.ExceptionName + " := recover().(*Variant); " + try.ExceptionName + " != nil {\n"

		str += turing_parser.PrintIndent(indent + 2)
		str += "ret_ = func(this *" + goFile.ThisName + ") *Variant {\n"

		str += turing_parser.PrintIndent(indent + 3)
		str += "// catch block:\n"
		for i := 0; i < try.CatchBlock.Opcodes.Size(); i++ {
			tmp_, _ := try.CatchBlock.Opcodes.Get(i)
			tmp := tmp_.(*turing_parser.JSOpcode)
			str += turing_parser.PrintIndent(indent + 3)
			str += g.BuildGOBlock(tmp, goFile, indent+4)
			turing_parser.AddSemiIfNeeded(&str)
		}

		str += turing_parser.PrintIndent(indent + 3)
		str += "return nil\n"
		str += turing_parser.PrintIndent(indent + 2)
		str += "}(this)\n"

		str += turing_parser.PrintIndent(indent + 1)
		str += "}\n"

		str += turing_parser.PrintIndent(indent)
		str += "}()\n"

		str += turing_parser.PrintIndent(indent)
		str += "// try block:\n"
		for i := 0; i < try.TryBlock.Opcodes.Size(); i++ {
			tmp_, _ := try.TryBlock.Opcodes.Get(i)
			tmp := tmp_.(*turing_parser.JSOpcode)
			str += turing_parser.PrintIndent(indent)
			str += g.BuildGOBlock(tmp, goFile, indent+1)
			turing_parser.AddSemiIfNeeded(&str)
		}

		str += turing_parser.PrintIndent(indent)
		str += "return nil\n"
		str += turing_parser.PrintIndent(indent - 1)
		str += "}(this)\n"

		str += turing_parser.PrintIndent(indent - 1)
		str += "if ret__ != nil {\n"
		str += turing_parser.PrintIndent(indent)
		str += "return ret__;\n "
		str += turing_parser.PrintIndent(indent - 1)
		str += "}\n"
		str += turing_parser.PrintIndent(indent - 1)
		str += "}\n"

		// todo use return values
		// todo finally block
	} else if op.Type == turing_parser.JsNew {
		new := op.Value.(*turing_parser.JSNew)
		if new.Constr.Func.Type == turing_parser.JsVarName {
			str += "New" // our contructors must be NewSomething
		}
		str += g.BuildGOBlock(&turing_parser.JSOpcode{Type: turing_parser.JsFuncCall, Value: new.Constr}, goFile, indent+1)
	} else if op.Type == turing_parser.JsDelete {
		del := op.Value.(*turing_parser.JSDelete)
		str += g.BuildGOBlock(&turing_parser.JSOpcode{Type: turing_parser.JsEquation, Value: del.Value}, goFile, indent+1)
		str += " = MkUndefined()"
		turing_parser.AddSemiIfNeeded(&str)
	}

	return str
}
