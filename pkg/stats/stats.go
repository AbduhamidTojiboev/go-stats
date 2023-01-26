package stats

import "github.com/AbduhamidTojiboev/go-bank/v2/pkg/types"

func Avg(payments []types.Payment) types.Money {
	var total types.Money = 0
	for _, payment := range payments {
		total += payment.Amount
	}

	return total / types.Money(len(payments))
}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	var total types.Money = 0

	for _, payment := range payments {
		if payment.Category == category {
			total += payment.Amount
		}
	}

	return total
}
