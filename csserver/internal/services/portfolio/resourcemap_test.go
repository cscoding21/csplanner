package portfolio_test

import (
	"csserver/internal/services/portfolio"
	"csserver/internal/testobjects"
	"fmt"
	"testing"
)

func TestFlattenPortfolio(t *testing.T) {
	rm := testobjects.GetResourceMap()
	port := testobjects.GetTestPortfolio()

	allocationMap, err := portfolio.FlattenPortfolio(port, rm)
	if err != nil {
		t.Error(err)
	}

	for _, m := range allocationMap {
		fmt.Println(m)
	}
}
