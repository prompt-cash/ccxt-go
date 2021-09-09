package turing_parser

///////////////////////////////////////////////////////////////////////////////////////////////////////
// Port of the Turing Parser to go

import (
	"github.com/emirpasic/gods/lists/arraylist"
	"strings"
)

///////////////////////////////////////////////////////////////////////////////////////////////////////
// SExpression

type ExpressionType uint8

const (
	None ExpressionType = iota
	Word
	Operator
	String
	Multi
	Separator
	WhiteSpace
	Comment
)

type Expression struct {
	Type ExpressionType

	Word  *string
	Multi *Expressions
}

func (exp *Expression) GetDbg() string {
	if exp.Word != nil {
		return *exp.Word
	}
	return exp.Multi.dbg
}

func (exp *Expression) IsWord() bool {
	return exp.Type == Word
}

func (exp *Expression) IsString() bool {
	return exp.Type == String
}

func (exp *Expression) IsNumber() bool {
	return exp.Type == Word && IsNumber((*exp.Word)[0])
}

func (exp *Expression) IsSeparator() bool {
	return exp.Type == Separator
}

func (exp *Expression) IsMulti() bool {
	return exp.Type == Multi
}

func (exp *Expression) IsOperator() bool {
	return exp.Type == Operator
}

func NewWordExpression(word string) *Expression {
	exp := &Expression{}
	if IsOperator(word[0]) {
		exp.Type = Operator
	} else if IsSeparator(word[0]) {
		exp.Type = Separator
	} else if word[0] == '"' || word[0] == '\'' {
		exp.Type = String
	} else {
		exp.Type = Word
	}
	exp.Word = &word
	return exp
}

func NewMultiExpression(multi *Expressions) *Expression {
	exp := &Expression{}
	exp.Type = Multi
	exp.Multi = multi
	return exp
}

///////////////////////////////////////////////////////////////////////////////////////////////////////
// CExpressions

type Expressions struct {
	Brace       byte
	Expressions arraylist.List

	dbg string
}

func (exps *Expressions) GetExp(index int) *Expression { // Note: the original cpp code used **Expression !!!
	exp, found := exps.Expressions.Get(index)
	if !found {
		return nil
	}
	return exp.(*Expression)
}

func (exps *Expressions) GetExpStr(index int) string {
	exp := exps.GetExp(index)
	if exp == nil || exp.IsMulti() {
		return ""
	}
	return *exp.Word
}

func (exps *Expressions) GetExpMulti(index int) *Expressions {
	exp := exps.GetExp(index)
	if exp == nil || !exp.IsMulti() {
		return nil
	}
	return exp.Multi
}

func (exps *Expressions) GetLastWord() string {
	exp := exps.GetExp(exps.Count() - 1)
	if exp == nil || exp.IsMulti() {
		return ""
	}
	return *exp.Word
}

func (exps *Expressions) IsExpMulti(index int) bool {
	exp := exps.GetExp(index)
	return exp != nil && exp.IsMulti()
}

func (exps *Expressions) IsExpOp(index int) bool {
	exp := exps.GetExp(index)
	return exp != nil && exp.IsOperator()
}

func (exps *Expressions) IsExpStr(index int) bool {
	exp := exps.GetExp(index)
	return exp != nil && exp.IsString()
}

func (exps *Expressions) SubordinateExp(index int, togo int) bool {
	if togo == 0 {
		return true
	}
	if index < 0 || index+togo > exps.Expressions.Size() {
		return false
	}

	subExp := &Expressions{}
	subExp.Brace = '('
	for ; togo > 0; togo-- {
		exp, _ := exps.Expressions.Get(index)
		subExp.Expressions.Add(exp)
		exps.Expressions.Remove(index)
	}
	exps.Expressions.Insert(index, NewMultiExpression(subExp))
	return true
}

func (exps *Expressions) Add(exp *Expression) {
	exps.Expressions.Add(exp)
}

func (exps *Expressions) Count() int {
	return exps.Expressions.Size()
}

func (exps *Expressions) Del(index int) {
	exps.Expressions.Remove(index)
}

func (exps *Expressions) GetBrace(opening bool) string {
	if opening {
		if exps.Brace != 0 {
			return string(exps.Brace)
		}
		return "<"
	} else {
		switch exps.Brace {
		case '(':
			return ")"
		case '[':
			return "]"
		case '{':
			return "}"
		case 0:
			break
		default:
			return string(exps.Brace)
		}

		return ">"
	}
}

