package portfolio_test

import (
	"csserver/internal/testobjects"
	"fmt"
	"testing"
)

func TestGetDateRange(t *testing.T) {
	p := testobjects.GetTestPortfolio()

	if len(p.Schedule) == 0 {
		t.Error("test schedule did not load correctly")
	}

	start, end := p.GetDateRange()

	fmt.Printf("%v - %v\n", start, end)
	fmt.Printf("Project count - %v\n", len(p.ProjectMap))
	fmt.Printf("Schedule count - %v\n", len(p.Schedule))
}

func TestBalancePortfolio(t *testing.T) {
	// p := testobjects.GetTestPortfolio()
	// ctx := context.Background()

}
