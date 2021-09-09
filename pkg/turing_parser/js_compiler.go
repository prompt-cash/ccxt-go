package turing_parser

import (
	"fmt"
	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/prompt-cash/js-transpiler/log"
)

type JSCompiler struct {
	logger log.Logger
}

func NewJSCompiler(logger log.Logger) (*JSCompiler, error) {
	return &JSCompiler{
		logger: logger,
	}, nil
}

type JSOpcodeType uint8

const (
	JsBlock JSOpcodeType = iota
	JsFunction
	JsVarInit
	//JsAssign
	JsVarName
	JsConstVal
	JsMember
	JsObject
	JsArray
	JsOperator
	JsEquation
	JsFuncCall
	JsCondition
	JsFor
	JsForEach
	JsReturn
	JsThrow
	JsTryCatch
	JsNew
	JsDelete
)

type JSOpcode struct {
	Type  JSOpcodeType
	Value interface{}
}

type JSBlock struct {
	Opcodes *arraylist.List // type JSOpcode
}

type JSFunction struct {
	Name      string
	Arguments *arraylist.List // type JSNameValue
	FuncBody  *arraylist.List // type JSOpcode
}

type JSEquation struct {
	Params *arraylist.List // type JSOpcode: JsVarName, JsConstVal, JsOperator, JsEquation, JsFuncCall
}

type JSOperator struct {
	Op string
}

type JSVarName struct {
	Name string
}

type JSConstVal struct {
	Value string
}

type JSMember struct {
	Object *JSOpcode
	Member *JSEquation
}

type JSNameValue struct {
	Name  string
	Value *JSEquation
}

type JSObject struct {
	Params *arraylist.List // type JSNameValue
}

type JSArray struct {
	Params *arraylist.List // type JsEquation
}

type JSCondition struct {
	Condition *JSEquation
	CodeBlock *JSBlock
	ElseBlock *JSBlock
}

type JSFor struct {
	Init      *JSOpcode
	Condition *JSEquation
	Next      *JSEquation

	CodeBlock *JSBlock
}

type JSForEach struct {
	Name string
	In   *JSEquation

	CodeBlock *JSBlock
}

type JSFuncCall struct {
	Func      *JSOpcode       // JSVarName, JSMember
	Arguments *arraylist.List // type JsEquation
}

type JSVarInit struct {
	Names     []string
	Value     *JSEquation
	Const     bool
	MultiType int
}

/*type JSAssign struct {
	Name  string
	Type  string
	Names *arraylist.List // type JsEquation
	Value *JSEquation
}*/

type JSReturn struct {
	Value *JSEquation
}

type JSThrow struct {
	Value *JSEquation
}

type JSTryCatch struct {
	TryBlock      *JSBlock
	ExceptionName string
	CatchBlock    *JSBlock
	FinallyBlock  *JSBlock
}

type JSNew struct {
	Constr *JSFuncCall
}

type JSDelete struct {
	Value *JSEquation // shoudl contain one JSVarName, JSMember
}

func IsAssigmnent(str string) bool {
	return str == "=" || str == "+=" || str == "-=" || str == "*=" || str == "/=" || str == "%=" || str == "<<=" || str == ">>=" || str == "|=" || str == "&=" || str == "^="
}

func IsComparison(str string) bool {
	return str == "==" || str == "!=" || str == "===" || str == "!==" || str == "<=" || str == ">=" || str == "<" || str == ">"
}

func (this *JSCompiler) CompileJsArguments(block *Expressions) *arraylist.List {

	args := &arraylist.List{}

	if block.Count() > 0 && block.GetExpStr(1) != "," {
		args.Add(this.CompileJsEquation(block, 0))
		return args
	}

	for i := 0; i < block.Count(); i++ {

		exp := block.GetExp(i)

		args.Add(this.CompileJsEquationEx(exp))

		i++
		if i < block.Count() && block.GetExpStr(i) != "," {
			this.logger.Errorf("ERROR invalid array\n")
			break
		}
	}

	return args
}

func (this *JSCompiler) CompileJsMember(blocks *Expressions, pFrom *int) *JSEquation {

	var eq *JSEquation = nil

	if len(blocks.GetExpStr(*pFrom)) > 0 && blocks.GetExpStr(*pFrom)[0] == '.' {

		eq = &JSEquation{
			Params: &arraylist.List{},
		}

		// todo split by .

		jsConstVal := &JSConstVal{Value: "\"" + blocks.GetExpStr(*pFrom)[1:] + "\""}

		eq.Params.Add(&JSOpcode{Type: JsConstVal, Value: jsConstVal})

		*pFrom++

	} else if blocks.IsExpMulti(*pFrom) && blocks.GetExpMulti(*pFrom).Brace == '[' {

		eq = this.CompileJsEquation(blocks.GetExpMulti(*pFrom), 0)

		*pFrom++
	}

	return eq
}

