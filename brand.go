package cards

import (
	"strconv"
)

const (
	// BrandVisa defines Visa card brand
	BrandVisa = "visa"
	// BrandMasterCard defines MasterCard card brand
	BrandMasterCard = "mastercard"
	// BrandMaestro defines Maestro card brand
	BrandMaestro = "maestro"
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

// BrandCheck defines brand name and check function
type BrandCheck struct {
	Brand string
	Check func(string) bool
}

// BrandChecks defines an array of BrandCheck
var BrandChecks = []BrandCheck{
	{BrandVisa, isVisa},
	{BrandMasterCard, isMasterCard},
	{BrandUnionPay, isUnionPay},
	{BrandDiners, isDiners},
	{BrandDiscover, isDiscover},
	{BrandAmericanExpress, isAmericanExpress},
	{BrandJCB, isJCB},

	// Maestro must be after Discover
	{BrandMaestro, isMaestro},
}

// Brand returns the corresponding card brand
func Brand(number string) string {
	for _, brandCheck := range BrandChecks {
		if brandCheck.Check(number) {
			return brandCheck.Brand
		}
	}

	return BrandOthers
}

func isVisa(number string) bool {
	return number[:1] == "4" && len(number) <= 19
}

func isMasterCard(number string) bool {
	if len(number) != 16 {
		return false
	}

	firstDigit := string(number[:1])
	if firstDigit == "5" {
		bin, _ := strconv.Atoi(string(number[:6]))
		return (bin >= 510000) && (bin <= 559999)
	} else if firstDigit == "2" {
		bin, _ := strconv.Atoi(string(number[:6]))
		return (bin >= 222100) && (bin <= 272099)
	}

	return false
}

func isUnionPay(number string) bool {
	len := len(number)
	if len < 16 || len > 19 {
		return false
	}

	if string(number[:2]) == "81" {
		bin, _ := strconv.Atoi(string(number[:6]))
		return (bin >= 810000) && (bin <= 817199)
	}

	return string(number[:2]) == "62"
}

func isDiners(number string) bool {
	if len(number) != 14 {
		return false
	}

	bin, _ := strconv.Atoi(string(number[:6]))
	return (bin >= 300000 && bin <= 305999) ||
		(bin >= 309500 && bin <= 309599) ||
		(bin >= 360000 && bin <= 369999) ||
		(bin >= 380000 && bin <= 399999)
}

func isDiscover(number string) bool {
	if len(number) != 16 {
		return false
	}

	if string(number[:4]) == "6011" {
		return true
	}

	if string(number[:2]) == "64" || string(number[:2]) == "65" {
		return true
	}

	return false
}

func isAmericanExpress(number string) bool {
	if len(number) != 15 {
		return false
	}

	return (string(number[:2]) == "34") || (string(number[:2]) == "37")
}

func isJCB(number string) bool {
	digits, _ := strconv.Atoi(string(number[:4]))
	return (digits >= 3528) && (digits <= 3589)
}

func isMaestro(number string) bool {
	if len(number) > 19 {
		return false
	}

	digits, _ := strconv.Atoi(string(number[:2]))
	return (digits == 50) ||
		((digits >= 56) && (digits <= 64)) ||
		((digits >= 66) && (digits <= 69))
}
