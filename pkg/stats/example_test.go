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
