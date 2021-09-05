package ccxt

import ()

type Binance struct {
	*ExchangeBase
}

var _ Exchange = (*Binance)(nil)

func init() {
	exchange := &Binance{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Binance) Describe(goArgs ...*Variant) *Variant {
	return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
		"id":   MkString("binance"),
		"name": MkString("Binance"),
		"countries": MkArray(&VarArray{
			MkString("JP"),
			MkString("MT"),
		}),
		"rateLimit": MkInteger(50),
		"certified": MkBool(true),
		"pro":       MkBool(true),
		"has": MkMap(&VarMap{
			"cancelAllOrders":        MkBool(true),
			"cancelOrder":            MkBool(true),
			"CORS":                   MkBool(false),
			"createOrder":            MkBool(true),
			"fetchCurrencies":        MkBool(true),
			"fetchBalance":           MkBool(true),
			"fetchBidsAsks":          MkBool(true),
			"fetchClosedOrders":      MkString("emulated"),
			"fetchDepositAddress":    MkBool(true),
			"fetchDeposits":          MkBool(true),
			"fetchFundingFees":       MkBool(true),
			"fetchFundingHistory":    MkBool(true),
			"fetchFundingRate":       MkBool(true),
			"fetchFundingRates":      MkBool(true),
			"fetchIsolatedPositions": MkBool(true),
			"fetchMarkets":           MkBool(true),
			"fetchMyTrades":          MkBool(true),
			"fetchOHLCV":             MkBool(true),
			"fetchOpenOrders":        MkBool(true),
			"fetchOrder":             MkBool(true),
			"fetchOrders":            MkBool(true),
			"fetchOrderBook":         MkBool(true),
			"fetchPositions":         MkBool(true),
			"fetchStatus":            MkBool(true),
			"fetchTicker":            MkBool(true),
			"fetchTickers":           MkBool(true),
			"fetchTime":              MkBool(true),
			"fetchTrades":            MkBool(true),
			"fetchTradingFee":        MkBool(true),
			"fetchTradingFees":       MkBool(true),
			"fetchTransactions":      MkBool(false),
			"fetchWithdrawals":       MkBool(true),
			"setLeverage":            MkBool(true),
			"setMarginMode":          MkBool(true),
			"withdraw":               MkBool(true),
			"transfer":               MkBool(true),
			"fetchTransfers":         MkBool(true),
		}),
		"timeframes": MkMap(&VarMap{
			"1m":  MkString("1m"),
			"3m":  MkString("3m"),
			"5m":  MkString("5m"),
			"15m": MkString("15m"),
			"30m": MkString("30m"),
			"1h":  MkString("1h"),
			"2h":  MkString("2h"),
			"4h":  MkString("4h"),
			"6h":  MkString("6h"),
			"8h":  MkString("8h"),
			"12h": MkString("12h"),
			"1d":  MkString("1d"),
			"3d":  MkString("3d"),
			"1w":  MkString("1w"),
			"1M":  MkString("1M"),
		}),
		"urls": MkMap(&VarMap{
			"logo": MkString("https://user-images.githubusercontent.com/1294454/29604020-d5483cdc-87ee-11e7-94c7-d1a8d9169293.jpg"),
			"test": MkMap(&VarMap{
				"dapiPublic":    MkString("https://testnet.binancefuture.com/dapi/v1"),
				"dapiPrivate":   MkString("https://testnet.binancefuture.com/dapi/v1"),
				"fapiPublic":    MkString("https://testnet.binancefuture.com/fapi/v1"),
				"fapiPrivate":   MkString("https://testnet.binancefuture.com/fapi/v1"),
				"fapiPrivateV2": MkString("https://testnet.binancefuture.com/fapi/v2"),
				"public":        MkString("https://testnet.binance.vision/api/v3"),
				"private":       MkString("https://testnet.binance.vision/api/v3"),
				"v3":            MkString("https://testnet.binance.vision/api/v3"),
				"v1":            MkString("https://testnet.binance.vision/api/v1"),
			}),
			"api": MkMap(&VarMap{
				"wapi":          MkString("https://api.binance.com/wapi/v3"),
				"sapi":          MkString("https://api.binance.com/sapi/v1"),
				"dapiPublic":    MkString("https://dapi.binance.com/dapi/v1"),
				"dapiPrivate":   MkString("https://dapi.binance.com/dapi/v1"),
				"dapiPrivateV2": MkString("https://dapi.binance.com/dapi/v2"),
				"dapiData":      MkString("https://dapi.binance.com/futures/data"),
				"fapiPublic":    MkString("https://fapi.binance.com/fapi/v1"),
				"fapiPrivate":   MkString("https://fapi.binance.com/fapi/v1"),
				"fapiData":      MkString("https://fapi.binance.com/futures/data"),
				"fapiPrivateV2": MkString("https://fapi.binance.com/fapi/v2"),
				"public":        MkString("https://api.binance.com/api/v3"),
				"private":       MkString("https://api.binance.com/api/v3"),
				"v3":            MkString("https://api.binance.com/api/v3"),
				"v1":            MkString("https://api.binance.com/api/v1"),
			}),
			"www": MkString("https://www.binance.com"),
			"referral": MkMap(&VarMap{
				"url":      MkString("https://www.binance.com/en/register?ref=BLEJC98C"),
				"discount": MkNumber(0.2),
			}),
			"doc": MkArray(&VarArray{
				MkString("https://binance-docs.github.io/apidocs/spot/en"),
			}),
			"api_management": MkString("https://www.binance.com/en/usercenter/settings/api-management"),
			"fees":           MkString("https://www.binance.com/en/fee/schedule"),
		}),
		"depth": MkInteger(1),
		"api": MkMap(&VarMap{
			"sapi": MkMap(&VarMap{
				"get": MkMap(&VarMap{
					"accountSnapshot":                       MkInteger(1),
					"system/status":                         MkInteger(1),
					"margin/asset":                          MkInteger(1),
					"margin/pair":                           MkInteger(1),
					"margin/allAssets":                      MkInteger(1),
					"margin/allPairs":                       MkInteger(1),
					"margin/priceIndex":                     MkInteger(1),
					"asset/assetDividend":                   MkInteger(1),
					"asset/dribblet":                        MkInteger(1),
					"asset/transfer":                        MkInteger(1),
					"asset/assetDetail":                     MkInteger(1),
					"asset/tradeFee":                        MkInteger(1),
					"asset/get-funding-asset":               MkInteger(1),
					"margin/loan":                           MkInteger(1),
					"margin/repay":                          MkInteger(1),
					"margin/account":                        MkInteger(1),
					"margin/transfer":                       MkInteger(1),
					"margin/interestHistory":                MkInteger(1),
					"margin/forceLiquidationRec":            MkInteger(1),
					"margin/order":                          MkInteger(1),
					"margin/openOrders":                     MkInteger(1),
					"margin/allOrders":                      MkInteger(1),
					"margin/myTrades":                       MkInteger(1),
					"margin/maxBorrowable":                  MkInteger(5),
					"margin/maxTransferable":                MkInteger(5),
					"margin/isolated/transfer":              MkInteger(1),
					"margin/isolated/account":               MkInteger(1),
					"margin/isolated/pair":                  MkInteger(1),
					"margin/isolated/allPairs":              MkInteger(1),
					"margin/interestRateHistory":            MkInteger(1),
					"margin/orderList":                      MkInteger(2),
					"margin/allOrderList":                   MkInteger(10),
					"margin/openOrderList":                  MkInteger(3),
					"fiat/orders":                           MkInteger(1),
					"fiat/payments":                         MkInteger(1),
					"futures/transfer":                      MkInteger(5),
					"futures/loan/borrow/history":           MkInteger(1),
					"futures/loan/repay/history":            MkInteger(1),
					"futures/loan/wallet":                   MkInteger(1),
					"futures/loan/configs":                  MkInteger(1),
					"futures/loan/calcAdjustLevel":          MkInteger(1),
					"futures/loan/calcMaxAdjustAmount":      MkInteger(1),
					"futures/loan/adjustCollateral/history": MkInteger(1),
					"futures/loan/liquidationHistory":       MkInteger(1),
					"capital/config/getall":                 MkInteger(1),
					"capital/deposit/address":               MkInteger(1),
					"capital/deposit/hisrec":                MkInteger(1),
					"capital/deposit/subAddress":            MkInteger(1),
					"capital/deposit/subHisrec":             MkInteger(1),
					"capital/withdraw/history":              MkInteger(1),
					"bnbBurn":                               MkInteger(1),
					"sub-account/assets":                    MkInteger(1),
					"sub-account/futures/account":           MkInteger(1),
					"sub-account/futures/accountSummary":    MkInteger(20),
					"sub-account/futures/positionRisk":      MkInteger(1),
					"sub-account/futures/internalTransfer":  MkInteger(1),
					"sub-account/list":                      MkInteger(1),
					"sub-account/margin/account":            MkInteger(1),
					"sub-account/margin/accountSummary":     MkInteger(1),
					"sub-account/spotSummary":               MkInteger(5),
					"sub-account/status":                    MkInteger(1),
					"sub-account/sub/transfer/history":      MkInteger(1),
					"sub-account/transfer/subUserHistory":   MkInteger(1),
					"sub-account/universalTransfer":         MkInteger(1),
					"lending/daily/product/list":            MkInteger(1),
					"lending/daily/userLeftQuota":           MkInteger(1),
					"lending/daily/userRedemptionQuota":     MkInteger(1),
					"lending/daily/token/position":          MkInteger(1),
					"lending/union/account":                 MkInteger(1),
					"lending/union/purchaseRecord":          MkInteger(1),
					"lending/union/redemptionRecord":        MkInteger(1),
					"lending/union/interestHistory":         MkInteger(1),
					"lending/project/list":                  MkInteger(1),
					"lending/project/position/list":         MkInteger(1),
					"mining/pub/algoList":                   MkInteger(1),
					"mining/pub/coinList":                   MkInteger(1),
					"mining/worker/detail":                  MkInteger(5),
					"mining/worker/list":                    MkInteger(5),
					"mining/payment/list":                   MkInteger(5),
					"mining/statistics/user/status":         MkInteger(5),
					"mining/statistics/user/list":           MkInteger(5),
					"bswap/pools":                           MkInteger(1),
					"bswap/liquidity": MkMap(&VarMap{
						"cost":     MkInteger(1),
						"noPoolId": MkInteger(10),
					}),
					"bswap/liquidityOps":                          MkInteger(2),
					"bswap/quote":                                 MkInteger(2),
					"bswap/swap":                                  MkInteger(1),
					"blvt/tokenInfo":                              MkInteger(1),
					"blvt/subscribe/record":                       MkInteger(1),
					"blvt/redeem/record":                          MkInteger(1),
					"blvt/userLimit":                              MkInteger(1),
					"apiReferral/ifNewUser":                       MkInteger(1),
					"apiReferral/customization":                   MkInteger(1),
					"apiReferral/userCustomization":               MkInteger(1),
					"apiReferral/rebate/recentRecord":             MkInteger(1),
					"apiReferral/rebate/historicalRecord":         MkInteger(1),
					"apiReferral/kickback/recentRecord":           MkInteger(1),
					"apiReferral/kickback/historicalRecord":       MkInteger(1),
					"broker/subAccountApi":                        MkInteger(1),
					"broker/subAccount":                           MkInteger(1),
					"broker/subAccountApi/commission/futures":     MkInteger(1),
					"broker/subAccountApi/commission/coinFutures": MkInteger(1),
					"broker/info":                                 MkInteger(1),
					"broker/transfer":                             MkInteger(1),
					"broker/transfer/futures":                     MkInteger(1),
					"broker/rebate/recentRecord":                  MkInteger(1),
					"broker/rebate/historicalRecord":              MkInteger(1),
					"broker/subAccount/bnbBurn/status":            MkInteger(1),
					"broker/subAccount/depositHist":               MkInteger(1),
					"broker/subAccount/spotSummary":               MkInteger(1),
					"broker/subAccount/marginSummary":             MkInteger(1),
					"broker/subAccount/futuresSummary":            MkInteger(1),
					"broker/rebate/futures/recentRecord":          MkInteger(1),
					"broker/subAccountApi/ipRestriction":          MkInteger(1),
					"broker/universalTransfer":                    MkInteger(1),
					"account/apiRestrictions":                     MkInteger(1),
					"managed-subaccount/asset":                    MkInteger(1),
				}),
				"post": MkMap(&VarMap{
					"asset/dust":                                        MkInteger(1),
					"asset/transfer":                                    MkInteger(1),
					"asset/get-funding-asset":                           MkInteger(1),
					"account/disableFastWithdrawSwitch":                 MkInteger(1),
					"account/enableFastWithdrawSwitch":                  MkInteger(1),
					"capital/withdraw/apply":                            MkInteger(1),
					"margin/transfer":                                   MkInteger(1),
					"margin/loan":                                       MkInteger(1),
					"margin/repay":                                      MkInteger(1),
					"margin/order":                                      MkInteger(4),
					"margin/order/oco":                                  MkInteger(1),
					"margin/isolated/create":                            MkInteger(1),
					"margin/isolated/transfer":                          MkInteger(1),
					"bnbBurn":                                           MkInteger(1),
					"sub-account/margin/transfer":                       MkInteger(1),
					"sub-account/margin/enable":                         MkInteger(1),
					"sub-account/futures/enable":                        MkInteger(1),
					"sub-account/futures/transfer":                      MkInteger(1),
					"sub-account/futures/internalTransfer":              MkInteger(1),
					"sub-account/transfer/subToSub":                     MkInteger(1),
					"sub-account/transfer/subToMaster":                  MkInteger(1),
					"sub-account/universalTransfer":                     MkInteger(1),
					"managed-subaccount/deposit":                        MkInteger(1),
					"managed-subaccount/withdraw":                       MkInteger(1),
					"userDataStream":                                    MkInteger(1),
					"userDataStream/isolated":                           MkInteger(1),
					"futures/transfer":                                  MkInteger(1),
					"futures/loan/borrow":                               MkInteger(20),
					"futures/loan/repay":                                MkInteger(20),
					"futures/loan/adjustCollateral":                     MkInteger(20),
					"lending/customizedFixed/purchase":                  MkInteger(1),
					"lending/daily/purchase":                            MkInteger(1),
					"lending/daily/redeem":                              MkInteger(1),
					"bswap/liquidityAdd":                                MkInteger(2),
					"bswap/liquidityRemove":                             MkInteger(2),
					"bswap/swap":                                        MkInteger(2),
					"blvt/subscribe":                                    MkInteger(1),
					"blvt/redeem":                                       MkInteger(1),
					"apiReferral/customization":                         MkInteger(1),
					"apiReferral/userCustomization":                     MkInteger(1),
					"apiReferral/rebate/historicalRecord":               MkInteger(1),
					"apiReferral/kickback/historicalRecord":             MkInteger(1),
					"broker/subAccount":                                 MkInteger(1),
					"broker/subAccount/margin":                          MkInteger(1),
					"broker/subAccount/futures":                         MkInteger(1),
					"broker/subAccountApi":                              MkInteger(1),
					"broker/subAccountApi/permission":                   MkInteger(1),
					"broker/subAccountApi/commission":                   MkInteger(1),
					"broker/subAccountApi/commission/futures":           MkInteger(1),
					"broker/subAccountApi/commission/coinFutures":       MkInteger(1),
					"broker/transfer":                                   MkInteger(1),
					"broker/transfer/futures":                           MkInteger(1),
					"broker/rebate/historicalRecord":                    MkInteger(1),
					"broker/subAccount/bnbBurn/spot":                    MkInteger(1),
					"broker/subAccount/bnbBurn/marginInterest":          MkInteger(1),
					"broker/subAccount/blvt":                            MkInteger(1),
					"broker/subAccountApi/ipRestriction":                MkInteger(1),
					"broker/subAccountApi/ipRestriction/ipList":         MkInteger(1),
					"broker/universalTransfer":                          MkInteger(1),
					"broker/subAccountApi/permission/universalTransfer": MkInteger(1),
					"broker/subAccountApi/permission/vanillaOptions":    MkInteger(1),
				}),
				"put": MkMap(&VarMap{
					"userDataStream":          MkInteger(1),
					"userDataStream/isolated": MkInteger(1),
				}),
				"delete": MkMap(&VarMap{
					"margin/openOrders":                         MkInteger(1),
					"margin/order":                              MkInteger(1),
					"margin/orderList":                          MkInteger(1),
					"userDataStream":                            MkInteger(1),
					"userDataStream/isolated":                   MkInteger(1),
					"broker/subAccountApi":                      MkInteger(1),
					"broker/subAccountApi/ipRestriction/ipList": MkInteger(1),
				}),
			}),
			"wapi": MkMap(&VarMap{
				"post": MkMap(&VarMap{
					"withdraw":             MkInteger(1),
					"sub-account/transfer": MkInteger(1),
				}),
				"get": MkMap(&VarMap{
					"depositHistory":               MkInteger(1),
					"withdrawHistory":              MkInteger(1),
					"depositAddress":               MkInteger(1),
					"accountStatus":                MkInteger(1),
					"systemStatus":                 MkInteger(1),
					"apiTradingStatus":             MkInteger(1),
					"userAssetDribbletLog":         MkInteger(1),
					"tradeFee":                     MkInteger(1),
					"assetDetail":                  MkInteger(1),
					"sub-account/list":             MkInteger(1),
					"sub-account/transfer/history": MkInteger(1),
					"sub-account/assets":           MkInteger(1),
				}),
			}),
			"dapiPublic": MkMap(&VarMap{"get": MkMap(&VarMap{
				"ping":         MkInteger(1),
				"time":         MkInteger(1),
				"exchangeInfo": MkInteger(1),
				"depth": MkMap(&VarMap{
					"cost": MkInteger(2),
					"byLimit": MkArray(&VarArray{
						MkArray(&VarArray{
							MkInteger(50),
							MkInteger(2),
						}),
						MkArray(&VarArray{
							MkInteger(100),
							MkInteger(5),
						}),
						MkArray(&VarArray{
							MkInteger(500),
							MkInteger(10),
						}),
						MkArray(&VarArray{
							MkInteger(1000),
							MkInteger(20),
						}),
					}),
				}),
				"trades":           MkInteger(1),
				"historicalTrades": MkInteger(20),
				"aggTrades":        MkInteger(20),
				"premiumIndex":     MkInteger(10),
				"fundingRate":      MkInteger(1),
				"klines": MkMap(&VarMap{
					"cost": MkInteger(1),
					"byLimit": MkArray(&VarArray{
						MkArray(&VarArray{
							MkInteger(99),
							MkInteger(1),
						}),
						MkArray(&VarArray{
							MkInteger(499),
							MkInteger(2),
						}),
						MkArray(&VarArray{
							MkInteger(1000),
							MkInteger(5),
						}),
						MkArray(&VarArray{
							MkInteger(10000),
							MkInteger(10),
						}),
					}),
				}),
				"continuousKlines": MkMap(&VarMap{
					"cost": MkInteger(1),
					"byLimit": MkArray(&VarArray{
						MkArray(&VarArray{
							MkInteger(99),
							MkInteger(1),
						}),
						MkArray(&VarArray{
							MkInteger(499),
							MkInteger(2),
						}),
						MkArray(&VarArray{
							MkInteger(1000),
							MkInteger(5),
						}),
						MkArray(&VarArray{
							MkInteger(10000),
							MkInteger(10),
						}),
					}),
				}),
				"indexPriceKlines": MkMap(&VarMap{
					"cost": MkInteger(1),
					"byLimit": MkArray(&VarArray{
						MkArray(&VarArray{
							MkInteger(99),
							MkInteger(1),
						}),
						MkArray(&VarArray{
							MkInteger(499),
							MkInteger(2),
						}),
						MkArray(&VarArray{
							MkInteger(1000),
							MkInteger(5),
						}),
						MkArray(&VarArray{
							MkInteger(10000),
							MkInteger(10),
						}),
					}),
				}),
				"markPriceKlines": MkMap(&VarMap{
					"cost": MkInteger(1),
					"byLimit": MkArray(&VarArray{
						MkArray(&VarArray{
							MkInteger(99),
							MkInteger(1),
						}),
						MkArray(&VarArray{
							MkInteger(499),
							MkInteger(2),
						}),
						MkArray(&VarArray{
							MkInteger(1000),
							MkInteger(5),
						}),
						MkArray(&VarArray{
							MkInteger(10000),
							MkInteger(10),
						}),
					}),
				}),
				"ticker/24hr": MkMap(&VarMap{
					"cost":     MkInteger(1),
					"noSymbol": MkInteger(40),
				}),
				"ticker/price": MkMap(&VarMap{
					"cost":     MkInteger(1),
					"noSymbol": MkInteger(2),
				}),
				"ticker/bookTicker": MkMap(&VarMap{
					"cost":     MkInteger(1),
					"noSymbol": MkInteger(2),
				}),
				"openInterest": MkInteger(1),
			})}),
			"dapiData": MkMap(&VarMap{"get": MkMap(&VarMap{
				"openInterestHist":            MkInteger(1),
				"topLongShortAccountRatio":    MkInteger(1),
				"topLongShortPositionRatio":   MkInteger(1),
				"globalLongShortAccountRatio": MkInteger(1),
				"takerBuySellVol":             MkInteger(1),
				"basis":                       MkInteger(1),
			})}),
			"dapiPrivate": MkMap(&VarMap{
				"get": MkMap(&VarMap{
					"positionSide/dual": MkInteger(30),
					"order":             MkInteger(1),
					"openOrder":         MkInteger(1),
					"openOrders": MkMap(&VarMap{
						"cost":     MkInteger(1),
						"noSymbol": MkInteger(5),
					}),
					"allOrders": MkMap(&VarMap{
						"cost":     MkInteger(20),
						"noSymbol": MkInteger(40),
					}),
					"balance":                MkInteger(1),
					"account":                MkInteger(5),
					"positionMargin/history": MkInteger(1),
					"positionRisk":           MkInteger(1),
					"userTrades": MkMap(&VarMap{
						"cost":     MkInteger(20),
						"noSymbol": MkInteger(40),
					}),
					"income":          MkInteger(20),
					"leverageBracket": MkInteger(1),
					"forceOrders": MkMap(&VarMap{
						"cost":     MkInteger(20),
						"noSymbol": MkInteger(50),
					}),
					"adlQuantile": MkInteger(5),
				}),
				"post": MkMap(&VarMap{
					"positionSide/dual":  MkInteger(1),
					"order":              MkInteger(4),
					"batchOrders":        MkInteger(5),
					"countdownCancelAll": MkInteger(10),
					"leverage":           MkInteger(1),
					"marginType":         MkInteger(1),
					"positionMargin":     MkInteger(1),
					"listenKey":          MkInteger(1),
				}),
				"put": MkMap(&VarMap{"listenKey": MkInteger(1)}),
				"delete": MkMap(&VarMap{
					"order":         MkInteger(1),
					"allOpenOrders": MkInteger(1),
					"batchOrders":   MkInteger(5),
					"listenKey":     MkInteger(1),
				}),
			}),
			"dapiPrivateV2": MkMap(&VarMap{"get": MkMap(&VarMap{"leverageBracket": MkInteger(1)})}),
			"fapiPublic": MkMap(&VarMap{"get": MkMap(&VarMap{
				"ping":         MkInteger(1),
				"time":         MkInteger(1),
				"exchangeInfo": MkInteger(1),
				"depth": MkMap(&VarMap{
					"cost": MkInteger(2),
					"byLimit": MkArray(&VarArray{
						MkArray(&VarArray{
							MkInteger(50),
							MkInteger(2),
						}),
						MkArray(&VarArray{
							MkInteger(100),
							MkInteger(5),
						}),
						MkArray(&VarArray{
							MkInteger(500),
							MkInteger(10),
						}),
						MkArray(&VarArray{
							MkInteger(1000),
							MkInteger(20),
						}),
					}),
				}),
				"trades":           MkInteger(1),
				"historicalTrades": MkInteger(20),
				"aggTrades":        MkInteger(20),
				"klines": MkMap(&VarMap{
					"cost": MkInteger(1),
					"byLimit": MkArray(&VarArray{
						MkArray(&VarArray{
							MkInteger(99),
							MkInteger(1),
						}),
						MkArray(&VarArray{
							MkInteger(499),
							MkInteger(2),
						}),
						MkArray(&VarArray{
							MkInteger(1000),
							MkInteger(5),
						}),
						MkArray(&VarArray{
							MkInteger(10000),
							MkInteger(10),
						}),
					}),
				}),
				"continuousKlines": MkMap(&VarMap{
					"cost": MkInteger(1),
					"byLimit": MkArray(&VarArray{
						MkArray(&VarArray{
							MkInteger(99),
							MkInteger(1),
						}),
						MkArray(&VarArray{
							MkInteger(499),
							MkInteger(2),
						}),
						MkArray(&VarArray{
							MkInteger(1000),
							MkInteger(5),
						}),
						MkArray(&VarArray{
							MkInteger(10000),
							MkInteger(10),
						}),
					}),
				}),
				"fundingRate":  MkInteger(1),
				"premiumIndex": MkInteger(1),
				"ticker/24hr": MkMap(&VarMap{
					"cost":     MkInteger(1),
					"noSymbol": MkInteger(40),
				}),
				"ticker/price": MkMap(&VarMap{
					"cost":     MkInteger(1),
					"noSymbol": MkInteger(2),
				}),
				"ticker/bookTicker": MkMap(&VarMap{
					"cost":     MkInteger(1),
					"noSymbol": MkInteger(2),
				}),
				"openInterest": MkInteger(1),
				"indexInfo":    MkInteger(1),
				"apiTradingStatus": MkMap(&VarMap{
					"cost":     MkInteger(1),
					"noSymbol": MkInteger(10),
				}),
				"lvtKlines": MkInteger(1),
			})}),
			"fapiData": MkMap(&VarMap{"get": MkMap(&VarMap{
				"openInterestHist":            MkInteger(1),
				"topLongShortAccountRatio":    MkInteger(1),
				"topLongShortPositionRatio":   MkInteger(1),
				"globalLongShortAccountRatio": MkInteger(1),
				"takerlongshortRatio":         MkInteger(1),
			})}),
			"fapiPrivate": MkMap(&VarMap{
				"get": MkMap(&VarMap{
					"forceOrders": MkMap(&VarMap{
						"cost":     MkInteger(20),
						"noSymbol": MkInteger(50),
					}),
					"allOrders":                     MkInteger(5),
					"openOrder":                     MkInteger(1),
					"openOrders":                    MkInteger(1),
					"order":                         MkInteger(1),
					"account":                       MkInteger(5),
					"balance":                       MkInteger(5),
					"leverageBracket":               MkInteger(1),
					"positionMargin/history":        MkInteger(1),
					"positionRisk":                  MkInteger(5),
					"positionSide/dual":             MkInteger(30),
					"userTrades":                    MkInteger(5),
					"income":                        MkInteger(30),
					"commissionRate":                MkInteger(20),
					"apiTradingStatus":              MkInteger(1),
					"multiAssetsMargin":             MkInteger(30),
					"apiReferral/ifNewUser":         MkInteger(1),
					"apiReferral/customization":     MkInteger(1),
					"apiReferral/userCustomization": MkInteger(1),
					"apiReferral/traderNum":         MkInteger(1),
					"apiReferral/overview":          MkInteger(1),
					"apiReferral/tradeVol":          MkInteger(1),
					"apiReferral/rebateVol":         MkInteger(1),
					"apiReferral/traderSummary":     MkInteger(1),
					"adlQuantile":                   MkInteger(5),
				}),
				"post": MkMap(&VarMap{
					"batchOrders":                   MkInteger(5),
					"positionSide/dual":             MkInteger(1),
					"positionMargin":                MkInteger(1),
					"marginType":                    MkInteger(1),
					"order":                         MkInteger(4),
					"leverage":                      MkInteger(1),
					"listenKey":                     MkInteger(1),
					"countdownCancelAll":            MkInteger(10),
					"multiAssetsMargin":             MkInteger(1),
					"apiReferral/customization":     MkInteger(1),
					"apiReferral/userCustomization": MkInteger(1),
				}),
				"put": MkMap(&VarMap{"listenKey": MkInteger(1)}),
				"delete": MkMap(&VarMap{
					"batchOrders":   MkInteger(1),
					"order":         MkInteger(1),
					"allOpenOrders": MkInteger(1),
					"listenKey":     MkInteger(1),
				}),
			}),
			"fapiPrivateV2": MkMap(&VarMap{"get": MkMap(&VarMap{
				"account":      MkInteger(1),
				"balance":      MkInteger(1),
				"positionRisk": MkInteger(1),
			})}),
			"v3": MkMap(&VarMap{"get": MkMap(&VarMap{
				"ticker/price": MkMap(&VarMap{
					"cost":     MkInteger(1),
					"noSymbol": MkInteger(2),
				}),
				"ticker/bookTicker": MkMap(&VarMap{
					"cost":     MkInteger(1),
					"noSymbol": MkInteger(2),
				}),
			})}),
			"public": MkMap(&VarMap{
				"get": MkMap(&VarMap{
					"ping": MkInteger(1),
					"time": MkInteger(1),
					"depth": MkMap(&VarMap{
						"cost": MkInteger(1),
						"byLimit": MkArray(&VarArray{
							MkArray(&VarArray{
								MkInteger(100),
								MkInteger(1),
							}),
							MkArray(&VarArray{
								MkInteger(500),
								MkInteger(5),
							}),
							MkArray(&VarArray{
								MkInteger(1000),
								MkInteger(10),
							}),
							MkArray(&VarArray{
								MkInteger(5000),
								MkInteger(50),
							}),
						}),
					}),
					"trades":           MkInteger(1),
					"aggTrades":        MkInteger(1),
					"historicalTrades": MkInteger(5),
					"klines":           MkInteger(1),
					"ticker/24hr": MkMap(&VarMap{
						"cost":     MkInteger(1),
						"noSymbol": MkInteger(40),
					}),
					"ticker/price": MkMap(&VarMap{
						"cost":     MkInteger(1),
						"noSymbol": MkInteger(2),
					}),
					"ticker/bookTicker": MkMap(&VarMap{
						"cost":     MkInteger(1),
						"noSymbol": MkInteger(2),
					}),
					"exchangeInfo": MkInteger(10),
				}),
				"put":    MkMap(&VarMap{"userDataStream": MkInteger(1)}),
				"post":   MkMap(&VarMap{"userDataStream": MkInteger(1)}),
				"delete": MkMap(&VarMap{"userDataStream": MkInteger(1)}),
			}),
			"private": MkMap(&VarMap{
				"get": MkMap(&VarMap{
					"allOrderList":  MkInteger(10),
					"openOrderList": MkInteger(3),
					"orderList":     MkInteger(2),
					"order":         MkInteger(2),
					"openOrders": MkMap(&VarMap{
						"cost":     MkInteger(3),
						"noSymbol": MkInteger(40),
					}),
					"allOrders": MkInteger(10),
					"account":   MkInteger(10),
					"myTrades":  MkInteger(10),
				}),
				"post": MkMap(&VarMap{
					"order/oco":  MkInteger(1),
					"order":      MkInteger(4),
					"order/test": MkInteger(1),
				}),
				"delete": MkMap(&VarMap{
					"openOrders": MkInteger(1),
					"orderList":  MkInteger(1),
					"order":      MkInteger(1),
				}),
			}),
		}),
		"fees": MkMap(&VarMap{
			"trading": MkMap(&VarMap{
				"feeSide":    MkString("get"),
				"tierBased":  MkBool(false),
				"percentage": MkBool(true),
				"taker":      this.ParseNumber(MkString("0.001")),
				"maker":      this.ParseNumber(MkString("0.001")),
			}),
			"future": MkMap(&VarMap{"trading": MkMap(&VarMap{
				"feeSide":    MkString("quote"),
				"tierBased":  MkBool(true),
				"percentage": MkBool(true),
				"taker":      this.ParseNumber(MkString("0.000400")),
				"maker":      this.ParseNumber(MkString("0.000200")),
				"tiers": MkMap(&VarMap{
					"taker": MkArray(&VarArray{
						MkArray(&VarArray{
							this.ParseNumber(MkString("0")),
							this.ParseNumber(MkString("0.000400")),
						}),
						MkArray(&VarArray{
							this.ParseNumber(MkString("250")),
							this.ParseNumber(MkString("0.000400")),
						}),
						MkArray(&VarArray{
							this.ParseNumber(MkString("2500")),
							this.ParseNumber(MkString("0.000350")),
						}),
						MkArray(&VarArray{
							this.ParseNumber(MkString("7500")),
							this.ParseNumber(MkString("0.000320")),
						}),
						MkArray(&VarArray{
							this.ParseNumber(MkString("22500")),
							this.ParseNumber(MkString("0.000300")),
						}),
						MkArray(&VarArray{
							this.ParseNumber(MkString("50000")),
							this.ParseNumber(MkString("0.000270")),
						}),
						MkArray(&VarArray{
							this.ParseNumber(MkString("100000")),
							this.ParseNumber(MkString("0.000250")),
						}),
						MkArray(&VarArray{
							this.ParseNumber(MkString("200000")),
							this.ParseNumber(MkString("0.000220")),
						}),
						MkArray(&VarArray{
							this.ParseNumber(MkString("400000")),
							this.ParseNumber(MkString("0.000200")),
						}),
						MkArray(&VarArray{
							this.ParseNumber(MkString("750000")),
							this.ParseNumber(MkString("0.000170")),
						}),
					}),
					"maker": MkArray(&VarArray{
						MkArray(&VarArray{
							this.ParseNumber(MkString("0")),
							this.ParseNumber(MkString("0.000200")),
						}),
						MkArray(&VarArray{
							this.ParseNumber(MkString("250")),
							this.ParseNumber(MkString("0.000160")),
						}),
						MkArray(&VarArray{
							this.ParseNumber(MkString("2500")),
							this.ParseNumber(MkString("0.000140")),
						}),
						MkArray(&VarArray{
							this.ParseNumber(MkString("7500")),
							this.ParseNumber(MkString("0.000120")),
						}),
						MkArray(&VarArray{
							this.ParseNumber(MkString("22500")),
							this.ParseNumber(MkString("0.000100")),
						}),
						MkArray(&VarArray{
							this.ParseNumber(MkString("50000")),
							this.ParseNumber(MkString("0.000080")),
						}),
						MkArray(&VarArray{
							this.ParseNumber(MkString("100000")),
							this.ParseNumber(MkString("0.000060")),
						}),
						MkArray(&VarArray{
							this.ParseNumber(MkString("200000")),
							this.ParseNumber(MkString("0.000040")),
						}),
						MkArray(&VarArray{
							this.ParseNumber(MkString("400000")),
							this.ParseNumber(MkString("0.000020")),
						}),
						MkArray(&VarArray{
							this.ParseNumber(MkString("750000")),
							this.ParseNumber(MkString("0")),
						}),
					}),
				}),
			})}),
			"delivery": MkMap(&VarMap{"trading": MkMap(&VarMap{
				"feeSide":    MkString("base"),
				"tierBased":  MkBool(true),
				"percentage": MkBool(true),
				"taker":      this.ParseNumber(MkString("0.000500")),
				"maker":      this.ParseNumber(MkString("0.000100")),
				"tiers": MkMap(&VarMap{
					"taker": MkArray(&VarArray{
						MkArray(&VarArray{
							this.ParseNumber(MkString("0")),
							this.ParseNumber(MkString("0.000500")),
						}),
						MkArray(&VarArray{
							this.ParseNumber(MkString("250")),
							this.ParseNumber(MkString("0.000450")),
						}),
						MkArray(&VarArray{
							this.ParseNumber(MkString("2500")),
							this.ParseNumber(MkString("0.000400")),
						}),
						MkArray(&VarArray{
							this.ParseNumber(MkString("7500")),
							this.ParseNumber(MkString("0.000300")),
						}),
						MkArray(&VarArray{
							this.ParseNumber(MkString("22500")),
							this.ParseNumber(MkString("0.000250")),
						}),
						MkArray(&VarArray{
							this.ParseNumber(MkString("50000")),
							this.ParseNumber(MkString("0.000240")),
						}),
						MkArray(&VarArray{
							this.ParseNumber(MkString("100000")),
							this.ParseNumber(MkString("0.000240")),
						}),
						MkArray(&VarArray{
							this.ParseNumber(MkString("200000")),
							this.ParseNumber(MkString("0.000240")),
						}),
						MkArray(&VarArray{
							this.ParseNumber(MkString("400000")),
							this.ParseNumber(MkString("0.000240")),
						}),
						MkArray(&VarArray{
							this.ParseNumber(MkString("750000")),
							this.ParseNumber(MkString("0.000240")),
						}),
					}),
					"maker": MkArray(&VarArray{
						MkArray(&VarArray{
							this.ParseNumber(MkString("0")),
							this.ParseNumber(MkString("0.000100")),
						}),
						MkArray(&VarArray{
							this.ParseNumber(MkString("250")),
							this.ParseNumber(MkString("0.000080")),
						}),
						MkArray(&VarArray{
							this.ParseNumber(MkString("2500")),
							this.ParseNumber(MkString("0.000050")),
						}),
						MkArray(&VarArray{
							this.ParseNumber(MkString("7500")),
							this.ParseNumber(MkString("0.0000030")),
						}),
						MkArray(&VarArray{
							this.ParseNumber(MkString("22500")),
							this.ParseNumber(MkString("0")),
						}),
						MkArray(&VarArray{
							this.ParseNumber(MkString("50000")),
							this.ParseNumber(MkString("-0.000050")),
						}),
						MkArray(&VarArray{
							this.ParseNumber(MkString("100000")),
							this.ParseNumber(MkString("-0.000060")),
						}),
						MkArray(&VarArray{
							this.ParseNumber(MkString("200000")),
							this.ParseNumber(MkString("-0.000070")),
						}),
						MkArray(&VarArray{
							this.ParseNumber(MkString("400000")),
							this.ParseNumber(MkString("-0.000080")),
						}),
						MkArray(&VarArray{
							this.ParseNumber(MkString("750000")),
							this.ParseNumber(MkString("-0.000090")),
						}),
					}),
				}),
			})}),
		}),
		"commonCurrencies": MkMap(&VarMap{
			"BCC":  MkString("BCC"),
			"YOYO": MkString("YOYOW"),
		}),
		"options": MkMap(&VarMap{
			"fetchCurrencies":                     MkBool(false),
			"defaultTimeInForce":                  MkString("GTC"),
			"defaultType":                         MkString("spot"),
			"hasAlreadyAuthenticatedSuccessfully": MkBool(false),
			"warnOnFetchOpenOrdersWithoutSymbol":  MkBool(true),
			"recvWindow":                          OpMulti(MkInteger(5), MkInteger(1000)),
			"timeDifference":                      MkInteger(0),
			"adjustForTimeDifference":             MkBool(false),
			"parseOrderToPrecision":               MkBool(false),
			"newOrderRespType": MkMap(&VarMap{
				"market": MkString("FULL"),
				"limit":  MkString("FULL"),
			}),
			"quoteOrderQty": MkBool(true),
			"broker": MkMap(&VarMap{
				"spot":     MkString("x-R4BD3S82"),
				"margin":   MkString("x-R4BD3S82"),
				"future":   MkString("x-xcKtGhcu"),
				"delivery": MkString("x-xcKtGhcu"),
			}),
			"accountsByType": MkMap(&VarMap{
				"main":     MkString("MAIN"),
				"spot":     MkString("MAIN"),
				"margin":   MkString("MARGIN"),
				"future":   MkString("UMFUTURE"),
				"delivery": MkString("CMFUTURE"),
				"mining":   MkString("MINING"),
			}),
			"typesByAccount": MkMap(&VarMap{
				"MAIN":     MkString("spot"),
				"MARGIN":   MkString("margin"),
				"UMFUTURE": MkString("future"),
				"CMFUTURE": MkString("delivery"),
				"MINING":   MkString("mining"),
			}),
			"legalMoney": MkMap(&VarMap{
				"MXN": MkBool(true),
				"UGX": MkBool(true),
				"SEK": MkBool(true),
				"CHF": MkBool(true),
				"VND": MkBool(true),
				"AED": MkBool(true),
				"DKK": MkBool(true),
				"KZT": MkBool(true),
				"HUF": MkBool(true),
				"PEN": MkBool(true),
				"PHP": MkBool(true),
				"USD": MkBool(true),
				"TRY": MkBool(true),
				"EUR": MkBool(true),
				"NGN": MkBool(true),
				"PLN": MkBool(true),
				"BRL": MkBool(true),
				"ZAR": MkBool(true),
				"KES": MkBool(true),
				"ARS": MkBool(true),
				"RUB": MkBool(true),
				"AUD": MkBool(true),
				"NOK": MkBool(true),
				"CZK": MkBool(true),
				"GBP": MkBool(true),
				"UAH": MkBool(true),
				"GHS": MkBool(true),
				"HKD": MkBool(true),
				"CAD": MkBool(true),
				"INR": MkBool(true),
				"JPY": MkBool(true),
				"NZD": MkBool(true),
			}),
		}),
		"exceptions": MkMap(&VarMap{
			"exact": MkMap(&VarMap{
				"System abnormality":                                     ExchangeError,
				"You are not authorized to execute this request.":        PermissionDenied,
				"API key does not exist":                                 AuthenticationError,
				"Order would trigger immediately.":                       OrderImmediatelyFillable,
				"Stop price would trigger immediately.":                  OrderImmediatelyFillable,
				"Order would immediately match and take.":                OrderImmediatelyFillable,
				"Account has insufficient balance for requested action.": InsufficientFunds,
				"Rest API trading is not enabled.":                       ExchangeNotAvailable,
				"You don't have permission.":                             PermissionDenied,
				"Market is closed.":                                      ExchangeNotAvailable,
				"Too many requests.":                                     DDoSProtection,
				"-1000":                                                  ExchangeNotAvailable,
				"-1001":                                                  ExchangeNotAvailable,
				"-1002":                                                  AuthenticationError,
				"-1003":                                                  RateLimitExceeded,
				"-1013":                                                  InvalidOrder,
				"-1015":                                                  RateLimitExceeded,
				"-1016":                                                  ExchangeNotAvailable,
				"-1020":                                                  BadRequest,
				"-1021":                                                  InvalidNonce,
				"-1022":                                                  AuthenticationError,
				"-1100":                                                  BadRequest,
				"-1101":                                                  BadRequest,
				"-1102":                                                  BadRequest,
				"-1103":                                                  BadRequest,
				"-1104":                                                  BadRequest,
				"-1105":                                                  BadRequest,
				"-1106":                                                  BadRequest,
				"-1111":                                                  BadRequest,
				"-1112":                                                  InvalidOrder,
				"-1114":                                                  BadRequest,
				"-1115":                                                  BadRequest,
				"-1116":                                                  BadRequest,
				"-1117":                                                  BadRequest,
				"-1118":                                                  BadRequest,
				"-1119":                                                  BadRequest,
				"-1120":                                                  BadRequest,
				"-1121":                                                  BadSymbol,
				"-1125":                                                  AuthenticationError,
				"-1127":                                                  BadRequest,
				"-1128":                                                  BadRequest,
				"-1130":                                                  BadRequest,
				"-1131":                                                  BadRequest,
				"-2008":                                                  AuthenticationError,
				"-2010":                                                  ExchangeError,
				"-2011":                                                  OrderNotFound,
				"-2013":                                                  OrderNotFound,
				"-2014":                                                  AuthenticationError,
				"-2015":                                                  AuthenticationError,
				"-2019":                                                  InsufficientFunds,
				"-3005":                                                  InsufficientFunds,
				"-3006":                                                  InsufficientFunds,
				"-3008":                                                  InsufficientFunds,
				"-3010":                                                  ExchangeError,
				"-3015":                                                  ExchangeError,
				"-3022":                                                  AccountSuspended,
				"-4028":                                                  BadRequest,
				"-3020":                                                  InsufficientFunds,
				"-3041":                                                  InsufficientFunds,
				"-5013":                                                  InsufficientFunds,
				"-11008":                                                 InsufficientFunds,
				"-4051":                                                  InsufficientFunds,
			}),
			"broad": MkMap(&VarMap{
				"has no operation privilege": PermissionDenied,
				"MAX_POSITION":               InvalidOrder,
			}),
		}),
	}))
}

