package stats

import (
	"fmt"
	"github.com/AbduhamidTojiboev/go-bank/v2/pkg/types"
	"reflect"
	"testing"
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

func TestPeriodsDynamic_empty(t *testing.T) {
	paymentFirst := map[types.Category]types.Money{}
	paymentSecond := map[types.Category]types.Money{}

	result := PeriodsDynamic(paymentFirst, paymentSecond)

	if len(result) != 0 {
		t.Error("result len != 0")
	}
}

func TestPeriodsDynamic_success(t *testing.T) {
	paymentFirst := map[types.Category]types.Money{
		"auto": 10,
		"food": 20,
	}
	paymentSecond := map[types.Category]types.Money{
		"auto": 5,
		"food": 3,
	}
	expected := map[types.Category]types.Money{
		"auto": -5,
		"food": -17,
	}

	result := PeriodsDynamic(paymentFirst, paymentSecond)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}
}

func TestPeriodsDynamic_success_second(t *testing.T) {
	paymentFirst := map[types.Category]types.Money{
		"auto": 10,
	}
	paymentSecond := map[types.Category]types.Money{
		"auto": 5,
		"food": 3,
	}
	expected := map[types.Category]types.Money{
		"auto": -5,
		"food": 3,
	}

	result := PeriodsDynamic(paymentFirst, paymentSecond)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}
}
