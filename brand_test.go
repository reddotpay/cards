package cards_test

import (
	"testing"

	"github.com/reddotpay/cards"
	"github.com/stretchr/testify/assert"
)

func TestBrand_Brand_Visa(t *testing.T) {
	assert.Equal(t, cards.BrandVisa, cards.Brand("4111111111111111"))
	assert.Equal(t, cards.BrandVisa, cards.Brand("4444333322221111"))
	assert.Equal(t, cards.BrandVisa, cards.Brand("4911830000000"))
	assert.Equal(t, cards.BrandVisa, cards.Brand("4917610000000000"))
}

func TestBrand_Brand_MasterCard(t *testing.T) {
	assert.Equal(t, cards.BrandMasterCard, cards.Brand("5105105105105100"))
	assert.Equal(t, cards.BrandMasterCard, cards.Brand("2223000048400011"))
}

func TestBrand_Brand_JCB(t *testing.T) {
	assert.Equal(t, cards.BrandJCB, cards.Brand("3530111333300000"))
}

func TestBrand_Brand_AmericanExpress(t *testing.T) {
	assert.Equal(t, cards.BrandAmericanExpress, cards.Brand("371449635398431"))
}

func TestBrand_Brand_Diners(t *testing.T) {
	assert.Equal(t, cards.BrandDiners, cards.Brand("30569309025904"))
}

func TestBrand_Brand_Discover(t *testing.T) {
	assert.Equal(t, cards.BrandDiscover, cards.Brand("6011000990139424"))
	assert.Equal(t, cards.BrandDiscover, cards.Brand("6011000400000000"))
	assert.Equal(t, cards.BrandDiscover, cards.Brand("6011601160116611"))
	assert.Equal(t, cards.BrandDiscover, cards.Brand("6445644564456445"))
	assert.Equal(t, cards.BrandDiscover, cards.Brand("6011111111111117"))
}

func TestBrand_Brand_UnionPay(t *testing.T) {
	assert.Equal(t, cards.BrandUnionPay, cards.Brand("6293077338195127"))
	assert.Equal(t, cards.BrandUnionPay, cards.Brand("8171999927660000"))
}

func TestBrand_Brand_Maestro(t *testing.T) {
	assert.Equal(t, cards.BrandMaestro, cards.Brand("6759649826438453"))
	assert.Equal(t, cards.BrandMaestro, cards.Brand("6799990100000000019"))
}

func TestBrand_Brand_Others(t *testing.T) {
	assert.Equal(t, cards.BrandOthers, cards.Brand("Z293077338195127"))
}