func (this *Binance) CostToPrecision(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	cost := GoGetArg(goArgs, 1, MkUndefined())
	_ = cost
	return this.DecimalToPrecision(cost, TRUNCATE, *(*(*(*this.At(MkString("markets"))).At(symbol)).At(MkString("precision"))).At(MkString("quote")), *this.At(MkString("precisionMode")), *this.At(MkString("paddingMode")))
}

func (this *Binance) CurrencyToPrecision(goArgs ...*Variant) *Variant {
	currency := GoGetArg(goArgs, 0, MkUndefined())
	_ = currency
	fee := GoGetArg(goArgs, 1, MkUndefined())
	_ = fee
	if IsTrue(OpHasMember(MkString("info"), *(*this.At(MkString("currencies"))).At(currency))) {
		return this.DecimalToPrecision(fee, TRUNCATE, *(*(*this.At(MkString("currencies"))).At(currency)).At(MkString("precision")), *this.At(MkString("precisionMode")), *this.At(MkString("paddingMode")))
	} else {
		return this.NumberToString(fee)
	}
	return MkUndefined()
}

func (this *Binance) Nonce(goArgs ...*Variant) *Variant {
	return OpSub(this.Milliseconds(), *(*this.At(MkString("options"))).At(MkString("timeDifference")))
}