func (this *JSCompiler) CompileJsMembers(prev *JSOpcode, blocks *Expressions, pFrom *int) *JSOpcode {

	obj := &JSMember{
		Object: prev,
	}

	for {
		eq := this.CompileJsMember(blocks, pFrom)
		if eq == nil {
			break
		}

		if obj.Member != nil {
			prev = &JSOpcode{Type: JsMember, Value: obj}

			obj = &JSMember{
				Object: prev,
			}
		}

		obj.Member = eq
	}

	if obj.Member == nil {
		return nil
	}

	return &JSOpcode{Type: JsMember, Value: obj}
}

func (this *JSCompiler) TryCompileJsFuncCall(block *Expressions, pos *int) *JSOpcode {

	j := *pos
	if block.GetExpStr(j) == "await" {
		j++
	}

	exp := block.GetExp(j)
	if exp == nil {
		return nil
	}

	// bla() -> FuncCall {Func=bla, args}
	// bla.blup() -> FuncCall {Func=bla.blup, args}
	// bla["bl" + "up"]() -> FuncCall {Func= Member{Object= bla, Member = "bl" + "up"}, args}
	// bla["bl" + "up"]["test"]() -> FuncCall {Func= Member{Object= Member{Object= bla, Member = "bl" + "up"}, Member = "test"}, args}
	// ...().foo	-> Member { Object = FuncCall {...}, Member = foo}
	// ...().foo()  -> FuncCall {Func = Member { Object = FuncCall {...}, Member = foo}, args}
	// ...()[bar]() -> FuncCall {Func = Member { Object = FuncCall {...}, Member = bar}, args}

	var opcode *JSOpcode

	if exp.IsMulti() && exp.Multi.Brace == '[' {
		opcode = this.CompileJsArray(exp.Multi)
	} else {
		opcode = this.CompileJsOpCode(exp)
		if opcode == nil || opcode.Type != JsVarName {
			return nil
		}
	}
	found := false
next:

	if (block.IsExpMulti(j+1) && block.GetExpMulti(j+1).Brace == '[') || (len(block.GetExpStr(j+1)) > 0 && block.GetExpStr(j + 1)[0] == '.') {
		j++
		opcode = this.CompileJsMembers(opcode, block, &j)
		j--
	}

	// is it a function call
	if !block.IsExpMulti(j+1) || block.GetExpMulti(j+1).Brace != '(' {
		if found {
			*pos = j
			return opcode
		} else {
			return nil
		}
	}

	//this.logger.Errorf("found fucntion call, name %s\n", block.GetExpStr(j))

	callArgs := block.GetExpMulti(j + 1)
	if callArgs == nil {
		this.logger.Errorf("ERROR invalid function call %s\n", block.dbg)
		return nil
	}

	j += 1

	jsCall := &JSFuncCall{Func: opcode, Arguments: this.CompileJsArguments(callArgs)}
	opcode = &JSOpcode{Type: JsFuncCall, Value: jsCall}

	found = true
	goto next
}

func (this *JSCompiler) CompileJsEquationEx(exp *Expression) *JSEquation {

	if exp == nil {
		return nil
	}

	if exp.IsMulti() {
		return this.CompileJsEquation(exp.Multi, 0)
	}

	eq := &JSEquation{
		Params: &arraylist.List{},
	}

	opcode := this.CompileJsOpCode(exp)
	if opcode != nil {
		eq.Params.Add(opcode)
	} else {
		this.logger.Errorf("ERROR invalid opcode\n")
	}

	return eq
}

func (this *JSCompiler) CompileJsOpCode(exp *Expression) *JSOpcode {

	// operator
	if exp.IsOperator() {
		jsOperator := &JSOperator{Op: *exp.Word}
		return &JSOpcode{Type: JsOperator, Value: jsOperator}
	}

	// constant
	if exp.IsString() || exp.IsNumber() {
		jsConstVal := &JSConstVal{Value: *exp.Word}
		return &JSOpcode{Type: JsConstVal, Value: jsConstVal}
	}

	// special letter operator
	if exp.IsWord() && (*exp.Word == "in" || *exp.Word == "typeof" || *exp.Word == "instanceof") {
		jsOperator := &JSOperator{Op: *exp.Word}
		return &JSOpcode{Type: JsOperator, Value: jsOperator}
	}

	// variable name
	if exp.IsWord() {
		jsVarName := &JSVarName{Name: *exp.Word}
		return &JSOpcode{Type: JsVarName, Value: jsVarName}
	}

	return nil
}

func (this *JSCompiler) CompileJsObjParam(exp *Expressions, sep string) *JSNameValue {

	if exp == nil || (exp.Count() > 1 && exp.GetExpStr(1) != sep) {
		this.logger.Errorf("ERROR invalid object parameter\n")
		return nil
	}

	name := exp.GetExpStr(0)
	return &JSNameValue{Name: name, Value: this.CompileJsEquation(exp, 2)}
}

func (this *JSCompiler) CompileJsObjectEx(blocks *Expressions, sep string) *arraylist.List {

	params := &arraylist.List{}

	if blocks.Count() > 0 && (blocks.GetExpStr(1) == sep || blocks.Count() == 1) { // special case, only one member
		params.Add(this.CompileJsObjParam(blocks, sep)) // TODO fix me mind the coma
		return params
	}

	for i := 0; i < blocks.Count(); i++ {

		exp := blocks.GetExpMulti(i)

		param := this.CompileJsObjParam(exp, sep)
		if param != nil {
			params.Add(param)
		}

		i++
		if i < blocks.Count() && blocks.GetExpStr(i) != "," {
			this.logger.Errorf("ERROR invalid object\n")
			break
		}
	}

	return params
}

