package ccxt

///////////////////////////////////////////////////////////////////////////////////////////////
// exceptions

type VError struct {
	Type    string
	Message string
}

func NewVError(err string, v ...*Variant) *Variant {
	str := ""
	for i := 0; i < len(v); i++ {
		if i > 0 {
			str += " "
		}
		str += v[i].ToStr()
	}
	return &Variant{Type: Error, Value: VError{Type: err, Message: str}}
}

var ExchangeError *Variant = NewExchangeError()
var PermissionDenied *Variant = NewPermissionDenied()
var NullResponse *Variant = NewNullResponse()
var BadResponse *Variant = NewBadResponse()
var AuthenticationError *Variant = NewAuthenticationError()
var OrderImmediatelyFillable *Variant = NewOrderImmediatelyFillable()
var InsufficientFunds *Variant = NewInsufficientFunds()
var ExchangeNotAvailable *Variant = NewExchangeNotAvailable()
var DDoSProtection *Variant = NewDDoSProtection()
var RateLimitExceeded *Variant = NewRateLimitExceeded()
var BadRequest *Variant = NewBadRequest()
var InvalidNonce *Variant = NewInvalidNonce()
var NetworkError *Variant = NewNetworkError()
var OrderNotFound *Variant = NewOrderNotFound()
var AccountSuspended *Variant = NewAccountSuspended()
var BadSymbol *Variant = NewBadSymbol()
var InvalidOrder *Variant = NewInvalidOrder()
var DuplicateOrderId *Variant = NewDuplicateOrderId()
var NotSupported *Variant = NewNotSupported()
var RequestTimeout *Variant = NewRequestTimeout()
var OnMaintenance *Variant = NewOnMaintenance()
var CancelPending *Variant = NewCancelPending()
var ArgumentsRequired *Variant = NewArgumentsRequired()
var InvalidAddress *Variant = NewInvalidAddress()
var ExceptionClass *Variant = NewExceptionClass()
var AddressPending *Variant = NewAddressPending()

func NewError(v ...*Variant) *Variant {
	return NewVError("Generic", v...)
}
func NewErrorClass(v ...*Variant) *Variant {
	return NewVError("Generic", v...)
}
func NewException(v ...*Variant) *Variant {
	return NewVError("Generic", v...)
}
func NewBadResponse(v ...*Variant) *Variant {
	return NewVError("BadResponse", v...)
}
func NewExchangeError(v ...*Variant) *Variant {
	return NewVError("ExchangeError", v...)
}
func NewPermissionDenied(v ...*Variant) *Variant {
	return NewVError("PermissionDenied", v...)
}
func NewAuthenticationError(v ...*Variant) *Variant {
	return NewVError("AuthenticationError", v...)
}
func NewOrderImmediatelyFillable(v ...*Variant) *Variant {
	return NewVError("OrderImmediatelyFillable", v...)
}
func NewInsufficientFunds(v ...*Variant) *Variant {
	return NewVError("InsufficientFunds", v...)
}
func NewExchangeNotAvailable(v ...*Variant) *Variant {
	return NewVError("ExchangeNotAvailable", v...)
}
func NewDDoSProtection(v ...*Variant) *Variant {
	return NewVError("DDoSProtection", v...)
}
func NewRateLimitExceeded(v ...*Variant) *Variant {
	return NewVError("RateLimitExceeded", v...)
}
func NewBadRequest(v ...*Variant) *Variant {
	return NewVError("BadRequest", v...)
}
func NewInvalidNonce(v ...*Variant) *Variant {
	return NewVError("InvalidNonce", v...)
}
func NewOrderNotFound(v ...*Variant) *Variant {
	return NewVError("OrderNotFound", v...)
}
func NewAccountSuspended(v ...*Variant) *Variant {
	return NewVError("AccountSuspended", v...)
}
func NewBadSymbol(v ...*Variant) *Variant {
	return NewVError("BadSymbol", v...)
}
func NewDuplicateOrderId(v ...*Variant) *Variant {
	return NewVError("DuplicateOrderId", v...)
}
func NewInvalidOrder(v ...*Variant) *Variant {
	return NewVError("InvalidOrder", v...)
}
func NewNotSupported(v ...*Variant) *Variant {
	return NewVError("NotSupported", v...)
}
func NewNetworkError(v ...*Variant) *Variant {
	return NewVError("NetworkError", v...)
}
func NewRequestTimeout(v ...*Variant) *Variant {
	return NewVError("RequestTimeout", v...)
}
func NewOnMaintenance(v ...*Variant) *Variant {
	return NewVError("OnMaintenance", v...)
}
func NewCancelPending(v ...*Variant) *Variant {
	return NewVError("CancelPending", v...)
}
func NewArgumentsRequired(v ...*Variant) *Variant {
	return NewVError("ArgumentsRequired", v...)
}
func NewInvalidAddress(v ...*Variant) *Variant {
	return NewVError("InvalidAddress", v...)
}
func NewExceptionClass(v ...*Variant) *Variant {
	return NewVError("ExceptionClass", v...)
}
func NewAddressPending(v ...*Variant) *Variant {
	return NewVError("AddressPending", v...)
}
func NewNullResponse(v ...*Variant) *Variant {
	return NewVError("NullResponse", v...)
}

func (this *ExchangeBase) ThrowBroadlyMatchedException(exact *Variant, str *Variant, message *Variant) {

	err := exact.At(str)
	if err != nil && (*err).Type != Undefined {
		panic((*err).ToStr())
	}
}

func (this *ExchangeBase) ThrowExactlyMatchedException(broad *Variant, str *Variant, message *Variant) {
	err := this.FindBroadlyMatchedKey(broad, str)
	if err != nil && err.Type != Undefined {
		panic(err.ToStr())
	}
}