func (this *Binance) FetchTime(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	type_ := this.SafeString2(*this.At(MkString("options")), MkString("fetchTime"), MkString("defaultType"), MkString("spot"))
	_ = type_
	method := MkString("publicGetTime")
	_ = method
	if IsTrue(OpEq2(type_, MkString("future"))) {
		method = MkString("fapiPublicGetTime")
	} else {
		if IsTrue(OpEq2(type_, MkString("delivery"))) {
			method = MkString("dapiPublicGetTime")
		}
	}
	response := this.Call(method, params)
	_ = response
	return this.SafeInteger(response, MkString("serverTime"))
}

func (this *Binance) LoadTimeDifference(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	serverTime := this.FetchTime(params)
	_ = serverTime
	after := this.Milliseconds()
	_ = after
	*(*this.At(MkString("options"))).At(MkString("timeDifference")) = OpSub(after, serverTime)
	return *(*this.At(MkString("options"))).At(MkString("timeDifference"))
}

func (this *Binance) FetchCurrencies(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	fetchCurrenciesEnabled := this.SafeValue(*this.At(MkString("options")), MkString("fetchCurrencies"))
	_ = fetchCurrenciesEnabled
	if IsTrue(OpNot(fetchCurrenciesEnabled)) {
		return MkUndefined()
	}
	if IsTrue(OpNot(this.CheckRequiredCredentials(MkBool(false)))) {
		return MkUndefined()
	}
	apiBackup := this.SafeString(*this.At(MkString("urls")), MkString("apiBackup"))
	_ = apiBackup
	if IsTrue(OpNotEq2(apiBackup, MkUndefined())) {
		return MkUndefined()
	}
	response := this.Call(MkString("sapiGetCapitalConfigGetall"), params)
	_ = response
	result := MkMap(&VarMap{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, response.Length)); OpIncr(&i) {
		entry := *(response).At(i)
		_ = entry
		id := this.SafeString(entry, MkString("coin"))
		_ = id
		name := this.SafeString(entry, MkString("name"))
		_ = name
		code := this.SafeCurrencyCode(id)
		_ = code
		precision := OpCopy(MkUndefined())
		_ = precision
		isWithdrawEnabled := OpCopy(MkBool(true))
		_ = isWithdrawEnabled
		isDepositEnabled := OpCopy(MkBool(true))
		_ = isDepositEnabled
		networkList := this.SafeValue(entry, MkString("networkList"), MkArray(&VarArray{}))
		_ = networkList
		fees := MkMap(&VarMap{})
		_ = fees
		fee := OpCopy(MkUndefined())
		_ = fee
		for j := MkInteger(0); IsTrue(OpLw(j, networkList.Length)); OpIncr(&j) {
			networkItem := *(networkList).At(j)
			_ = networkItem
			network := this.SafeString(networkItem, MkString("network"))
			_ = network
			withdrawFee := this.SafeNumber(networkItem, MkString("withdrawFee"))
			_ = withdrawFee
			depositEnable := this.SafeValue(networkItem, MkString("depositEnable"))
			_ = depositEnable
			withdrawEnable := this.SafeValue(networkItem, MkString("withdrawEnable"))
			_ = withdrawEnable
			isDepositEnabled = OpOr(isDepositEnabled, depositEnable)
			isWithdrawEnabled = OpOr(isWithdrawEnabled, withdrawEnable)
			*(fees).At(network) = OpCopy(withdrawFee)
			isDefault := this.SafeValue(networkItem, MkString("isDefault"))
			_ = isDefault
			if IsTrue(OpOr(isDefault, OpEq2(fee, MkUndefined()))) {
				fee = OpCopy(withdrawFee)
			}
		}
		trading := this.SafeValue(entry, MkString("trading"))
		_ = trading
		active := OpAnd(isWithdrawEnabled, OpAnd(isDepositEnabled, trading))
		_ = active
		*(result).At(code) = MkMap(&VarMap{
			"id":        id,
			"name":      name,
			"code":      code,
			"precision": precision,
			"info":      entry,
			"active":    active,
			"fee":       fee,
			"fees":      fees,
			"limits":    *this.At(MkString("limits")),
		})
	}
	return result
}

func (this *Binance) FetchMarkets(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	defaultType := this.SafeString2(*this.At(MkString("options")), MkString("fetchMarkets"), MkString("defaultType"), MkString("spot"))
	_ = defaultType
	type_ := this.SafeString(params, MkString("type"), defaultType)
	_ = type_
	query := this.Omit(params, MkString("type"))
	_ = query
	if IsTrue(OpAnd(OpNotEq2(type_, MkString("spot")), OpAnd(OpNotEq2(type_, MkString("future")), OpAnd(OpNotEq2(type_, MkString("margin")), OpNotEq2(type_, MkString("delivery")))))) {
		panic(NewExchangeError(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" does not support '"), OpAdd(type_, MkString("' type, set exchange.options['defaultType'] to 'spot', 'margin', 'delivery' or 'future'"))))))
	}
	method := MkString("publicGetExchangeInfo")
	_ = method
	if IsTrue(OpEq2(type_, MkString("future"))) {
		method = MkString("fapiPublicGetExchangeInfo")
	} else {
		if IsTrue(OpEq2(type_, MkString("delivery"))) {
			method = MkString("dapiPublicGetExchangeInfo")
		}
	}
	response := this.Call(method, query)
	_ = response
	if IsTrue(*(*this.At(MkString("options"))).At(MkString("adjustForTimeDifference"))) {
		this.LoadTimeDifference()
	}
	markets := this.SafeValue(response, MkString("symbols"), MkArray(&VarArray{}))
	_ = markets
	result := MkArray(&VarArray{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, markets.Length)); OpIncr(&i) {
		market := *(markets).At(i)
		_ = market
		spot := OpEq2(type_, MkString("spot"))
		_ = spot
		future := OpEq2(type_, MkString("future"))
		_ = future
		delivery := OpEq2(type_, MkString("delivery"))
		_ = delivery
		id := this.SafeString(market, MkString("symbol"))
		_ = id
		lowercaseId := this.SafeStringLower(market, MkString("symbol"))
		_ = lowercaseId
		baseId := this.SafeString(market, MkString("baseAsset"))
		_ = baseId
		quoteId := this.SafeString(market, MkString("quoteAsset"))
		_ = quoteId
		base := this.SafeCurrencyCode(baseId)
		_ = base
		quote := this.SafeCurrencyCode(quoteId)
		_ = quote
		contractType := this.SafeString(market, MkString("contractType"))
		_ = contractType
		idSymbol := OpAnd(OpOr(future, delivery), OpNotEq2(contractType, MkString("PERPETUAL")))
		_ = idSymbol
		symbol := OpCopy(MkUndefined())
		_ = symbol
		expiry := OpCopy(MkUndefined())
		_ = expiry
		if IsTrue(idSymbol) {
			symbol = OpCopy(id)
			expiry = this.SafeInteger(market, MkString("deliveryDate"))
		} else {
			symbol = OpAdd(base, OpAdd(MkString("/"), quote))
		}
		filters := this.SafeValue(market, MkString("filters"), MkArray(&VarArray{}))
		_ = filters
		filtersByType := this.IndexBy(filters, MkString("filterType"))
		_ = filtersByType
		precision := MkMap(&VarMap{
			"base":   this.SafeInteger(market, MkString("baseAssetPrecision")),
			"quote":  this.SafeInteger(market, MkString("quotePrecision")),
			"amount": this.SafeInteger(market, MkString("quantityPrecision")),
			"price":  this.SafeInteger(market, MkString("pricePrecision")),
		})
		_ = precision
		status := this.SafeString2(market, MkString("status"), MkString("contractStatus"))
		_ = status
		active := OpEq2(status, MkString("TRADING"))
		_ = active
		margin := this.SafeValue(market, MkString("isMarginTradingAllowed"), MkBool(false))
		_ = margin
		contractSize := OpCopy(MkUndefined())
		_ = contractSize
		fees := OpCopy(*this.At(MkString("fees")))
		_ = fees
		if IsTrue(OpOr(future, delivery)) {
			contractSize = this.SafeString(market, MkString("contractSize"), MkString("1"))
			fees = *(*this.At(MkString("fees"))).At(type_)
		}
		maker := *(*(fees).At(MkString("trading"))).At(MkString("maker"))
		_ = maker
		taker := *(*(fees).At(MkString("trading"))).At(MkString("taker"))
		_ = taker
		entry := MkMap(&VarMap{
			"id":             id,
			"lowercaseId":    lowercaseId,
			"symbol":         symbol,
			"base":           base,
			"quote":          quote,
			"baseId":         baseId,
			"quoteId":        quoteId,
			"info":           market,
			"spot":           spot,
			"type":           type_,
			"margin":         margin,
			"future":         future,
			"delivery":       delivery,
			"linear":         future,
			"inverse":        delivery,
			"expiry":         expiry,
			"expiryDatetime": this.Iso8601(expiry),
			"active":         active,
			"precision":      precision,
			"contractSize":   contractSize,
			"maker":          maker,
			"taker":          taker,
			"limits": MkMap(&VarMap{
				"amount": MkMap(&VarMap{
					"min": MkUndefined(),
					"max": MkUndefined(),
				}),
				"price": MkMap(&VarMap{
					"min": MkUndefined(),
					"max": MkUndefined(),
				}),
				"cost": MkMap(&VarMap{
					"min": MkUndefined(),
					"max": MkUndefined(),
				}),
			}),
		})
		_ = entry
		if IsTrue(OpHasMember(MkString("PRICE_FILTER"), filtersByType)) {
			filter := this.SafeValue(filtersByType, MkString("PRICE_FILTER"), MkMap(&VarMap{}))
			_ = filter
			tickSize := this.SafeString(filter, MkString("tickSize"))
			_ = tickSize
			*(*(entry).At(MkString("precision"))).At(MkString("price")) = this.PrecisionFromString(tickSize)
			*(*(entry).At(MkString("limits"))).At(MkString("price")) = MkMap(&VarMap{
				"min": this.SafeNumber(filter, MkString("minPrice")),
				"max": this.SafeNumber(filter, MkString("maxPrice")),
			})
			*(*(entry).At(MkString("precision"))).At(MkString("price")) = this.PrecisionFromString(*(filter).At(MkString("tickSize")))
		}
		if IsTrue(OpHasMember(MkString("LOT_SIZE"), filtersByType)) {
			filter := this.SafeValue(filtersByType, MkString("LOT_SIZE"), MkMap(&VarMap{}))
			_ = filter
			stepSize := this.SafeString(filter, MkString("stepSize"))
			_ = stepSize
			*(*(entry).At(MkString("precision"))).At(MkString("amount")) = this.PrecisionFromString(stepSize)
			*(*(entry).At(MkString("limits"))).At(MkString("amount")) = MkMap(&VarMap{
				"min": this.SafeNumber(filter, MkString("minQty")),
				"max": this.SafeNumber(filter, MkString("maxQty")),
			})
		}
		if IsTrue(OpHasMember(MkString("MARKET_LOT_SIZE"), filtersByType)) {
			filter := this.SafeValue(filtersByType, MkString("MARKET_LOT_SIZE"), MkMap(&VarMap{}))
			_ = filter
			*(*(entry).At(MkString("limits"))).At(MkString("market")) = MkMap(&VarMap{
				"min": this.SafeNumber(filter, MkString("minQty")),
				"max": this.SafeNumber(filter, MkString("maxQty")),
			})
		}
		if IsTrue(OpHasMember(MkString("MIN_NOTIONAL"), filtersByType)) {
			filter := this.SafeValue(filtersByType, MkString("MIN_NOTIONAL"), MkMap(&VarMap{}))
			_ = filter
			*(*(*(entry).At(MkString("limits"))).At(MkString("cost"))).At(MkString("min")) = this.SafeNumber2(filter, MkString("minNotional"), MkString("notional"))
		}
		result.Push(entry)
	}
	return result
}

func (this *Binance) FetchBalance(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	defaultType := this.SafeString2(*this.At(MkString("options")), MkString("fetchBalance"), MkString("defaultType"), MkString("spot"))
	_ = defaultType
	type_ := this.SafeString(params, MkString("type"), defaultType)
	_ = type_
	method := MkString("privateGetAccount")
	_ = method
	if IsTrue(OpEq2(type_, MkString("future"))) {
		options := this.SafeValue(*this.At(MkString("options")), MkString("future"), MkMap(&VarMap{}))
		_ = options
		fetchBalanceOptions := this.SafeValue(options, MkString("fetchBalance"), MkMap(&VarMap{}))
		_ = fetchBalanceOptions
		method = this.SafeString(fetchBalanceOptions, MkString("method"), MkString("fapiPrivateV2GetAccount"))
	} else {
		if IsTrue(OpEq2(type_, MkString("delivery"))) {
			options := this.SafeValue(*this.At(MkString("options")), MkString("delivery"), MkMap(&VarMap{}))
			_ = options
			fetchBalanceOptions := this.SafeValue(options, MkString("fetchBalance"), MkMap(&VarMap{}))
			_ = fetchBalanceOptions
			method = this.SafeString(fetchBalanceOptions, MkString("method"), MkString("dapiPrivateGetAccount"))
		} else {
			if IsTrue(OpEq2(type_, MkString("margin"))) {
				method = MkString("sapiGetMarginAccount")
			} else {
				if IsTrue(OpEq2(type_, MkString("savings"))) {
					method = MkString("sapiGetLendingUnionAccount")
				} else {
					if IsTrue(OpEq2(type_, MkString("pay"))) {
						method = MkString("sapiPostAssetGetFundingAsset")
					}
				}
			}
		}
	}
	query := this.Omit(params, MkString("type"))
	_ = query
	response := this.Call(method, query)
	_ = response
	result := MkMap(&VarMap{"info": response})
	_ = result
	timestamp := OpCopy(MkUndefined())
	_ = timestamp
	if IsTrue(OpOr(OpEq2(type_, MkString("spot")), OpEq2(type_, MkString("margin")))) {
		timestamp = this.SafeInteger(response, MkString("updateTime"))
		balances := this.SafeValue2(response, MkString("balances"), MkString("userAssets"), MkArray(&VarArray{}))
		_ = balances
		for i := MkInteger(0); IsTrue(OpLw(i, balances.Length)); OpIncr(&i) {
			balance := *(balances).At(i)
			_ = balance
			currencyId := this.SafeString(balance, MkString("asset"))
			_ = currencyId
			code := this.SafeCurrencyCode(currencyId)
			_ = code
			account := this.Account()
			_ = account
			*(account).At(MkString("free")) = this.SafeString(balance, MkString("free"))
			*(account).At(MkString("used")) = this.SafeString(balance, MkString("locked"))
			*(result).At(code) = OpCopy(account)
		}
	} else {
		if IsTrue(OpEq2(type_, MkString("savings"))) {
			positionAmountVos := this.SafeValue(response, MkString("positionAmountVos"))
			_ = positionAmountVos
			for i := MkInteger(0); IsTrue(OpLw(i, positionAmountVos.Length)); OpIncr(&i) {
				entry := *(positionAmountVos).At(i)
				_ = entry
				currencyId := this.SafeString(entry, MkString("asset"))
				_ = currencyId
				code := this.SafeCurrencyCode(currencyId)
				_ = code
				account := this.Account()
				_ = account
				usedAndTotal := this.SafeString(entry, MkString("amount"))
				_ = usedAndTotal
				*(account).At(MkString("total")) = OpCopy(usedAndTotal)
				*(account).At(MkString("used")) = OpCopy(usedAndTotal)
				*(result).At(code) = OpCopy(account)
			}
		} else {
			if IsTrue(OpEq2(type_, MkString("pay"))) {
				for i := MkInteger(0); IsTrue(OpLw(i, response.Length)); OpIncr(&i) {
					entry := *(response).At(i)
					_ = entry
					account := this.Account()
					_ = account
					currencyId := this.SafeString(entry, MkString("asset"))
					_ = currencyId
					code := this.SafeCurrencyCode(currencyId)
					_ = code
					*(account).At(MkString("free")) = this.SafeString(entry, MkString("free"))
					frozen := this.SafeString(entry, MkString("freeze"))
					_ = frozen
					withdrawing := this.SafeString(entry, MkString("withdrawing"))
					_ = withdrawing
					locked := this.SafeString(entry, MkString("locked"))
					_ = locked
					*(account).At(MkString("used")) = Precise.StringAdd(frozen, Precise.StringAdd(locked, withdrawing))
					*(result).At(code) = OpCopy(account)
				}
			} else {
				balances := OpCopy(response)
				_ = balances
				if IsTrue(OpNot(GoIsArray(response))) {
					balances = this.SafeValue(response, MkString("assets"), MkArray(&VarArray{}))
				}
				for i := MkInteger(0); IsTrue(OpLw(i, balances.Length)); OpIncr(&i) {
					balance := *(balances).At(i)
					_ = balance
					currencyId := this.SafeString(balance, MkString("asset"))
					_ = currencyId
					code := this.SafeCurrencyCode(currencyId)
					_ = code
					account := this.Account()
					_ = account
					*(account).At(MkString("free")) = this.SafeString(balance, MkString("availableBalance"))
					*(account).At(MkString("used")) = this.SafeString(balance, MkString("initialMargin"))
					*(account).At(MkString("total")) = this.SafeString2(balance, MkString("marginBalance"), MkString("balance"))
					*(result).At(code) = OpCopy(account)
				}
			}
		}
	}
	*(result).At(MkString("timestamp")) = OpCopy(timestamp)
	*(result).At(MkString("datetime")) = this.Iso8601(timestamp)
	return this.ParseBalance(result)
}

func (this *Binance) FetchOrderBook(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	limit := GoGetArg(goArgs, 1, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"symbol": *(market).At(MkString("id"))})
	_ = request
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("limit")) = OpCopy(limit)
	}
	method := MkString("publicGetDepth")
	_ = method
	if IsTrue(*(market).At(MkString("linear"))) {
		method = MkString("fapiPublicGetDepth")
	} else {
		if IsTrue(*(market).At(MkString("inverse"))) {
			method = MkString("dapiPublicGetDepth")
		}
	}
	response := this.Call(method, this.Extend(request, params))
	_ = response
	timestamp := this.SafeInteger(response, MkString("T"))
	_ = timestamp
	orderbook := this.ParseOrderBook(response, symbol, timestamp)
	_ = orderbook
	*(orderbook).At(MkString("nonce")) = this.SafeInteger(response, MkString("lastUpdateId"))
	return orderbook
}

func (this *Binance) ParseTicker(goArgs ...*Variant) *Variant {
	ticker := GoGetArg(goArgs, 0, MkUndefined())
	_ = ticker
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	timestamp := this.SafeInteger(ticker, MkString("closeTime"))
	_ = timestamp
	marketId := this.SafeString(ticker, MkString("symbol"))
	_ = marketId
	symbol := this.SafeSymbol(marketId, market)
	_ = symbol
	last := this.SafeNumber(ticker, MkString("lastPrice"))
	_ = last
	isCoinm := OpHasMember(MkString("baseVolume"), ticker)
	_ = isCoinm
	baseVolume := OpCopy(MkUndefined())
	_ = baseVolume
	quoteVolume := OpCopy(MkUndefined())
	_ = quoteVolume
	if IsTrue(isCoinm) {
		baseVolume = this.SafeNumber(ticker, MkString("baseVolume"))
		quoteVolume = this.SafeNumber(ticker, MkString("volume"))
	} else {
		baseVolume = this.SafeNumber(ticker, MkString("volume"))
		quoteVolume = this.SafeNumber(ticker, MkString("quoteVolume"))
	}
	return MkMap(&VarMap{
		"symbol":        symbol,
		"timestamp":     timestamp,
		"datetime":      this.Iso8601(timestamp),
		"high":          this.SafeNumber(ticker, MkString("highPrice")),
		"low":           this.SafeNumber(ticker, MkString("lowPrice")),
		"bid":           this.SafeNumber(ticker, MkString("bidPrice")),
		"bidVolume":     this.SafeNumber(ticker, MkString("bidQty")),
		"ask":           this.SafeNumber(ticker, MkString("askPrice")),
		"askVolume":     this.SafeNumber(ticker, MkString("askQty")),
		"vwap":          this.SafeNumber(ticker, MkString("weightedAvgPrice")),
		"open":          this.SafeNumber(ticker, MkString("openPrice")),
		"close":         last,
		"last":          last,
		"previousClose": this.SafeNumber(ticker, MkString("prevClosePrice")),
		"change":        this.SafeNumber(ticker, MkString("priceChange")),
		"percentage":    this.SafeNumber(ticker, MkString("priceChangePercent")),
		"average":       MkUndefined(),
		"baseVolume":    baseVolume,
		"quoteVolume":   quoteVolume,
		"info":          ticker,
	})
}

