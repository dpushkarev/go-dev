package main

import (
	"fmt"
)

type Payer interface {
	Pay(int) error
}

type Walet struct {
	Cache int
}

func (w *Walet) Pay(amount int) error {
	if w.Cache < amount {
		return fmt.Errorf("no money")
	}
	w.Cache -= amount
	return nil
}

func Buy(p Payer)  {
	err := p.Pay(10)
	if err != nil {
		panic(err)
	}
	fmt.Println("Спасибо за покупку")
}

func main() {
	myWallet := &Walet{Cache:11}
	Buy(myWallet)
}

