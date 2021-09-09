package transpiler

import (
	"fmt"
	"github.com/prompt-cash/js-transpiler/log"
	"github.com/prompt-cash/js-transpiler/pkg/turing_parser"
	"github.com/prompt-cash/js-transpiler/utils"
	"io/ioutil"
	"path"
	"path/filepath"
	"strings"
)

type TranspileParams struct {
	Input  string
	Output string
}

type TranspilerOptions struct {
	CCxtDir   string
	CCxtGoDir string
}

type Transpiler struct {
	ccxtDir   string
	ccxtGoDir string
	logger    log.Logger
}

func NewTranspiler(options *TranspilerOptions, logger log.Logger) (*Transpiler, error) {
	if options == nil {
		options = &TranspilerOptions{}
	}

	transpiler := &Transpiler{
		ccxtDir:   options.CCxtDir,
		ccxtGoDir: options.CCxtGoDir,
		logger: logger.WithFields(log.Fields{
			"module": "transpiler",
		}),
	}

	err := utils.EnsureDirectoryExists(transpiler.ccxtGoDir)
	if err != nil {
		return nil, err
	}

	return transpiler, nil
}

func (t Transpiler) TranspileFiles() error {

	if t.ccxtDir[len(t.ccxtDir)-3:] == ".js" {

		output, err := t.Transpile(t.ccxtDir)
		if err != nil {
			return err
		}

		fileName := filepath.Base(t.ccxtDir)
		moduleName := fileName[:len(fileName)-len(filepath.Ext(fileName))]

		fmt.Printf("Tanspiled: %s\n", fileName)

		out_data := []byte(output)
		err = ioutil.WriteFile(path.Join(t.ccxtGoDir, fmt.Sprintf("ex_%s.go", moduleName)), out_data, 0644)

	} else {

		files, err := ioutil.ReadDir(t.ccxtDir)
		if err != nil {
			return err
		}
		for i := 0; i < len(files); i++ {
			file := files[i]
			if file.IsDir() || file.Name()[len(file.Name())-3:] != ".js" {
				continue
			}

			filePath := path.Join(t.ccxtDir, file.Name())
			output, err := t.Transpile(filePath)
			if err != nil {
				t.logger.Errorf("Error transpiling file %s: %+v", filePath, err)
				continue
			}

			fileName := file.Name()
			moduleName := fileName[:len(fileName)-len(filepath.Ext(fileName))]

			fmt.Printf("Tanspiled: %s\n", fileName)

			out_data := []byte(output)
			err = ioutil.WriteFile(path.Join(t.ccxtGoDir, fmt.Sprintf("ex_%s.go", moduleName)), out_data, 0644)
		}

	}
	return nil
}

type JsPatch struct {
	Find    string
	Replace string
}

type JsPatches []JsPatch

var PatchList = map[string]JsPatches{}

func init() {
	PatchList["bitmart.js"] = JsPatches{JsPatch{
		"currency = this.currenc (code);", "currency = this.currency (code);", // typo in the original code
	}}
	PatchList["btcturk.js"] = JsPatches{JsPatch{
		"price = this.parseNumber (precisePrice);", "price = this.parseNumber (precisePrice.toString());", // go does nto do impliite tostring conversion
	}}
	PatchList["okex3.js"] = JsPatches{JsPatch{
		"settlementCurrencyСode", "settlementCurrencyCode", // the C is some non ascii variation
	}}
	PatchList["therock.js"] = JsPatches{JsPatch{
		"headers = (headers === undefined) ? {} : headers;", "headers = (headers === undefined) ? ({}) : headers;", // the parser always breaks an expression at } so we need a workaround
	}}
}

func (t Transpiler) Transpile(s string) (string, error) {

	// first we read the js file
	content, err := ioutil.ReadFile(s)
	if err != nil {
		return "", err
	}

	str := string(content)

	name := filepath.Base(s)

	// our parser is good but not perfect, a few special cases need special handling
	// apply the required patches
	// TODO pull these changes to the official ccxt repor
	patches := PatchList[name]
	if patches != nil {
		for i := 0; i < len(patches); i++ {
			str = strings.ReplaceAll(str, patches[i].Find, patches[i].Replace)
		}
	}

	// parse the js code using ago port of an old C/C++ based generic script parser
	blocks := turing_parser.ParseFile(str)
	turing_parser.AddDebugStrings(blocks)

	jsCompiler, _ := turing_parser.NewJSCompiler(t.logger)

	moduleExport := jsCompiler.GetModuleExport(blocks)

	moduleName := moduleExport.GetExpStr(3)
	moduleBody := moduleExport.GetExpMulti(6)

	// to make the transpilation easier we compile the parsed javascript into a sort of byte code
	js := jsCompiler.CompileJsBlocks(moduleBody)

	goFile := &GoFile{
		Package: moduleName,
		Imports: []string{
			//"MyBase",
		},
	}

	generator, err := NewGoGenerator(allBaseMethods(), t.logger)
	if err != nil {
		return "", err
	}

	// do the transpilation
	output := generator.Genrate(js, goFile)
	//fmt.Println(output)

	return output, nil
}