func (this *Binance) FetchStatus(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("sapiGetSystemStatus"), params)
	_ = response
	status := this.SafeString(response, MkString("status"))
	_ = status
	if IsTrue(OpNotEq2(status, MkUndefined())) {
		status = OpTrinary(OpEq2(status, MkString("0")), MkString("ok"), MkString("maintenance"))
		*this.At(MkString("status")) = this.Extend(*this.At(MkString("status")), MkMap(&VarMap{
			"status":  status,
			"updated": this.Milliseconds(),
		}))
	}
	return *this.At(MkString("status"))
}

func (this *Binance) FetchTicker(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"symbol": *(market).At(MkString("id"))})
	_ = request
	method := MkString("publicGetTicker24hr")
	_ = method
	if IsTrue(*(market).At(MkString("linear"))) {
		method = MkString("fapiPublicGetTicker24hr")
	} else {
		if IsTrue(*(market).At(MkString("inverse"))) {
			method = MkString("dapiPublicGetTicker24hr")
		}
	}
	response := this.Call(method, this.Extend(request, params))
	_ = response
	if IsTrue(GoIsArray(response)) {
		firstTicker := this.SafeValue(response, MkInteger(0), MkMap(&VarMap{}))
		_ = firstTicker
		return this.ParseTicker(firstTicker, market)
	}
	return this.ParseTicker(response, market)
}

func (this *Binance) FetchBidsAsks(goArgs ...*Variant) *Variant {
	symbols := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbols
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	defaultType := this.SafeString2(*this.At(MkString("options")), MkString("fetchBidsAsks"), MkString("defaultType"), MkString("spot"))
	_ = defaultType
	type_ := this.SafeString(params, MkString("type"), defaultType)
	_ = type_
	query := this.Omit(params, MkString("type"))
	_ = query
	method := OpCopy(MkUndefined())
	_ = method
	if IsTrue(OpEq2(type_, MkString("future"))) {
		method = MkString("fapiPublicGetTickerBookTicker")
	} else {
		if IsTrue(OpEq2(type_, MkString("delivery"))) {
			method = MkString("dapiPublicGetTickerBookTicker")
		} else {
			method = MkString("publicGetTickerBookTicker")
		}
	}
	response := this.Call(method, query)
	_ = response
	return this.ParseTickers(response, symbols)
}

func (this *Binance) FetchTickers(goArgs ...*Variant) *Variant {
	symbols := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbols
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	defaultType := this.SafeString2(*this.At(MkString("options")), MkString("fetchTickers"), MkString("defaultType"), MkString("spot"))
	_ = defaultType
	type_ := this.SafeString(params, MkString("type"), defaultType)
	_ = type_
	query := this.Omit(params, MkString("type"))
	_ = query
	defaultMethod := OpCopy(MkUndefined())
	_ = defaultMethod
	if IsTrue(OpEq2(type_, MkString("future"))) {
		defaultMethod = MkString("fapiPublicGetTicker24hr")
	} else {
		if IsTrue(OpEq2(type_, MkString("delivery"))) {
			defaultMethod = MkString("dapiPublicGetTicker24hr")
		} else {
			defaultMethod = MkString("publicGetTicker24hr")
		}
	}
	method := this.SafeString(*this.At(MkString("options")), MkString("fetchTickersMethod"), defaultMethod)
	_ = method
	response := this.Call(method, query)
	_ = response
	return this.ParseTickers(response, symbols)
}

func (this *Binance) ParseOHLCV(goArgs ...*Variant) *Variant {
	ohlcv := GoGetArg(goArgs, 0, MkUndefined())
	_ = ohlcv
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	return MkArray(&VarArray{
		this.SafeInteger(ohlcv, MkInteger(0)),
		this.SafeNumber(ohlcv, MkInteger(1)),
		this.SafeNumber(ohlcv, MkInteger(2)),
		this.SafeNumber(ohlcv, MkInteger(3)),
		this.SafeNumber(ohlcv, MkInteger(4)),
		this.SafeNumber(ohlcv, MkInteger(5)),
	})
}

func (this *Binance) FetchOHLCV(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	timeframe := GoGetArg(goArgs, 1, MkString("1m"))
	_ = timeframe
	since := GoGetArg(goArgs, 2, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 3, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 4, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	defaultLimit := MkInteger(500)
	_ = defaultLimit
	maxLimit := MkInteger(1500)
	_ = maxLimit
	limit = OpTrinary(OpEq2(limit, MkUndefined()), defaultLimit, Math.Min(limit, maxLimit))
	request := MkMap(&VarMap{
		"symbol":   *(market).At(MkString("id")),
		"interval": *(*this.At(MkString("timeframes"))).At(timeframe),
		"limit":    limit,
	})
	_ = request
	duration := this.ParseTimeframe(timeframe)
	_ = duration
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("startTime")) = OpCopy(since)
		if IsTrue(OpGt(since, MkInteger(0))) {
			endTime := this.Sum(since, OpSub(OpMulti(limit, OpMulti(duration, MkInteger(1000))), MkInteger(1)))
			_ = endTime
			now := this.Milliseconds()
			_ = now
			*(request).At(MkString("endTime")) = Math.Min(now, endTime)
		}
	}
	method := MkString("publicGetKlines")
	_ = method
	if IsTrue(*(market).At(MkString("linear"))) {
		method = MkString("fapiPublicGetKlines")
	} else {
		if IsTrue(*(market).At(MkString("inverse"))) {
			method = MkString("dapiPublicGetKlines")
		}
	}
	response := this.Call(method, this.Extend(request, params))
	_ = response
	return this.ParseOHLCVs(response, market, timeframe, since, limit)
}

func (this *Binance) ParseTrade(goArgs ...*Variant) *Variant {
	trade := GoGetArg(goArgs, 0, MkUndefined())
	_ = trade
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	if IsTrue(OpHasMember(MkString("isDustTrade"), trade)) {
		return this.ParseDustTrade(trade, market)
	}
	timestamp := this.SafeInteger2(trade, MkString("T"), MkString("time"))
	_ = timestamp
	priceString := this.SafeString2(trade, MkString("p"), MkString("price"))
	_ = priceString
	amountString := this.SafeString2(trade, MkString("q"), MkString("qty"))
	_ = amountString
	price := this.ParseNumber(priceString)
	_ = price
	amount := this.ParseNumber(amountString)
	_ = amount
	marketId := this.SafeString(trade, MkString("symbol"))
	_ = marketId
	symbol := this.SafeSymbol(marketId, market)
	_ = symbol
	costString := Precise.StringMul(priceString, amountString)
	_ = costString
	cost := this.ParseNumber(costString)
	_ = cost
	id := this.SafeString2(trade, MkString("t"), MkString("a"))
	_ = id
	id = this.SafeString(trade, MkString("id"), id)
	side := OpCopy(MkUndefined())
	_ = side
	orderId := this.SafeString(trade, MkString("orderId"))
	_ = orderId
	if IsTrue(OpHasMember(MkString("m"), trade)) {
		side = OpTrinary(*(trade).At(MkString("m")), MkString("sell"), MkString("buy"))
	} else {
		if IsTrue(OpHasMember(MkString("isBuyerMaker"), trade)) {
			side = OpTrinary(*(trade).At(MkString("isBuyerMaker")), MkString("sell"), MkString("buy"))
		} else {
			if IsTrue(OpHasMember(MkString("side"), trade)) {
				side = this.SafeStringLower(trade, MkString("side"))
			} else {
				if IsTrue(OpHasMember(MkString("isBuyer"), trade)) {
					side = OpTrinary(*(trade).At(MkString("isBuyer")), MkString("buy"), MkString("sell"))
				}
			}
		}
	}
	fee := OpCopy(MkUndefined())
	_ = fee
	if IsTrue(OpHasMember(MkString("commission"), trade)) {
		fee = MkMap(&VarMap{
			"cost":     this.SafeNumber(trade, MkString("commission")),
			"currency": this.SafeCurrencyCode(this.SafeString(trade, MkString("commissionAsset"))),
		})
	}
	takerOrMaker := OpCopy(MkUndefined())
	_ = takerOrMaker
	if IsTrue(OpHasMember(MkString("isMaker"), trade)) {
		takerOrMaker = OpTrinary(*(trade).At(MkString("isMaker")), MkString("maker"), MkString("taker"))
	}
	if IsTrue(OpHasMember(MkString("maker"), trade)) {
		takerOrMaker = OpTrinary(*(trade).At(MkString("maker")), MkString("maker"), MkString("taker"))
	}
	return MkMap(&VarMap{
		"info":         trade,
		"timestamp":    timestamp,
		"datetime":     this.Iso8601(timestamp),
		"symbol":       symbol,
		"id":           id,
		"order":        orderId,
		"type":         MkUndefined(),
		"side":         side,
		"takerOrMaker": takerOrMaker,
		"price":        price,
		"amount":       amount,
		"cost":         cost,
		"fee":          fee,
	})
}

func (this *Binance) FetchTrades(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"symbol": *(market).At(MkString("id"))})
	_ = request
	defaultType := this.SafeString2(*this.At(MkString("options")), MkString("fetchTrades"), MkString("defaultType"), MkString("spot"))
	_ = defaultType
	type_ := this.SafeString(params, MkString("type"), defaultType)
	_ = type_
	query := this.Omit(params, MkString("type"))
	_ = query
	defaultMethod := OpCopy(MkUndefined())
	_ = defaultMethod
	if IsTrue(OpEq2(type_, MkString("future"))) {
		defaultMethod = MkString("fapiPublicGetAggTrades")
	} else {
		if IsTrue(OpEq2(type_, MkString("delivery"))) {
			defaultMethod = MkString("dapiPublicGetAggTrades")
		} else {
			defaultMethod = MkString("publicGetAggTrades")
		}
	}
	method := this.SafeString(*this.At(MkString("options")), MkString("fetchTradesMethod"), defaultMethod)
	_ = method
	if IsTrue(OpEq2(method, MkString("publicGetAggTrades"))) {
		if IsTrue(OpNotEq2(since, MkUndefined())) {
			*(request).At(MkString("startTime")) = OpCopy(since)
			*(request).At(MkString("endTime")) = this.Sum(since, MkInteger(3600000))
		}
		if IsTrue(OpEq2(type_, MkString("future"))) {
			method = MkString("fapiPublicGetAggTrades")
		} else {
			if IsTrue(OpEq2(type_, MkString("delivery"))) {
				method = MkString("dapiPublicGetAggTrades")
			}
		}
	} else {
		if IsTrue(OpEq2(method, MkString("publicGetHistoricalTrades"))) {
			if IsTrue(OpEq2(type_, MkString("future"))) {
				method = MkString("fapiPublicGetHistoricalTrades")
			} else {
				if IsTrue(OpEq2(type_, MkString("delivery"))) {
					method = MkString("dapiPublicGetHistoricalTrades")
				}
			}
		}
	}
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("limit")) = OpCopy(limit)
	}
	response := this.Call(method, this.Extend(request, query))
	_ = response
	return this.ParseTrades(response, market, since, limit)
}

func (this *Binance) ParseOrderStatus(goArgs ...*Variant) *Variant {
	status := GoGetArg(goArgs, 0, MkUndefined())
	_ = status
	statuses := MkMap(&VarMap{
		"NEW":              MkString("open"),
		"PARTIALLY_FILLED": MkString("open"),
		"FILLED":           MkString("closed"),
		"CANCELED":         MkString("canceled"),
		"PENDING_CANCEL":   MkString("canceling"),
		"REJECTED":         MkString("rejected"),
		"EXPIRED":          MkString("expired"),
	})
	_ = statuses
	return this.SafeString(statuses, status, status)
}

func (this *Binance) ParseOrder(goArgs ...*Variant) *Variant {
	order := GoGetArg(goArgs, 0, MkUndefined())
	_ = order
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	status := this.ParseOrderStatus(this.SafeString(order, MkString("status")))
	_ = status
	marketId := this.SafeString(order, MkString("symbol"))
	_ = marketId
	symbol := this.SafeSymbol(marketId, market)
	_ = symbol
	filledString := this.SafeString(order, MkString("executedQty"), MkString("0"))
	_ = filledString
	filled := this.ParseNumber(filledString)
	_ = filled
	filledFloat := parseFloat(filledString)
	_ = filledFloat
	timestamp := OpCopy(MkUndefined())
	_ = timestamp
	lastTradeTimestamp := OpCopy(MkUndefined())
	_ = lastTradeTimestamp
	if IsTrue(OpHasMember(MkString("time"), order)) {
		timestamp = this.SafeInteger(order, MkString("time"))
	} else {
		if IsTrue(OpHasMember(MkString("transactTime"), order)) {
			timestamp = this.SafeInteger(order, MkString("transactTime"))
		} else {
			if IsTrue(OpHasMember(MkString("updateTime"), order)) {
				if IsTrue(OpEq2(status, MkString("open"))) {
					if IsTrue(OpGt(filledFloat, MkInteger(0))) {
						lastTradeTimestamp = this.SafeInteger(order, MkString("updateTime"))
					} else {
						timestamp = this.SafeInteger(order, MkString("updateTime"))
					}
				}
			}
		}
	}
	averageString := this.SafeString(order, MkString("avgPrice"))
	_ = averageString
	average := this.ParseNumber(this.OmitZero(averageString))
	_ = average
	priceString := this.SafeString(order, MkString("price"))
	_ = priceString
	price := this.ParseNumber(this.OmitZero(priceString))
	_ = price
	amount := this.SafeNumber(order, MkString("origQty"))
	_ = amount
	cost := this.SafeNumber2(order, MkString("cummulativeQuoteQty"), MkString("cumQuote"))
	_ = cost
	id := this.SafeString(order, MkString("orderId"))
	_ = id
	type_ := this.SafeStringLower(order, MkString("type"))
	_ = type_
	side := this.SafeStringLower(order, MkString("side"))
	_ = side
	fills := this.SafeValue(order, MkString("fills"), MkArray(&VarArray{}))
	_ = fills
	trades := this.ParseTrades(fills, market)
	_ = trades
	clientOrderId := this.SafeString(order, MkString("clientOrderId"))
	_ = clientOrderId
	timeInForce := this.SafeString(order, MkString("timeInForce"))
	_ = timeInForce
	postOnly := OpOr(OpEq2(type_, MkString("limit_maker")), OpEq2(timeInForce, MkString("GTX")))
	_ = postOnly
	if IsTrue(OpEq2(type_, MkString("limit_maker"))) {
		type_ = MkString("limit")
	}
	stopPriceString := this.SafeString(order, MkString("stopPrice"))
	_ = stopPriceString
	stopPrice := this.ParseNumber(this.OmitZero(stopPriceString))
	_ = stopPrice
	return this.SafeOrder(MkMap(&VarMap{
		"info":               order,
		"id":                 id,
		"clientOrderId":      clientOrderId,
		"timestamp":          timestamp,
		"datetime":           this.Iso8601(timestamp),
		"lastTradeTimestamp": lastTradeTimestamp,
		"symbol":             symbol,
		"type":               type_,
		"timeInForce":        timeInForce,
		"postOnly":           postOnly,
		"side":               side,
		"price":              price,
		"stopPrice":          stopPrice,
		"amount":             amount,
		"cost":               cost,
		"average":            average,
		"filled":             filled,
		"remaining":          MkUndefined(),
		"status":             status,
		"fee":                MkUndefined(),
		"trades":             trades,
	}))
}

func (this *Binance) CreateOrder(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	type_ := GoGetArg(goArgs, 1, MkUndefined())
	_ = type_
	side := GoGetArg(goArgs, 2, MkUndefined())
	_ = side
	amount := GoGetArg(goArgs, 3, MkUndefined())
	_ = amount
	price := GoGetArg(goArgs, 4, MkUndefined())
	_ = price
	params := GoGetArg(goArgs, 5, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	defaultType := this.SafeString2(*this.At(MkString("options")), MkString("createOrder"), MkString("defaultType"), MkString("spot"))
	_ = defaultType
	orderType := this.SafeString(params, MkString("type"), defaultType)
	_ = orderType
	clientOrderId := this.SafeString2(params, MkString("newClientOrderId"), MkString("clientOrderId"))
	_ = clientOrderId
	params = this.Omit(params, MkArray(&VarArray{
		MkString("type"),
		MkString("newClientOrderId"),
		MkString("clientOrderId"),
	}))
	method := MkString("privatePostOrder")
	_ = method
	if IsTrue(OpEq2(orderType, MkString("future"))) {
		method = MkString("fapiPrivatePostOrder")
	} else {
		if IsTrue(OpEq2(orderType, MkString("delivery"))) {
			method = MkString("dapiPrivatePostOrder")
		} else {
			if IsTrue(OpEq2(orderType, MkString("margin"))) {
				method = MkString("sapiPostMarginOrder")
			}
		}
	}
	if IsTrue(*(market).At(MkString("spot"))) {
		test := this.SafeValue(params, MkString("test"), MkBool(false))
		_ = test
		if IsTrue(test) {
			OpAddAssign(&method, MkString("Test"))
		}
		params = this.Omit(params, MkString("test"))
	}
	uppercaseType := type_.ToUpperCase()
	_ = uppercaseType
	validOrderTypes := this.SafeValue(*(market).At(MkString("info")), MkString("orderTypes"))
	_ = validOrderTypes
	if IsTrue(OpNot(this.InArray(uppercaseType, validOrderTypes))) {
		panic(NewInvalidOrder(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), OpAdd(type_, OpAdd(MkString(" is not a valid order type in market "), symbol))))))
	}
	request := MkMap(&VarMap{
		"symbol": *(market).At(MkString("id")),
		"type":   uppercaseType,
		"side":   side.ToUpperCase(),
	})
	_ = request
	if IsTrue(OpEq2(clientOrderId, MkUndefined())) {
		broker := this.SafeValue(*this.At(MkString("options")), MkString("broker"))
		_ = broker
		if IsTrue(OpNotEq2(broker, MkUndefined())) {
			brokerId := this.SafeString(broker, orderType)
			_ = brokerId
			if IsTrue(OpNotEq2(brokerId, MkUndefined())) {
				*(request).At(MkString("newClientOrderId")) = OpAdd(brokerId, this.Uuid22())
			}
		}
	} else {
		*(request).At(MkString("newClientOrderId")) = OpCopy(clientOrderId)
	}
	if IsTrue(OpOr(OpEq2(orderType, MkString("spot")), OpEq2(orderType, MkString("margin")))) {
		*(request).At(MkString("newOrderRespType")) = this.SafeValue(*(*this.At(MkString("options"))).At(MkString("newOrderRespType")), type_, MkString("RESULT"))
	} else {
		*(request).At(MkString("newOrderRespType")) = MkString("RESULT")
	}
	timeInForceIsRequired := OpCopy(MkBool(false))
	_ = timeInForceIsRequired
	priceIsRequired := OpCopy(MkBool(false))
	_ = priceIsRequired
	stopPriceIsRequired := OpCopy(MkBool(false))
	_ = stopPriceIsRequired
	quantityIsRequired := OpCopy(MkBool(false))
	_ = quantityIsRequired
	if IsTrue(OpEq2(uppercaseType, MkString("MARKET"))) {
		quoteOrderQty := this.SafeValue(*this.At(MkString("options")), MkString("quoteOrderQty"), MkBool(false))
		_ = quoteOrderQty
		if IsTrue(quoteOrderQty) {
			quoteOrderQty := this.SafeNumber(params, MkString("quoteOrderQty"))
			_ = quoteOrderQty
			precision := *(*(market).At(MkString("precision"))).At(MkString("price"))
			_ = precision
			if IsTrue(OpNotEq2(quoteOrderQty, MkUndefined())) {
				*(request).At(MkString("quoteOrderQty")) = this.DecimalToPrecision(quoteOrderQty, TRUNCATE, precision, *this.At(MkString("precisionMode")))
				params = this.Omit(params, MkString("quoteOrderQty"))
			} else {
				if IsTrue(OpNotEq2(price, MkUndefined())) {
					*(request).At(MkString("quoteOrderQty")) = this.DecimalToPrecision(OpMulti(amount, price), TRUNCATE, precision, *this.At(MkString("precisionMode")))
				} else {
					quantityIsRequired = OpCopy(MkBool(true))
				}
			}
		} else {
			quantityIsRequired = OpCopy(MkBool(true))
		}
	} else {
		if IsTrue(OpEq2(uppercaseType, MkString("LIMIT"))) {
			priceIsRequired = OpCopy(MkBool(true))
			timeInForceIsRequired = OpCopy(MkBool(true))
			quantityIsRequired = OpCopy(MkBool(true))
		} else {
			if IsTrue(OpOr(OpEq2(uppercaseType, MkString("STOP_LOSS")), OpEq2(uppercaseType, MkString("TAKE_PROFIT")))) {
				stopPriceIsRequired = OpCopy(MkBool(true))
				quantityIsRequired = OpCopy(MkBool(true))
				if IsTrue(OpOr(*(market).At(MkString("linear")), *(market).At(MkString("inverse")))) {
					priceIsRequired = OpCopy(MkBool(true))
				}
			} else {
				if IsTrue(OpOr(OpEq2(uppercaseType, MkString("STOP_LOSS_LIMIT")), OpEq2(uppercaseType, MkString("TAKE_PROFIT_LIMIT")))) {
					quantityIsRequired = OpCopy(MkBool(true))
					stopPriceIsRequired = OpCopy(MkBool(true))
					priceIsRequired = OpCopy(MkBool(true))
					timeInForceIsRequired = OpCopy(MkBool(true))
				} else {
					if IsTrue(OpEq2(uppercaseType, MkString("LIMIT_MAKER"))) {
						priceIsRequired = OpCopy(MkBool(true))
						quantityIsRequired = OpCopy(MkBool(true))
					} else {
						if IsTrue(OpEq2(uppercaseType, MkString("STOP"))) {
							quantityIsRequired = OpCopy(MkBool(true))
							stopPriceIsRequired = OpCopy(MkBool(true))
							priceIsRequired = OpCopy(MkBool(true))
						} else {
							if IsTrue(OpOr(OpEq2(uppercaseType, MkString("STOP_MARKET")), OpEq2(uppercaseType, MkString("TAKE_PROFIT_MARKET")))) {
								closePosition := this.SafeValue(params, MkString("closePosition"))
								_ = closePosition
								if IsTrue(OpEq2(closePosition, MkUndefined())) {
									quantityIsRequired = OpCopy(MkBool(true))
								}
								stopPriceIsRequired = OpCopy(MkBool(true))
							} else {
								if IsTrue(OpEq2(uppercaseType, MkString("TRAILING_STOP_MARKET"))) {
									quantityIsRequired = OpCopy(MkBool(true))
									callbackRate := this.SafeNumber(params, MkString("callbackRate"))
									_ = callbackRate
									if IsTrue(OpEq2(callbackRate, MkUndefined())) {
										panic(NewInvalidOrder(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" createOrder() requires a callbackRate extra param for a "), OpAdd(type_, MkString(" order"))))))
									}
								}
							}
						}
					}
				}
			}
		}
	}
	if IsTrue(quantityIsRequired) {
		*(request).At(MkString("quantity")) = this.AmountToPrecision(symbol, amount)
	}
	if IsTrue(priceIsRequired) {
		if IsTrue(OpEq2(price, MkUndefined())) {
			panic(NewInvalidOrder(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" createOrder() requires a price argument for a "), OpAdd(type_, MkString(" order"))))))
		}
		*(request).At(MkString("price")) = this.PriceToPrecision(symbol, price)
	}
	if IsTrue(timeInForceIsRequired) {
		*(request).At(MkString("timeInForce")) = *(*this.At(MkString("options"))).At(MkString("defaultTimeInForce"))
	}
	if IsTrue(stopPriceIsRequired) {
		stopPrice := this.SafeNumber(params, MkString("stopPrice"))
		_ = stopPrice
		if IsTrue(OpEq2(stopPrice, MkUndefined())) {
			panic(NewInvalidOrder(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" createOrder() requires a stopPrice extra param for a "), OpAdd(type_, MkString(" order"))))))
		} else {
			params = this.Omit(params, MkString("stopPrice"))
			*(request).At(MkString("stopPrice")) = this.PriceToPrecision(symbol, stopPrice)
		}
	}
	response := this.Call(method, this.Extend(request, params))
	_ = response
	return this.ParseOrder(response, market)
}