///////////////////////////////////////////////////////////////////////////////////////////////////////
// CTuringOps

/*
*	Operators:		"*+-/%&^|!=<>:?~."
* 	Separator:		",;"
*	Brackets:		"()[]{}"
*	Numbers:		"0123456789"
*	Unused:			"#$@"
 */

func IsOperator(c byte) bool {
	//ops := "*+-/%&^|!=<>:?~."
	ops := "*+-/%&^|!=<>:?~"
	return strings.Index(ops, string(c)) != -1
}

func IsNumber(c byte) bool {
	ops := "0123456789" // TODO: Urgent fix 123.0 !!! . is an operator and causes string splitting, dont!
	return strings.Index(ops, string(c)) != -1
}

func IsSeparator(c byte) bool {
	ws := ";,"
	return strings.Index(ws, string(c)) != -1
}

func IsWhiteSpace(c byte) bool {
	ws := " \t\r\n"
	return strings.Index(ws, string(c)) != -1
}

///////////////////////////////////////////////////////////////////////////////////////////////////////
// CTuringParser

/*func GetBlock(str string, start int) (*Expressions, int) {

	var block *Expressions = &Expressions{}

	dbgLine := str[start:]
	_ = dbgLine

	i := start
	strLen := len(str)
	for ; i < strLen; i++ {

		dbgLine2 := str[i:]
		_ = dbgLine2

		exps, j := GetExpressions(str, i)
		if exps != nil {
			i = j
		} else {
			return nil, i
		}

		//if exps.Count() > 0 || block.Count() == 0 {
		block.Add(NewMultiExpression(exps))
		//}

		if i >= strLen || str[i] == ')' || str[i] == ']' || str[i] == '}' {
			break
		}
		//if str[i] == ',' {
		//	block.Add(NewWordExpression(","))
		//}
	}

	if block.Expressions.Size() == 1 {
		return block.GetExpMulti(0), i
	}
	return block, i
}*/

/*func GetBlock(str string, start int) (*Expressions, int) {

	var exps *Expressions = nil
	var block *Expressions = nil

	i := start
	strLen := len(str)
	for ; i < strLen; i++ {

		if exps != nil {
			if block == nil {
				block = &Expressions{}
			}
			block.Add(NewMultiExpression(exps))
		}

		var j int
		exps, j = GetExpressions(str, i)
		if exps != nil {
			i = j
		} else {
			return nil, i
		}

		if i >= strLen || str[i] == ')' || str[i] == ']' || str[i] == '}' {
			break
		}
	}

	if block != nil {
		if exps.Count() > 0 {
			block.Add(NewMultiExpression(exps))
		}
		return block, i
	}
	return exps, i
}*/

func GetBlock(str string, start int) (*Expressions, int) {

	var exps *Expressions = nil
	var block *Expressions = nil
	addComa := false

	i := start
	strLen := len(str)
	for ; i < strLen; i++ {

		if exps != nil {
			if block == nil {
				block = &Expressions{}
			}
			block.Add(NewMultiExpression(exps))

			if addComa {
				block.Add(NewWordExpression(","))
			}
		}
		addComa = false

		var j int
		exps, j = GetExpressions(str, i)
		if exps != nil {
			i = j
		} else {
			return nil, i
		}

		if i >= strLen || str[i] == ')' || str[i] == ']' || str[i] == '}' {
			break
		}

		if str[i] == ',' {
			addComa = true
		}
	}

	if block != nil {
		if exps.Count() > 0 {
			block.Add(NewMultiExpression(exps))

			if addComa {
				block.Add(NewWordExpression(","))
			}
		}
		return block, i
	}
	return exps, i
}