func (this *JSCompiler) CompileJsObject(blocks *Expressions) *JSOpcode {

	obj := &JSObject{
		Params: this.CompileJsObjectEx(blocks, ":"),
	}

	return &JSOpcode{Type: JsObject, Value: obj}
}

func (this *JSCompiler) CompileJsArray(blocks *Expressions) *JSOpcode {

	obj := &JSArray{
		Params: this.CompileJsArguments(blocks),
	}

	return &JSOpcode{Type: JsArray, Value: obj}
}

func (this *JSCompiler) CompileJsEquation(blocks *Expressions, from int) *JSEquation {

	eq := &JSEquation{
		Params: &arraylist.List{},
	}

	//OrderEquation(blocks, from, false);
	//AddDebugStrings(blocks)

	for i := from; i < blocks.Count(); i++ {

		// function call
		opcode := this.TryCompileJsFuncCall(blocks, &i)
		if opcode != nil {
			eq.Params.Add(opcode)
			continue
		}

		exp := blocks.GetExp(i)

		if exp.IsMulti() {

			// equation
			if exp.Multi.Brace == '(' {
				eq := this.CompileJsEquation(exp.Multi, 0)
				opcode = &JSOpcode{Type: JsEquation, Value: eq}
			} else

			// object
			if exp.Multi.Brace == '{' {
				opcode = this.CompileJsObject(exp.Multi)
			} else

			// array
			if exp.Multi.Brace == '[' {
				opcode = this.CompileJsArray(exp.Multi)
			}

		} else {

			// new operator
			if exp.IsWord() && *exp.Word == "new" {

				pos := i + 1
				opcode = this.TryCompileJsFuncCall(blocks, &pos)
				if opcode != nil {
					i = pos

					jsNew := &JSNew{}

					jsNew.Constr = opcode.Value.(*JSFuncCall)

					opcode = &JSOpcode{Type: JsNew, Value: jsNew}

				} else {
					this.logger.Errorf("ERROR invalid new\n")
				}

			} else

			// opcode
			{
				opcode = this.CompileJsOpCode(exp)
				if opcode != nil && opcode.Type == JsVarName && blocks.IsExpMulti(i+1) && blocks.GetExpMulti(i+1).Brace == '[' {
					i++
					opcode = this.CompileJsMembers(opcode, blocks, &i)
					i--
				}
			}

		}

		if opcode != nil {
			eq.Params.Add(opcode)
			continue
		}

		this.logger.Errorf("ERROR invalid opcode\n")
	}

	//this.logger.Errorf("%s\n", blocks.dbg)

	/*dump := PrintExpressionImpl(blocks, from, 0)
	fmt.Println(dump)*/

	return eq
}

func (this *JSCompiler) CompileSimpleJsBlock(block *Expressions) *JSOpcode {

	/*pos := 0
	funcCall := TryCompileJsFuncCall(block, &pos)
	if funcCall != nil{
		return funcCall
	}*/

	j := 0
	str := block.GetExpStr(j)

	// is it a variable initialization
	if str == "const" || str == "let" {
		//this.logger.Errorf("found variable initialization\n")

		if block.Count() <= j+3 {
			this.logger.Errorf("ERROR invalid variable initialization %s\n", block.dbg)
			return nil
		}

		jsVarInit := &JSVarInit{Names: []string{}, Value: this.CompileJsEquation(block, j+3), Const: str == "const"}

		if block.IsExpMulti(j + 1) {
			exp := block.GetExpMulti(j + 1)
			if exp.Brace == '[' {
				jsVarInit.MultiType = 1
			} else if exp.Brace == '{' {
				jsVarInit.MultiType = 2
			} else {
				this.logger.Errorf("ERROR invalid variable multi initialization %s\n", block.dbg)
				return nil
			}

			// todo: add more plausibility checks
			for k := 0; k < exp.Count(); k += 2 { // skip ,'s
				exp2 := exp.GetExpMulti(k)
				if exp2 != nil {
					jsVarInit.Names = append(jsVarInit.Names, exp2.GetExpStr(0))
				}
			}
		} else {
			varName := block.GetExpStr(j + 1)
			if varName == "" {
				this.logger.Errorf("ERROR invalid variable initialization %s\n", block.dbg)
				return nil
			}
			jsVarInit.Names = append(jsVarInit.Names, varName)
		}

		return &JSOpcode{Type: JsVarInit, Value: jsVarInit}
	}

	/*// is is a variable assignment, simple foo.bar = x or a complex foo["bar"] = x
	if len(str) > 0 {
		j++
		k := j
		eq := 0
		for ; j < block.Count(); j++ {
			if IsAssigmnent(block.GetExpStr(j)) {
				eq = j
				break
			}
			if !block.IsExpMulti(j) || block.GetExpMulti(j).Brace != '[' {
				break
			}
		}

		if eq != 0 {
			names := &arraylist.List{}

			if eq > k {
				//this.logger.Errorf("found member asignment\n")

				for ; k < eq; k++ {
					names.Add(CompileJsEquation(block.GetExpMulti(k), 0))
				}
			} else {
				//this.logger.Errorf("found variable asignment\n")
			}

			jsAssign := &JSAssign{Name: str, Type: block.GetExpStr(eq), Value: CompileJsEquation(block, eq+1)}
			if names.Size() > 0 {
				jsAssign.Names = names
			}
			return &JSOpcode{Type: JsAssign, Value: jsAssign}

			return nil
		}
	}
	*/

	eq := this.CompileJsEquation(block, 0)
	if eq != nil {
		return &JSOpcode{Type: JsEquation, Value: eq}
	}

	this.logger.Errorf("ERROR encountered unidentified block %s\n", block.dbg)
	return nil
}