/*func init() {
	baseType := reflect.TypeOf(&ccxt.ExchangeBase{})

	// get all exchange base functions
	for i := 0; i < baseType.NumMethod(); i++ {
		method := baseType.Method(i)
		//ExchangeBaseMethods[method.Name] = true
		fmt.Printf("BaseMethods[\"" +method.Name+ "\"] = true\n")
	}

	// get all exchange base variables
	//baseType2 := reflect.TypeOf(ccxt.ExchangeBase{})
	//for i := 0; i < baseType2.NumField(); i++ {
	//	field := baseType2.Field(i)
	//	ExchangeBaseMembers[field.Name] = true
	//}
}*/

func allBaseMethods() map[string]bool {

	//
	// list of all known methods defined on the ExchangeBase object
	// all calls to non listed members of "this" will be refactored to use this.Call("name", ...)
	//
	// TODO keep this up to date i.e. in sync with ExchangeBase
	//

	var BaseMethods = map[string]bool{}
	BaseMethods["Account"] = true
	BaseMethods["AmountToPrecision"] = true
	BaseMethods["ArrayConcat"] = true
	BaseMethods["At"] = true
	BaseMethods["Base16ToBinary"] = true
	BaseMethods["Base58ToBinary"] = true
	BaseMethods["Base64ToBinary"] = true
	BaseMethods["BaseDescribe"] = true
	BaseMethods["BinaryConcat"] = true
	BaseMethods["BinaryConcatArray"] = true
	BaseMethods["BinaryToBase16"] = true
	BaseMethods["BinaryToBase64"] = true
	BaseMethods["BuildApiPath"] = true
	BaseMethods["BuildOHLCVC"] = true
	BaseMethods["CalculateFee"] = true
	BaseMethods["Call"] = true
	BaseMethods["CancelOrder"] = true
	BaseMethods["CancelUnifiedOrder"] = true
	BaseMethods["Capitalize"] = true
	BaseMethods["CheckAddress"] = true
	BaseMethods["CheckRequiredCredentials"] = true
	BaseMethods["CheckRequiredDependencies"] = true
	BaseMethods["CommonCurrencyCode"] = true
	BaseMethods["ConvertTradingViewToOHLCV"] = true
	BaseMethods["CostToPrecision"] = true
	BaseMethods["CreateLimitBuyOrder"] = true
	BaseMethods["CreateLimitOrder"] = true
	BaseMethods["CreateLimitSellOrder"] = true
	BaseMethods["CreateMarketBuyOrder"] = true
	BaseMethods["CreateMarketOrder"] = true
	BaseMethods["CreateMarketSellOrder"] = true
	BaseMethods["CreateOrder"] = true
	BaseMethods["Currency"] = true
	BaseMethods["CurrencyFromPrecision"] = true
	BaseMethods["CurrencyToPrecision"] = true
	BaseMethods["DecimalToPrecision"] = true
	BaseMethods["Decode"] = true
	BaseMethods["DeepExtend"] = true
	BaseMethods["Ecdsa"] = true
	BaseMethods["Eddsa"] = true
	BaseMethods["Encode"] = true
	BaseMethods["ExecuteRestRequest"] = true
	BaseMethods["Extend"] = true
	BaseMethods["ExtractParams"] = true
	BaseMethods["FeeToPrecision"] = true
	BaseMethods["Fetch"] = true
	BaseMethods["Fetch2"] = true
	BaseMethods["FetchClosedOrders"] = true
	BaseMethods["FetchDeposits"] = true
	BaseMethods["FetchMyTrades"] = true
	BaseMethods["FetchOpenOrders"] = true
	BaseMethods["FetchOrder"] = true
	BaseMethods["FetchOrders"] = true
	BaseMethods["FetchTickers"] = true
	BaseMethods["FetchTransactions"] = true
	BaseMethods["FetchUnifiedOrder"] = true
	BaseMethods["FetchWithdrawals"] = true
	BaseMethods["FilterBy"] = true
	BaseMethods["FilterByArray"] = true
	BaseMethods["FilterByCurrencySinceLimit"] = true
	BaseMethods["FilterBySinceLimit"] = true
	BaseMethods["FilterBySymbol"] = true
	BaseMethods["FilterBySymbolSinceLimit"] = true
	BaseMethods["FilterByValueSinceLimit"] = true
	BaseMethods["FindBroadlyMatchedKey"] = true
	BaseMethods["Flatten"] = true
	BaseMethods["FromPrecision"] = true
	BaseMethods["FuturesTransfer"] = true
	BaseMethods["GroupBy"] = true
	BaseMethods["HandleErrors"] = true
	BaseMethods["Hash"] = true
	BaseMethods["HashMessage"] = true
	BaseMethods["Hmac"] = true
	BaseMethods["Id"] = true
	BaseMethods["ImplodeHostname"] = true
	BaseMethods["ImplodeParams"] = true
	BaseMethods["InArray"] = true
	BaseMethods["IndexBy"] = true
	BaseMethods["IsEmpty"] = true
	BaseMethods["IsJsonEncodedObject"] = true
	BaseMethods["Iso8601"] = true
	BaseMethods["Json"] = true
	BaseMethods["Jwt"] = true
	BaseMethods["Keysort"] = true
	BaseMethods["LoadAccounts"] = true
	BaseMethods["LoadMarkets"] = true
	BaseMethods["Market"] = true
	BaseMethods["MarketId"] = true
	BaseMethods["MarketIds"] = true
	BaseMethods["Microseconds"] = true
	BaseMethods["Milliseconds"] = true
	BaseMethods["Nonce"] = true
	BaseMethods["NumberToBE"] = true
	BaseMethods["NumberToLE"] = true
	BaseMethods["NumberToString"] = true
	BaseMethods["Oath"] = true
	BaseMethods["Omit"] = true
	BaseMethods["OmitZero"] = true
	BaseMethods["Ordered"] = true
	BaseMethods["Parse8601"] = true
	BaseMethods["ParseBalance"] = true
	BaseMethods["ParseBidAsk"] = true
	BaseMethods["ParseBidsAsks"] = true
	BaseMethods["ParseDepositAddresses"] = true
	BaseMethods["ParseJson"] = true
	BaseMethods["ParseLedger"] = true
	BaseMethods["ParseNumber"] = true
	BaseMethods["ParseOHLCV"] = true
	BaseMethods["ParseOHLCVs"] = true
	BaseMethods["ParseOrderBook"] = true
	BaseMethods["ParseOrders"] = true
	BaseMethods["ParsePrecision"] = true
	BaseMethods["ParseSafeNumber"] = true
	BaseMethods["ParseTickers"] = true
	BaseMethods["ParseTimeframe"] = true
	BaseMethods["ParseTrades"] = true
	BaseMethods["ParseTradingViewOHLCV"] = true
	BaseMethods["ParseTransactions"] = true
	BaseMethods["ParseTransfers"] = true
	BaseMethods["PrecisionFromString"] = true
	BaseMethods["PriceFromPrecision"] = true
	BaseMethods["PriceToPrecision"] = true
	BaseMethods["Rawencode"] = true
	BaseMethods["ReduceFeesByCurrency"] = true
	BaseMethods["Remove0xPrefix"] = true
	BaseMethods["Request"] = true
	BaseMethods["Rsa"] = true
	BaseMethods["SafeCurrency"] = true
	BaseMethods["SafeCurrencyCode"] = true
	BaseMethods["SafeFloat"] = true
	BaseMethods["SafeFloat2"] = true
	BaseMethods["SafeInteger"] = true
	BaseMethods["SafeInteger2"] = true
	BaseMethods["SafeIntegerProduct"] = true
	BaseMethods["SafeIntegerProduct2"] = true
	BaseMethods["SafeMarket"] = true
	BaseMethods["SafeNumber"] = true
	BaseMethods["SafeNumber2"] = true
	BaseMethods["SafeOrder"] = true
	BaseMethods["SafeString"] = true
	BaseMethods["SafeString2"] = true
	BaseMethods["SafeStringLower"] = true
	BaseMethods["SafeStringLower2"] = true
	BaseMethods["SafeStringUpper"] = true
	BaseMethods["SafeStringUpper2"] = true
	BaseMethods["SafeSymbol"] = true
	BaseMethods["SafeTimestamp"] = true
	BaseMethods["SafeTimestamp2"] = true
	BaseMethods["SafeValue"] = true
	BaseMethods["SafeValue2"] = true
	BaseMethods["Seconds"] = true
	BaseMethods["SetHeaders"] = true
	BaseMethods["SetMarkets"] = true
	BaseMethods["Setup"] = true
	BaseMethods["Sign"] = true
	BaseMethods["SignHash"] = true
	BaseMethods["SignMessage"] = true
	BaseMethods["SignMessageString"] = true
	BaseMethods["SortBy"] = true
	BaseMethods["StringToBase64"] = true
	BaseMethods["StringToBinary"] = true
	BaseMethods["Strip"] = true
	BaseMethods["Sum"] = true
	BaseMethods["ThrowBroadlyMatchedException"] = true
	BaseMethods["ThrowExactlyMatchedException"] = true
	BaseMethods["ToArray"] = true
	BaseMethods["ToMilliseconds"] = true
	BaseMethods["ToPrecision"] = true
	BaseMethods["TryCallAPI"] = true
	BaseMethods["Urlencode"] = true
	BaseMethods["UrlencodeWithArrayRepeat"] = true
	BaseMethods["Uuid"] = true
	BaseMethods["Uuid16"] = true
	BaseMethods["Uuid22"] = true
	BaseMethods["Uuidv1"] = true
	BaseMethods["VCall"] = true
	BaseMethods["Vwap"] = true
	BaseMethods["Ymd"] = true
	BaseMethods["Ymdhms"] = true
	return BaseMethods
}
