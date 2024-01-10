package main

import (
	"fmt"
	"sync"
)

type Income struct {
	Source string
	Amount int
}

var wg sync.WaitGroup

func main() {

	var bankBalance int
	var balance sync.Mutex

	fmt.Println("Initial bank balance", bankBalance)

	incomes := []Income{
		{Source: "Main Job", Amount: 600},
		{Source: "Gifts", Amount: 20},
		{Source: "Part time Job", Amount: 60},
		{Source: "Main Job", Amount: 100},
	}

	wg.Add(len(incomes))

	for i, income := range incomes {

		go func(i int, income Income) {
			defer wg.Done()

			for week := 1; week <= 52; week++ {
				balance.Lock()
				temp := bankBalance
				temp += income.Amount
				bankBalance = temp
				balance.Unlock()

				fmt.Printf("On week %v, you earned $%v.00 from %v \n", week, income.Amount, income.Source)
			}
		}(i, income)
	}

	wg.Wait()

	fmt.Println("Final bank balance: ", bankBalance)
}