func (this *JSCompiler) CompileJsFuncBody(blocks *Expressions) *arraylist.List {

	js := &arraylist.List{}

	for i := 0; i < blocks.Count(); i++ {

		block := blocks.GetExpMulti(i)
		if block == nil {
			this.logger.Errorf("ERROR initial block invalid\n", blocks.dbg)
			continue
		}

		j := 0
		str := block.GetExpStr(j)

		var prevIf *JSCondition = nil
	else_if:
		// is it an if
		if str == "if" || prevIf != nil {
			//this.logger.Errorf("found if\n")

			if prevIf != nil {
				j += 1
			}

			ifCond := block.GetExpMulti(j + 1)
			ifBlock := block.GetExpMulti(j + 2)
			if ifCond == nil || ifBlock == nil {
				this.logger.Errorf("ERROR invalid if statement %s\n", block.dbg)
				continue
			}

			jsCondition := &JSCondition{
				Condition: this.CompileJsEquation(ifCond, 0),
				CodeBlock: &JSBlock{Opcodes: this.CompileJsFuncBody(ifBlock)},
			}

			if prevIf != nil {
				prevIf.ElseBlock = &JSBlock{Opcodes: &arraylist.List{}}
				prevIf.ElseBlock.Opcodes.Add(&JSOpcode{Type: JsCondition, Value: jsCondition})
			} else {
				js.Add(&JSOpcode{Type: JsCondition, Value: jsCondition})
			}
			continue
		}

		// is it an if
		if str == "else" {
			//this.logger.Errorf("found else\n")

			prev_, _ := js.Get(js.Size() - 1)
			prev := prev_.(*JSOpcode)

			if prev.Type != JsCondition {
				this.logger.Errorf("ERROR unassiciated else statement %s\n", block.dbg)
				continue
			}
			prevIf = prev.Value.(*JSCondition)

			for prevIf.ElseBlock != nil {

				prev_, _ = prevIf.ElseBlock.Opcodes.Get(0)
				prev = prev_.(*JSOpcode)

				if prev.Type != JsCondition || prevIf.ElseBlock.Opcodes.Size() != 1 {
					this.logger.Errorf("ERROR unassiciated else statement %s\n", block.dbg)
					continue
				}

				prevIf = prev.Value.(*JSCondition)
			}

			if block.GetExpStr(j+1) == "if" {
				goto else_if
			}

			elseBlock := block.GetExpMulti(j + 1)
			if elseBlock == nil {
				this.logger.Errorf("ERROR invalid else statement %s\n", block.dbg)
				continue
			}

			prevIf.ElseBlock = &JSBlock{Opcodes: this.CompileJsFuncBody(elseBlock)}
			continue
		}

		// is it a for
		if str == "for" {
			//this.logger.Errorf("found for\n")

			forBody := block.GetExpMulti(j + 1)
			forBlock := block.GetExpMulti(j + 2)
			if forBody == nil || forBlock == nil {
				this.logger.Errorf("ERROR invalid for statement %s\n", block.dbg)
				continue
			}

			forInit := forBody.GetExpMulti(0)
			forCond := forBody.GetExpMulti(1)
			forNext := forBody.GetExpMulti(2)

			if forInit != nil && forCond != nil && forNext != nil { // for (let i=0; i < 123; i++) {}

				// todo handle for (let i = 1, j = 2; i < j; i+=2, j++) needs parser support!!!
				jsFor := &JSFor{
					Init:      this.CompileSimpleJsBlock(forInit),
					Condition: this.CompileJsEquation(forCond, 0),
					Next:      this.CompileJsEquation(forNext, 0),
					CodeBlock: &JSBlock{Opcodes: this.CompileJsFuncBody(forBlock)},
				}
				js.Add(&JSOpcode{Type: JsFor, Value: jsFor})
				continue
			} else if forBody.GetExpStr(1) == "in" || forBody.GetExpStr(2) == "in" { // for (const bla in blup) {}

				m := 0
				if forBody.GetExpStr(2) == "in" {
					m = 1
				}

				jsForEach := &JSForEach{
					Name: forBody.GetExpStr(m),
					In:   this.CompileJsEquation(forBody, m+2),
				}
				js.Add(&JSOpcode{Type: JsForEach, Value: jsForEach})
				continue
			} else {
				this.logger.Errorf("ERROR unknown for statement %s\n", block.dbg)
				continue
			}
		}

		// is it a try catch
		if str == "try" {

			tryBlock := block.GetExpMulti(j + 1)
			nextBlock := blocks.GetExpMulti(i + 1)
			if tryBlock == nil || nextBlock == nil {
				this.logger.Errorf("ERROR invalid try catch statement %s\n", block.dbg)
				continue
			}

			catchWord := nextBlock.GetExpStr(0)
			catchBracket := nextBlock.GetExpMulti(1)
			catchBlock := nextBlock.GetExpMulti(2)
			if catchWord != "catch" || catchBlock == nil {
				this.logger.Errorf("ERROR invalid try catch statement %s\n", block.dbg)
				continue
			}
			i++ // consume next block

			jsTryCatch := &JSTryCatch{
				TryBlock:      &JSBlock{Opcodes: this.CompileJsFuncBody(tryBlock)},
				ExceptionName: catchBracket.GetExpStr(0),
				CatchBlock:    &JSBlock{Opcodes: this.CompileJsFuncBody(catchBlock)},
				//FinallyBlock:
			}

			nextBlock = blocks.GetExpMulti(i + 1)
			if nextBlock != nil {
				finallyWord := nextBlock.GetExpStr(0)
				finallyBlock := nextBlock.GetExpMulti(1)
				if finallyWord == "finally" && finallyBlock != nil {
					i++ // consume next block

					jsTryCatch.FinallyBlock = &JSBlock{Opcodes: this.CompileJsFuncBody(finallyBlock)}
				}
			}

			js.Add(&JSOpcode{Type: JsTryCatch, Value: jsTryCatch})
			continue
		}

		// is it a return
		if str == "return" {
			//this.logger.Errorf("found return\n")

			var value *JSEquation = nil
			if j+1 < block.Count() {
				value = this.CompileJsEquation(block, j+1)
			}

			jsReturn := &JSReturn{Value: value}
			js.Add(&JSOpcode{Type: JsReturn, Value: jsReturn})
			continue
		}

		// is it a throw
		if str == "throw" {
			//this.logger.Errorf("found return\n")

			var value *JSEquation = nil
			if j+1 < block.Count() {
				value = this.CompileJsEquation(block, 1)
			}

			jsThrow := &JSThrow{Value: value}
			js.Add(&JSOpcode{Type: JsThrow, Value: jsThrow})
			continue
		}

		// is it a delete
		if str == "delete" {
			//this.logger.Errorf("found return\n")

			var value *JSEquation = nil
			if j+1 < block.Count() {
				value = this.CompileJsEquation(block, 1)
			}

			jsDelete := &JSDelete{Value: value}
			js.Add(&JSOpcode{Type: JsDelete, Value: jsDelete})
			continue
		}

		op := this.CompileSimpleJsBlock(block)
		if op != nil {
			js.Add(op)
		}
	}

	return js
}

