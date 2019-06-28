# cards
--
    import "github.com/reddotpay/cards"


## Usage

```go
const (
	// BrandVisa defines Visa card brand
	BrandVisa = "visa"
	// BrandMasterCard defines MasterCard card brand
	BrandMasterCard = "mastercard"
	// BrandUnionPay defines UnionPay card brand
	BrandUnionPay = "unionpay"
	// BrandDiners defines Diners card brand
	BrandDiners = "diners"
	// BrandDiscover defines Discover card brand
	BrandDiscover = "discover"
	// BrandAmericanExpress defines American Express card brand
	BrandAmericanExpress = "amex"
	// BrandJCB defines JCB card brand
	BrandJCB = "jcb"
	// BrandOthers defines undefined brands
	BrandOthers = "others"
)
```

```go
var BrandCheck = map[string]func(string) bool{
	BrandVisa:            isVisa,
	BrandMasterCard:      isMasterCard,
	BrandUnionPay:        isUnionPay,
	BrandDiners:          isDiners,
	BrandDiscover:        isDiscover,
	BrandAmericanExpress: isAmericanExpress,
	BrandJCB:             isJCB,
}
```
BrandCheck defines a map of functions to check the various card brands

#### func  Brand

```go
func Brand(number string) string
```
Brand returns the corresponding card brand

#### func  ValidateNumber

```go
func ValidateNumber(number string) bool
```
ValidateNumber returns true if `number` satisfy the Luhn algorithm