func (this *Binance) FetchOrder(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	if IsTrue(OpEq2(symbol, MkUndefined())) {
		panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" fetchOrder() requires a symbol argument"))))
	}
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	defaultType := this.SafeString2(*this.At(MkString("options")), MkString("fetchOrder"), MkString("defaultType"), MkString("spot"))
	_ = defaultType
	type_ := this.SafeString(params, MkString("type"), defaultType)
	_ = type_
	method := MkString("privateGetOrder")
	_ = method
	if IsTrue(OpEq2(type_, MkString("future"))) {
		method = MkString("fapiPrivateGetOrder")
	} else {
		if IsTrue(OpEq2(type_, MkString("delivery"))) {
			method = MkString("dapiPrivateGetOrder")
		} else {
			if IsTrue(OpEq2(type_, MkString("margin"))) {
				method = MkString("sapiGetMarginOrder")
			}
		}
	}
	request := MkMap(&VarMap{"symbol": *(market).At(MkString("id"))})
	_ = request
	clientOrderId := this.SafeValue2(params, MkString("origClientOrderId"), MkString("clientOrderId"))
	_ = clientOrderId
	if IsTrue(OpNotEq2(clientOrderId, MkUndefined())) {
		*(request).At(MkString("origClientOrderId")) = OpCopy(clientOrderId)
	} else {
		*(request).At(MkString("orderId")) = OpCopy(id)
	}
	query := this.Omit(params, MkArray(&VarArray{
		MkString("type"),
		MkString("clientOrderId"),
		MkString("origClientOrderId"),
	}))
	_ = query
	response := this.Call(method, this.Extend(request, query))
	_ = response
	return this.ParseOrder(response, market)
}

func (this *Binance) FetchOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	if IsTrue(OpEq2(symbol, MkUndefined())) {
		panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" fetchOrders() requires a symbol argument"))))
	}
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	defaultType := this.SafeString2(*this.At(MkString("options")), MkString("fetchOrders"), MkString("defaultType"), MkString("spot"))
	_ = defaultType
	type_ := this.SafeString(params, MkString("type"), defaultType)
	_ = type_
	method := MkString("privateGetAllOrders")
	_ = method
	if IsTrue(OpEq2(type_, MkString("future"))) {
		method = MkString("fapiPrivateGetAllOrders")
	} else {
		if IsTrue(OpEq2(type_, MkString("delivery"))) {
			method = MkString("dapiPrivateGetAllOrders")
		} else {
			if IsTrue(OpEq2(type_, MkString("margin"))) {
				method = MkString("sapiGetMarginAllOrders")
			}
		}
	}
	request := MkMap(&VarMap{"symbol": *(market).At(MkString("id"))})
	_ = request
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("startTime")) = OpCopy(since)
	}
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("limit")) = OpCopy(limit)
	}
	query := this.Omit(params, MkString("type"))
	_ = query
	response := this.Call(method, this.Extend(request, query))
	_ = response
	return this.ParseOrders(response, market, since, limit)
}

func (this *Binance) FetchOpenOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := OpCopy(MkUndefined())
	_ = market
	query := OpCopy(MkUndefined())
	_ = query
	type_ := OpCopy(MkUndefined())
	_ = type_
	request := MkMap(&VarMap{})
	_ = request
	if IsTrue(OpNotEq2(symbol, MkUndefined())) {
		market = this.Market(symbol)
		*(request).At(MkString("symbol")) = *(market).At(MkString("id"))
		defaultType := this.SafeString2(*this.At(MkString("options")), MkString("fetchOpenOrders"), MkString("defaultType"), MkString("spot"))
		_ = defaultType
		type_ = this.SafeString(params, MkString("type"), defaultType)
		query = this.Omit(params, MkString("type"))
	} else {
		if IsTrue(*(*this.At(MkString("options"))).At(MkString("warnOnFetchOpenOrdersWithoutSymbol"))) {
			symbols := OpCopy(*this.At(MkString("symbols")))
			_ = symbols
			numSymbols := OpCopy(symbols.Length)
			_ = numSymbols
			fetchOpenOrdersRateLimit := parseInt(OpDiv(numSymbols, MkInteger(2)))
			_ = fetchOpenOrdersRateLimit
			panic(NewExchangeError(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" fetchOpenOrders WARNING: fetching open orders without specifying a symbol is rate-limited to one call per "), OpAdd(fetchOpenOrdersRateLimit.ToString(), OpAdd(MkString(" seconds. Do not call this method frequently to avoid ban. Set "), OpAdd(*this.At(MkString("id")), MkString(".options[\"warnOnFetchOpenOrdersWithoutSymbol\"] = false to suppress this warning message."))))))))
		} else {
			defaultType := this.SafeString2(*this.At(MkString("options")), MkString("fetchOpenOrders"), MkString("defaultType"), MkString("spot"))
			_ = defaultType
			type_ = this.SafeString(params, MkString("type"), defaultType)
			query = this.Omit(params, MkString("type"))
		}
	}
	method := MkString("privateGetOpenOrders")
	_ = method
	if IsTrue(OpEq2(type_, MkString("future"))) {
		method = MkString("fapiPrivateGetOpenOrders")
	} else {
		if IsTrue(OpEq2(type_, MkString("delivery"))) {
			method = MkString("dapiPrivateGetOpenOrders")
		} else {
			if IsTrue(OpEq2(type_, MkString("margin"))) {
				method = MkString("sapiGetMarginOpenOrders")
			}
		}
	}
	response := this.Call(method, this.Extend(request, query))
	_ = response
	return this.ParseOrders(response, market, since, limit)
}

func (this *Binance) FetchClosedOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	orders := this.FetchOrders(symbol, since, limit, params)
	_ = orders
	return this.FilterBy(orders, MkString("status"), MkString("closed"))
}

func (this *Binance) CancelOrder(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	if IsTrue(OpEq2(symbol, MkUndefined())) {
		panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" cancelOrder() requires a symbol argument"))))
	}
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	defaultType := this.SafeString2(*this.At(MkString("options")), MkString("fetchOpenOrders"), MkString("defaultType"), MkString("spot"))
	_ = defaultType
	type_ := this.SafeString(params, MkString("type"), defaultType)
	_ = type_
	origClientOrderId := this.SafeValue2(params, MkString("origClientOrderId"), MkString("clientOrderId"))
	_ = origClientOrderId
	request := MkMap(&VarMap{"symbol": *(market).At(MkString("id"))})
	_ = request
	if IsTrue(OpEq2(origClientOrderId, MkUndefined())) {
		*(request).At(MkString("orderId")) = OpCopy(id)
	} else {
		*(request).At(MkString("origClientOrderId")) = OpCopy(origClientOrderId)
	}
	method := MkString("privateDeleteOrder")
	_ = method
	if IsTrue(OpEq2(type_, MkString("future"))) {
		method = MkString("fapiPrivateDeleteOrder")
	} else {
		if IsTrue(OpEq2(type_, MkString("delivery"))) {
			method = MkString("dapiPrivateDeleteOrder")
		} else {
			if IsTrue(OpEq2(type_, MkString("margin"))) {
				method = MkString("sapiDeleteMarginOrder")
			}
		}
	}
	query := this.Omit(params, MkArray(&VarArray{
		MkString("type"),
		MkString("origClientOrderId"),
		MkString("clientOrderId"),
	}))
	_ = query
	response := this.Call(method, this.Extend(request, query))
	_ = response
	return this.ParseOrder(response)
}

func (this *Binance) CancelAllOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	if IsTrue(OpEq2(symbol, MkUndefined())) {
		panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" cancelAllOrders() requires a symbol argument"))))
	}
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"symbol": *(market).At(MkString("id"))})
	_ = request
	defaultType := this.SafeString2(*this.At(MkString("options")), MkString("cancelAllOrders"), MkString("defaultType"), MkString("spot"))
	_ = defaultType
	type_ := this.SafeString(params, MkString("type"), defaultType)
	_ = type_
	query := this.Omit(params, MkString("type"))
	_ = query
	method := MkString("privateDeleteOpenOrders")
	_ = method
	if IsTrue(OpEq2(type_, MkString("margin"))) {
		method = MkString("sapiDeleteMarginOpenOrders")
	} else {
		if IsTrue(OpEq2(type_, MkString("future"))) {
			method = MkString("fapiPrivateDeleteAllOpenOrders")
		} else {
			if IsTrue(OpEq2(type_, MkString("delivery"))) {
				method = MkString("dapiPrivateDeleteAllOpenOrders")
			}
		}
	}
	response := this.Call(method, this.Extend(request, query))
	_ = response
	if IsTrue(GoIsArray(response)) {
		return this.ParseOrders(response, market)
	} else {
		return response
	}
	return MkUndefined()
}

func (this *Binance) FetchMyTrades(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	if IsTrue(OpEq2(symbol, MkUndefined())) {
		panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" fetchMyTrades() requires a symbol argument"))))
	}
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	defaultType := this.SafeString2(*this.At(MkString("options")), MkString("fetchMyTrades"), MkString("defaultType"), MkString("spot"))
	_ = defaultType
	type_ := this.SafeString(params, MkString("type"), defaultType)
	_ = type_
	params = this.Omit(params, MkString("type"))
	method := OpCopy(MkUndefined())
	_ = method
	if IsTrue(OpEq2(type_, MkString("spot"))) {
		method = MkString("privateGetMyTrades")
	} else {
		if IsTrue(OpEq2(type_, MkString("margin"))) {
			method = MkString("sapiGetMarginMyTrades")
		} else {
			if IsTrue(OpEq2(type_, MkString("future"))) {
				method = MkString("fapiPrivateGetUserTrades")
			} else {
				if IsTrue(OpEq2(type_, MkString("delivery"))) {
					method = MkString("dapiPrivateGetUserTrades")
				}
			}
		}
	}
	request := MkMap(&VarMap{"symbol": *(market).At(MkString("id"))})
	_ = request
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("startTime")) = OpCopy(since)
	}
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("limit")) = OpCopy(limit)
	}
	response := this.Call(method, this.Extend(request, params))
	_ = response
	return this.ParseTrades(response, market, since, limit)
}

func (this *Binance) FetchMyDustTrades(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{})
	_ = request
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("startTime")) = OpCopy(since)
		*(request).At(MkString("endTime")) = this.Sum(since, MkInteger(7776000000))
	}
	response := this.Call(MkString("sapiGetAssetDribblet"), this.Extend(request, params))
	_ = response
	results := this.SafeValue(response, MkString("userAssetDribblets"), MkArray(&VarArray{}))
	_ = results
	rows := this.SafeInteger(response, MkString("total"), MkInteger(0))
	_ = rows
	data := MkArray(&VarArray{})
	_ = data
	for i := MkInteger(0); IsTrue(OpLw(i, rows)); OpIncr(&i) {
		logs := this.SafeValue(*(results).At(i), MkString("userAssetDribbletDetails"), MkArray(&VarArray{}))
		_ = logs
		for j := MkInteger(0); IsTrue(OpLw(j, logs.Length)); OpIncr(&j) {
			*(*(logs).At(j)).At(MkString("isDustTrade")) = OpCopy(MkBool(true))
			data.Push(*(logs).At(j))
		}
	}
	trades := this.ParseTrades(data, MkUndefined(), since, limit)
	_ = trades
	return this.FilterBySinceLimit(trades, since, limit)
}

func (this *Binance) ParseDustTrade(goArgs ...*Variant) *Variant {
	trade := GoGetArg(goArgs, 0, MkUndefined())
	_ = trade
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	orderId := this.SafeString(trade, MkString("transId"))
	_ = orderId
	timestamp := this.Parse8601(this.SafeString(trade, MkString("operateTime")))
	_ = timestamp
	currencyId := this.SafeString(trade, MkString("fromAsset"))
	_ = currencyId
	tradedCurrency := this.SafeCurrencyCode(currencyId)
	_ = tradedCurrency
	bnb := this.Currency(MkString("BNB"))
	_ = bnb
	earnedCurrency := *(bnb).At(MkString("code"))
	_ = earnedCurrency
	applicantSymbol := OpAdd(earnedCurrency, OpAdd(MkString("/"), tradedCurrency))
	_ = applicantSymbol
	tradedCurrencyIsQuote := OpCopy(MkBool(false))
	_ = tradedCurrencyIsQuote
	if IsTrue(OpHasMember(applicantSymbol, *this.At(MkString("markets")))) {
		tradedCurrencyIsQuote = OpCopy(MkBool(true))
	}
	feeCostString := this.SafeString(trade, MkString("serviceChargeAmount"))
	_ = feeCostString
	fee := MkMap(&VarMap{
		"currency": earnedCurrency,
		"cost":     this.ParseNumber(feeCostString),
	})
	_ = fee
	symbol := OpCopy(MkUndefined())
	_ = symbol
	amountString := OpCopy(MkUndefined())
	_ = amountString
	costString := OpCopy(MkUndefined())
	_ = costString
	side := OpCopy(MkUndefined())
	_ = side
	if IsTrue(tradedCurrencyIsQuote) {
		symbol = OpCopy(applicantSymbol)
		amountString = Precise.StringAdd(this.SafeString(trade, MkString("transferedAmount")), feeCostString)
		costString = this.SafeString(trade, MkString("amount"))
		side = MkString("buy")
	} else {
		symbol = OpAdd(tradedCurrency, OpAdd(MkString("/"), earnedCurrency))
		amountString = this.SafeString(trade, MkString("amount"))
		costString = Precise.StringAdd(this.SafeString(trade, MkString("transferedAmount")), feeCostString)
		side = MkString("sell")
	}
	priceString := OpCopy(MkUndefined())
	_ = priceString
	if IsTrue(OpNotEq2(costString, MkUndefined())) {
		if IsTrue(amountString) {
			priceString = Precise.StringDiv(costString, amountString)
		}
	}
	id := OpCopy(MkUndefined())
	_ = id
	amount := this.ParseNumber(amountString)
	_ = amount
	price := this.ParseNumber(priceString)
	_ = price
	cost := this.ParseNumber(costString)
	_ = cost
	type_ := OpCopy(MkUndefined())
	_ = type_
	takerOrMaker := OpCopy(MkUndefined())
	_ = takerOrMaker
	return MkMap(&VarMap{
		"id":           id,
		"timestamp":    timestamp,
		"datetime":     this.Iso8601(timestamp),
		"symbol":       symbol,
		"order":        orderId,
		"type":         type_,
		"takerOrMaker": takerOrMaker,
		"side":         side,
		"amount":       amount,
		"price":        price,
		"cost":         cost,
		"fee":          fee,
		"info":         trade,
	})
}

func (this *Binance) FetchDeposits(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	currency := OpCopy(MkUndefined())
	_ = currency
	response := OpCopy(MkUndefined())
	_ = response
	request := MkMap(&VarMap{})
	_ = request
	legalMoney := this.SafeValue(*this.At(MkString("options")), MkString("legalMoney"), MkMap(&VarMap{}))
	_ = legalMoney
	if IsTrue(OpHasMember(code, legalMoney)) {
		if IsTrue(OpNotEq2(code, MkUndefined())) {
			currency = this.Currency(code)
		}
		*(request).At(MkString("transactionType")) = MkInteger(0)
		if IsTrue(OpNotEq2(since, MkUndefined())) {
			*(request).At(MkString("beginTime")) = OpCopy(since)
		}
		raw := this.Call(MkString("sapiGetFiatOrders"), this.Extend(request, params))
		_ = raw
		response = this.SafeValue(raw, MkString("data"))
	} else {
		if IsTrue(OpNotEq2(code, MkUndefined())) {
			currency = this.Currency(code)
			*(request).At(MkString("coin")) = *(currency).At(MkString("id"))
		}
		if IsTrue(OpNotEq2(since, MkUndefined())) {
			*(request).At(MkString("startTime")) = OpCopy(since)
			*(request).At(MkString("endTime")) = this.Sum(since, MkInteger(7776000000))
		}
		if IsTrue(OpNotEq2(limit, MkUndefined())) {
			*(request).At(MkString("limit")) = OpCopy(limit)
		}
		response = this.Call(MkString("sapiGetCapitalDepositHisrec"), this.Extend(request, params))
	}
	return this.ParseTransactions(response, currency, since, limit)
}

func (this *Binance) FetchWithdrawals(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	legalMoney := this.SafeValue(*this.At(MkString("options")), MkString("legalMoney"), MkMap(&VarMap{}))
	_ = legalMoney
	request := MkMap(&VarMap{})
	_ = request
	response := OpCopy(MkUndefined())
	_ = response
	currency := OpCopy(MkUndefined())
	_ = currency
	if IsTrue(OpHasMember(code, legalMoney)) {
		if IsTrue(OpNotEq2(code, MkUndefined())) {
			currency = this.Currency(code)
		}
		*(request).At(MkString("transactionType")) = MkInteger(1)
		if IsTrue(OpNotEq2(since, MkUndefined())) {
			*(request).At(MkString("beginTime")) = OpCopy(since)
		}
		raw := this.Call(MkString("sapiGetFiatOrders"), this.Extend(request, params))
		_ = raw
		response = this.SafeValue(raw, MkString("data"))
	} else {
		if IsTrue(OpNotEq2(code, MkUndefined())) {
			currency = this.Currency(code)
			*(request).At(MkString("coin")) = *(currency).At(MkString("id"))
		}
		if IsTrue(OpNotEq2(since, MkUndefined())) {
			*(request).At(MkString("startTime")) = OpCopy(since)
			*(request).At(MkString("endTime")) = this.Sum(since, MkInteger(7776000000))
		}
		if IsTrue(OpNotEq2(limit, MkUndefined())) {
			*(request).At(MkString("limit")) = OpCopy(limit)
		}
		response = this.Call(MkString("sapiGetCapitalWithdrawHistory"), this.Extend(request, params))
	}
	return this.ParseTransactions(response, currency, since, limit)
}

func (this *Binance) ParseTransactionStatusByType(goArgs ...*Variant) *Variant {
	status := GoGetArg(goArgs, 0, MkUndefined())
	_ = status
	type_ := GoGetArg(goArgs, 1, MkUndefined())
	_ = type_
	statusesByType := MkMap(&VarMap{
		"deposit": MkMap(&VarMap{
			"0":             MkString("pending"),
			"1":             MkString("ok"),
			"Processing":    MkString("pending"),
			"Failed":        MkString("failed"),
			"Successful":    MkString("ok"),
			"Refunding":     MkString("canceled"),
			"Refunded":      MkString("canceled"),
			"Refund Failed": MkString("failed"),
		}),
		"withdrawal": MkMap(&VarMap{
			"0":             MkString("pending"),
			"1":             MkString("canceled"),
			"2":             MkString("pending"),
			"3":             MkString("failed"),
			"4":             MkString("pending"),
			"5":             MkString("failed"),
			"6":             MkString("ok"),
			"Processing":    MkString("pending"),
			"Failed":        MkString("failed"),
			"Successful":    MkString("ok"),
			"Refunding":     MkString("canceled"),
			"Refunded":      MkString("canceled"),
			"Refund Failed": MkString("failed"),
		}),
	})
	_ = statusesByType
	statuses := this.SafeValue(statusesByType, type_, MkMap(&VarMap{}))
	_ = statuses
	return this.SafeString(statuses, status, status)
}

