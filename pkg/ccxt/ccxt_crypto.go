package ccxt

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	b64 "encoding/base64"
	"fmt"
	"golang.org/x/crypto/sha3"
)

//
// this file implements cryptographic functions
//

func (this *ExchangeBase) Hash(goArgs ...*Variant) *Variant {
	request := GoGetArg(goArgs, 0, MkUndefined())
	hash := GoGetArg(goArgs, 1, MkString("md5"))
	digest := GoGetArg(goArgs, 2, MkString("hex"))
	options := MkMap(&VarMap{})
	if IsTrue(OpEq2(hash, MkString("keccak"))) {
		hash = MkString("SHA3")
		*(options).At(MkString("outputLength")) = MkInteger(256)
	}

	result := []byte{}
	if hash.ToStr() == "md5" {
		val := md5.Sum([]byte(request.ToStr()))
		result = val[:]
	} else if hash.ToStr() == "sha1" {
		val := sha1.Sum([]byte(request.ToStr()))
		result = val[:]
	} else if hash.ToStr() == "sha256" {
		val := sha256.Sum256([]byte(request.ToStr()))
		result = val[:]
	} else if hash.ToStr() == "sha512" {
		val := sha512.Sum512([]byte(request.ToStr()))
		result = val[:]
	} else {
		panic(fmt.Sprintf("Unknown hash algorithm: %s", hash.ToStr()))
	}

	if digest.ToStr() == "hex" {
		return MkString(fmt.Sprintf("%x", result))
	} else if digest.ToStr() == "base64" {
		return MkString(b64.StdEncoding.EncodeToString(result))
	} else {
		panic(fmt.Sprintf("Unknown digest format: %s", hash.ToStr()))
	}
}

func (this *ExchangeBase) Hmac(goArgs ...*Variant) *Variant {
	request := GoGetArg(goArgs, 0, MkUndefined())
	secret := GoGetArg(goArgs, 1, MkUndefined())
	hash := GoGetArg(goArgs, 2, MkString("sha256"))
	digest := GoGetArg(goArgs, 3, MkString("hex"))

	key := []byte(secret.ToStr()) // todo is that right does that need decoding?

	result := []byte{}
	if hash.ToStr() == "md5" {
		mac := hmac.New(md5.New, key)
		mac.Write([]byte(request.ToStr()))
		result = mac.Sum(nil)
	} else if hash.ToStr() == "sha1" {
		mac := hmac.New(sha1.New, key)
		mac.Write([]byte(request.ToStr()))
		result = mac.Sum(nil)
	} else if hash.ToStr() == "sha384" {
		mac := hmac.New(sha3.New384, key) // todo: is that right?
		mac.Write([]byte(request.ToStr()))
		result = mac.Sum(nil)
	} else if hash.ToStr() == "sha256" {
		mac := hmac.New(sha256.New, key)
		mac.Write([]byte(request.ToStr()))
		result = mac.Sum(nil)
	} else if hash.ToStr() == "sha512" {
		mac := hmac.New(sha512.New, key)
		mac.Write([]byte(request.ToStr()))
		result = mac.Sum(nil)
	} else {
		panic(fmt.Sprintf("Unknown hash algorithm: %s", hash.ToStr()))
	}

	if digest.ToStr() == "hex" {
		return MkString(fmt.Sprintf("%x", result))
	} else if digest.ToStr() == "base64" {
		return MkString(b64.StdEncoding.EncodeToString(result))
	} else {
		panic(fmt.Sprintf("Unknown digest format: %s", hash.ToStr()))
	}
}

func (this *ExchangeBase) HashMessage(goArgs ...*Variant) *Variant {
	// transpiled function
	message := GoGetArg(goArgs, 0, MkUndefined())
	binaryMessage := this.Base16ToBinary(this.Remove0xPrefix(message))
	prefix := this.StringToBinary(OpAdd(MkString("\x19Ethereum Signed Message:\n"), *binaryMessage.At(MkString("sigBytes"))))
	return OpAdd(MkString("0x"), this.Hash(this.BinaryConcat(prefix, binaryMessage), MkString("keccak"), MkString("hex")))
}

func (this *ExchangeBase) SignHash(goArgs ...*Variant) *Variant {
	// transpiled function
	hash := GoGetArg(goArgs, 0, MkUndefined())
	privateKey := GoGetArg(goArgs, 1, MkUndefined())
	signature := this.Ecdsa(hash.Slice(OpNeg(MkInteger(64))), privateKey.Slice(OpNeg(MkInteger(64))), MkString("secp256k1"), MkUndefined())
	return MkMap(&VarMap{
		"r": OpAdd(MkString("0x"), *(signature).At(MkString("r"))),
		"s": OpAdd(MkString("0x"), *(signature).At(MkString("s"))),
		"v": OpAdd(MkInteger(27), *(signature).At(MkString("v"))),
	})
}

func (this *ExchangeBase) SignMessage(goArgs ...*Variant) *Variant {
	// transpiled function
	message := GoGetArg(goArgs, 0, MkUndefined())
	privateKey := GoGetArg(goArgs, 1, MkUndefined())
	return this.SignHash(this.HashMessage(message), privateKey.Slice(OpNeg(MkInteger(64))))
}

func (this *ExchangeBase) SignMessageString(goArgs ...*Variant) *Variant {
	// transpiled function
	message := GoGetArg(goArgs, 0, MkUndefined())
	privateKey := GoGetArg(goArgs, 1, MkUndefined())
	signature := this.SignMessage(message, privateKey)
	return OpAdd(*(signature).At(MkString("r")), OpAdd(this.Remove0xPrefix(*(signature).At(MkString("s"))), this.BinaryToBase16(this.NumberToBE(*(signature).At(MkString("v"))))))
}

func (this *ExchangeBase) Eddsa(v ...*Variant) *Variant {
	panic("not implemented function") // todo implement
}

func (this *ExchangeBase) Ecdsa(v ...*Variant) *Variant {
	panic("not implemented function") // todo implement
}

func (this *ExchangeBase) Rsa(v ...*Variant) *Variant {
	panic("not implemented function") // todo implement
}

func (this *ExchangeBase) Jwt(v ...*Variant) *Variant {
	panic("not implemented function") // todo implement
}

func (this *ExchangeBase) Oath(v ...*Variant) *Variant {
	panic("not implemented function") // todo implement
}