func GetExpressions(line string, start int) (*Expressions, int) {
	var exps *Expressions = &Expressions{}
	var exp string = ""

	dbgLine := line[start:]
	_ = dbgLine

	var esc bool = false
	var prime byte = 0
	var expType ExpressionType = None

	lineLen := len(line)
	i := start
	for ; i < lineLen; i++ {
		var char byte = line[i]

		dbgStr := line[i:]
		_ = dbgStr

		var end int = 0
		if prime != 0 { // inside a string
			if esc {
				esc = false

				/*switch char {
				case '\\':
					exp += "\\"
					break
				case '\'':
					exp += "'"
					break
				case '"':
					exp += "\""
					break
				case 'a':
					exp += "\a"
					break
				case 'b':
					exp += "\b"
					break
				case 'f':
					exp += "\f"
					break
				case 'n':
					exp += "\n"
					break
				case 'r':
					exp += "\r"
					break
				case 't':
					exp += "\t"
					break
				case 'v':
					exp += "\v"
					break
				}*/
				exp += "\\" + string(char)
				continue
			} else if char == '\\' {
				esc = true
				continue
			}

			if char == prime {
				prime = 0
				end = 1
			}
		} else if char == '"' || char == '\'' { // begin of a string
			prime = char
			end = -1
			expType = None
		} else if char == '(' || char == '[' || char == '{' {
			end = 3 // end current word and parse the sub expression/block
		} else if char == ')' || char == ']' || char == '}' {
			break
		} else {
			newType := None
			if IsOperator(char) {
				newType = Operator
			} else if IsSeparator(char) {
				newType = Separator
				end = 4
			} else if IsWhiteSpace(char) {
				newType = WhiteSpace
				end = 2
			}

			if newType != expType {
				expType = newType
				if expType == None || expType == Operator {
					end = -1
				}
			}
		}

		if end == 0 || end == 1 {
			exp += string(char)
		}
		if end != 0 {

			if len(exp) > 0 {

				if exp == "//" || exp == "/*" { // comments
					for ; i < lineLen; i++ {
						if exp == "//" && (line[i] == '\r' || line[i] == '\n') { // skip single lien comment, keep line break
							i--
							break
						}
						if exp == "/*" && (line[i-1] == '*' && line[i] == '/') { // skip multi lien comment
							break
						}
					}
					exp = ""
					continue
				} else {
					exps.Add(NewWordExpression(exp))
					exp = ""
				}
			}

			if end == 4 { // separator
				//exps.Add(NewWordExpression(string(char)))
				//if char == ';' { // break on ; but not on ,
				break
				//}
			}

			if end == 3 {
				subExps, j := GetBlock(line, i+1)
				if subExps == nil {
					return nil, i // pass the error
				} else {
					i = j // +1 done by the for
					subExps.Brace = char
					exps.Add(NewMultiExpression(subExps))
					if char == '{' {
						i++
						break
					}
				}
			}

			if end == -1 {
				exp += string(char)
			}
		}
	}

	if len(exp) > 0 {
		exps.Add(NewWordExpression(exp))
	}

	if prime != 0 { // if the string is not complete its a irrecoverable error
		return nil, i
	}
	return exps, i
}

func EscapeString(str string) string {
	printer := ""
	lineLen := len(str) - 1
	for i := 1; i < lineLen; i++ {
		var char byte = str[i]

		switch char {
		case '\\':
			printer += "\\\\"
			break
		case '\'':
			printer += "'"
			break
		case '"':
			printer += "\\\""
			break
		case '\a':
			printer += "\\a"
			break
		case '\b':
			printer += "\\b"
			break
		case '\f':
			printer += "\\f"
			break
		case '\n':
			printer += "\\n"
			break
		case '\r':
			printer += "\\r"
			break
		case '\t':
			printer += "\\t"
			break
		case '\v':
			printer += "\\v"
			break
		default:
			printer += string(char)
		}
	}
	return "\"" + printer + "\""
}

func PrintIndent(indent int) string {
	printer := ""
	for i := 0; i < indent; i++ {
		printer += "  "
	}
	return printer
}