func (this *Binance) ParseTransaction(goArgs ...*Variant) *Variant {
	transaction := GoGetArg(goArgs, 0, MkUndefined())
	_ = transaction
	currency := GoGetArg(goArgs, 1, MkUndefined())
	_ = currency
	id := this.SafeString2(transaction, MkString("id"), MkString("orderNo"))
	_ = id
	address := this.SafeString(transaction, MkString("address"))
	_ = address
	tag := this.SafeString(transaction, MkString("addressTag"))
	_ = tag
	if IsTrue(OpNotEq2(tag, MkUndefined())) {
		if IsTrue(OpLw(tag.Length, MkInteger(1))) {
			tag = OpCopy(MkUndefined())
		}
	}
	txid := this.SafeString(transaction, MkString("txId"))
	_ = txid
	if IsTrue(OpAnd(OpNotEq2(txid, MkUndefined()), OpGtEq(txid.IndexOf(MkString("Internal transfer ")), MkInteger(0)))) {
		txid = txid.Slice(MkInteger(18))
	}
	currencyId := this.SafeString2(transaction, MkString("coin"), MkString("fiatCurrency"))
	_ = currencyId
	code := this.SafeCurrencyCode(currencyId, currency)
	_ = code
	timestamp := OpCopy(MkUndefined())
	_ = timestamp
	insertTime := this.SafeInteger2(transaction, MkString("insertTime"), MkString("createTime"))
	_ = insertTime
	applyTime := this.Parse8601(this.SafeString(transaction, MkString("applyTime")))
	_ = applyTime
	type_ := this.SafeString(transaction, MkString("type"))
	_ = type_
	if IsTrue(OpEq2(type_, MkUndefined())) {
		if IsTrue(OpAnd(OpNotEq2(insertTime, MkUndefined()), OpEq2(applyTime, MkUndefined()))) {
			type_ = MkString("deposit")
			timestamp = OpCopy(insertTime)
		} else {
			if IsTrue(OpAnd(OpEq2(insertTime, MkUndefined()), OpNotEq2(applyTime, MkUndefined()))) {
				type_ = MkString("withdrawal")
				timestamp = OpCopy(applyTime)
			}
		}
	}
	status := this.ParseTransactionStatusByType(this.SafeString(transaction, MkString("status")), type_)
	_ = status
	amount := this.SafeNumber(transaction, MkString("amount"))
	_ = amount
	feeCost := this.SafeNumber2(transaction, MkString("transactionFee"), MkString("totalFee"))
	_ = feeCost
	fee := OpCopy(MkUndefined())
	_ = fee
	if IsTrue(OpNotEq2(feeCost, MkUndefined())) {
		fee = MkMap(&VarMap{
			"currency": code,
			"cost":     feeCost,
		})
	}
	updated := this.SafeInteger2(transaction, MkString("successTime"), MkString("updateTime"))
	_ = updated
	internal := this.SafeInteger(transaction, MkString("transferType"), MkBool(false))
	_ = internal
	internal = OpTrinary(internal, MkBool(true), MkBool(false))
	return MkMap(&VarMap{
		"info":        transaction,
		"id":          id,
		"txid":        txid,
		"timestamp":   timestamp,
		"datetime":    this.Iso8601(timestamp),
		"address":     address,
		"addressTo":   address,
		"addressFrom": MkUndefined(),
		"tag":         tag,
		"tagTo":       tag,
		"tagFrom":     MkUndefined(),
		"type":        type_,
		"amount":      amount,
		"currency":    code,
		"status":      status,
		"updated":     updated,
		"internal":    internal,
		"fee":         fee,
	})
}

func (this *Binance) ParseTransferStatus(goArgs ...*Variant) *Variant {
	status := GoGetArg(goArgs, 0, MkUndefined())
	_ = status
	statuses := MkMap(&VarMap{"CONFIRMED": MkString("ok")})
	_ = statuses
	return this.SafeString(statuses, status, status)
}

func (this *Binance) ParseTransfer(goArgs ...*Variant) *Variant {
	transfer := GoGetArg(goArgs, 0, MkUndefined())
	_ = transfer
	currency := GoGetArg(goArgs, 1, MkUndefined())
	_ = currency
	id := this.SafeString(transfer, MkString("tranId"))
	_ = id
	currencyId := this.SafeString(transfer, MkString("asset"))
	_ = currencyId
	code := this.SafeCurrencyCode(currencyId, currency)
	_ = code
	amount := this.SafeNumber(transfer, MkString("amount"))
	_ = amount
	type_ := this.SafeString(transfer, MkString("type"))
	_ = type_
	fromAccount := OpCopy(MkUndefined())
	_ = fromAccount
	toAccount := OpCopy(MkUndefined())
	_ = toAccount
	typesByAccount := this.SafeValue(*this.At(MkString("options")), MkString("typesByAccount"), MkMap(&VarMap{}))
	_ = typesByAccount
	if IsTrue(OpNotEq2(type_, MkUndefined())) {
		parts := type_.Split(MkString("_"))
		_ = parts
		fromAccount = this.SafeValue(parts, MkInteger(0))
		toAccount = this.SafeValue(parts, MkInteger(1))
		fromAccount = this.SafeString(typesByAccount, fromAccount, fromAccount)
		toAccount = this.SafeString(typesByAccount, toAccount, toAccount)
	}
	timestamp := this.SafeInteger(transfer, MkString("timestamp"))
	_ = timestamp
	status := this.ParseTransferStatus(this.SafeString(transfer, MkString("status")))
	_ = status
	return MkMap(&VarMap{
		"info":        transfer,
		"id":          id,
		"timestamp":   timestamp,
		"datetime":    this.Iso8601(timestamp),
		"currency":    code,
		"amount":      amount,
		"fromAccount": fromAccount,
		"toAccount":   toAccount,
		"status":      status,
	})
}

func (this *Binance) ParseIncome(goArgs ...*Variant) *Variant {
	income := GoGetArg(goArgs, 0, MkUndefined())
	_ = income
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	marketId := this.SafeString(income, MkString("symbol"))
	_ = marketId
	symbol := this.SafeSymbol(marketId, market)
	_ = symbol
	amount := this.SafeNumber(income, MkString("income"))
	_ = amount
	currencyId := this.SafeString(income, MkString("asset"))
	_ = currencyId
	code := this.SafeCurrencyCode(currencyId)
	_ = code
	id := this.SafeString(income, MkString("tranId"))
	_ = id
	timestamp := this.SafeInteger(income, MkString("time"))
	_ = timestamp
	return MkMap(&VarMap{
		"info":      income,
		"symbol":    symbol,
		"code":      code,
		"timestamp": timestamp,
		"datetime":  this.Iso8601(timestamp),
		"id":        id,
		"amount":    amount,
	})
}

func (this *Binance) ParseIncomes(goArgs ...*Variant) *Variant {
	incomes := GoGetArg(goArgs, 0, MkUndefined())
	_ = incomes
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	since := GoGetArg(goArgs, 2, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 3, MkUndefined())
	_ = limit
	result := MkArray(&VarArray{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, incomes.Length)); OpIncr(&i) {
		entry := *(incomes).At(i)
		_ = entry
		parsed := this.ParseIncome(entry, market)
		_ = parsed
		result.Push(parsed)
	}
	return this.FilterBySinceLimit(result, since, limit, MkString("timestamp"))
}

func (this *Binance) Transfer(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	amount := GoGetArg(goArgs, 1, MkUndefined())
	_ = amount
	fromAccount := GoGetArg(goArgs, 2, MkUndefined())
	_ = fromAccount
	toAccount := GoGetArg(goArgs, 3, MkUndefined())
	_ = toAccount
	params := GoGetArg(goArgs, 4, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	currency := this.Currency(code)
	_ = currency
	type_ := this.SafeString(params, MkString("type"))
	_ = type_
	if IsTrue(OpEq2(type_, MkUndefined())) {
		accountsByType := this.SafeValue(*this.At(MkString("options")), MkString("accountsByType"), MkMap(&VarMap{}))
		_ = accountsByType
		fromId := this.SafeString(accountsByType, fromAccount, fromAccount)
		_ = fromId
		toId := this.SafeString(accountsByType, toAccount, toAccount)
		_ = toId
		if IsTrue(OpEq2(fromId, MkUndefined())) {
			keys := GoGetKeys(accountsByType)
			_ = keys
			panic(NewExchangeError(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" fromAccount must be one of "), keys.Join(MkString(", "))))))
		}
		if IsTrue(OpEq2(toId, MkUndefined())) {
			keys := GoGetKeys(accountsByType)
			_ = keys
			panic(NewExchangeError(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" toAccount must be one of "), keys.Join(MkString(", "))))))
		}
		type_ = OpAdd(fromId, OpAdd(MkString("_"), toId))
	}
	request := MkMap(&VarMap{
		"asset":  *(currency).At(MkString("id")),
		"amount": this.CurrencyToPrecision(code, amount),
		"type":   type_,
	})
	_ = request
	response := this.Call(MkString("sapiPostAssetTransfer"), this.Extend(request, params))
	_ = response
	transfer := this.ParseTransfer(response, currency)
	_ = transfer
	return this.Extend(transfer, MkMap(&VarMap{
		"amount":      amount,
		"currency":    code,
		"fromAccount": fromAccount,
		"toAccount":   toAccount,
	}))
}

func (this *Binance) FetchTransfers(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	currency := this.Currency(code)
	_ = currency
	defaultType := this.SafeString2(*this.At(MkString("options")), MkString("fetchTransfers"), MkString("defaultType"), MkString("spot"))
	_ = defaultType
	fromAccount := this.SafeString(params, MkString("fromAccount"), defaultType)
	_ = fromAccount
	defaultTo := OpTrinary(OpEq2(fromAccount, MkString("future")), MkString("spot"), MkString("future"))
	_ = defaultTo
	toAccount := this.SafeString(params, MkString("toAccount"), defaultTo)
	_ = toAccount
	type_ := this.SafeString(params, MkString("type"))
	_ = type_
	accountsByType := this.SafeValue(*this.At(MkString("options")), MkString("accountsByType"), MkMap(&VarMap{}))
	_ = accountsByType
	fromId := this.SafeString(accountsByType, fromAccount)
	_ = fromId
	toId := this.SafeString(accountsByType, toAccount)
	_ = toId
	if IsTrue(OpEq2(type_, MkUndefined())) {
		if IsTrue(OpEq2(fromId, MkUndefined())) {
			keys := GoGetKeys(accountsByType)
			_ = keys
			panic(NewExchangeError(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" fromAccount parameter must be one of "), keys.Join(MkString(", "))))))
		}
		if IsTrue(OpEq2(toId, MkUndefined())) {
			keys := GoGetKeys(accountsByType)
			_ = keys
			panic(NewExchangeError(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" toAccount parameter must be one of "), keys.Join(MkString(", "))))))
		}
		type_ = OpAdd(fromId, OpAdd(MkString("_"), toId))
	}
	request := MkMap(&VarMap{"type": type_})
	_ = request
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("startTime")) = OpCopy(since)
	}
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("size")) = OpCopy(limit)
	}
	response := this.Call(MkString("sapiGetAssetTransfer"), this.Extend(request, params))
	_ = response
	rows := this.SafeValue(response, MkString("rows"), MkArray(&VarArray{}))
	_ = rows
	return this.ParseTransfers(rows, currency, since, limit)
}

func (this *Binance) FetchDepositAddress(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	currency := this.Currency(code)
	_ = currency
	request := MkMap(&VarMap{"coin": *(currency).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("sapiGetCapitalDepositAddress"), this.Extend(request, params))
	_ = response
	address := this.SafeString(response, MkString("address"))
	_ = address
	tag := this.SafeString(response, MkString("tag"))
	_ = tag
	this.CheckAddress(address)
	return MkMap(&VarMap{
		"currency": code,
		"address":  address,
		"tag":      tag,
		"info":     response,
	})
}

func (this *Binance) FetchFundingFees(goArgs ...*Variant) *Variant {
	codes := GoGetArg(goArgs, 0, MkUndefined())
	_ = codes
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("sapiGetCapitalConfigGetall"), params)
	_ = response
	withdrawFees := MkMap(&VarMap{})
	_ = withdrawFees
	for i := MkInteger(0); IsTrue(OpLw(i, response.Length)); OpIncr(&i) {
		entry := *(response).At(i)
		_ = entry
		currencyId := this.SafeString(entry, MkString("coin"))
		_ = currencyId
		code := this.SafeCurrencyCode(currencyId)
		_ = code
		networkList := this.SafeValue(entry, MkString("networkList"))
		_ = networkList
		*(withdrawFees).At(code) = MkMap(&VarMap{})
		for j := MkInteger(0); IsTrue(OpLw(j, networkList.Length)); OpIncr(&j) {
			networkEntry := *(networkList).At(j)
			_ = networkEntry
			networkId := this.SafeString(networkEntry, MkString("network"))
			_ = networkId
			networkCode := this.SafeCurrencyCode(networkId)
			_ = networkCode
			fee := this.SafeNumber(networkEntry, MkString("withdrawFee"))
			_ = fee
			*(*(withdrawFees).At(code)).At(networkCode) = OpCopy(fee)
		}
	}
	return MkMap(&VarMap{
		"withdraw": withdrawFees,
		"deposit":  MkMap(&VarMap{}),
		"info":     response,
	})
}

func (this *Binance) Withdraw(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	amount := GoGetArg(goArgs, 1, MkUndefined())
	_ = amount
	address := GoGetArg(goArgs, 2, MkUndefined())
	_ = address
	tag := GoGetArg(goArgs, 3, MkUndefined())
	_ = tag
	params := GoGetArg(goArgs, 4, MkMap(&VarMap{}))
	_ = params
	this.CheckAddress(address)
	this.LoadMarkets()
	currency := this.Currency(code)
	_ = currency
	request := MkMap(&VarMap{
		"coin":    *(currency).At(MkString("id")),
		"address": address,
		"amount":  amount,
	})
	_ = request
	if IsTrue(OpNotEq2(tag, MkUndefined())) {
		*(request).At(MkString("addressTag")) = OpCopy(tag)
	}
	response := this.Call(MkString("sapiPostCapitalWithdrawApply"), this.Extend(request, params))
	_ = response
	return MkMap(&VarMap{
		"info": response,
		"id":   this.SafeString(response, MkString("id")),
	})
}

func (this *Binance) ParseTradingFee(goArgs ...*Variant) *Variant {
	fee := GoGetArg(goArgs, 0, MkUndefined())
	_ = fee
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	marketId := this.SafeString(fee, MkString("symbol"))
	_ = marketId
	symbol := this.SafeSymbol(marketId)
	_ = symbol
	return MkMap(&VarMap{
		"info":   fee,
		"symbol": symbol,
		"maker":  this.SafeNumber(fee, MkString("makerCommission")),
		"taker":  this.SafeNumber(fee, MkString("takerCommission")),
	})
}

func (this *Binance) FetchTradingFee(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"symbol": *(market).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("sapiGetAssetTradeFee"), this.Extend(request, params))
	_ = response
	first := this.SafeValue(response, MkInteger(0), MkMap(&VarMap{}))
	_ = first
	return this.ParseTradingFee(first)
}

func (this *Binance) FetchTradingFees(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	method := OpCopy(MkUndefined())
	_ = method
	defaultType := this.SafeString2(*this.At(MkString("options")), MkString("fetchFundingRates"), MkString("defaultType"), MkString("future"))
	_ = defaultType
	type_ := this.SafeString(params, MkString("type"), defaultType)
	_ = type_
	query := this.Omit(params, MkString("type"))
	_ = query
	if IsTrue(OpOr(OpEq2(type_, MkString("spot")), OpEq2(type_, MkString("margin")))) {
		method = MkString("sapiGetAssetTradeFee")
	} else {
		if IsTrue(OpEq2(type_, MkString("future"))) {
			method = MkString("fapiPrivateGetAccount")
		} else {
			if IsTrue(OpEq2(type_, MkString("delivery"))) {
				method = MkString("dapiPrivateGetAccount")
			}
		}
	}
	response := this.Call(method, query)
	_ = response
	if IsTrue(OpOr(OpEq2(type_, MkString("spot")), OpEq2(type_, MkString("margin")))) {
		result := MkMap(&VarMap{})
		_ = result
		for i := MkInteger(0); IsTrue(OpLw(i, response.Length)); OpIncr(&i) {
			fee := this.ParseTradingFee(*(response).At(i))
			_ = fee
			symbol := *(fee).At(MkString("symbol"))
			_ = symbol
			*(result).At(symbol) = OpCopy(fee)
		}
		return result
	} else {
		if IsTrue(OpEq2(type_, MkString("future"))) {
			symbols := GoGetKeys(*this.At(MkString("markets")))
			_ = symbols
			result := MkMap(&VarMap{})
			_ = result
			feeTier := this.SafeInteger(response, MkString("feeTier"))
			_ = feeTier
			feeTiers := *(*(*(*this.At(MkString("fees"))).At(type_)).At(MkString("trading"))).At(MkString("tiers"))
			_ = feeTiers
			maker := *(*(*(feeTiers).At(MkString("maker"))).At(feeTier)).At(MkInteger(1))
			_ = maker
			taker := *(*(*(feeTiers).At(MkString("taker"))).At(feeTier)).At(MkInteger(1))
			_ = taker
			for i := MkInteger(0); IsTrue(OpLw(i, symbols.Length)); OpIncr(&i) {
				symbol := *(symbols).At(i)
				_ = symbol
				*(result).At(symbol) = MkMap(&VarMap{
					"info":   MkMap(&VarMap{"feeTier": feeTier}),
					"symbol": symbol,
					"maker":  maker,
					"taker":  taker,
				})
			}
			return result
		} else {
			if IsTrue(OpEq2(type_, MkString("delivery"))) {
				symbols := GoGetKeys(*this.At(MkString("markets")))
				_ = symbols
				result := MkMap(&VarMap{})
				_ = result
				feeTier := this.SafeInteger(response, MkString("feeTier"))
				_ = feeTier
				feeTiers := *(*(*(*this.At(MkString("fees"))).At(type_)).At(MkString("trading"))).At(MkString("tiers"))
				_ = feeTiers
				maker := *(*(*(feeTiers).At(MkString("maker"))).At(feeTier)).At(MkInteger(1))
				_ = maker
				taker := *(*(*(feeTiers).At(MkString("taker"))).At(feeTier)).At(MkInteger(1))
				_ = taker
				for i := MkInteger(0); IsTrue(OpLw(i, symbols.Length)); OpIncr(&i) {
					symbol := *(symbols).At(i)
					_ = symbol
					*(result).At(symbol) = MkMap(&VarMap{
						"info":   MkMap(&VarMap{"feeTier": feeTier}),
						"symbol": symbol,
						"maker":  maker,
						"taker":  taker,
					})
				}
				return result
			}
		}
	}
	return MkUndefined()
}

func (this *Binance) FuturesTransfer(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	amount := GoGetArg(goArgs, 1, MkUndefined())
	_ = amount
	type_ := GoGetArg(goArgs, 2, MkUndefined())
	_ = type_
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	if IsTrue(OpOr(OpLw(type_, MkInteger(1)), OpGt(type_, MkInteger(4)))) {
		panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" type must be between 1 and 4"))))
	}
	this.LoadMarkets()
	currency := this.Currency(code)
	_ = currency
	request := MkMap(&VarMap{
		"asset":  *(currency).At(MkString("id")),
		"amount": amount,
		"type":   type_,
	})
	_ = request
	response := this.Call(MkString("sapiPostFuturesTransfer"), this.Extend(request, params))
	_ = response
	return this.ParseTransfer(response, currency)
}

func (this *Binance) FetchFundingRate(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"symbol": *(market).At(MkString("id"))})
	_ = request
	method := OpCopy(MkUndefined())
	_ = method
	if IsTrue(*(market).At(MkString("linear"))) {
		method = MkString("fapiPublicGetPremiumIndex")
	} else {
		if IsTrue(*(market).At(MkString("inverse"))) {
			method = MkString("dapiPublicGetPremiumIndex")
		} else {
			panic(NewNotSupported(OpAdd(*this.At(MkString("id")), MkString(" setMarginMode() supports linear and inverse contracts only"))))
		}
	}
	response := this.Call(method, this.Extend(request, params))
	_ = response
	return this.ParseFundingRate(response)
}

func (this *Binance) FetchFundingRates(goArgs ...*Variant) *Variant {
	symbols := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbols
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	method := OpCopy(MkUndefined())
	_ = method
	defaultType := this.SafeString2(*this.At(MkString("options")), MkString("fetchFundingRates"), MkString("defaultType"), MkString("future"))
	_ = defaultType
	type_ := this.SafeString(params, MkString("type"), defaultType)
	_ = type_
	query := this.Omit(params, MkString("type"))
	_ = query
	if IsTrue(OpEq2(type_, MkString("future"))) {
		method = MkString("fapiPublicGetPremiumIndex")
	} else {
		if IsTrue(OpEq2(type_, MkString("delivery"))) {
			method = MkString("dapiPublicGetPremiumIndex")
		} else {
			panic(NewNotSupported(OpAdd(*this.At(MkString("id")), MkString(" setMarginMode() supports linear and inverse contracts only"))))
		}
	}
	response := this.Call(method, query)
	_ = response
	result := MkArray(&VarArray{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, response.Length)); OpIncr(&i) {
		entry := *(response).At(i)
		_ = entry
		parsed := this.ParseFundingRate(entry)
		_ = parsed
		result.Push(parsed)
	}
	return this.FilterByArray(result, MkString("symbol"), symbols)
}

