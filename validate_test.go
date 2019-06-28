package cards_test

import (
	"testing"

	"github.com/reddotpay/cards"
	"github.com/stretchr/testify/assert"
)

func TestCard_ValidateNumber_Visa_4111111111111111(t *testing.T) {
	assert.True(t, cards.ValidateNumber("4111111111111111"))
}

func TestCard_ValidateNumber_Visa_4012888888881881(t *testing.T) {
	assert.True(t, cards.ValidateNumber("4012888888881881"))
}

func TestCard_ValidateNumber_AmericanExpress_378282246310005(t *testing.T) {
	assert.True(t, cards.ValidateNumber("378282246310005"))
}

func TestCard_ValidateNumber_AmericanExpress_371449635398431(t *testing.T) {
	assert.True(t, cards.ValidateNumber("371449635398431"))
}

func TestCard_ValidateNumber_AmericanExpressCorporate_378734493671000(t *testing.T) {
	assert.True(t, cards.ValidateNumber("378734493671000"))
}

func TestCard_ValidateNumber_DinersClub_30569309025904(t *testing.T) {
	assert.True(t, cards.ValidateNumber("30569309025904"))
}

func TestCard_ValidateNumber_DinersClub_38520000023237(t *testing.T) {
	assert.True(t, cards.ValidateNumber("38520000023237"))
}

func TestCard_ValidateNumber_Discover_6011000990139424(t *testing.T) {
	assert.True(t, cards.ValidateNumber("6011000990139424"))
}

func TestCard_ValidateNumber_JCB_3530111333300000(t *testing.T) {
	assert.True(t, cards.ValidateNumber("3530111333300000"))
}

func TestCard_ValidateNumber_JCB_3566002020360505(t *testing.T) {
	assert.True(t, cards.ValidateNumber("3566002020360505"))
}

func TestCard_ValidateNumber_MasterCard_5555555555554444(t *testing.T) {
	assert.True(t, cards.ValidateNumber("5555555555554444"))
}

func TestCard_ValidateNumber_MasterCard_5105105105105100(t *testing.T) {
	assert.True(t, cards.ValidateNumber("5105105105105100"))
}

func TestCard_ValidateNumber_UnionPay_6293077338195127(t *testing.T) {
	assert.True(t, cards.ValidateNumber("6293077338195127"))
}

func TestCard_ValidateNumber_InvalidCards(t *testing.T) {
	assert.False(t, cards.ValidateNumber("5111111111111111"))
	assert.False(t, cards.ValidateNumber("41111111111111111"))
}

func BenchmarkCard_ValidateNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cards.ValidateNumber("4111111111111111")
	}
}