func PrintBlocks(exps *Expressions, indent int, aux int) string {
	printer := ""
	/*if exps.Brace == '{' {
		printer += "\r\n"
		printer += PrintIndent(indent)
	}*/

	isJson := (exps.Count() >= 3 && exps.GetExpStr(1) == ":")

	for i := 0; i < exps.Count(); i++ {
		exp := exps.GetExp(i)

		/*if !isJson && inFor && i > 0 {
			printer += ";"
		}*/

		if exp.IsMulti() {
			if exp.Multi != nil {

				newAux := 0
				if i > 0 && exps.GetExpStr(i-1) == "for" {
					newAux = 1
				}
				if exp.Multi.Count() > 2 && exp.Multi.Brace == '[' && exp.Multi.GetExpStr(1) == "," {
					newAux = 2
				}
				if exp.Multi.Count() > 1 && exp.Multi.Brace == '{' {
					newAux = 3
				}

				printer += exp.Multi.GetBrace(true)
				if newAux == 3 || newAux == 2 {
					printer += "\r\n"
				}
				printer += PrintBlocks(exp.Multi, indent+1, newAux)
				if newAux == 3 || newAux == 2 {
					if printer[len(printer)-1:] != "\n" {
						printer += "\r\n"
					}
				}
				printer += exp.Multi.GetBrace(false)
				if newAux == 3 && exps.GetExpStr(i+1) != ";" && exps.GetExpStr(i+1) != "," && exps.GetExpStr(i+1) != "else" {
					printer += "\r\n"
				}
			}
		} else {
			if exp.IsString() {
				printer += EscapeString(*exp.Word)
			} else {
				printer += *exp.Word
				if aux != 1 && *exp.Word == ";" {
					printer += "\r\n"
				}
			}
			if (isJson || aux == 2) && *exp.Word == "," {
				printer += "\r\n" // + PrintIndent(indent)
			}
		}

		/*if !isJson && exps.Brace == '{' {
			//printer += ";\r\n"
			printer += "\r\n"
		} else*/{
			if exp.IsWord() && exps.Count() > i+1 && (exps.GetExp(i+1).IsWord() || exps.GetExp(i+1).IsString()) {
				printer += " " // add a space if there are two words following each other
			}
		}
	}
	return printer
}

func PrintAllBlocks(exps *Expressions) string {
	return PrintBlocks(exps, 0, 0)
}

func PrintExpressionImpl(exps *Expressions, from int, indent int) string {

	printer := PrintIndent(indent)

	for i := from; i < exps.Count(); i++ {
		exp := exps.GetExp(i)
		if len(printer) > 0 && printer[len(printer)-1] != ' ' {
			printer += " "
		}

		if exp.IsMulti() {
			if exp.Multi != nil {
				printer += "\n" + PrintIndent(indent) + exp.Multi.GetBrace(true) + "\n" + PrintIndent(indent)
				printer += PrintExpressionImpl(exp.Multi, 0, indent+1)
				printer += "\n" + PrintIndent(indent) + exp.Multi.GetBrace(false) + "\n"
			}
		} else if exp.IsString() {
			printer += EscapeString(*exp.Word)
		} else {
			printer += *exp.Word
		}
	}

	return printer
}

func PrintExpression(exps *Expressions) string {
	return PrintExpressionImpl(exps, 0, 0)
}

///////////////////////////////////////////////////////////////////////////////////////////////////////
//

func ParseFile(fileData string) *Expressions {

	//blocks,_ := GetBlosks(fileData, 0)
	block, _ := GetBlock(fileData, 0)
	//block.Brace = '{'

	//dump := PrintExpressions(block)
	//dump := DumpExpressions(block, 0);
	//fmt.Println(dump)

	return block

	/*i := 0
	lineLen := len(fileData)
	for i < lineLen {

		subStr := fileData[i:]
		_ = subStr
		exps, j := GetExpressions(fileData, i, false)
		i = j

		/dump := turing_parser.DumpExpressions(exps);
		fmt.Println(dump)/

		newLine := PrintExpressions(exps)
		fmt.Println(newLine)

		if fileData[i] == ';' || fileData[i] == '\r' || fileData[i] == '\n' {
			i++
		} else if fileData[i] == '{' {
			i++
			fmt.Println("{")
			j := ParseFile(fileData, i)
			i = j
			fmt.Println("}")
		} else if fileData[i] == '}' {
			break
		}
	}

	i--
	return i*/

	/*dump := turing_parser.DumpExpressions(exps);



	fmt.Println(line)
	fmt.Println(dump)
	fmt.Println(newLine)*/

}

func AddDebugStrings(blocks *Expressions) {
	for i := 0; i < blocks.Count(); i++ {
		exp := blocks.GetExp(i)
		if exp.IsMulti() {
			exp.Multi.dbg = PrintAllBlocks(exp.Multi)
			AddDebugStrings(exp.Multi)
		}
	}
}

///////////////////////////////////////////////////////////////////////////////////////////////////////
//

/**
* This function is used to sort equation execution in mathematically correct order:
*	Hierarchy:
*		1) * / !
*		2) + - =
*		3) == != < > <= >= ~~
*		4) && ||
*		5) : ?
 */

