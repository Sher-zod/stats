package stats

import (
	"fmt"

	"github.com/Sher-zod/bank/pkg/types"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			ID:       1,
			Amount:   100_00,
			Category: "A",
		},
		{
			ID:       2,
			Amount:   200_00,
			Category: "B",
		},
		{
			ID:       3,
			Amount:   300_00,
			Category: "C",
		},
	}

	fmt.Println(Avg(payments))

	//Output: 20000
}

func ExampleTotalInCategory() {
	payments := []types.Payment{
		{
			ID:       2,
			Amount:   53_00,
			Category: "C",
		},
		{
			ID:       1,
			Amount:   51_00,
			Category: "C",
		},
		{
			ID:       3,
			Amount:   52_00,
			Category: "R",
		},
	}

	fmt.Println(TotalInCategory(payments, "C"))

	//Output: 10400
}