func (this *JSCompiler) TryCompileJsFunction(exp *Expression) *JSFunction {

	if !exp.IsMulti() {
		return nil
	}

	block := exp.Multi

	j := 0
	if block.GetExpStr(j) == "async" {
		j++
	}

	// is it a function declaration
	if block.GetExpStr(j) == "" || !block.IsExpMulti(j+1) || block.GetExpMulti(j+1).Brace != '(' || !block.IsExpMulti(j+2) || block.GetExpMulti(j+2).Brace != '{' {
		return nil
	}

	name := block.GetExpStr(j)
	args := this.CompileJsObjectEx(block.GetExpMulti(j+1), "=")
	body := this.CompileJsFuncBody(block.GetExpMulti(j + 2))

	return &JSFunction{Name: name, Arguments: args, FuncBody: body}
}

func (this *JSCompiler) CompileJsBlocks(blocks *Expressions) *arraylist.List { // list of JSOpcode's

	js := &arraylist.List{}

	for i := 0; i < blocks.Count(); i++ {

		exp := blocks.GetExp(i)

		fx := this.TryCompileJsFunction(exp)
		if fx != nil {
			js.Add(&JSOpcode{Type: JsFunction, Value: fx})
			continue
		}

		this.logger.Errorf("ERROR unexpected JS block %s", exp.GetDbg())
	}

	return js
}

func (this *JSCompiler) GetModuleExport(blocks *Expressions) *Expressions {

	for i := 0; i < blocks.Count(); i++ {
		exp := blocks.GetExp(i)
		if !exp.IsMulti() {
			continue
		}

		str1 := exp.Multi.GetExpStr(0)
		str2 := exp.Multi.GetExpStr(1)
		if str1 == "module.exports" && str2 == "=" {
			return exp.Multi
		}
	}

	return nil
}

func (this *JSCompiler) CompileJS(blocks *Expressions) *arraylist.List {

	moduleExport := this.GetModuleExport(blocks)

	//moduleName := moduleExport.GetExpStr(3)
	moduleBody := moduleExport.GetExpMulti(6)

	return this.CompileJsBlocks(moduleBody)
	//return CompileJsFuncBody(blocks);
}

///////////////////////////////////////////////////////////////////////////////////////////////////////
// Go port from C/C++ of turing parsers order equation function adapted to use JS Opcodes