/*var OpLevels int = 7

func GetOperatorLevel(str string) int {
	if len(str) == 0 {
		return 0 // this is a multi expression
	}
	if str == "*" || str == "/" || str == "!" { // high prio math op
		return 6
	}
	if str == "+" || str == "-" || str == "=" { // low prio math op
		return 5
	}
	if str == "&" || str == "|" { // str op
		return 4
	}
	if str == "==" || str == "!=" || str == "~=" || str == "~~" || str == "<" || str == ">" || str == "<=" || str == ">=" { // comparation
		return 3
	}
	if str == "&&" || str == "||" { // logic op
		return 2
	}
	if str == ":" || str == "?" { // other op
		return 1
	}
	return 0 // not a operator
}*/

/*func OrderEquation(Expressions *Expressions, bSub bool) bool {

	// Make sure normal equations a = b & a & c work proeprly and dont overwrite inproperly
	Operator := Expressions.GetExpStr(1)
	if !bSub && Expressions.Count() > 3 && (Operator == "=" || Operator == ":=") {
		Expressions.SubordinateExp(2, Expressions.Count()-2)
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
	for i := 0; i < Expressions.Count()+1; i++ {
		ExpStr := ""
		if i < Expressions.Count() {
			Expression := Expressions.GetExp(i)

			if Expression.IsMulti() {
				if i > 0 && !Expressions.IsExpOp(i-1) { // this is a function inside a equation

					/if (!m_compile_equations)
					{
						// for function calls when nativly executing equations, subordinate arguments
						int arg_couner = 1;
						for (int j = 0; j < Expression->Exp.Multi->Count();)
						{
							if (Expression->Exp.Multi->GetExpStr(j) == L",") // for arg1, , arg3
							{
								Expression->Exp.Multi->Del(j);
								arg_couner++;
								continue;
							}

							int ToGo = 0;
							while ((Expression->Exp.Multi->GetExpStr(j + ++ToGo) != L",") && (j + ToGo < Expression->Exp.Multi->Count()));
							if (!Expression->Exp.Multi->SubordinateExp(j, ToGo))
								return false;

							SExpression* SubExpression = *Expression->Exp.Multi->GetExp(j);
							Operator = SubExpression->Exp.Multi->GetExpStr(1);
							if (Operator != L"=" && Operator != L":=")
							{
								SubExpression->Exp.Multi->Insert(new SExpression(L"." + int2wstring(arg_couner)));
								SubExpression->Exp.Multi->Insert(new SExpression(L"="), 1);
							}
							arg_couner++;

							if (!OrderEquation(SubExpression->Exp.Multi, false))
								return false;
							Expression->Exp.Multi->Del(j-- + 1);

							j+=2;
						}
					}/

					// put a () around the entire function call in order not to cofuse the later sorting process
					if Expressions.Count() > 2 {
						i--
						if !Expressions.SubordinateExp(i, 2) {
							return false
						}
					}
					continue
				}

				if !OrderEquation(Expression.Multi, true) {
					return false
				}
				continue
			}

			if !Expression.IsOperator() {
				continue
			}
			ExpStr = *Expression.Word
		}

		if (bSub && i < 0) || (!bSub && i < 2) {
			continue
		}

		if i == Prev+1 { // two operators in a row
			ToGo := 1
			Index := i
			for {
				tmp := !Expressions.IsExpOp(i + ToGo)
				ToGo++
				if tmp {
					break
				}
			}

			if ToGo == Expressions.Count() {
				ToGo--
				Index++
			}
			if ToGo > 1 {
				i--
				if !Expressions.SubordinateExp(Index, ToGo) {
					return false
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
			return false
		}
		//ASSERT(Level < OpLevels);
		order_back, _ := Order.Get(Order.Size() - 1)
		if Level != order_back.(int) {
			order_back, _ = Order.Get(Order.Size() - 1)
			if Level < order_back.(int) { // step down, sobordinate expression
				Index_, _ := Start.Get(order_back.(int))
				Index := Index_.(int) - 1
				Order.Remove(order_back.(int))
				ToGo := i - Index
				if (bSub && ToGo < Expressions.Count()) || (!bSub && ToGo < (Expressions.Count()-2)) {
					i = Index - 1
					if !Expressions.SubordinateExp(Index, ToGo) {
						return false
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
	i := 0
	if !bSub {
		i = 2
	}
	for ; i < Expressions.Count(); i++ {
		Expression := Expressions.GetExp(i)
		if Expression.IsMulti() {
			continue
		}

		if i > 0 && *Expression.Word == "=" {
			Expressions.SubordinateExp(i-1, Expressions.Count()-(i-1))
		}
	}

	return true
}*/
