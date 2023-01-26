package stats

import (
	"fmt"
	"github.com/AbduhamidTojiboev/go-bank/v2/pkg/types"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			ID:       1,
			Amount:   10_00,
			Category: "qr",
		},
		{
			ID:       2,
			Amount:   150_00,
			Category: "qr",
		},
		{
			ID:       3,
			Amount:   110_00,
			Category: "shop",
		},
	}

	fmt.Println(Avg(payments))
	//Output: 9000
}

func ExampleTotalInCategory() {
	payments := []types.Payment{
		{
			ID:       1,
			Amount:   10_00,
			Category: "qr",
		},
		{
			ID:       2,
			Amount:   150_00,
			Category: "qr",
		},
		{
			ID:       3,
			Amount:   110_00,
			Category: "shop",
		},
	}

	fmt.Println(TotalInCategory(payments, "qr"))
	//Output: 16000
}
