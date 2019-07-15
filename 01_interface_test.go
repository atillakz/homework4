package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestPayWallet(t *testing.T) {

	amount := 50

	payment := 60

	w := &Wallet{amount}

	w.Pay(payment)

	assert.Equal(t, w.funds, amount-payment)

}

func TestPayByWalletGivesError(t *testing.T) {
	amount := 50
	payment := 60
	w := &Wallet{amount}
	err := w.Pay(payment)
	assert.Error(t, err)
}

func TestBuyByWallet(t *testing.T) {
	amount := 50
	payment := 60
	w := &Wallet{amount}

	_ = Buy(w, payment)

	assert.Equal(t, w.funds, amount-payment)
}

func TestBuyByCreditCard(t *testing.T) {
	amount := 100
	payment := 50

	c := &CreditCard{amount, "Zhannur", time.Now(), nil}

	_ = Buy(c, payment)

	assert.Equal(t, c.funds, amount-payment)
}

func TestCheckAndBuy(t *testing.T) {

	amount := 50
	b := &Bitcoin{amount, "Rus"}

	CheckAndBuy(b, amount)

	assert.Equal(t, b.funds, 0)
}

func TestCheckPaymentType(t *testing.T) {
	amount := 100

	c := &Bitcoin{amount, "Rusya"}

	s := CheckPaymentType(c)

	fmt.Println(s)
}

func TestBitcoin_Pay(t *testing.T) {

	amount := 50

	payment := 20

	b := &Bitcoin{amount, "Rusya"}

	b.Pay(payment)

	assert.Equal(t, b.funds, amount-payment)

}

func TestConverter(t *testing.T) {

	amount := 200

	amountInTenge := amount * 400

	b := &Bitcoin{amount, "Rusya"}

	b.Tracker(400)

	assert.Equal(t, b.funds, amountInTenge)

}