func (this *Binance) ParseFundingRate(goArgs ...*Variant) *Variant {
	premiumIndex := GoGetArg(goArgs, 0, MkUndefined())
	_ = premiumIndex
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	timestamp := this.SafeInteger(premiumIndex, MkString("time"))
	_ = timestamp
	marketId := this.SafeString(premiumIndex, MkString("symbol"))
	_ = marketId
	symbol := this.SafeSymbol(marketId, market)
	_ = symbol
	markPrice := this.SafeNumber(premiumIndex, MkString("markPrice"))
	_ = markPrice
	indexPrice := this.SafeNumber(premiumIndex, MkString("indexPrice"))
	_ = indexPrice
	interestRate := this.SafeNumber(premiumIndex, MkString("interestRate"))
	_ = interestRate
	fundingRate := this.SafeNumber(premiumIndex, MkString("lastFundingRate"))
	_ = fundingRate
	nextFundingTime := this.SafeInteger(premiumIndex, MkString("nextFundingTime"))
	_ = nextFundingTime
	return MkMap(&VarMap{
		"info":                 premiumIndex,
		"symbol":               symbol,
		"markPrice":            markPrice,
		"indexPrice":           indexPrice,
		"interestRate":         interestRate,
		"timestamp":            timestamp,
		"datetime":             this.Iso8601(timestamp),
		"fundingRate":          fundingRate,
		"nextFundingTimestamp": nextFundingTime,
		"nextFundingDatetime":  this.Iso8601(nextFundingTime),
	})
}

func (this *Binance) ParseAccountPositions(goArgs ...*Variant) *Variant {
	account := GoGetArg(goArgs, 0, MkUndefined())
	_ = account
	positions := this.SafeValue(account, MkString("positions"))
	_ = positions
	assets := this.SafeValue(account, MkString("assets"))
	_ = assets
	balances := MkMap(&VarMap{})
	_ = balances
	for i := MkInteger(0); IsTrue(OpLw(i, assets.Length)); OpIncr(&i) {
		entry := *(assets).At(i)
		_ = entry
		currencyId := this.SafeString(entry, MkString("asset"))
		_ = currencyId
		code := this.SafeCurrencyCode(currencyId)
		_ = code
		crossWalletBalance := this.SafeString(entry, MkString("crossWalletBalance"))
		_ = crossWalletBalance
		crossUnPnl := this.SafeString(entry, MkString("crossUnPnl"))
		_ = crossUnPnl
		*(balances).At(code) = MkMap(&VarMap{
			"crossMargin":        Precise.StringAdd(crossWalletBalance, crossUnPnl),
			"crossWalletBalance": crossWalletBalance,
		})
	}
	result := MkArray(&VarArray{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, positions.Length)); OpIncr(&i) {
		position := *(positions).At(i)
		_ = position
		marketId := this.SafeString(position, MkString("symbol"))
		_ = marketId
		market := this.SafeMarket(marketId)
		_ = market
		code := OpTrinary(OpEq2(*(*this.At(MkString("options"))).At(MkString("defaultType")), MkString("future")), *(market).At(MkString("quote")), *(market).At(MkString("base")))
		_ = code
		parsed := this.ParsePosition(this.Extend(position, MkMap(&VarMap{
			"crossMargin":        *(*(balances).At(code)).At(MkString("crossMargin")),
			"crossWalletBalance": *(*(balances).At(code)).At(MkString("crossWalletBalance")),
		})), market)
		_ = parsed
		result.Push(parsed)
	}
	return result
}

func (this *Binance) ParsePosition(goArgs ...*Variant) *Variant {
	position := GoGetArg(goArgs, 0, MkUndefined())
	_ = position
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	marketId := this.SafeString(position, MkString("symbol"))
	_ = marketId
	market = this.SafeMarket(marketId, market)
	symbol := *(market).At(MkString("symbol"))
	_ = symbol
	leverageString := this.SafeString(position, MkString("leverage"))
	_ = leverageString
	leverage := parseInt(leverageString)
	_ = leverage
	initialMarginString := this.SafeString(position, MkString("initialMargin"))
	_ = initialMarginString
	initialMargin := this.ParseNumber(initialMarginString)
	_ = initialMargin
	initialMarginPercentageString := Precise.StringDiv(MkString("1"), leverageString, MkInteger(8))
	_ = initialMarginPercentageString
	rational := OpEq2(OpMod(MkInteger(1000), leverage), MkInteger(0))
	_ = rational
	if IsTrue(OpNot(rational)) {
		initialMarginPercentageString = Precise.StringDiv(Precise.StringAdd(initialMarginPercentageString, MkString("1e-8")), MkString("1"), MkInteger(8))
	}
	usdm := OpHasMember(MkString("notional"), position)
	_ = usdm
	maintenanceMarginString := this.SafeString(position, MkString("maintMargin"))
	_ = maintenanceMarginString
	maintenanceMargin := this.ParseNumber(maintenanceMarginString)
	_ = maintenanceMargin
	entryPriceString := this.SafeString(position, MkString("entryPrice"))
	_ = entryPriceString
	entryPrice := this.ParseNumber(entryPriceString)
	_ = entryPrice
	notionalString := this.SafeString2(position, MkString("notional"), MkString("notionalValue"))
	_ = notionalString
	notionalStringAbs := Precise.StringAbs(notionalString)
	_ = notionalStringAbs
	notionalFloat := parseFloat(notionalString)
	_ = notionalFloat
	notionalFloatAbs := parseFloat(notionalStringAbs)
	_ = notionalFloatAbs
	notional := this.ParseNumber(Precise.StringAbs(notionalString))
	_ = notional
	contractsString := this.SafeString(position, MkString("positionAmt"))
	_ = contractsString
	contractsStringAbs := Precise.StringAbs(contractsString)
	_ = contractsStringAbs
	if IsTrue(OpEq2(contractsString, MkUndefined())) {
		entryNotional := Precise.StringMul(Precise.StringMul(leverageString, initialMarginString), entryPriceString)
		_ = entryNotional
		contractsString = Precise.StringDiv(entryNotional, *(market).At(MkString("contractSize")))
		contractsStringAbs = Precise.StringDiv(Precise.StringAdd(contractsString, MkString("0.5")), MkString("1"), MkInteger(0))
	}
	contracts := this.ParseNumber(contractsStringAbs)
	_ = contracts
	leverageBrackets := this.SafeValue(*this.At(MkString("options")), MkString("leverageBrackets"), MkMap(&VarMap{}))
	_ = leverageBrackets
	leverageBracket := this.SafeValue(leverageBrackets, symbol, MkArray(&VarArray{}))
	_ = leverageBracket
	maintenanceMarginPercentageString := OpCopy(MkUndefined())
	_ = maintenanceMarginPercentageString
	for i := MkInteger(0); IsTrue(OpLw(i, leverageBracket.Length)); OpIncr(&i) {
		bracket := *(leverageBracket).At(i)
		_ = bracket
		if IsTrue(OpLw(notionalFloatAbs, *(bracket).At(MkInteger(0)))) {
			break
		}
		maintenanceMarginPercentageString = *(bracket).At(MkInteger(1))
	}
	maintenanceMarginPercentage := this.ParseNumber(maintenanceMarginPercentageString)
	_ = maintenanceMarginPercentage
	unrealizedPnlString := this.SafeString(position, MkString("unrealizedProfit"))
	_ = unrealizedPnlString
	unrealizedPnl := this.ParseNumber(unrealizedPnlString)
	_ = unrealizedPnl
	timestamp := this.SafeInteger(position, MkString("updateTime"))
	_ = timestamp
	if IsTrue(OpEq2(timestamp, MkInteger(0))) {
		timestamp = OpCopy(MkUndefined())
	}
	isolated := this.SafeValue(position, MkString("isolated"))
	_ = isolated
	marginType := OpCopy(MkUndefined())
	_ = marginType
	collateralString := OpCopy(MkUndefined())
	_ = collateralString
	walletBalance := OpCopy(MkUndefined())
	_ = walletBalance
	if IsTrue(isolated) {
		marginType = MkString("isolated")
		walletBalance = this.SafeString(position, MkString("isolatedWallet"))
		collateralString = Precise.StringAdd(walletBalance, unrealizedPnlString)
	} else {
		marginType = MkString("cross")
		walletBalance = this.SafeString(position, MkString("crossWalletBalance"))
		collateralString = this.SafeString(position, MkString("crossMargin"))
	}
	collateral := this.ParseNumber(collateralString)
	_ = collateral
	marginRatio := OpCopy(MkUndefined())
	_ = marginRatio
	side := OpCopy(MkUndefined())
	_ = side
	percentage := OpCopy(MkUndefined())
	_ = percentage
	liquidationPriceStringRaw := OpCopy(MkUndefined())
	_ = liquidationPriceStringRaw
	liquidationPrice := OpCopy(MkUndefined())
	_ = liquidationPrice
	if IsTrue(OpEq2(notionalFloat, MkNumber(0.0))) {
		entryPrice = OpCopy(MkUndefined())
	} else {
		side = OpTrinary(OpLw(notionalFloat, MkInteger(0)), MkString("short"), MkString("long"))
		marginRatio = this.ParseNumber(Precise.StringDiv(maintenanceMarginString, collateralString, MkInteger(4)))
		percentage = this.ParseNumber(Precise.StringMul(Precise.StringDiv(unrealizedPnlString, initialMarginString, MkInteger(4)), MkString("100")))
		if IsTrue(usdm) {
			onePlusMaintenanceMarginPercentageString := OpCopy(MkUndefined())
			_ = onePlusMaintenanceMarginPercentageString
			entryPriceSignString := OpCopy(entryPriceString)
			_ = entryPriceSignString
			if IsTrue(OpEq2(side, MkString("short"))) {
				onePlusMaintenanceMarginPercentageString = Precise.StringAdd(MkString("1"), maintenanceMarginPercentageString)
			} else {
				onePlusMaintenanceMarginPercentageString = Precise.StringAdd(MkString("-1"), maintenanceMarginPercentageString)
				entryPriceSignString = Precise.StringMul(MkString("-1"), entryPriceSignString)
			}
			leftSide := Precise.StringDiv(walletBalance, Precise.StringMul(contractsStringAbs, onePlusMaintenanceMarginPercentageString))
			_ = leftSide
			rightSide := Precise.StringDiv(entryPriceSignString, onePlusMaintenanceMarginPercentageString)
			_ = rightSide
			liquidationPriceStringRaw = Precise.StringAdd(leftSide, rightSide)
		} else {
			onePlusMaintenanceMarginPercentageString := OpCopy(MkUndefined())
			_ = onePlusMaintenanceMarginPercentageString
			entryPriceSignString := OpCopy(entryPriceString)
			_ = entryPriceSignString
			if IsTrue(OpEq2(side, MkString("short"))) {
				onePlusMaintenanceMarginPercentageString = Precise.StringSub(MkString("1"), maintenanceMarginPercentageString)
			} else {
				onePlusMaintenanceMarginPercentageString = Precise.StringSub(MkString("-1"), maintenanceMarginPercentageString)
				entryPriceSignString = Precise.StringMul(MkString("-1"), entryPriceSignString)
			}
			size := Precise.StringMul(contractsStringAbs, *(market).At(MkString("contractSize")))
			_ = size
			leftSide := Precise.StringMul(size, onePlusMaintenanceMarginPercentageString)
			_ = leftSide
			rightSide := Precise.StringSub(Precise.StringMul(Precise.StringDiv(MkString("1"), entryPriceSignString), size), walletBalance)
			_ = rightSide
			liquidationPriceStringRaw = Precise.StringDiv(leftSide, rightSide)
		}
		pricePrecision := *(*(market).At(MkString("precision"))).At(MkString("price"))
		_ = pricePrecision
		pricePrecisionPlusOne := OpAdd(pricePrecision, MkInteger(1))
		_ = pricePrecisionPlusOne
		pricePrecisionPlusOneString := pricePrecisionPlusOne.ToString()
		_ = pricePrecisionPlusOneString
		rounder := NewPrecise(OpAdd(MkString("5e-"), pricePrecisionPlusOneString))
		_ = rounder
		rounderString := rounder.ToString()
		_ = rounderString
		liquidationPriceRoundedString := Precise.StringAdd(rounderString, liquidationPriceStringRaw)
		_ = liquidationPriceRoundedString
		truncatedLiquidationPrice := Precise.StringDiv(liquidationPriceRoundedString, MkString("1"), pricePrecision)
		_ = truncatedLiquidationPrice
		if IsTrue(OpEq2(*(truncatedLiquidationPrice).At(MkInteger(0)), MkString("-"))) {
			truncatedLiquidationPrice = OpCopy(MkUndefined())
		}
		liquidationPrice = this.ParseNumber(truncatedLiquidationPrice)
	}
	return MkMap(&VarMap{
		"info":                        position,
		"symbol":                      symbol,
		"timestamp":                   timestamp,
		"datetime":                    this.Iso8601(timestamp),
		"initialMargin":               initialMargin,
		"initialMarginPercentage":     this.ParseNumber(initialMarginPercentageString),
		"maintenanceMargin":           maintenanceMargin,
		"maintenanceMarginPercentage": maintenanceMarginPercentage,
		"entryPrice":                  entryPrice,
		"notional":                    notional,
		"leverage":                    leverage,
		"unrealizedPnl":               unrealizedPnl,
		"contracts":                   contracts,
		"contractSize":                this.ParseNumber(*(market).At(MkString("contractSize"))),
		"marginRatio":                 marginRatio,
		"liquidationPrice":            liquidationPrice,
		"markPrice":                   MkUndefined(),
		"collateral":                  collateral,
		"marginType":                  marginType,
		"side":                        side,
		"percentage":                  percentage,
	})
}

func (this *Binance) ParsePositionRisk(goArgs ...*Variant) *Variant {
	position := GoGetArg(goArgs, 0, MkUndefined())
	_ = position
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	marketId := this.SafeString(position, MkString("symbol"))
	_ = marketId
	market = this.SafeMarket(marketId, market)
	symbol := *(market).At(MkString("symbol"))
	_ = symbol
	leverageBrackets := this.SafeValue(*this.At(MkString("options")), MkString("leverageBrackets"), MkMap(&VarMap{}))
	_ = leverageBrackets
	leverageBracket := this.SafeValue(leverageBrackets, symbol, MkArray(&VarArray{}))
	_ = leverageBracket
	notionalString := this.SafeString2(position, MkString("notional"), MkString("notionalValue"))
	_ = notionalString
	notionalStringAbs := Precise.StringAbs(notionalString)
	_ = notionalStringAbs
	notionalFloatAbs := parseFloat(notionalStringAbs)
	_ = notionalFloatAbs
	notionalFloat := parseFloat(notionalString)
	_ = notionalFloat
	maintenanceMarginPercentageString := OpCopy(MkUndefined())
	_ = maintenanceMarginPercentageString
	for i := MkInteger(0); IsTrue(OpLw(i, leverageBracket.Length)); OpIncr(&i) {
		bracket := *(leverageBracket).At(i)
		_ = bracket
		if IsTrue(OpLw(notionalFloatAbs, *(bracket).At(MkInteger(0)))) {
			break
		}
		maintenanceMarginPercentageString = *(bracket).At(MkInteger(1))
	}
	notional := this.ParseNumber(notionalStringAbs)
	_ = notional
	contractsAbs := Precise.StringAbs(this.SafeString(position, MkString("positionAmt")))
	_ = contractsAbs
	contracts := this.ParseNumber(contractsAbs)
	_ = contracts
	unrealizedPnlString := this.SafeString(position, MkString("unRealizedProfit"))
	_ = unrealizedPnlString
	unrealizedPnl := this.ParseNumber(unrealizedPnlString)
	_ = unrealizedPnl
	leverageString := this.SafeString(position, MkString("leverage"))
	_ = leverageString
	leverage := parseInt(leverageString)
	_ = leverage
	liquidationPrice := this.SafeNumber(position, MkString("liquidationPrice"))
	_ = liquidationPrice
	collateralString := this.SafeString(position, MkString("isolatedMargin"))
	_ = collateralString
	collateralFloat := parseFloat(collateralString)
	_ = collateralFloat
	collateral := this.ParseNumber(collateralString)
	_ = collateral
	markPriceString := this.SafeString(position, MkString("markPrice"))
	_ = markPriceString
	markPriceFloat := parseFloat(markPriceString)
	_ = markPriceFloat
	markPrice := OpCopy(MkUndefined())
	_ = markPrice
	if IsTrue(OpNotEq2(markPriceFloat, MkNumber(0.0))) {
		markPrice = this.ParseNumber(markPriceString)
	}
	entryPrice := this.SafeNumber(position, MkString("entryPrice"))
	_ = entryPrice
	timestamp := this.SafeInteger(position, MkString("updateTime"))
	_ = timestamp
	maintenanceMarginPercentage := this.ParseNumber(maintenanceMarginPercentageString)
	_ = maintenanceMarginPercentage
	maintenanceMarginString := Precise.StringMul(maintenanceMarginPercentageString, notionalStringAbs)
	_ = maintenanceMarginString
	maintenanceMargin := this.ParseNumber(maintenanceMarginString)
	_ = maintenanceMargin
	initialMarginPercentageString := Precise.StringDiv(MkString("1"), leverageString, MkInteger(8))
	_ = initialMarginPercentageString
	rational := OpEq2(OpMod(MkInteger(1000), leverage), MkInteger(0))
	_ = rational
	if IsTrue(OpNot(rational)) {
		initialMarginPercentageString = Precise.StringAdd(initialMarginPercentageString, MkString("1e-8"))
	}
	initialMarginString := Precise.StringDiv(Precise.StringMul(notionalStringAbs, initialMarginPercentageString), MkString("1"), MkInteger(8))
	_ = initialMarginString
	initialMargin := this.ParseNumber(initialMarginString)
	_ = initialMargin
	marginRatio := OpCopy(MkUndefined())
	_ = marginRatio
	side := OpCopy(MkUndefined())
	_ = side
	percentage := OpCopy(MkUndefined())
	_ = percentage
	if IsTrue(OpEq2(collateralFloat, MkNumber(0.0))) {
		liquidationPrice = OpCopy(MkUndefined())
	} else {
		marginRatio = this.ParseNumber(Precise.StringDiv(maintenanceMarginString, collateralString, MkInteger(4)))
		side = OpTrinary(OpLw(notionalFloat, MkInteger(0)), MkString("short"), MkString("long"))
		percentage = this.ParseNumber(Precise.StringMul(Precise.StringDiv(unrealizedPnlString, initialMarginString, MkInteger(4)), MkString("100")))
	}
	marginType := this.SafeString(position, MkString("marginType"))
	_ = marginType
	if IsTrue(OpEq2(marginType, MkString("cross"))) {
		liquidationPrice = OpCopy(MkUndefined())
	}
	return MkMap(&VarMap{
		"info":                        position,
		"symbol":                      symbol,
		"contracts":                   contracts,
		"unrealizedPnl":               unrealizedPnl,
		"leverage":                    leverage,
		"liquidationPrice":            liquidationPrice,
		"collateral":                  collateral,
		"notional":                    notional,
		"markPrice":                   markPrice,
		"entryPrice":                  entryPrice,
		"timestamp":                   timestamp,
		"initialMargin":               initialMargin,
		"initialMarginPercentage":     this.ParseNumber(initialMarginPercentageString),
		"maintenanceMargin":           maintenanceMargin,
		"maintenanceMarginPercentage": maintenanceMarginPercentage,
		"marginRatio":                 marginRatio,
		"datetime":                    this.Iso8601(timestamp),
		"marginType":                  marginType,
		"side":                        side,
		"percentage":                  percentage,
	})
}

func (this *Binance) LoadLeverageBrackets(goArgs ...*Variant) *Variant {
	reload := GoGetArg(goArgs, 0, MkBool(false))
	_ = reload
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	leverageBrackets := this.SafeValue(*this.At(MkString("options")), MkString("leverageBrackets"))
	_ = leverageBrackets
	if IsTrue(OpOr(OpEq2(leverageBrackets, MkUndefined()), reload)) {
		method := OpCopy(MkUndefined())
		_ = method
		defaultType := this.SafeString2(*this.At(MkString("options")), MkString("fetchPositions"), MkString("defaultType"), MkString("future"))
		_ = defaultType
		type_ := this.SafeString(params, MkString("type"), defaultType)
		_ = type_
		query := this.Omit(params, MkString("type"))
		_ = query
		if IsTrue(OpEq2(type_, MkString("future"))) {
			method = MkString("fapiPrivateGetLeverageBracket")
		} else {
			if IsTrue(OpEq2(type_, MkString("delivery"))) {
				method = MkString("dapiPrivateV2GetLeverageBracket")
			} else {
				panic(NewNotSupported(OpAdd(*this.At(MkString("id")), MkString(" loadLeverageBrackets() supports linear and inverse contracts only"))))
			}
		}
		response := this.Call(method, query)
		_ = response
		*(*this.At(MkString("options"))).At(MkString("leverageBrackets")) = MkMap(&VarMap{})
		for i := MkInteger(0); IsTrue(OpLw(i, response.Length)); OpIncr(&i) {
			entry := *(response).At(i)
			_ = entry
			marketId := this.SafeString(entry, MkString("symbol"))
			_ = marketId
			symbol := this.SafeSymbol(marketId)
			_ = symbol
			brackets := this.SafeValue(entry, MkString("brackets"))
			_ = brackets
			result := MkArray(&VarArray{})
			_ = result
			for j := MkInteger(0); IsTrue(OpLw(j, brackets.Length)); OpIncr(&j) {
				bracket := *(brackets).At(j)
				_ = bracket
				floorValue := this.SafeFloat2(bracket, MkString("notionalFloor"), MkString("qtyFloor"))
				_ = floorValue
				maintenanceMarginPercentage := this.SafeString(bracket, MkString("maintMarginRatio"))
				_ = maintenanceMarginPercentage
				result.Push(MkArray(&VarArray{
					floorValue,
					maintenanceMarginPercentage,
				}))
			}
			*(*(*this.At(MkString("options"))).At(MkString("leverageBrackets"))).At(symbol) = OpCopy(result)
		}
	}
	return *(*this.At(MkString("options"))).At(MkString("leverageBrackets"))
}

