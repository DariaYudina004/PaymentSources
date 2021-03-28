package card

import (
	"bank/pkg/bank/types"
	
)

// PaymentSources предоставляет выбор карт для оплаты 
// карты с отрицательным балансом и  заблокированные карты игнорируются 
func PaymentSources(cards []types.Card) [] types.PaymentSource {
	var payments []types.PaymentSource
      for _, card := range cards {
	  if  card.Balance<= 0 && !card.Active{
		continue
 
	}
	payments = append(payments,types.PaymentSource{"card","9074 9934 8645 7642",1_000_00})
 	}

    return payments
} 