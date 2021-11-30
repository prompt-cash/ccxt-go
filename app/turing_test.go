package app

import (
	"fmt"
	"github.com/prompt-cash/ccxt-go/pkg/turing_parser"
	"io/ioutil"
	"testing"
)

func TestParseScript(t *testing.T) {
	//getContextForTesting(t) // needed to load config

	//line := "test.bla(blup[x+(1+2)]); if(test.foo) {return 123;}"
	//fileData := "for (let i = 0, i < codes.length, i++) {\r\nconst code = codes[i];\r\naccount['free'] = this.safeString (balances, currencyId);\r\n}\r\n"

	//turing_parser.ParseFile(fileData, 0)

	//content, err := ioutil.ReadFile("..\\test\\test.js")
	content, err := ioutil.ReadFile("F:\\Projects\\PromptCash\\ccxt-master\\js\\coincheck.js")
	//ccxtDir := getCcxtDir()
	//content, err := ioutil.ReadFile(path.Join(ccxtDir, "js", "coincheck.js"))

	if err != nil {
		t.Log(err)
	}

	str := string(content)
	//str := "test.bla(blup[x+(1+2)]);temp.foo(bar[x+(1+2)]); if(x=123) {window.alert('test');} return 123;"
	//str := "{bla; }"
	//str := "wnd.alert('test');for(let i=0; i<10; i++) {j++;}"
	blocks := turing_parser.ParseFile(str)

	turing_parser.AddDebugStrings(blocks)

	dump := turing_parser.PrintExpression(blocks)
	fmt.Println(dump)

	js := turing_parser.CompileJS(blocks)
	//js := turing_parser.CompileJsFuncBody(blocks)

	output := turing_parser.BuildJS(js)

	_ = output

	out_data := []byte(output)
	ioutil.WriteFile("F:\\Projects\\PromptCash\\ccxt-master\\js\\coincheck.tmp", out_data, 0644)

	return

	//turing_parser.AddDebugStrings(blocks)

	//dump := turing_parser.PrintAllBlocks(blocks)
	//fmt.Println(dump)
}