func (this *Binance) FetchPositions(goArgs ...*Variant) *Variant {
	symbols := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbols
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	this.LoadLeverageBrackets()
	method := OpCopy(MkUndefined())
	_ = method
	defaultType := this.SafeString2(*this.At(MkString("options")), MkString("fetchPositions"), MkString("defaultType"), MkString("future"))
	_ = defaultType
	type_ := this.SafeString(params, MkString("type"), defaultType)
	_ = type_
	query := this.Omit(params, MkString("type"))
	_ = query
	if IsTrue(OpEq2(type_, MkString("future"))) {
		method = MkString("fapiPrivateGetAccount")
	} else {
		if IsTrue(OpEq2(type_, MkString("delivery"))) {
			method = MkString("dapiPrivateGetAccount")
		} else {
			panic(NewNotSupported(OpAdd(*this.At(MkString("id")), MkString(" fetchPositions() supports linear and inverse contracts only"))))
		}
	}
	account := this.Call(method, query)
	_ = account
	result := this.ParseAccountPositions(account)
	_ = result
	return this.FilterByArray(result, MkString("symbol"), symbols, MkBool(false))
}

func (this *Binance) FetchIsolatedPositions(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	this.LoadLeverageBrackets()
	request := MkMap(&VarMap{})
	_ = request
	market := OpCopy(MkUndefined())
	_ = market
	method := OpCopy(MkUndefined())
	_ = method
	defaultType := MkString("future")
	_ = defaultType
	if IsTrue(OpNotEq2(symbol, MkUndefined())) {
		market = this.Market(symbol)
		*(request).At(MkString("symbol")) = *(market).At(MkString("id"))
		if IsTrue(*(market).At(MkString("linear"))) {
			defaultType = MkString("future")
		} else {
			if IsTrue(*(market).At(MkString("inverse"))) {
				defaultType = MkString("delivery")
			} else {
				panic(NewNotSupported(OpAdd(*this.At(MkString("id")), MkString(" fetchIsolatedPositions() supports linear and inverse contracts only"))))
			}
		}
	}
	defaultType = this.SafeString2(*this.At(MkString("options")), MkString("fetchIsolatedPositions"), MkString("defaultType"), defaultType)
	type_ := this.SafeString(params, MkString("type"), defaultType)
	_ = type_
	params = this.Omit(params, MkString("type"))
	if IsTrue(OpOr(OpEq2(type_, MkString("future")), OpEq2(type_, MkString("linear")))) {
		method = MkString("fapiPrivateGetPositionRisk")
	} else {
		if IsTrue(OpOr(OpEq2(type_, MkString("delivery")), OpEq2(type_, MkString("inverse")))) {
			method = MkString("dapiPrivateGetPositionRisk")
		} else {
			panic(NewNotSupported(OpAdd(*this.At(MkString("id")), MkString(" fetchIsolatedPositions() supports linear and inverse contracts only"))))
		}
	}
	response := this.Call(method, this.Extend(request, params))
	_ = response
	if IsTrue(OpEq2(symbol, MkUndefined())) {
		result := MkArray(&VarArray{})
		_ = result
		for i := MkInteger(0); IsTrue(OpLw(i, response.Length)); OpIncr(&i) {
			parsed := this.ParsePositionRisk(*(response).At(i), market)
			_ = parsed
			if IsTrue(OpEq2(*(parsed).At(MkString("marginType")), MkString("isolated"))) {
				result.Push(parsed)
			}
		}
		return result
	} else {
		return this.ParsePositionRisk(this.SafeValue(response, MkInteger(0)), market)
	}
	return MkUndefined()
}

func (this *Binance) FetchFundingHistory(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := OpCopy(MkUndefined())
	_ = market
	method := OpCopy(MkUndefined())
	_ = method
	defaultType := MkString("future")
	_ = defaultType
	request := MkMap(&VarMap{"incomeType": MkString("FUNDING_FEE")})
	_ = request
	if IsTrue(OpNotEq2(symbol, MkUndefined())) {
		market = this.Market(symbol)
		*(request).At(MkString("symbol")) = *(market).At(MkString("id"))
		if IsTrue(*(market).At(MkString("linear"))) {
			defaultType = MkString("future")
		} else {
			if IsTrue(*(market).At(MkString("inverse"))) {
				defaultType = MkString("delivery")
			} else {
				panic(NewNotSupported(OpAdd(*this.At(MkString("id")), MkString(" fetchFundingHistory() supports linear and inverse contracts only"))))
			}
		}
	}
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("startTime")) = OpCopy(since)
	}
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("limit")) = OpCopy(limit)
	}
	defaultType = this.SafeString2(*this.At(MkString("options")), MkString("fetchFundingHistory"), MkString("defaultType"), defaultType)
	type_ := this.SafeString(params, MkString("type"), defaultType)
	_ = type_
	params = this.Omit(params, MkString("type"))
	if IsTrue(OpOr(OpEq2(type_, MkString("future")), OpEq2(type_, MkString("linear")))) {
		method = MkString("fapiPrivateGetIncome")
	} else {
		if IsTrue(OpOr(OpEq2(type_, MkString("delivery")), OpEq2(type_, MkString("inverse")))) {
			method = MkString("dapiPrivateGetIncome")
		} else {
			panic(NewNotSupported(OpAdd(*this.At(MkString("id")), MkString(" fetchFundingHistory() supports linear and inverse contracts only"))))
		}
	}
	response := this.Call(method, this.Extend(request, params))
	_ = response
	return this.ParseIncomes(response, market, since, limit)
}

func (this *Binance) SetLeverage(goArgs ...*Variant) *Variant {
	leverage := GoGetArg(goArgs, 0, MkUndefined())
	_ = leverage
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	if IsTrue(OpEq2(symbol, MkUndefined())) {
		panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" setLeverage() requires a symbol argument"))))
	}
	if IsTrue(OpOr(OpLw(leverage, MkInteger(1)), OpGt(leverage, MkInteger(125)))) {
		panic(NewBadRequest(OpAdd(*this.At(MkString("id")), MkString(" leverage should be between 1 and 125"))))
	}
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	method := OpCopy(MkUndefined())
	_ = method
	if IsTrue(*(market).At(MkString("linear"))) {
		method = MkString("fapiPrivatePostLeverage")
	} else {
		if IsTrue(*(market).At(MkString("inverse"))) {
			method = MkString("dapiPrivatePostLeverage")
		} else {
			panic(NewNotSupported(OpAdd(*this.At(MkString("id")), MkString(" setLeverage() supports linear and inverse contracts only"))))
		}
	}
	request := MkMap(&VarMap{
		"symbol":   *(market).At(MkString("id")),
		"leverage": leverage,
	})
	_ = request
	return this.Call(method, this.Extend(request, params))
}

func (this *Binance) SetMarginMode(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	marginType := GoGetArg(goArgs, 1, MkUndefined())
	_ = marginType
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	marginType = marginType.ToUpperCase()
	if IsTrue(OpAnd(OpNotEq2(marginType, MkString("ISOLATED")), OpNotEq2(marginType, MkString("CROSSED")))) {
		panic(NewBadRequest(OpAdd(*this.At(MkString("id")), MkString(" marginType must be either isolated or crossed"))))
	}
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	method := OpCopy(MkUndefined())
	_ = method
	if IsTrue(*(market).At(MkString("linear"))) {
		method = MkString("fapiPrivatePostMarginType")
	} else {
		if IsTrue(*(market).At(MkString("inverse"))) {
			method = MkString("dapiPrivatePostMarginType")
		} else {
			panic(NewNotSupported(OpAdd(*this.At(MkString("id")), MkString(" setMarginMode() supports linear and inverse contracts only"))))
		}
	}
	request := MkMap(&VarMap{
		"symbol":     *(market).At(MkString("id")),
		"marginType": marginType,
	})
	_ = request
	return this.Call(method, this.Extend(request, params))
}

func (this *Binance) Sign(goArgs ...*Variant) *Variant {
	path := GoGetArg(goArgs, 0, MkUndefined())
	_ = path
	api := GoGetArg(goArgs, 1, MkString("public"))
	_ = api
	method := GoGetArg(goArgs, 2, MkString("GET"))
	_ = method
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	headers := GoGetArg(goArgs, 4, MkUndefined())
	_ = headers
	body := GoGetArg(goArgs, 5, MkUndefined())
	_ = body
	if IsTrue(OpNot(OpHasMember(api, *(*this.At(MkString("urls"))).At(MkString("api"))))) {
		panic(NewNotSupported(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" does not have a testnet/sandbox URL for "), OpAdd(api, MkString(" endpoints"))))))
	}
	url := *(*(*this.At(MkString("urls"))).At(MkString("api"))).At(api)
	_ = url
	OpAddAssign(&url, OpAdd(MkString("/"), path))
	if IsTrue(OpEq2(api, MkString("wapi"))) {
		OpAddAssign(&url, MkString(".html"))
	}
	if IsTrue(OpEq2(path, MkString("historicalTrades"))) {
		if IsTrue(*this.At(MkString("apiKey"))) {
			headers = MkMap(&VarMap{"X-MBX-APIKEY": *this.At(MkString("apiKey"))})
		} else {
			panic(NewAuthenticationError(OpAdd(*this.At(MkString("id")), MkString(" historicalTrades endpoint requires `apiKey` credential"))))
		}
	}
	userDataStream := OpOr(OpEq2(path, MkString("userDataStream")), OpEq2(path, MkString("listenKey")))
	_ = userDataStream
	if IsTrue(userDataStream) {
		if IsTrue(*this.At(MkString("apiKey"))) {
			headers = MkMap(&VarMap{
				"X-MBX-APIKEY": *this.At(MkString("apiKey")),
				"Content-Type": MkString("application/x-www-form-urlencoded"),
			})
			if IsTrue(OpNotEq2(method, MkString("GET"))) {
				body = this.Urlencode(params)
			}
		} else {
			panic(NewAuthenticationError(OpAdd(*this.At(MkString("id")), MkString(" userDataStream endpoint requires `apiKey` credential"))))
		}
	} else {
		if IsTrue(OpOr(OpEq2(api, MkString("private")), OpOr(OpEq2(api, MkString("sapi")), OpOr(OpAnd(OpEq2(api, MkString("wapi")), OpNotEq2(path, MkString("systemStatus"))), OpOr(OpEq2(api, MkString("dapiPrivate")), OpOr(OpEq2(api, MkString("dapiPrivateV2")), OpOr(OpEq2(api, MkString("fapiPrivate")), OpEq2(api, MkString("fapiPrivateV2"))))))))) {
			this.CheckRequiredCredentials()
			query := OpCopy(MkUndefined())
			_ = query
			recvWindow := this.SafeInteger(*this.At(MkString("options")), MkString("recvWindow"), MkInteger(5000))
			_ = recvWindow
			if IsTrue(OpAnd(OpEq2(api, MkString("sapi")), OpEq2(path, MkString("asset/dust")))) {
				query = this.UrlencodeWithArrayRepeat(this.Extend(MkMap(&VarMap{
					"timestamp":  this.Nonce(),
					"recvWindow": recvWindow,
				}), params))
			} else {
				if IsTrue(OpOr(OpEq2(path, MkString("batchOrders")), OpGtEq(path.IndexOf(MkString("sub-account")), MkInteger(0)))) {
					query = this.Rawencode(this.Extend(MkMap(&VarMap{
						"timestamp":  this.Nonce(),
						"recvWindow": recvWindow,
					}), params))
				} else {
					query = this.Urlencode(this.Extend(MkMap(&VarMap{
						"timestamp":  this.Nonce(),
						"recvWindow": recvWindow,
					}), params))
				}
			}
			signature := this.Hmac(this.Encode(query), this.Encode(*this.At(MkString("secret"))))
			_ = signature
			OpAddAssign(&query, OpAdd(MkString("&"), OpAdd(MkString("signature="), signature)))
			headers = MkMap(&VarMap{"X-MBX-APIKEY": *this.At(MkString("apiKey"))})
			if IsTrue(OpOr(OpEq2(method, MkString("GET")), OpOr(OpEq2(method, MkString("DELETE")), OpEq2(api, MkString("wapi"))))) {
				OpAddAssign(&url, OpAdd(MkString("?"), query))
			} else {
				body = OpCopy(query)
				*(headers).At(MkString("Content-Type")) = MkString("application/x-www-form-urlencoded")
			}
		} else {
			if IsTrue(*(GoGetKeys(params)).At(MkString("length"))) {
				OpAddAssign(&url, OpAdd(MkString("?"), this.Urlencode(params)))
			}
		}
	}
	return MkMap(&VarMap{
		"url":     url,
		"method":  method,
		"body":    body,
		"headers": headers,
	})
}

func (this *Binance) HandleErrors(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	reason := GoGetArg(goArgs, 1, MkUndefined())
	_ = reason
	url := GoGetArg(goArgs, 2, MkUndefined())
	_ = url
	method := GoGetArg(goArgs, 3, MkUndefined())
	_ = method
	headers := GoGetArg(goArgs, 4, MkUndefined())
	_ = headers
	body := GoGetArg(goArgs, 5, MkUndefined())
	_ = body
	response := GoGetArg(goArgs, 6, MkUndefined())
	_ = response
	requestHeaders := GoGetArg(goArgs, 7, MkUndefined())
	_ = requestHeaders
	requestBody := GoGetArg(goArgs, 8, MkUndefined())
	_ = requestBody
	if IsTrue(OpOr(OpEq2(code, MkInteger(418)), OpEq2(code, MkInteger(429)))) {
		panic(NewDDoSProtection(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), OpAdd(code.ToString(), OpAdd(MkString(" "), OpAdd(reason, OpAdd(MkString(" "), body))))))))
	}
	if IsTrue(OpGtEq(code, MkInteger(400))) {
		if IsTrue(OpGtEq(body.IndexOf(MkString("Price * QTY is zero or less")), MkInteger(0))) {
			panic(NewInvalidOrder(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" order cost = amount * price is zero or less "), body))))
		}
		if IsTrue(OpGtEq(body.IndexOf(MkString("LOT_SIZE")), MkInteger(0))) {
			panic(NewInvalidOrder(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" order amount should be evenly divisible by lot size "), body))))
		}
		if IsTrue(OpGtEq(body.IndexOf(MkString("PRICE_FILTER")), MkInteger(0))) {
			panic(NewInvalidOrder(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" order price is invalid, i.e. exceeds allowed price precision, exceeds min price or max price limits or is invalid float value in general, use this.priceToPrecision (symbol, amount) "), body))))
		}
	}
	if IsTrue(OpEq2(response, MkUndefined())) {
		return MkUndefined()
	}
	success := this.SafeValue(response, MkString("success"), MkBool(true))
	_ = success
	if IsTrue(OpNot(success)) {
		message := this.SafeString(response, MkString("msg"))
		_ = message
		parsedMessage := OpCopy(MkUndefined())
		_ = parsedMessage
		if IsTrue(OpNotEq2(message, MkUndefined())) {
			{
				ret__ := func(this *Binance) (ret_ *Variant) {
					defer func() {
						if e := recover().(*Variant); e != nil {
							ret_ = func(this *Binance) *Variant {
								// catch block:
								parsedMessage = OpCopy(MkUndefined())
								return nil
							}(this)
						}
					}()
					// try block:
					parsedMessage = JSON.Parse(message)
					return nil
				}(this)
				if ret__ != nil {
					return ret__
				}
			}
			if IsTrue(OpNotEq2(parsedMessage, MkUndefined())) {
				response = OpCopy(parsedMessage)
			}
		}
	}
	message := this.SafeString(response, MkString("msg"))
	_ = message
	if IsTrue(OpNotEq2(message, MkUndefined())) {
		this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions"))).At(MkString("exact")), message, OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), message)))
		this.ThrowBroadlyMatchedException(*(*this.At(MkString("exceptions"))).At(MkString("broad")), message, OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), message)))
	}
	error := this.SafeString(response, MkString("code"))
	_ = error
	if IsTrue(OpNotEq2(error, MkUndefined())) {
		if IsTrue(OpOr(OpEq2(error, MkString("200")), Precise.StringEquals(error, MkString("0")))) {
			return MkUndefined()
		}
		if IsTrue(OpAnd(OpEq2(error, MkString("-2015")), *(*this.At(MkString("options"))).At(MkString("hasAlreadyAuthenticatedSuccessfully")))) {
			panic(NewDDoSProtection(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" temporary banned: "), body))))
		}
		feedback := OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), body))
		_ = feedback
		this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions"))).At(MkString("exact")), error, feedback)
		panic(NewExchangeError(feedback))
	}
	if IsTrue(OpNot(success)) {
		panic(NewExchangeError(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), body))))
	}
	return MkUndefined()
}

func (this *Binance) CalculateRateLimiterCost(goArgs ...*Variant) *Variant {
	api := GoGetArg(goArgs, 0, MkUndefined())
	_ = api
	method := GoGetArg(goArgs, 1, MkUndefined())
	_ = method
	path := GoGetArg(goArgs, 2, MkUndefined())
	_ = path
	params := GoGetArg(goArgs, 3, MkUndefined())
	_ = params
	config := GoGetArg(goArgs, 4, MkMap(&VarMap{}))
	_ = config
	context := GoGetArg(goArgs, 5, MkMap(&VarMap{}))
	_ = context
	if IsTrue(OpAnd(OpHasMember(MkString("noSymbol"), config), OpNot(OpHasMember(MkString("symbol"), params)))) {
		return *(config).At(MkString("noSymbol"))
	} else {
		if IsTrue(OpAnd(OpHasMember(MkString("noPoolId"), config), OpNot(OpHasMember(MkString("poolId"), params)))) {
			return *(config).At(MkString("noPoolId"))
		} else {
			if IsTrue(OpAnd(OpHasMember(MkString("byLimit"), config), OpHasMember(MkString("limit"), params))) {
				limit := *(params).At(MkString("limit"))
				_ = limit
				byLimit := *(config).At(MkString("byLimit"))
				_ = byLimit
				for i := MkInteger(0); IsTrue(OpLw(i, byLimit.Length)); OpIncr(&i) {
					entry := *(byLimit).At(i)
					_ = entry
					if IsTrue(OpLwEq(limit, *(entry).At(MkInteger(0)))) {
						return *(entry).At(MkInteger(1))
					}
				}
			}
		}
	}
	return this.SafeInteger(config, MkString("cost"), MkInteger(1))
}

func (this *Binance) Request(goArgs ...*Variant) *Variant {
	path := GoGetArg(goArgs, 0, MkUndefined())
	_ = path
	api := GoGetArg(goArgs, 1, MkString("public"))
	_ = api
	method := GoGetArg(goArgs, 2, MkString("GET"))
	_ = method
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	headers := GoGetArg(goArgs, 4, MkUndefined())
	_ = headers
	body := GoGetArg(goArgs, 5, MkUndefined())
	_ = body
	config := GoGetArg(goArgs, 6, MkMap(&VarMap{}))
	_ = config
	context := GoGetArg(goArgs, 7, MkMap(&VarMap{}))
	_ = context
	response := this.Fetch2(path, api, method, params, headers, body, config, context)
	_ = response
	if IsTrue(OpOr(OpEq2(api, MkString("private")), OpEq2(api, MkString("wapi")))) {
		*(*this.At(MkString("options"))).At(MkString("hasAlreadyAuthenticatedSuccessfully")) = OpCopy(MkBool(true))
	}
	return response
}

func (this *Binance) ModifyMarginHelper(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	amount := GoGetArg(goArgs, 1, MkUndefined())
	_ = amount
	addOrReduce := GoGetArg(goArgs, 2, MkUndefined())
	_ = addOrReduce
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	defaultType := this.SafeString(*this.At(MkString("options")), MkString("defaultType"), MkString("future"))
	_ = defaultType
	if IsTrue(OpEq2(defaultType, MkString("spot"))) {
		defaultType = MkString("future")
	}
	type_ := this.SafeString(params, MkString("type"), defaultType)
	_ = type_
	if IsTrue(OpOr(OpEq2(type_, MkString("margin")), OpEq2(type_, MkString("spot")))) {
		panic(NewNotSupported(OpAdd(*this.At(MkString("id")), MkString(" add / reduce margin only supported with type future or delivery"))))
	}
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{
		"type":   addOrReduce,
		"symbol": *(market).At(MkString("id")),
		"amount": amount,
	})
	_ = request
	method := OpCopy(MkUndefined())
	_ = method
	code := OpCopy(MkUndefined())
	_ = code
	if IsTrue(OpEq2(type_, MkString("future"))) {
		method = MkString("fapiPrivatePostPositionMargin")
		code = *(market).At(MkString("quote"))
	} else {
		method = MkString("dapiPrivatePostPositionMargin")
		code = *(market).At(MkString("base"))
	}
	response := this.Call(method, this.Extend(request, params))
	_ = response
	rawType := this.SafeInteger(response, MkString("type"))
	_ = rawType
	resultType := OpTrinary(OpEq2(rawType, MkInteger(1)), MkString("add"), MkString("reduce"))
	_ = resultType
	resultAmount := this.SafeNumber(response, MkString("amount"))
	_ = resultAmount
	errorCode := this.SafeString(response, MkString("code"))
	_ = errorCode
	status := OpTrinary(OpEq2(errorCode, MkString("200")), MkString("ok"), MkString("failed"))
	_ = status
	return MkMap(&VarMap{
		"info":   response,
		"type":   resultType,
		"amount": resultAmount,
		"code":   code,
		"symbol": *(market).At(MkString("symbol")),
		"status": status,
	})
}

func (this *Binance) ReduceMargin(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	amount := GoGetArg(goArgs, 1, MkUndefined())
	_ = amount
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	return this.ModifyMarginHelper(symbol, amount, MkInteger(2), params)
}

func (this *Binance) AddMargin(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	amount := GoGetArg(goArgs, 1, MkUndefined())
	_ = amount
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	return this.ModifyMarginHelper(symbol, amount, MkInteger(1), params)
}
