package main

import (
	"fmt"
	"time"
)

func main()  {


}


type Wallet struct {

	funds int
}

func (w *Wallet) Pay(amount int) error  {

	if (amount > w.funds){

		return fmt.Errorf("No money no honey")
	}

	w.funds -= amount

	return nil

}

type Payer interface {

	Pay(int) error
}

func Buy(p Payer, amount int) error  {

	err := p.Pay(amount)

	if err != nil{
		return err
	}

	return nil
	
}

type CreditCard struct {
	funds      int
	owner      string
	expireTime time.Time
	bonuses    []Bonus
}

type Bonus struct {

	name string
	description string
}

func (c *CreditCard) Pay(amount int) error{

	if amount > c.funds{

		return fmt.Errorf("No money No honey")
	}

	c.funds -= amount

	return nil
}


type Funder interface {
	GetFunds() int
}

type PayFunder interface {
	Payer
	Funder
}


func (w *Wallet) GetFunds() int {
	return w.funds
}

func (c *CreditCard) GetFunds() int {
	return c.funds
}

func (b *Bitcoin) GetFunds() int {
	return b.funds
}

func (p *Bitcoin) Tracker(exchange int) error {

	p.funds *= exchange

	return nil
}

func CheckAndBuy(p PayFunder, amount int) error {
	if p.GetFunds() <= 0 {
		fmt.Println("пополните счет")
		return nil
	}

	err := p.Pay(amount)
	if err != nil {
		return err
	}
	return nil
}

func CheckPaymentType(p Payer) interface{}  {

	switch p {
	case p.(*Wallet):
		fmt.Println("Paying with cash")
	case p.(*CreditCard):
		fmt.Println("Paying with credit card")
	case p.(*Bitcoin):
		fmt.Println("Paying with Bitcoin")
	}
	return nil

}


type Bitcoin struct {

	funds int

	owner string

}

func (b *Bitcoin) Pay(amount int)  error {

	if (amount > b.funds){
		return fmt.Errorf("No money No honey")
	}

	b.funds -= amount

	return nil

}

