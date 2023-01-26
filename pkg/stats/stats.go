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

func PeriodsDynamic(first map[types.Category]types.Money,
	second map[types.Category]types.Money) map[types.Category]types.Money {
	result := make(map[types.Category]types.Money)

	for category, money := range first {
		result[category] = second[category] - money
	}

	for category, money := range second {
		result[category] = money - first[category]
	}

	return result
}