func SubordinateEquation(exps *arraylist.List, index int, togo int) bool {
	if togo == 0 {
		return true
	}
	if index < 0 || index+togo > exps.Size() {
		return false
	}

	subExp := &JSEquation{
		Params: &arraylist.List{},
	}
	for ; togo > 0; togo-- {
		exp, _ := exps.Get(index)
		subExp.Params.Add(exp)
		exps.Remove(index)
	}
	exps.Insert(index, &JSOpcode{Type: JsEquation, Value: subExp})
	return true
}

func IsAssignmentOp(str string) bool {
	if str == "=" || str == "+=" || str == "-=" || str == "*=" || str == "/=" || str == "%=" || str == "||=" || str == "&&=" || str == "|=" || str == "&=" || str == "^=" || str == "<<=" || str == ">>=" {
		return true
	}
	return false
}

var OpLevels int = 9

func GetOperatorLevel(str string) int {
	if len(str) == 0 {
		return 0 // this is a multi expression
	}
	if str == "!" || str == "~" || str == "++" || str == "--" || str == "in" || str == "typeof" || str == "instanceof" {
		return 8
	}
	if str == "*" || str == "/" || str == "%" { // high prio math op
		return 7
	}
	if str == "+" || str == "-" { // low prio math op
		return 6
	}
	if IsAssignmentOp(str) { // assignment op
		return 5
	}
	if str == "&" || str == "|" || str == "^" || str == "<<" || str == ">>" { // bitwise op
		return 4
	}
	if str == "==" || str == "!=" || str == "===" || str == "!==" || str == "<" || str == ">" || str == "<=" || str == ">=" { // comparation
		return 3
	}
	if str == "&&" || str == "||" { // logic op
		return 2
	}
	if str == ":" || str == "?" { // other op
		return 1
	}
	return 0 // not a operator
}

func GoOrderEquation(eq *arraylist.List, bSub bool) *arraylist.List {

	tmp := *eq
	Expressions := &tmp

	// Note: if sub is false the first 2 entries will _not_ be re ordered!!!

	// Make sure normal equations a = b & a & c work proeprly and dont overwrite inproperly
	if !bSub && Expressions.Size() > 3 {
		exp_, _ := Expressions.Get(1)
		exp := exp_.(*JSOpcode)
		if exp != nil && exp.Type == JsOperator {
			if IsAssignmentOp(exp.Value.(*JSOperator).Op) {
				SubordinateEquation(Expressions, 2, Expressions.Size()-2)
			}
		}
	}

	Prev := 1
	if bSub {
		Prev = -1
	}
	Order := &arraylist.List{}
	Order.Add(0)
	Start := &arraylist.List{}
	for i := 0; i < OpLevels; i++ {
		Start.Add(0)
	}
	for i := 0; i < Expressions.Size()+1; i++ {
		ExpStr := ""
		if i < Expressions.Size() {
			Expression_, _ := Expressions.Get(i)
			Expression := Expression_.(*JSOpcode)

			if Expression.Type == JsEquation {

				eq := Expression.Value.(*JSEquation)
				if eq.Params.Size() > 1 {
					newParams := GoOrderEquation(eq.Params, true)
					if newParams == nil {
						return nil
					}
					eq.Params = newParams
				}
				continue
			}

			if Expression.Type != JsOperator {

				// special case for x++ or y--
				if Expressions.Size() > 2 && i+1 < Expressions.Size() {
					exp_, _ := Expressions.Get(i + 1)
					exp := exp_.(*JSOpcode)
					if exp.Type == JsOperator && (exp.Value.(*JSOperator).Op == "++" || exp.Value.(*JSOperator).Op == "--") {
						if !SubordinateEquation(Expressions, i, 2) {
							return nil
						}
					}
				}
				// note: ++x and --y are already covered by the regular behaviour

				continue
			}

			op := Expression.Value.(*JSOperator)
			ExpStr = op.Op
		}

		if (bSub && i < 0) || (!bSub && i < 2) {
			continue
		}

		if i == Prev+1 { // two operators in a row
			ToGo := 1
			Index := i
			for {
				exp_, _ := Expressions.Get(i + ToGo)
				ToGo++
				if exp_ == nil {
					break
				}
				exp := exp_.(*JSOpcode)
				if exp.Type != JsOperator {
					break
				}
			}

			if ToGo == Expressions.Size() {
				ToGo--
				Index++
			}
			if ToGo > 1 {
				i--
				if !SubordinateEquation(Expressions, Index, ToGo) {
					return nil
				}
				continue
			}
		}

		Prev = i
		if i == 0 {
			continue
		}

		Level := GetOperatorLevel(ExpStr)
		if Level < 0 {
			return nil
		}
		//ASSERT(Level < OpLevels);
		order_back, _ := Order.Get(Order.Size() - 1)
		if Level != order_back.(int) {
			if Level < order_back.(int) { // step down, sobordinate expression
				Index_, _ := Start.Get(order_back.(int))
				Index := Index_.(int) - 1
				Order.Remove(Order.Size() - 1)
				ToGo := i - Index
				if (bSub && ToGo < Expressions.Size()) || (!bSub && ToGo < Expressions.Size()-2) {
					i = Index - 1
					if !SubordinateEquation(Expressions, Index, ToGo) {
						return nil
					}
					continue
				}
			}

			order_back, _ = Order.Get(Order.Size() - 1)
			if Level > order_back.(int) { // step up
				Order.Add(Level)
				Start.Remove(Level)
				Start.Insert(Level, i)
			}
		}
	}

	// make sure stuff like a = b = c + d results in -> a = (b = c + d), in order to set a and b to the value c + d
	i := 2
	if bSub {
		i = 0
	}
	for ; i < Expressions.Size(); i++ {
		exp_, _ := Expressions.Get(i)
		exp := exp_.(*JSOpcode)
		if exp.Type == JsEquation {
			continue
		}

		if i > 0 && exp.Type == JsOperator && IsAssignmentOp(exp.Value.(*JSOperator).Op) {
			SubordinateEquation(Expressions, i-1, Expressions.Size()-(i-1))
		}
	}

	return Expressions
}

