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

	fmt.Println(start)
	fmt.Println(end)
}

func TestGetCurrentPortfolio(t *testing.T) {

}
