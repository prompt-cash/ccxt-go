# CCXT Go
Transpiled [CCXT exchange library](https://github.com/ccxt/ccxt) from their original JavaScript source to Go (Golang).

## Features
- support 100+ cryptocurrency exchanges with a unified API
- fully implemented public and private APIs for fetching market data, trading, arbitrage, etc..


### Requirements
```
Go >= 1.13
```

### Installation
```
go get https://github.com/prompt-cash/ccxt-go
```
Take a look into [ccxt_test.go](pkg/ccxt/ccxt_test.go) for an example how to get started.


## Alpha Release
It’s all still very alpha, so no guarantees of any kind.
Especially some rarely used helper functions are not implemented yet and
just panic with a todo message. If you need these exchanges you will
need to implement those.

## Contact
[Website](https://prompt.cash/) -
[Twitter](https://twitter.com/CashPrompt) -
[Telegram](https://t.me/PromptCash) -
[YouTube](https://www.youtube.com/channel/UClfNVdL3T0RF6pF1yGi9teg)