func DumpEquation(eq *arraylist.List) {

	str := ""
	for i := 0; i < eq.Size(); i++ {
		cur_, _ := eq.Get(i)
		cur := cur_.(*JSOpcode)
		str += BuildJSBlock(cur, 0)
	}
	fmt.Printf("%s\n", str)
}

////////////////////////////////////////////////////////////////////////////////////////////
// rebuild JS, used for debuging only
//

func AddSemiIfNeeded(str *string) {
	if (*str)[len(*str)-1:] != "\n" {
		*str += ";\n"
	}
}

func BuildJSBlock(op *JSOpcode, indent int) string {

	inEq := (indent & 0x80000000) != 0
	indent &= 0x7FFFFFFF

	str := ""

	if op.Type == JsFunction {
		fx := op.Value.(*JSFunction)
		str += fx.Name
		str += "("
		for i := 0; i < fx.Arguments.Size(); i++ {
			if i > 0 {
				str += ", "
			}
			arg_, _ := fx.Arguments.Get(i)
			arg := arg_.(*JSNameValue)
			str += arg.Name
			if arg.Value.Params.Size() > 0 {
				str += " = "
				str += BuildJSBlock(&JSOpcode{Type: JsEquation, Value: arg.Value}, indent+1)
			}
		}
		str += ")"
		str += "{\n"
		for i := 0; i < fx.FuncBody.Size(); i++ {
			tmp_, _ := fx.FuncBody.Get(i)
			tmp := tmp_.(*JSOpcode)
			str += PrintIndent(indent)
			str += BuildJSBlock(tmp, indent+1)
			AddSemiIfNeeded(&str)
		}
		str += "}\n\n"
	} else if op.Type == JsVarInit {
		init := op.Value.(*JSVarInit)
		if init.Const {
			str += "const "
		} else {
			str += "let "
		}
		if len(init.Names) == 1 {
			str += init.Names[0]
			str += " = "
			str += BuildJSBlock(&JSOpcode{Type: JsEquation, Value: init.Value}, indent+1)
		} else {
			// TODO: !!!!! add this
		}
	} else if op.Type == JsVarName {
		val := op.Value.(*JSVarName)
		str += val.Name
		str += " "
	} else if op.Type == JsMember {
		val := op.Value.(*JSMember)
		str += BuildJSBlock(val.Object, indent+1)
		str += "["
		str += BuildJSBlock(&JSOpcode{Type: JsEquation, Value: val.Member}, indent+1)
		str += "]"
	} else if op.Type == JsConstVal {
		val := op.Value.(*JSConstVal)
		str += val.Value
		str += " "
	} else if op.Type == JsObject {
		obj := op.Value.(*JSObject)
		str += "{"
		if obj.Params.Size() > 1 {
			str += "\n"
		}
		for i := 0; i < obj.Params.Size(); i++ {
			//if i > 0 {
			//	str += ",\n"
			//}
			arg_, _ := obj.Params.Get(i)
			arg := arg_.(*JSNameValue)
			if obj.Params.Size() > 1 {
				str += PrintIndent(indent)
			}
			str += arg.Name
			str += ":"
			str += BuildJSBlock(&JSOpcode{Type: JsEquation, Value: arg.Value}, indent+1)
			if obj.Params.Size() > 1 {
				str += ",\n"
			}
		}
		if obj.Params.Size() > 1 {
			str += PrintIndent(indent)
		}
		str += "}"
	} else if op.Type == JsArray {
		arr := op.Value.(*JSArray)
		str += "["
		if arr.Params.Size() > 0 {
			str += "\n"
		}
		for i := 0; i < arr.Params.Size(); i++ {
			//if i > 0 {
			//	str += ",\n"
			//}
			arg_, _ := arr.Params.Get(i)
			arg := arg_.(*JSEquation)
			str += PrintIndent(indent)
			str += BuildJSBlock(&JSOpcode{Type: JsEquation, Value: arg}, indent+1)
			str += ",\n"
		}
		if arr.Params.Size() > 0 {
			str += PrintIndent(indent)
		}
		str += "]"
	} else if op.Type == JsOperator {
		val := op.Value.(*JSOperator)
		str += val.Op
		str += " "
	} else if op.Type == JsEquation {
		eq := op.Value.(*JSEquation)

		if inEq {
			str += "("
		}
		for i := 0; i < eq.Params.Size(); i++ {
			eq, _ := eq.Params.Get(i)
			str += BuildJSBlock(eq.(*JSOpcode), (indent+1)|0x80000000)
		}
		if inEq {
			str += ")"
		}

	} else if op.Type == JsFuncCall {
		fx := op.Value.(*JSFuncCall)
		str += BuildJSBlock(fx.Func, indent+1)
		//str += fx.Name
		str += "("
		for i := 0; i < fx.Arguments.Size(); i++ {
			if i > 0 {
				str += ", "
			}
			arg_, _ := fx.Arguments.Get(i)
			arg := arg_.(*JSEquation)
			str += BuildJSBlock(&JSOpcode{Type: JsEquation, Value: arg}, indent+1)
		}
		str += ")"
	} else if op.Type == JsCondition {
		cond := op.Value.(*JSCondition)
		//str += PrintIndent(indent)
		str += "if ("
		str += BuildJSBlock(&JSOpcode{Type: JsEquation, Value: cond.Condition}, indent+1)
		str += ")"
		str += "{\n"
		for i := 0; i < cond.CodeBlock.Opcodes.Size(); i++ {
			tmp_, _ := cond.CodeBlock.Opcodes.Get(i)
			tmp := tmp_.(*JSOpcode)
			str += PrintIndent(indent)
			str += BuildJSBlock(tmp, indent+1)
			AddSemiIfNeeded(&str)
		}
		str += PrintIndent(indent)
		str += "}\n"
		if cond.ElseBlock != nil {
			str += PrintIndent(indent)
			str += "else"
			str += "{\n"
			for i := 0; i < cond.ElseBlock.Opcodes.Size(); i++ {
				tmp_, _ := cond.ElseBlock.Opcodes.Get(i)
				tmp := tmp_.(*JSOpcode)
				str += PrintIndent(indent + 1)
				str += BuildJSBlock(tmp, indent+1+1)
				AddSemiIfNeeded(&str)
			}
			str += PrintIndent(indent)
			str += "}\n"
		}
	} else if op.Type == JsFor {
		loop := op.Value.(*JSFor)
		//str += PrintIndent(indent)
		str += "for ("
		str += BuildJSBlock(loop.Init, indent+1)
		str += "; "
		str += BuildJSBlock(&JSOpcode{Type: JsEquation, Value: loop.Condition}, indent+1)
		str += "; "
		str += BuildJSBlock(&JSOpcode{Type: JsEquation, Value: loop.Next}, indent+1)
		str += ")"
		str += "{\n"
		for i := 0; i < loop.CodeBlock.Opcodes.Size(); i++ {
			tmp_, _ := loop.CodeBlock.Opcodes.Get(i)
			tmp := tmp_.(*JSOpcode)
			str += PrintIndent(indent)
			str += BuildJSBlock(tmp, indent+1)
			AddSemiIfNeeded(&str)
		}
		str += "}\n"
	} else if op.Type == JsForEach {
		loop := op.Value.(*JSForEach)
		//str += PrintIndent(indent)
		str += "for ("
		str += loop.Name
		str += " in "
		str += BuildJSBlock(&JSOpcode{Type: JsEquation, Value: loop.In}, indent+1)
		str += ")"
		str += "{\n"
		for i := 0; i < loop.CodeBlock.Opcodes.Size(); i++ {
			tmp_, _ := loop.CodeBlock.Opcodes.Get(i)
			tmp := tmp_.(*JSOpcode)
			str += PrintIndent(indent)
			str += BuildJSBlock(tmp, indent+1)
			AddSemiIfNeeded(&str)
		}
		str += "}\n"
	} else if op.Type == JsReturn {
		ret := op.Value.(*JSReturn)
		str += "return "
		str += BuildJSBlock(&JSOpcode{Type: JsEquation, Value: ret.Value}, indent+1)
		AddSemiIfNeeded(&str)
	} else if op.Type == JsThrow {
		throw := op.Value.(*JSThrow)
		str += "throw "
		str += BuildJSBlock(&JSOpcode{Type: JsEquation, Value: throw.Value}, indent+1)
		AddSemiIfNeeded(&str)
	} else if op.Type == JsTryCatch {
		// TODO: !!!!! add this
	} else if op.Type == JsNew {
		//new := op.Value.(*JSNew)
		str += "new " // TODO: fix me
		//str += BuildJSBlock(&JSOpcode{Type: JsFunction, Value: new.Constr}, indent+1)
		//AddSemiIfNeeded(&str)
	} else if op.Type == JsDelete {
		del := op.Value.(*JSDelete)
		str += "delete "
		//str += BuildJSBlock(del.What, indent+1); // TODO: fix me
		str += BuildJSBlock(&JSOpcode{Type: JsEquation, Value: del.Value}, indent+1)
		AddSemiIfNeeded(&str)
	}

	return str
}

func BuildJS(js *arraylist.List) string {

	str := ""

	for i := 0; i < js.Size(); i++ {
		op, _ := js.Get(i)
		str += BuildJSBlock(op.(*JSOpcode), 1)
	}

	return str
}
