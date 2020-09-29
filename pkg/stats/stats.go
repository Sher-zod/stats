package stats

import (
	"github.com/Sher-zod/bank/pkg/types"
)

//Avg расчитовает среднюю сумму платежа ssss
func Avg(payments []types.Payment) types.Money {
	countPayments := types.Money(len(payments))
	sumPaymenys := types.Money(0)
	for _, payment := range payments {
		moneyPayments := payment.Amount
		sumPaymenys += moneyPayments
	}
	return sumPaymenys / countPayments
}
